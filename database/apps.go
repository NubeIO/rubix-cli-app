package dbase

import (
	"fmt"
	"gthub.com/NubeIO/rubix-cli-app/pkg/helpers/uuid"
	"gthub.com/NubeIO/rubix-cli-app/pkg/logger"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
	"gthub.com/NubeIO/rubix-cli-app/service/product"
)

func (d *DB) GetApps() ([]*apps.Store, error) {
	var m []*apps.Store
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) GetApp(uuid string) (*apps.Store, error) {
	var m *apps.Store
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetApp error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) CreateApp(app *apps.Store) (*apps.Store, error) {
	app.UUID = fmt.Sprintf("app_%s", uuid.SmallUUID())
	pro, err := product.Get()
	app.ProductType = pro.Type
	if err != nil {
		return nil, err
	}
	if err := d.DB.Create(&app).Error; err != nil {
		return nil, err
	} else {
		return app, nil
	}
}

func (d *DB) UpdateApp(uuid string, app *apps.Store) (*apps.Store, error) {
	var m *apps.Store
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(app)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return app, query.Error
	}
}

func (d *DB) DeleteApp(uuid string) (*DeleteMessage, error) {
	var m *apps.Store
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

type DeleteMessage struct {
	Message string `json:"message"`
}

// DropApps delete all.
func (d *DB) DropApps() (*DeleteMessage, error) {
	var m *apps.Store
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
