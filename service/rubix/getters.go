package rubix

import "fmt"

// GetAppInstallPath get the full app install path and version => /data/rubix-service/apps/install/wires-builds
func (inst *App) getAppPath(appName string) string {
	return fmt.Sprintf("%s/%s", DataDir, appName)
}

// GetAppInstallPath get the full app install path and version => /data/rubix-service/apps/install/wires-builds
func (inst *App) getAppInstallPath(appBuildName string) string {
	return fmt.Sprintf("%s/%s", AppsInstallDir, appBuildName)
}

// GetAppInstallPathAndVersion get the full app install path and version => /data/rubix-service/apps/install/wires-builds/v0.0.1
func (inst *App) getAppInstallPathAndVersion(appBuildName, version string) string {
	return fmt.Sprintf("%s/%s/%s", AppsInstallDir, appBuildName, version)
}
