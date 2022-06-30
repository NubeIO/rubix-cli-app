package database

import (
	"errors"
	"fmt"
	"github.com/NubeIO/edge/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"path"
	"time"

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
	logLevel := logger.Silent
	dbLogLevel := viper.GetString("database.loglevel")
	if dbLogLevel == "ERROR" {
		logLevel = logger.Error
	} else if dbLogLevel == "WARN" {
		logLevel = logger.Warn
	} else if dbLogLevel == "INFO" {
		logLevel = logger.Info
	}
	writer := io.MultiWriter(os.Stdout, getWriter())
	colorful := true
	if config.Config.Prod() {
		colorful = false
	}
	newDBLogger := logger.New(
		log.New(writer, "", log.LstdFlags), // io writers
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logLevel,         // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  colorful,         // Disable color
		},
	)

	var db = DB
	dbName := viper.GetString("database.name")
	driver := viper.GetString("database.driver")

	if driver == "" {
		driver = "sqlite"
	}
	switch driver {
	case "sqlite":
		connection := fmt.Sprintf("%s?_foreign_keys=on", path.Join(config.Config.GetAbsDataDir(), dbName))
		db, err = gorm.Open(sqlite.Open(connection), &gorm.Config{
			Logger: newDBLogger,
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
		})
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

func getWriter() io.Writer {
	fileLocation := fmt.Sprintf("%s/edge.db.log", config.Config.GetAbsDataDir())
	file, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return os.Stdout
	} else {
		return file
	}
}
