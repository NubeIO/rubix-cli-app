package dbase

import (
	"errors"
	"github.com/NubeIO/rubix-edge/pkg/model"
)

const deviceInfo = "device info"

func (db *DB) GetDeviceInfo(uuid string) (*model.DeviceInfo, error) {
	var m *model.DeviceInfo
	if err := db.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		return nil, handleNotFound(deviceInfo)
	}
	return m, nil
}

func (db *DB) GetAllDeviceInfos() ([]*model.DeviceInfo, error) {
	var m []*model.DeviceInfo
	if err := db.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (db *DB) AddDeviceInfo(body *model.DeviceInfo) (resp *model.DeviceInfo, err error) {
	infos, err := db.GetAllDeviceInfos()
	if err != nil {
		return nil, err
	}
	if infos != nil {
		return nil, errors.New("device info can only be added once")
	}
	if err := db.DB.Create(&body).Error; err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func (db *DB) UpdateDeviceInfo(uuid string, app *model.DeviceInfo) (*model.DeviceInfo, error) {
	var m *model.DeviceInfo
	query := db.DB.Where("uuid = ?", uuid).Find(&m).Updates(app)
	if query.Error != nil {
		return nil, handleNotFound(deviceInfo)
	} else {
		return app, query.Error
	}
}

func (db *DB) DeleteDeviceInfo(uuid string) (*DeleteMessage, error) {
	var m *model.DeviceInfo
	query := db.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}
