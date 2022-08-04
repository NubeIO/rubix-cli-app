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
func (inst *EdgeApps) RestoreBackup(back *installer.RestoreBackup) (*installer.RestoreResponse, error) {
	if back == nil {
		return nil, errors.New("RestoreBackup interface can not be empty")
	}
	if back.Destination == "" {
		return nil, errors.New("destination can not be empty")
	}
	resp := &installer.RestoreResponse{}
	restoreResp, err := inst.App.RestoreBackup(back)
	if err != nil {
		return nil, err
	}
	resp.TakeBackupPath = restoreResp.TakeBackupPath
	if back.RebootDevice {
		resp.Message = "device will reboot in 10 seconds"
		go reboot()
	}
	return resp, nil
}

// RestoreAppBackup restore a backup an app
func (inst *EdgeApps) RestoreAppBackup(back *installer.RestoreBackup) (*installer.RestoreResponse, error) {
	if back == nil {
		return nil, errors.New("RestoreBackup interface can not be empty")
	}
	if back.AppName == "" {
		return nil, errors.New("app name can not be empty")
	}
	if back.Destination == "" {
		return nil, errors.New("destination can not be empty")
	}
	resp := &installer.RestoreResponse{}
	restoreResp, err := inst.App.RestoreAppBackup(back)
	if err != nil {
		return nil, err
	}
	resp.TakeBackupPath = restoreResp.TakeBackupPath
	return resp, nil
}

func (inst *EdgeApps) FullBackUp(deiceName ...string) (*BackupResp, error) {
	path, err := inst.App.FullBackUp(deiceName...)
	return &BackupResp{BackupPath: path}, err
}

func (inst *EdgeApps) BackupApp(appName string, deiceName ...string) (*BackupResp, error) {
	path, err := inst.App.BackupApp(appName, deiceName...)
	return &BackupResp{BackupPath: path}, err
}

func (inst *EdgeApps) ListFullBackups() ([]string, error) {
	return inst.App.ListFullBackups()
}

func (inst *EdgeApps) ListAppBackupsDirs() ([]string, error) {
	return inst.App.ListAppBackupsDirs()
}

func (inst *EdgeApps) ListBackupsByApp(appName string) ([]string, error) {
	return inst.App.ListBackupsByApp(appName)
}

func (inst *EdgeApps) DeleteAllFullBackups() (*installer.MessageResponse, error) {
	return inst.App.DeleteAllFullBackups()
}

func (inst *EdgeApps) DeleteAllAppBackups() (*installer.MessageResponse, error) {
	return inst.App.DeleteAllAppBackups()
}

func (inst *EdgeApps) DeleteAppAllBackUpByName(appName string) (*installer.MessageResponse, error) {
	return inst.App.DeleteAppAllBackUpByName(appName)
}

func (inst *EdgeApps) DeleteAppOneBackUpByName(appName, backupFolder string) (*installer.MessageResponse, error) {
	return inst.App.DeleteAppOneBackUpByName(appName, backupFolder)
}

func (inst *EdgeApps) WipeBackups() (*installer.MessageResponse, error) {
	return inst.App.WipeBackups()
}
