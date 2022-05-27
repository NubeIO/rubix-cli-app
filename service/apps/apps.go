package apps

import (
	"context"
	"fmt"
	"github.com/NubeIO/git/pkg/git"
	"github.com/NubeIO/lib-systemctl-go/builder"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	log "github.com/sirupsen/logrus"
	"time"
)

var err error

func New(inst *Installer) *Installer {
	opts := &git.AssetOptions{
		Owner:    inst.Owner,
		Repo:     inst.Repo,
		Tag:      inst.Tag,
		Arch:     inst.Arch,
		DestPath: inst.DownloadPath,
		Target:   inst.DownloadPathSubDir,
	}
	ctx := context.Background()
	inst.gitClient = git.NewClient(inst.Token, opts, ctx)
	return inst
}

type RespDownload struct {
	AssetName string `json:"asset_name"`
	GitError  string `json:"git_error"`
}

type RespBuilder struct {
	BuilderErr string `json:"builder_err"`
}

type RespInstall struct {
	InstallResp *ctl.InstallResp `json:"install_resp"`
}

type Installer struct {
	Token              string                  `json:"token"`
	Owner              string                  `json:"owner"`
	Repo               string                  `json:"repo"`
	Arch               string                  `json:"arch"`
	Tag                string                  `json:"tag"`
	DownloadPath       string                  `json:"download_path"`         //home/user
	DownloadPathSubDir string                  `json:"download_path_sub_dir"` //home/user /bios
	ServiceFile        *builder.SystemDBuilder `json:"service"`
	gitClient          *git.Client
}

func (inst *Installer) GitDownload() (*RespDownload, error) {
	ret := &RespDownload{}
	//download and unzip to /data
	resp, err := inst.gitClient.DownloadOnly()
	if err != nil {
		ret.GitError = err.Error()
		return ret, err
	}
	ret.AssetName = resp.ReleaseAsset.GetName()
	return ret, nil
}

func (inst *Installer) GenerateServiceFile() (*RespBuilder, error) {
	ret := &RespBuilder{}
	newService := "nubeio-rubix-bios"
	description := "BIOS comes with default OS, non-upgradable"
	user := "root"
	directory := "/data/rubix-bios-app"
	execCmd := "/data/rubix-bios-app/rubix-bios -p 1615 -g /data/rubix-bios -d data -c config -a apps --prod --auth  --device-type amd64 --token 1234"

	bld := &builder.SystemDBuilder{
		Description:      description,
		User:             user,
		WorkingDirectory: directory,
		ExecStart:        execCmd,
		SyslogIdentifier: "rubix-bios",
		WriteFile: builder.WriteFile{
			Write:    true,
			FileName: newService,
			Path:     "/data",
		},
	}

	err = bld.Build()
	if err != nil {
		ret.BuilderErr = err.Error()
		return ret, err
	}
	return ret, nil
}

//InstallService a new linux service
//	- service: the service name (eg: rubix-bios)
//	- path: the service file path and name (eg: "/tmp/rubix-bios.service")
func (inst *Installer) InstallService(service, path string) (*RespInstall, error) {
	ret := &RespInstall{}
	//path := "/tmp/nubeio-rubix-bios.service"
	timeOut := 30
	ser := ctl.New(service, path)
	opts := systemctl.Options{Timeout: timeOut}
	installOpts := ctl.InstallOpts{
		Options: opts,
	}
	ser.InstallOpts = installOpts
	ret.InstallResp = ser.Install()
	fmt.Println("full install error", err)
	if err != nil {
		fmt.Println("full install error", err)
	}

	time.Sleep(8 * time.Second)

	status, err := systemctl.Status(service, systemctl.Options{})
	if err != nil {
		log.Errorf("service found: %s: %v", service, err)
	}
	fmt.Println(status)
	return ret, nil
}
