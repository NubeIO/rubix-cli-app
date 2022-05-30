package dbase

import (
	"errors"
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/pkg/helpers/uuid"
	"gthub.com/NubeIO/rubix-cli-app/pkg/logger"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

func (d *DB) GetApps() ([]*apps.InstalledApp, error) {
	var m []*apps.InstalledApp
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) GetApp(uuid string) (*apps.InstalledApp, error) {
	var m *apps.InstalledApp
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetAppByName(name string) (*apps.InstalledApp, error) {
	var m *apps.InstalledApp
	if err := d.DB.Where("name = ? ", name).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) AddApp(body *apps.InstalledApp) (*apps.InstalledApp, error) {
	store, err := d.GetAppImageByName(body.AppStoreName)
	if err != nil {
		return nil, errors.New("no app store is installed for this app")
	}
	body.UUID = fmt.Sprintf("app_%s", uuid.SmallUUID())
	body.AppStoreUUID = store.UUID
	body.AppStoreName = store.Name
	if err := d.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (d *DB) UpdateApp(uuid string, app *apps.InstalledApp) (*apps.InstalledApp, error) {
	var m *apps.InstalledApp
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(app)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return app, query.Error
	}
}

func (d *DB) DeleteApp(uuid string) (*DeleteMessage, error) {
	var m *apps.InstalledApp
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

func (d *DB) DropApps() (*DeleteMessage, error) {
	var m *apps.InstalledApp
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
