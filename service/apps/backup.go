package apps

import (
	"errors"
	"github.com/NubeIO/lib-rubix-installer/installer"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"time"
)

type BackupResp struct {
	BackupPath string
}

func reboot() {
	time.Sleep(10 * time.Second)
	log.Errorln("TODO implement reboot")
}

type RestoreBackup struct {
	AppName     string                `json:"app_name"`
	Destination string                `json:"destination"`
	DeviceName  string                `json:"device_name"`
	TakeBackup  bool                  `json:"take_backup"`
	File        *multipart.FileHeader `json:"file"`
}

// RestoreBackup restore a backup data dir /data
func (inst *EdgeApp) RestoreBackup(back *installer.RestoreBackup) (*installer.RestoreResponse, error) {
	if back == nil {
		return nil, errors.New("RestoreBackup interface can not be empty")
	}
	restoreResp, err := inst.App.RestoreBackup(back)
	if err != nil {
		return nil, err
	}
	if back.RebootDevice {
		restoreResp.Message = "device will reboot in 10 seconds"
		go reboot()
	}
	return restoreResp, nil
}

// RestoreAppBackup restore a backup an app
func (inst *EdgeApp) RestoreAppBackup(back *installer.RestoreBackup) (*installer.RestoreResponse, error) {
	if back == nil {
		return nil, errors.New("RestoreBackup interface can not be empty")
	}
	if back.AppName == "" {
		return nil, errors.New("app name can not be empty")
	}
	restoreResp, err := inst.App.RestoreAppBackup(back)
	if err != nil {
		return nil, err
	}
	return restoreResp, nil
}

func (inst *EdgeApp) FullBackUp(deiceName *string) (*BackupResp, error) {
	path, err := inst.App.FullBackUp(deiceName)
	return &BackupResp{BackupPath: path}, err
}

func (inst *EdgeApp) BackupApp(appName string, deiceName *string) (*BackupResp, error) {
	path, err := inst.App.BackupApp(appName, deiceName)
	return &BackupResp{BackupPath: path}, err
}

func (inst *EdgeApp) ListFullBackups() ([]string, error) {
	return inst.App.ListFullBackups()
}

func (inst *EdgeApp) ListAppsBackups() ([]string, error) {
	return inst.App.ListAppsBackups()
}

func (inst *EdgeApp) ListAppBackups(appName string) ([]string, error) {
	return inst.App.ListAppBackups(appName)
}

func (inst *EdgeApp) DeleteAllFullBackups() (*installer.MessageResponse, error) {
	return inst.App.DeleteAllFullBackups()
}

func (inst *EdgeApp) DeleteAllAppsBackups() (*installer.MessageResponse, error) {
	return inst.App.DeleteAllAppsBackups()
}

func (inst *EdgeApp) DeleteAllAppBackups(appName string) (*installer.MessageResponse, error) {
	return inst.App.DeleteAllAppBackups(appName)
}

func (inst *EdgeApp) DeleteOneAppBackup(appName, zipFile string) (*installer.MessageResponse, error) {
	return inst.App.DeleteOneAppBackup(appName, zipFile)
}

func (inst *EdgeApp) WipeBackups() (*installer.MessageResponse, error) {
	return inst.App.WipeBackups()
}
