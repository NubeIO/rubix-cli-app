package dbase

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

type InstallApp struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	Token   string `json:"token"`
}

type InstallResponse struct {
	ErrorMessage    string `json:"error_message"`
	GetAppFromStore string `json:"get_app_from_store"`
	MakeApp         string `json:"make_app"`
	MakeDownload    string `json:"make_download"`
	GitDownload     string `json:"git_download"`
	MakeInstallDir  string `json:"make_install_dir"`
	UnpackBuild     string `json:"unpack_build"`
	GenerateService string `json:"generate_service"`
	InstallService  string `json:"install_service"`
	CleanUp         string `json:"clean_up"`
}

// ok messages
const (
	selectAppStore  = "ok"
	makeDownload    = "ok"
	gitDownload     = "ok"
	makeNewApp      = "ok"
	makeInstallDir  = "ok"
	unpackBuild     = "ok"
	generateService = "ok"
	installService  = "ok"
	cleanUp         = "ok"
)

// not ok messages
const (
	selectAppStoreErr  = "this app is was not found in the app store, try flow-framework, rubix-wires"
	makeDownloadErr    = "issue on trying to make the path to download the zip folder"
	gitDownloadErr     = "error on git download"
	makeNewAppErr      = "failed to make a new app"
	makeInstallDirErr  = "unable to make the install dir for the app"
	unpackBuildErr     = "unable to unzip the build"
	generateServiceErr = "unable to make the app service file"
	installServiceErr  = "unable to install the app"
	cleanUpErr         = "unable to clean up the install"
)

func (db *DB) InstallApp(body *InstallApp) (*InstallResponse, error) {
	app, err := db.installApp(body)
	if err != nil {
		app.ErrorMessage = err.Error()
		return app, err
	}
	return app, err
}

func (db *DB) installApp(body *InstallApp) (*InstallResponse, error) {

	resp := &InstallResponse{
		ErrorMessage: "no error",
	}

	appStore, err := db.GetAppImageByName(body.AppName)
	if err != nil {
		resp.GetAppFromStore = selectAppStoreErr
		return resp, err
	}
	resp.GetAppFromStore = selectAppStore
	installedApp := &apps.InstalledApp{
		AppStoreName:     appStore.Name,
		AppStoreUUID:     appStore.UUID,
		InstalledVersion: body.Version,
	}

	app, err := db.AddApp(installedApp)
	if err != nil {
		resp.MakeApp = makeNewAppErr
		return resp, err
	}
	resp.MakeApp = makeNewApp
	log.Infof("add new app name:%s", app.AppStoreName)

	var inst = &apps.Apps{
		Token:   body.Token,
		Perm:    0700,
		Version: body.Version,
		App:     appStore,
	}
	newApp, err := apps.New(inst, appStore.Name)
	if err != nil {
		log.Errorln("new app: failed to init a new app", err)
		return resp, err
	}

	if err = inst.MakeDownloadDir(); err != nil {
		resp.MakeDownload = makeDownloadErr
		return resp, err
	}
	resp.MakeDownload = makeDownload
	if _, err = newApp.GitDownload(inst.App.DownloadPath); err != nil {
		log.Errorf("git: download error %s \n", err.Error())
		resp.GitDownload = gitDownloadErr
		return resp, err
	}
	resp.GitDownload = gitDownload
	if err = inst.MakeInstallDir(); err != nil {
		resp.MakeInstallDir = makeInstallDirErr
		return resp, err
	}
	resp.MakeInstallDir = makeInstallDir
	if err = inst.UnpackBuild(); err != nil {
		resp.UnpackBuild = unpackBuildErr
		return resp, err
	}
	resp.UnpackBuild = unpackBuild
	tmpFileDir := newApp.App.DownloadPath
	if _, err = newApp.GenerateServiceFile(newApp, tmpFileDir); err != nil {
		log.Errorf("make service file build: failed error:%s \n", err.Error())
		resp.GenerateService = generateServiceErr
		return resp, err
	}
	resp.GenerateService = generateService
	tmpServiceFile := fmt.Sprintf("%s/%s.service", tmpFileDir, newApp.App.ServiceName)
	if _, err = newApp.InstallService(newApp.App.ServiceName, tmpServiceFile); err != nil {
		resp.InstallService = installServiceErr
		return resp, err
	}
	resp.InstallService = installService
	if err = inst.CleanUp(); err != nil {
		resp.CleanUp = cleanUpErr
		return resp, err
	}
	resp.CleanUp = cleanUp
	return resp, err

}
