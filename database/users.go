package dbase

import (
	"github.com/NubeIO/edge/pkg/logger"
	"github.com/NubeIO/edge/pkg/model"
	"github.com/NubeIO/lib-uuid/uuid"
)

func (db *DB) GetUser(uuid string) (*model.User, error) {
	m := new(model.User)
	if err := db.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (db *DB) GetUsers() ([]*model.User, error) {
	var m []*model.User
	if err := db.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (db *DB) CreateUser(user *model.User) (*model.User, error) {
	user.UUID = uuid.ShortUUID("usr")
	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (db *DB) UpdateUser(uuid string, User *model.User) (*model.User, error) {
	m := new(model.User)
	query := db.DB.Where("uuid = ?", uuid).Find(&m).Updates(User)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return User, query.Error
	}
}

func (db *DB) DeleteUser(uuid string) (*DeleteMessage, error) {
	m := new(model.User)
	query := db.DB.Where("uuid = ? ", uuid).Delete(&m)
	return deleteResponse(query)
}

// DropUsers delete all.
func (db *DB) DropUsers() (*DeleteMessage, error) {
	var m *model.User
	query := db.DB.Where("1 = 1")
	query.Delete(&m)
	return deleteResponse(query)
}
