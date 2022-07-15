package rubix

import (
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type AppResponse struct {
	Name       string                 `json:"app"`
	Version    string                 `json:"version"`
	IsAService bool                   `json:"is_service"`
	AppStatus  *systemctl.SystemState `json:"app_status,omitempty"`
	Error      string                 `json:"error,omitempty"`
}

var DefaultTimeout = 30

var systemOpts = systemctl.Options{
	UserMode: false,
	Timeout:  DefaultTimeout,
}

func checkAppsDir(rootDir string) ([]AppResponse, error) {
	if rootDir == "" {
		rootDir = AppsInstallDir
	}
	var files []AppResponse
	app := AppResponse{}
	err := filepath.WalkDir(rootDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && strings.Count(p, string(os.PathSeparator)) == 6 {
			parts := strings.Split(p, "/")
			if len(parts) >= 5 { // app name
				app.Name = parts[5]
			}
			if len(parts) >= 6 { // version
				app.Version = parts[6]
			}
			app.AppStatus = nil
			files = append(files, app)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func checkAppsService(apps []AppResponse) ([]AppResponse, error) {
	var out []AppResponse
	for _, app := range apps {
		installed, err := systemctl.IsInstalled(app.Name, systemOpts)
		if installed {
			app.IsAService = true
			state, err := systemctl.State(app.Name, systemOpts)
			if err != nil {
				app.Error = err.Error()
			}
			app.AppStatus = &state
			out = append(out, app)
		} else {
			app.IsAService = false
			app.Error = err.Error()
			out = append(out, app)
		}

	}
	return out, nil
}
