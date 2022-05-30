package dbase

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

type unInstallLog struct {
	GetApp          string         `json:"get_app"`
	GetAppFromStore string         `json:"get_app_from_store"`
	InitApp         string         `json:"init_app"`
	RemoveService   string         `json:"remove_service"`
	DeleteFromDB    string         `json:"delete_from_db"`
	VersionRemoved  string         `json:"version_removed"`
	Service         *ctl.RemoveRes `json:"service"`
}

type UnInstallResponse struct {
	Message      string       `json:"message"`
	Error        string       `json:"error"`
	UnInstallLog unInstallLog `json:"log"`
}

func (db *DB) UnInstallApp(body *App) (*UnInstallResponse, error) {
	resp := &UnInstallResponse{}
	remove, err := db.unInstallApp(body)
	if err != nil {
		resp.Error = err.Error()
	}
	if remove.UnInstallLog.VersionRemoved == "" {
		resp.Message = "the app was not installed but run uninstall ok anyway"
	} else {
		resp.Message = fmt.Sprintf("removed app: %s version: %s", body.AppName, remove.UnInstallLog.VersionRemoved)
	}
	resp.Error = "no errors"
	resp.UnInstallLog = remove.UnInstallLog
	return resp, err
}

func (db *DB) unInstallApp(body *App) (*UnInstallResponse, error) {
	resp := &UnInstallResponse{}

	getApp, err := db.GetAppByName(body.AppName)
	if err != nil {
		resp.UnInstallLog.GetApp = "failed to get app, but will still try to uninstall"
	} else {
		resp.UnInstallLog.GetApp = "ok"
	}

	appStoreName := body.AppName
	if getApp != nil {
		appStoreName = getApp.AppStoreName
	}
	appStore, err := db.GetAppStoreByName(appStoreName)
	if err != nil {
		resp.UnInstallLog.GetAppFromStore = "failed to get service name from app store so exit"
		return resp, err
	}
	resp.UnInstallLog.GetAppFromStore = selectAppStore

	var inst = &apps.Apps{
		Token:   body.Token,
		Perm:    apps.Permission,
		Version: body.Version,
		App:     appStore,
	}
	app, err := apps.New(inst)
	if err != nil {
		log.Errorln("new app: failed to init a new app", err)
		resp.UnInstallLog.InitApp = "new app: failed to init a new app"
		return resp, err
	}
	resp.UnInstallLog.InitApp = "ok"
	service, err := app.UninstallService(appStore.ServiceName)
	if err != nil {
		resp.UnInstallLog.RemoveService = fmt.Sprintf("failed to remove service: %s", appStore.ServiceName)
		return nil, err
	}
	if service.Stop {
		resp.UnInstallLog.RemoveService = "ok"
	} else {
		resp.UnInstallLog.RemoveService = fmt.Sprintf("service was not found: %s", appStore.ServiceName)
	}
	resp.UnInstallLog.Service = service
	if getApp != nil {
		resp.UnInstallLog.VersionRemoved = getApp.InstalledVersion
		_, err = db.DeleteApp(getApp.UUID)
		if err != nil {
			resp.UnInstallLog.DeleteFromDB = "failed to delete the app from the db"
			return nil, err
		}
		resp.UnInstallLog.DeleteFromDB = "delete ok from the db"
	} else {
		resp.UnInstallLog.DeleteFromDB = "app was not found so it could not be deleted"
	}
	return resp, err

}
