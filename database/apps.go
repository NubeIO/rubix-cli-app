package dbase

import (
	"errors"
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/pkg/helpers/uuid"
	"gthub.com/NubeIO/rubix-cli-app/pkg/logger"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

func (db *DB) GetApps() ([]*apps.InstalledApp, error) {
	var m []*apps.InstalledApp
	if err := db.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (db *DB) GetApp(uuid string) (*apps.InstalledApp, error) {
	var m *apps.InstalledApp
	if err := db.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (db *DB) GetAppByName(name string) (*apps.InstalledApp, error) {
	var m *apps.InstalledApp
	if err := db.DB.Where("app_store_name = ? ", name).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (db *DB) AddApp(body *apps.InstalledApp) (resp *apps.InstalledApp, existingInstall bool, err error) {
	store, err := db.GetAppImageByName(body.AppStoreName)
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
	body.UUID = fmt.Sprintf("app_%s", uuid.SmallUUID())
	body.AppStoreUUID = store.UUID
	body.AppStoreName = store.Name
	if err := db.DB.Create(&body).Error; err != nil {
		return nil, false, err
	} else {
		return body, false, nil
	}
}

func (db *DB) UpdateApp(uuid string, app *apps.InstalledApp) (*apps.InstalledApp, error) {
	var m *apps.InstalledApp
	query := db.DB.Where("uuid = ?", uuid).Find(&m).Updates(app)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return app, query.Error
	}
}

func (db *DB) DeleteApp(uuid string) (*DeleteMessage, error) {
	var m *apps.InstalledApp
	query := db.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (db *DB) DropApps() (*DeleteMessage, error) {
	var m *apps.InstalledApp
	query := db.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
