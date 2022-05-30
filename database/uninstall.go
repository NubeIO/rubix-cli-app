package dbase

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/ctl"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

type UnInstallResponse struct {
	GetApp          string         `json:"get_app"`
	GetAppFromStore string         `json:"get_app_from_store"`
	InitApp         string         `json:"init_app"`
	RemoveService   string         `json:"remove_service"`
	DeleteFromDB    string         `json:"delete_from_db"`
	Error           string         `json:"error"`
	Service         *ctl.RemoveRes `json:"service"`
}

func (db *DB) UnInstallApp(body *App) (*UnInstallResponse, error) {
	resp, err := db.unInstallApp(body)
	if err != nil {
		resp.Error = err.Error()
	}
	return resp, err
}

func (db *DB) unInstallApp(body *App) (*UnInstallResponse, error) {
	resp := &UnInstallResponse{}
	getApp, err := db.GetAppByName(body.AppName)
	if err != nil {
		resp.GetApp = "failed to get app, but will still try to uninstall"
	}
	resp.GetApp = "ok"
	appStoreName := body.AppName
	if getApp != nil {
		appStoreName = getApp.AppStoreName
	}
	appStore, err := db.GetAppImageByName(appStoreName)
	if err != nil {
		resp.GetAppFromStore = "failed to get service name from app store so exit"
		return resp, err
	}
	resp.GetAppFromStore = selectAppStore

	var inst = &apps.Apps{
		Token:   body.Token,
		Perm:    0700,
		Version: body.Version,
		App:     appStore,
	}
	app, err := apps.New(inst)
	if err != nil {
		log.Errorln("new app: failed to init a new app", err)
		resp.InitApp = "new app: failed to init a new app"
		return resp, err
	}
	resp.InitApp = "ok"
	service, err := app.UninstallService(appStore.ServiceName)
	if err != nil {
		resp.RemoveService = fmt.Sprintf("failed to remove service: %s", appStore.ServiceName)
		return nil, err
	}
	resp.RemoveService = "ok"
	resp.Service = service
	if getApp != nil {
		_, err = db.DeleteApp(getApp.UUID)
		if err != nil {
			resp.DeleteFromDB = "failed to delete the app from the db"
			return nil, err
		}
		resp.DeleteFromDB = "delete ok from the db"
	} else {
		resp.DeleteFromDB = "app was not found so it could not be deleted"
	}
	return resp, err

}
