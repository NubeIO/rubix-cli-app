package dbase

import (
	"gthub.com/NubeIO/rubix-cli-app/pkg/config"
	"gthub.com/NubeIO/rubix-cli-app/pkg/logger"
	"gthub.com/NubeIO/rubix-cli-app/pkg/model"
)

func (d *DB) GetUser(uuid string) (*model.User, error) {
	m := new(model.User)
	if err := d.DB.Where("uuid = ? ", uuid).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		return nil, err
	}
	return m, nil
}

func (d *DB) GetUsers() ([]*model.User, error) {
	var m []*model.User
	if err := d.DB.Find(&m).Error; err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (d *DB) CreateUser(User *model.User) (*model.User, error) {
	User.UUID, _ = config.MakeUUID()
	if err := d.DB.Create(&User).Error; err != nil {
		return nil, err
	} else {
		return User, nil
	}
}

func (d *DB) UpdateUser(uuid string, User *model.User) (*model.User, error) {
	m := new(model.User)
	query := d.DB.Where("uuid = ?", uuid).Find(&m).Updates(User)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return User, query.Error
	}
}

func (d *DB) DeleteUser(uuid string) (ok bool, err error) {
	m := new(model.User)
	query := d.DB.Where("uuid = ? ", uuid).Delete(&m)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}

// DropUsers delete all.
func (d *DB) DropUsers() (bool, error) {
	var m *model.User
	query := d.DB.Where("1 = 1")
	query.Delete(&m)
	if query.Error != nil {
		return false, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return false, nil
	}
	return true, nil
}
