package apps

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
)

type BackupResp struct {
	BackupPath string
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
