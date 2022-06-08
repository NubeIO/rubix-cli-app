package installer

import (
	"github.com/NubeIO/lib-store/store"
	"github.com/NubeIO/rubix-cli-app/service/apps"
	log "github.com/sirupsen/logrus"
)

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
