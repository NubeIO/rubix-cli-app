package installer

import (
	"errors"
	"fmt"
	dbase "github.com/NubeIO/edge/database"
	"github.com/NubeIO/edge/service/apps"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type App struct {
	AppName            string `json:"app_name"`              // flow-framework, this is the name from the app-store
	Version            string `json:"version"`               // app version latest or v.0.0.1
	Token              string `json:"token"`                 // github token
	ManualInstall      bool   `json:"manual_install"`        // will not download from GitHub, and will use the app-store download path
	ManualAssetZipName string `json:"manual_asset_zip_name"` // flow-framework-0.5.5-1575cf89.amd64.zip
	ManualAssetTag     string `json:"-"`                     // this is the release tag as in v0.0.1
	Cleanup            bool   `json:"cleanup"`
}

type InstallResponse struct {
	Message    string     `json:"message"`
	Error      string     `json:"error"`
	InstallLog InstallLog `json:"log"`
}

type InstallLog struct {
	GetAppFromStore string `json:"get_app_from_store"`
	AppInstall      string `json:"-"`
	MakeDownload    string `json:"make_download"`
	ManualInstall   string `json:"manual_install"`
	GitDownload     string `json:"git_download"`
	MakeInstallDir  string `json:"make_install_dir"`
	UnpackBuild     string `json:"unpack_build"`
	GenerateService string `json:"generate_service"`
	InstallService  string `json:"install_service"`
	CleanUp         string `json:"clean_up"`
}

type Installer struct {
	DB *dbase.DB
}

func New(install *Installer) *Installer {
	return install
}

// ok messages
const (
	ok = "ok"
)

func (inst *Installer) GetInstallProgress(key string) (*InstallResponse, error) {
	key = fmt.Sprintf("install-%s", key)
	data, ok := progress.Get(key)
	if ok {
		parse := data.(*InstallResponse)
		return parse, nil
	}
	resp := &InstallResponse{
		Message: "not found able to find the app",
	}
	return resp, nil

}

func (inst *Installer) InstallApp(body *App) (*InstallResponse, error) {
	resp := &InstallResponse{}
	app, err := inst.installApp(body)
	if err != nil {
		resp.InstallLog = app.InstallLog
		resp.Message = fmt.Sprintf("install fail! %s", body.AppName)
		resp.Error = err.Error()
		return resp, err
	}
	resp.InstallLog = app.InstallLog
	resp.Error = "no errors"
	resp.Message = fmt.Sprintf("install ok! %s", app.InstallLog.AppInstall)
	return resp, err
}

func (inst *Installer) installApp(body *App) (*InstallResponse, error) {
	resp := &InstallResponse{}
	progressKey := fmt.Sprintf("install-%s", body.AppName)
	SetProgress(progressKey, resp)
	appStore, err := inst.DB.GetAppStoreByName(body.AppName)
	if err != nil {
		resp.InstallLog.GetAppFromStore = err.Error()
		return resp, err
	}

	if body.Version == "" {
		resp.InstallLog.MakeDownload = "app version can not be empty"
		SetProgress(progressKey, resp)
		return resp, errors.New("app version can not be empty")
	}

	resp.InstallLog.GetAppFromStore = ok
	installedApp := &apps.App{
		AppStoreName:     appStore.Name,
		AppStoreUUID:     appStore.UUID,
		InstalledVersion: body.Version,
	}

	var newApps = &apps.Apps{
		Token:   body.Token,
		Perm:    apps.Permission,
		Version: body.Version,
		App:     appStore,
	}
	newApp, err := apps.New(newApps)
	SetProgress(progressKey, resp)
	if err != nil {
		log.Errorln("new app: failed to init a new app", err)
		return resp, err
	}
	if err = newApps.MakeDownloadDir(); err != nil {
		resp.InstallLog.MakeDownload = "issue on trying to make the path to download the zip folder"
		SetProgress(progressKey, resp)
		return resp, err
	}
	assetTag := ""
	resp.InstallLog.MakeDownload = ok
	if body.ManualInstall { // manual installation
		if body.ManualAssetTag == "" {
			match, count, version, archMatch, arch := matchRepoName(body.ManualAssetZipName, newApp.App.Repo)
			if !match {
				resp.InstallLog.ManualInstall = fmt.Sprintf("failed on match uploaded app, match-count:%d zip file name:%s repo-name:%s arch%s", count, body.ManualAssetZipName, newApp.App.Repo, arch)
				return resp, errors.New(resp.InstallLog.ManualInstall)
			}
			if !archMatch {
				resp.InstallLog.ManualInstall = fmt.Sprintf("failed on match arch, zip file name:%s repo-name:%s arch%s", body.ManualAssetZipName, newApp.App.Repo, arch)
				return resp, errors.New(resp.InstallLog.ManualInstall)
			}
			assetTag = version
		} else {
			assetTag = body.ManualAssetTag
		}
		if body.ManualAssetZipName == "" {
			resp.InstallLog.ManualInstall = "zip folder name can not be empty, try flow-framework-0.5.5-1575cf89.amd64.zip"
			return resp, errors.New("zip folder name can not be empty, try flow-framework-0.5.5-1575cf89.amd64.zip")
		}
		if err = newApps.BuildExists(body.ManualAssetZipName); err != nil { //check if it is there
			resp.InstallLog.ManualInstall = err.Error()
			return resp, err
		}
		resp.InstallLog.ManualInstall = "found existing build zip folder"

		if assetTag == "" {
			resp.InstallLog.ManualInstall = "asset tag can not be empty, try v0.5.5"
			return resp, errors.New("asset tag can not be empty, try v0.5.5")
		}

		resp.InstallLog.GitDownload = "no download as was a manual installation"
	} else { // or download from GitHub
		download, err := newApp.GitDownload(newApps.App.DownloadPath)
		SetProgress(progressKey, resp)
		if err != nil {
			log.Errorf("git: download error %s \n", err.Error())
			resp.InstallLog.GitDownload = err.Error()
			SetProgress(progressKey, resp)
			return resp, err
		}
		assetTag = download.RepositoryRelease.GetTagName()
		resp.InstallLog.GitDownload = fmt.Sprintf("installed version: %s", assetTag)
	}

	SetProgress(progressKey, resp)
	if err = newApps.MakeInstallDir(); err != nil { // make the installation dir /data/rubix-apps/installed/flow-framework
		resp.InstallLog.MakeInstallDir = err.Error()
		SetProgress(progressKey, resp)
		return resp, err
	}
	resp.InstallLog.MakeInstallDir = ok
	SetProgress(progressKey, resp)
	if err = newApps.UnpackBuild(body.ManualAssetZipName); err != nil { // unzip from: /home/user/downloads  to: /data/rubix-apps/installed/flow-framework
		resp.InstallLog.UnpackBuild = err.Error()
		SetProgress(progressKey, resp)
		return resp, err
	}
	resp.InstallLog.UnpackBuild = ok
	tmpFileDir := newApp.App.DownloadPath
	SetProgress(progressKey, resp)
	if _, err = newApp.GenerateServiceFile(newApp, tmpFileDir); err != nil { // make systemd file
		log.Errorf("make service file build: failed error:%s \n", err.Error())
		resp.InstallLog.GenerateService = err.Error()
		SetProgress(progressKey, resp)
		return resp, err
	}
	resp.InstallLog.GenerateService = ok
	tmpServiceFile := fmt.Sprintf("%s/%s.service", tmpFileDir, newApp.App.ServiceName)
	SetProgress(progressKey, resp)
	if _, err = newApp.InstallService(newApp.App.ServiceName, tmpServiceFile); err != nil { // install the systemd service
		resp.InstallLog.InstallService = err.Error()
		SetProgress(progressKey, resp)
		return resp, err
	}
	resp.InstallLog.InstallService = ok
	SetProgress(progressKey, resp)
	if body.Cleanup {
		if err = newApps.CleanUp(body.ManualAssetZipName); err != nil { // delete tmp install dirs
			resp.InstallLog.CleanUp = err.Error()
			SetProgress(progressKey, resp)
			return resp, err
		}
		resp.InstallLog.CleanUp = ok
	} else {
		resp.InstallLog.CleanUp = "clean up was disabled"
	}

	installedApp.InstalledVersion = assetTag
	SetProgress(progressKey, resp)
	app, existingApp, err := inst.DB.AddApp(installedApp)
	if err != nil {
		resp.InstallLog.AppInstall = err.Error()
		SetProgress(progressKey, resp)
		return resp, err
	}
	if existingApp { // if it was existing app update the version
		existingVersion := app.InstalledVersion
		app.InstalledVersion = assetTag
		_, err := inst.DB.UpdateApp(app.UUID, app)
		SetProgress(progressKey, resp)
		if err != nil {
			resp.InstallLog.AppInstall = fmt.Sprintf("an existing app was installed error:%s", err.Error())
			SetProgress(progressKey, resp)
			return resp, err
		}
		resp.InstallLog.AppInstall = fmt.Sprintf("an existing app was installed upgraded from: %s to: %s", existingVersion, assetTag)
	} else {
		resp.InstallLog.AppInstall = "installed a new app"
		log.Infof(fmt.Sprintf("an existing app was installed upgraded from:%s to:%s", app.InstalledVersion, assetTag))
	}

	SetProgress(progressKey, resp)
	return resp, err
}

//matchRepoName get the tag name from the zip eg, wires-builds-0.5.5-1575cf89.amd64.zip => wires-builds
// 	returns
// 	- true if is a match if it is a match
// 	- match count
// 	- string version name
// 	- arch match
// 	- arch type
func matchRepoName(zipName, repoName string) (bool, int, string, bool, string) {
	parts := strings.Split(zipName, "-")
	repoNameParts := strings.Split(repoName, "-")
	count := 0
	version := ""
	arch := ""
	archMatch := false
	repoMatch := false
	for i, part := range parts {
		p := strings.Split(part, ".")
		// if len is 3 eg, 0.0.1
		isNum := 0
		if len(p) == 3 || len(p) == 4 {
			// check if they are numbers
			for _, s := range p {
				if _, err := strconv.Atoi(s); err == nil {
					isNum++
				}
			}
			if isNum == 3 {
				count = i
				version = part
				version = strings.Trim(version, ".zip")
			}
		}
	}
	match := 0
	for i := 0; i < count; i++ {
		if isMatch(parts, repoNameParts[i]) {
			match++
		}
	}
	if match == count {
		repoMatch = true
	}
	if repoName != "wires-builds" { //wires can run on any os
		arch, _ = getArch()
		if contains(parts, arch) {
			if repoName == "wires-builds" {

			}
			archMatch = true
		}
	} else {
		archMatch = true
	}
	return repoMatch, count, version, archMatch, arch
}

func getArch() (string, error) {
	arch, err := cmd.DetectArch()
	if err != nil {
		return "", err
	}
	return arch.ArchModel, err
}

func isMatch(s []string, term string) bool {
	count := 0
	for _, item := range s {
		if item == term {
			count++
			return true
		}
	}
	return false
}

func contains(s []string, term string) bool {
	count := 0
	for _, item := range s {
		if strings.Contains(item, term) {
			count++
			return true
		}
	}
	return false
}
