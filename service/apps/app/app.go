package app

import (
	"errors"
)

var this *Service

type App struct {
	InstallApp    string
	AppName       string // ff, flow-framework or flow
	Version       string //version must be the installed version as in v0.0.1
	RubixRootPath string // /data
	AppsPath      string // /data/rubix-apps/install
	AppPath       string // data/flow-framework
}

var (
	appName    = ""
	appVersion = ""
)

func New(app *App, rubixApp string) (*Service, error) {
	if app == nil {
		return nil, errors.New("type app must not be nil")
	}
	if app.Version == "" {
		return nil, errors.New("app version must not be nil")
	}
	service := &Service{}
	if app != nil { // override install dir
		if app.RubixRootPath != "" {
			rootDir = app.RubixRootPath
		}
		if app.AppsPath != "" {
			appsPath = app.AppsPath
		}
	}
	appName = rubixApp
	appVersion = app.Version
	this = service
	return this, nil
}
