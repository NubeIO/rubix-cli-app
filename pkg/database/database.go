package database

import (
	"errors"
	"fmt"
	"github.com/NubeIO/edge/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"

	"github.com/NubeIO/edge/pkg/model"
	"github.com/NubeIO/edge/service/apps"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// Setup opens a database and saves the reference to `Database` struct.
func Setup() error {
	var db = DB
	dbName := viper.GetString("database.name")
	driver := viper.GetString("database.driver")

	if driver == "" {
		driver = "sqlite"
	}
	connection := fmt.Sprintf("%s?_foreign_keys=on", path.Join(config.Config.GetAbsDataDir(), dbName))
	switch driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(connection), &gorm.Config{})
	default:
		return errors.New("unsupported database driver")
	}

	if err != nil {
		return err
	}

	// Auto migrate project models
	err = db.AutoMigrate(
		&model.User{},
		&apps.Store{},
		&apps.App{},
		&model.DeviceInfo{},
	)

	if err != nil {
		return err
	}
	DB = db
	return nil
}
