package apps

import (
	"context"
	"errors"
	"fmt"
	"github.com/NubeIO/git/pkg/git"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/service/apps/app"
	"os"
	"time"
)

var err error

type Apps struct {
	Token              string      `json:"git_token"`
	Version            string      `json:"tag"`
	AppName            string      `json:"app_name"`
	DownloadPath       string      `json:"download_path"`   // home/user/downloads
	RubixRootPath      string      `json:"rubix_root_path"` // /data
	InstallPath        string      `json:"install_path"`    // RubixRootPath/rubix-apps
	ServiceFileTmpPath string      `json:"service_file_tmp_path"`
	ServiceName        string      `json:"service_name"` // nubeio-rubix-wires
	Perm               os.FileMode `json:"-"`
	gitClient          *git.Client
	GeneratedApp       *app.Service `json:"-"`
}

func New(inst *Apps) (*Apps, error) {
	if inst == nil {
		return nil, errors.New("type apps must not be nil")
	}
	if inst.ServiceName == "" {
		return nil, errors.New("service-name must not be nil, try nubeio-rubix-wires")
	}
	if inst.Perm == 0 {
		inst.Perm = 0700
	}
	installer, err := app.New(&app.App{
		AppName:       inst.AppName,
		Version:       inst.Version,
		RubixRootPath: inst.RubixRootPath,
		InstallPath:   inst.InstallPath,
	})
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	selectApp, err := installer.SelectApp()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	opts := &git.AssetOptions{
		Owner: selectApp.Owner,
		Repo:  selectApp.Repo,
		Tag:   selectApp.Version,
		Arch:  selectApp.Arch,
	}
	ctx := context.Background()
	inst.gitClient = git.NewClient(inst.Token, opts, ctx)
	inst.GeneratedApp = selectApp
	return inst, err
}

type RespBuilder struct {
	BuilderErr string `json:"builder_err"`
}

func (inst *Apps) GitDownload(destination string) (*git.DownloadResponse, error) {
	return inst.gitClient.Download(destination)
}

func (inst *Apps) GenerateServiceFile(app *app.Service, tmpFilePath string) (*RespBuilder, error) {
	ret := &RespBuilder{}

	newService := app.ServiceName
	description := app.ServiceDescription
	user := app.RunAsUser
	directory := app.ServiceWorkingDirectory
	execCmd := app.ServiceExecStart

	bld := &builder.SystemDBuilder{
		ServiceName:      app.Name,
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

	err = bld.Build(0700)
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
