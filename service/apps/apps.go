package apps

import (
	"context"
	"errors"
	"github.com/NubeIO/git/pkg/git"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/service/apps/app"
	"os"
)

var err error

type Apps struct {
	Token              string      `json:"git_token"`
	Version            string      `json:"tag"`
	AppName            string      `json:"app_name"`
	DownloadPath       string      `json:"download_path"`   // home/user/downloads
	RubixRootPath      string      `json:"rubix_root_path"` // /data
	AppPath            string      `json:"app_path"`        // RubixRootPath/rubix-wires
	ServiceFileTmpPath string      `json:"service_file_tmp_path"`
	AssetZipName       string      `json:"asset_zip_name"`
	Perm               os.FileMode `json:"-"`
	gitClient          *git.Client
	GeneratedApp       *app.Service `json:"-"`
	serviceName        string       // nubeio-rubix-wires
}

func New(inst *Apps, rubixApp string) (*Apps, error) {
	if inst == nil {
		return nil, errors.New("type apps must not be nil")
	}
	if rubixApp == "" {
		return nil, errors.New("no app was passed in, try ff, flow or flow-framework")
	}
	if inst.Perm == 0 {
		inst.Perm = 0700
	}
	installer, err := app.New(&app.App{
		AppName:       inst.AppName,
		Version:       inst.Version,
		RubixRootPath: inst.RubixRootPath,
	}, rubixApp)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	selectApp, err := installer.SelectApp()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	if installer.ServiceName == "" {
		return nil, errors.New("service-name must not be nil, try nubeio-rubix-wires")
	}
	inst.serviceName = installer.ServiceName
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
