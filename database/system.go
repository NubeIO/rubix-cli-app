package dbase

import (
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

type SystemCtl struct {
	Action string    `json:" action"` // start. stop
	App    *apps.App `json:"app"`
}

func (db *DB) SystemCtlAction(body *SystemCtl) (*apps.SystemResponseChecks, error) {
	appStore, _, err := db.GetAppAndStore(body.App)
	if err != nil {
		return nil, err
	}
	inst := &apps.Apps{
		App: &apps.Store{
			ServiceName: appStore.ServiceName,
		},
	}
	app, err := apps.New(inst)
	status, err := app.SystemCtlStatus("isRunning", apps.DefaultTimeout)
	if err != nil {
		return nil, err
	}
	return status, err
}
