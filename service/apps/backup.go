package apps

type BackupResp struct {
	BackupPath string
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

func (inst *EdgeApps) FullBackUp(deiceName ...string) (*BackupResp, error) {
	path, err := inst.App.FullBackUp(deiceName...)
	return &BackupResp{BackupPath: path}, err
}

func (inst *EdgeApps) BackupApp(appName string, deiceName ...string) (*BackupResp, error) {
	path, err := inst.App.BackupApp(appName, deiceName...)
	return &BackupResp{BackupPath: path}, err
}
