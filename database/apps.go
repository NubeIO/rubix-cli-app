package dbase

import (
	"errors"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/lib-uuid/uuid"
	"gthub.com/NubeIO/rubix-cli-app/pkg/logger"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

func (db *DB) GetApps() ([]*apps.App, error) {
	var m []*apps.App
	if err := db.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

type AppStats struct {
	App   *apps.App             `json:"app"`
	Stats systemctl.SystemState `json:"stats"`
}

func (db *DB) AppStats(body *apps.App) (*AppStats, error) {
	var stats *AppStats
	appStore, getApp, err := db.GetAppAndStore(body)
	if err != nil {
		return nil, err
	}
	service, err := initAppService(appStore.ServiceName)
	if err != nil {
		return nil, err
	}
	status, err := service.ServiceStats(apps.DefaultTimeout)
	if err != nil {
		return nil, err
	}
	stats.App = getApp
	stats.Stats = status
	return stats, nil
}

func (db *DB) GetApp(uuid string) (*apps.App, error) {
	var m *apps.App
	if err := db.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (db *DB) GetAppAndStore(body *apps.App) (*apps.Store, *apps.App, error) {
	var app *apps.App
	if body.UUID == "" {
		appByName, err := db.GetAppByName(body.AppStoreName)
		if err != nil {
			return nil, nil, errors.New("app not found by name")
		}
		app = appByName
	} else {
		appById, err := db.GetApp(body.UUID)
		if err != nil {
			return nil, nil, errors.New("app not found by id")
		}
		app = appById
	}
	appStore, err := db.GetAppStore(app.AppStoreUUID)
	if err != nil {
		return nil, nil, errors.New("app store not found")
	}

	return appStore, app, nil
}

func (db *DB) GetAppByName(name string) (*apps.App, error) {
	var m *apps.App
	if err := db.DB.Where("app_store_name = ? ", name).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (db *DB) AddApp(body *apps.App) (resp *apps.App, existingInstall bool, err error) {
	store, err := db.GetAppStoreByName(body.AppStoreName)
	if err != nil {
		return nil, false, errors.New("no app store is installed for this app")
	}
	name, _ := db.GetAppByName(store.Name)
	if name != nil {
		return name, true, nil
	}
	if body.InstalledVersion == "" {
		return nil, false, errors.New("installed version can not be null")
	}
	body.UUID = uuid.ShortUUID("app")
	body.AppStoreUUID = store.UUID
	body.AppStoreName = store.Name
	if err := db.DB.Create(&body).Error; err != nil {
		return nil, false, err
	} else {
		return body, false, nil
	}
}

func (db *DB) UpdateApp(uuid string, app *apps.App) (*apps.App, error) {
	var m *apps.App
	query := db.DB.Where("uuid = ?", uuid).Find(&m).Updates(app)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return app, query.Error
	}
}

func (db *DB) DeleteApp(uuid string) (*DeleteMessage, error) {
	var m *apps.App
	query := db.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (db *DB) DropApps() (*DeleteMessage, error) {
	var m *apps.App
	query := db.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
