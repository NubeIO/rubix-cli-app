package apps

import (
	"fmt"
	"github.com/NubeIO/git/pkg/git"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/service/apps/app"
	"time"
)

/*
- check the download dir exists and make it if not
- download the app
- stop the app if it was already running
- check install dir exists and make if not then unzip the build
- generate the service file
- install the service
- clean up download dir
*/

var dirs = fileutils.New()

func (inst *Apps) MakeDownloadDir() error {
	if !dirs.DirExists(inst.DownloadPath) {
		log.Errorf("no dir exists %s \n", inst.DownloadPath)
		err := dirs.MkdirAll(inst.DownloadPath, inst.Perm)
		if err != nil {
			log.Errorf("unzip build: failed to make new dir %s \n", inst.DownloadPath)
			return err
		}
		log.Infof("unzip build: made new dir:%s \n", inst.DownloadPath)
	}
	log.Infof("unzip build: existing dir to download zip:%s \n", inst.DownloadPath)
	return nil
}

func (inst *Apps) MakeInstallDir() error {
	action, err := inst.Stop(defaultTimeout)
	if err != nil {
		log.Errorf("stop app:%s failed err:%s \n", inst.AppName, err.Error())
		return err
	}
	if action.Ok {
		log.Infof("stop app:%s  it was running \n", inst.AppName)
	} else {
		log.Infof("stop app:%s  failed or was not running msg:%s \n", inst.AppName, action.Message)
	}

	installPath := fmt.Sprintf(inst.GeneratedApp.AppsPath) // /data/rubix-apps/installed/flow-framework
	if !dirs.DirExists(installPath) {
		log.Errorf("no dir exists %s \n", installPath)
		err := dirs.MkdirAll(installPath, inst.Perm)
		if err != nil {
			log.Errorf("install dir: failed to make new dir %s \n", installPath)
			return err
		}
	}
	log.Infof("install dir: existing install dir existed:%s \n", installPath)
	return nil
}

func (inst *Apps) UnpackBuild() error {
	installPath := inst.GeneratedApp.AppsPath
	zipFileAndPath := fmt.Sprintf("%s/%s", inst.DownloadPath, inst.AssetZipName)
	_, err = dirs.UnZip(zipFileAndPath, installPath, inst.Perm)
	if err != nil {
		log.Errorf("unzip build: failed to unzip source:%s  dest:%s  error:%s\n", inst.DownloadPath, installPath, err.Error())
		return err
	} else {
		log.Infof("unzip build: existing install dir existed:%s \n", installPath)
	}
	return nil
}

func (inst *Apps) CleanUp() error {
	zipFileAndPath := fmt.Sprintf("%s/%s", inst.DownloadPath, inst.AssetZipName)
	err = dirs.Rm(zipFileAndPath)
	if err != nil {
		log.Errorf("delete zip: failed to unzip source:%s  error:%s\n", inst.DownloadPath, err.Error())
		return err
	} else {
		log.Infof("delete zip: ok:%s \n", inst.DownloadPath)
	}
	return nil
}

type RespBuilder struct {
	BuilderErr string `json:"builder_err"`
}

func (inst *Apps) GitDownload(destination string) (*git.DownloadResponse, error) {
	download, err := inst.gitClient.Download(destination)
	inst.AssetZipName = download.AssetName
	return download, err
}

func (inst *Apps) GenerateServiceFile(app *app.Service, tmpFilePath string) (*RespBuilder, error) {
	ret := &RespBuilder{}
	newService := app.ServiceName
	description := app.ServiceDescription
	user := app.RunAsUser
	directory := app.ServiceWorkingDirectory
	execCmd := app.ServiceExecStart

	bld := &builder.SystemDBuilder{
		ServiceName:      app.AppName,
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: newService,
		WriteFile: builder.WriteFile{
			Write:    true,
			FileName: newService,
			Path:     tmpFilePath,
		},
	}

	err = bld.Build(inst.Perm)
	if err != nil {
		ret.BuilderErr = err.Error()
		return ret, err
	}
	return ret, nil
}

//InstallService a new linux service
//	- service: the service name (eg: nubeio-rubix-wires)
//	- path: the service file path and name (eg: "/tmp/rubix-bios.service")
func (inst *Apps) InstallService(service, tmpServiceFile string) (*ctl.InstallResp, error) {
	resp := &ctl.InstallResp{}

	//path := "/tmp/nubeio-rubix-bios.service"
	timeOut := 30
	ser := ctl.New(service, tmpServiceFile)
	ser.InstallOpts = ctl.InstallOpts{
		Options: systemctl.Options{Timeout: timeOut},
	}
	err = ser.TransferFile()
	if err != nil {
		fmt.Println("full install error", err)
		return nil, err
	}

	resp = ser.Install()
	if err != nil {
		fmt.Println("full install error", err)
		return nil, err
	}
	time.Sleep(8 * time.Second)
	active, status, err := systemctl.IsRunning(service, systemctl.Options{})
	if err != nil {
		log.Errorf("service found or failed to check IsRunning: %s: %v", service, err)
		return nil, err
	}
	log.Infof("service: %s: isActive: %t status: %s", service, active, status)
	return resp, nil
}
