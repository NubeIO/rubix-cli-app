package dbase

import (
	"fmt"
	"github.com/NubeIO/edge/service/apps"
	"github.com/NubeIO/lib-store/store"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

type DeleteMessage struct {
	Message string `json:"message"`
}

func deleteResponse(query *gorm.DB) (*DeleteMessage, error) {
	msg := &DeleteMessage{
		Message: fmt.Sprintf("no record found, deleted count:%d", 0),
	}
	if query.Error != nil {
		return msg, query.Error
	}
	r := query.RowsAffected
	if r == 0 {
		return msg, query.Error
	}
	msg.Message = fmt.Sprintf("deleted count:%d", query.RowsAffected)
	return msg, nil
}

func initAppService(serviceName string) (*apps.Apps, error) {
	inst := &apps.Apps{
		App: &apps.Store{
			ServiceName: serviceName,
		},
	}
	app, err := apps.New(inst)
	return app, err
}

var progress = initStore()

func initStore() *store.Handler {
	return store.Init()
}

func SetProgress(key string, data interface{}) {
	progress.SetNoExpire(key, data)
}

func initApp(initApp *apps.Apps, appStore *apps.Store) (*apps.Apps, error) {
	var inst = &apps.Apps{
		Token:   initApp.Token,
		Perm:    apps.Permission,
		Version: initApp.Version,
		App:     appStore,
	}
	app, err := apps.New(inst)
	if err != nil {
		log.Errorln("new app: failed to init a new app", err)
		return app, err
	}
	return app, err
}
