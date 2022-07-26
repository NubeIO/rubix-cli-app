package apps

import (
	"context"
	"errors"
	"fmt"
	"github.com/NubeIO/git/pkg/git"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-edge/pkg/logger"
	"os"
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
	if inst.App.DownloadPath == "/tmp" {
		logger.Logger.Infof("make download dir: was tmp dir so skip \n")
		return nil
	}
	if !dirs.DirExists(inst.App.DownloadPath) {
		logger.Logger.Errorf("no dir exists %s \n", inst.App.DownloadPath)
		err := dirs.MkdirAll(inst.App.DownloadPath, os.FileMode(inst.Perm))
		if err != nil {
			logger.Logger.Errorf("make download dir:  failed to make new dir %s \n", inst.App.DownloadPath)
			return err
		}
		logger.Logger.Infof("make download dir:  made new dir:%s \n", inst.App.DownloadPath)
	}
	logger.Logger.Infof("make download dir:  existing dir to download zip:%s \n", inst.App.DownloadPath)
	return nil
}

func (inst *Apps) MakeInstallDir() error {
	//action, err := inst.Stop(DefaultTimeout)
	//if err != nil {
	//	logger.Logger.Errorf("stop app:%s failed err:%s \n", inst.App.Name, err.Error())
	//}
	//if action.Ok {
	//	logger.Logger.Infof("stop app:%s  it was running \n", inst.App.Name)
	//} else {
	//	logger.Logger.Infof("stop app:%s  failed or was not running msg:%s \n", inst.App.Name, action.Message)
	//}

	installPath := fmt.Sprintf(inst.App.AppsPath) // /data/rubix-apps/installed/flow-framework
	if !dirs.DirExists(installPath) {
		logger.Logger.Errorf("no dir exists %s \n", installPath)
		err := dirs.MkdirAll(installPath, os.FileMode(inst.Perm))
		if err != nil {
			logger.Logger.Errorf("install dir: failed to make new dir %s \n", installPath)
			return err
		}
	}
	logger.Logger.Infof("make-install-dir: existing install dir existed:%s \n", installPath)
	return nil
}

func (inst *Apps) BuildExists(manualAssetZipName string) error {
	assetZipName := ""
	if manualAssetZipName != "" {
		assetZipName = manualAssetZipName
	} else {
		assetZipName = inst.App.AssetZipName
	}
	if assetZipName == "" {
		return errors.New("asset zip folder name can not be empty")
	}

	installPath := inst.App.AppsPath
	zipFileAndPath := fmt.Sprintf("%s/%s", inst.App.DownloadPath, assetZipName)
	exists := fileutils.New().FileExists(zipFileAndPath)
	if !exists {
		logger.Logger.Errorf("unzip build: the existing downloaded build dose not exist source:%s  dest:%s\n", inst.App.DownloadPath, installPath)
		return errors.New(fmt.Sprintf("unzip build: the existing downloaded build dose not exist source:%s  dest:%s", inst.App.DownloadPath, installPath))
	}
	return nil
}

func (inst *Apps) UnpackBuild(manualAssetZipName string) error {
	assetZipName := ""
	if manualAssetZipName != "" {
		assetZipName = manualAssetZipName
	} else {
		assetZipName = inst.App.AssetZipName
	}
	if assetZipName == "" {
		return errors.New("asset zip folder name can not be empty")
	}

	installPath := inst.App.AppsPath
	zipFileAndPath := fmt.Sprintf("%s/%s", inst.App.DownloadPath, assetZipName)

	_, err = dirs.UnZip(zipFileAndPath, installPath, os.FileMode(inst.Perm))
	if err != nil {
		logger.Logger.Errorf("unzip build: failed to unzip source:%s  dest:%s  error:%s\n", inst.App.DownloadPath, installPath, err.Error())
		return err
	} else {
		logger.Logger.Infof("unzip build: from:%s \n", zipFileAndPath)
		logger.Logger.Infof("unzip build: to:%s \n", installPath)
	}
	return nil
}

func (inst *Apps) CleanUp(manualAssetZipName string) error {
	assetZipName := ""
	if manualAssetZipName != "" {
		assetZipName = manualAssetZipName
	} else {
		assetZipName = inst.App.AssetZipName
	}
	if assetZipName == "" {
		return errors.New("asset zip folder name can not be empty")
	}
	zipFileAndPath := fmt.Sprintf("%s/%s", inst.App.DownloadPath, assetZipName)
	err = dirs.Rm(zipFileAndPath)
	if err != nil {
		logger.Logger.Errorf("delete zip: failed to unzip source:%s  error:%s\n", inst.App.DownloadPath, err.Error())
		return err
	} else {
		logger.Logger.Infof("clean-up delete zip: ok:%s \n", inst.App.DownloadPath)
	}
	return nil
}

type RespBuilder struct {
	BuilderErr string `json:"builder_err"`
}

func (inst *Apps) GitDownload(destination string) (*git.DownloadResponse, error) {
	if inst.Token == "" {
		return nil, errors.New("git token can not be empty")
	}
	opts := &git.AssetOptions{
		Owner: inst.App.Owner,
		Repo:  inst.App.Repo,
		Tag:   inst.Version,
		Arch:  inst.App.Arch,
	}
	ctx := context.Background()
	gitClient = git.NewClient(inst.Token, opts, ctx)
	download, err := gitClient.Download(destination)
	if err != nil {
		return nil, err
	}
	inst.App.AssetZipName = download.AssetName
	logger.Logger.Infof("git downloaded full asset name: %s", download.AssetName)
	return download, err
}

func (inst *Apps) GenerateServiceFile(app *Apps, tmpFilePath string) (*RespBuilder, error) {
	ret := &RespBuilder{}
	newService := app.App.ServiceName
	description := app.App.ServiceDescription
	user := app.App.RunAsUser
	directory := app.App.ServiceWorkingDirectory
	execCmd := app.App.ServiceExecStart

	bld := &builder.SystemDBuilder{
		ServiceName:      app.App.ServiceName,
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

	err = bld.Build(os.FileMode(inst.Perm))
	if err != nil {
		ret.BuilderErr = err.Error()
		return ret, err
	}
	return ret, nil
}

// InstallService a new linux service
//	- service: the service name (eg: nubeio-rubix-wires)
//	- path: the service file path and name (eg: "/tmp/rubix-bios.service")
func (inst *Apps) InstallService(service, tmpServiceFile string) (*ctl.InstallResp, error) {
	resp := &ctl.InstallResp{}

	// path := "/tmp/nubeio-rubix-bios.service"
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
		logger.Logger.Errorf("service found or failed to check IsRunning: %s: %v", service, err)
		return nil, err
	}
	logger.Logger.Infof("service: %s: isActive: %t status: %s", service, active, status)
	return resp, nil
}
