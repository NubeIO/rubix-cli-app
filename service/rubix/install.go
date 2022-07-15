package rubix

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type App struct {
	Name        string `json:"app"`          // rubix-wires
	Version     string `json:"version"`      // v1.1.1
	DirName     string `json:"dir_name"`     // wires-builds
	ServiceName string `json:"service_name"` // nubeio-rubix-wires
}

//makeTmpDir  => /data/tmp
func makeTmpDir() error {
	if err := checkDir(DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", DataDir))
	}
	return makeDirectoryIfNotExists(TmpDir, os.FileMode(FilePerm))
}

//makeAppInstallDir  => /data/rubix-service/apps/install/flow-framework
func makeAppInstallDir(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkDir(DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", DataDir))
	}
	if err := checkDir(AppsInstallDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", AppsInstallDir))
	}
	return makeDirectoryIfNotExists(fmt.Sprintf("%s/%s", AppsInstallDir, appName), os.FileMode(FilePerm))
}

//makeAppVersionDir  => /data/rubix-service/apps/install/flow-framework/v0.0.1
func makeAppVersionDir(appName, version string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkVersion(version); err != nil {
		return err
	}
	if err := checkDir(DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", DataDir))
	}
	if err := checkDir(AppsInstallDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", AppsInstallDir))
	}
	appDir := fmt.Sprintf("%s/%s", AppsInstallDir, appName)
	if err := checkDir(appDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", appDir))
	}
	return makeDirectoryIfNotExists(fmt.Sprintf("%s/%s", appDir, version), os.FileMode(FilePerm))
}

//makeAppDir  => /data/flow-framework
func makeAppDir(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkDir(DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", DataDir))
	}
	return makeDirectoryIfNotExists(fmt.Sprintf("%s/%s", DataDir, appName), os.FileMode(FilePerm))
}

// makeDirectoryIfNotExists if not exist make dir
func makeDirectoryIfNotExists(path string, perm os.FileMode) error {
	if perm == 0 {
		perm = 0755
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|perm)
	}
	return nil
}

func emptyPath(path string) error {
	if path == "" {
		return errors.New("path can not be empty")
	}
	return nil
}

func checkDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
}

func checkVersion(version string) error {
	if version[0:1] != "v" { // make sure have a v at the start v0.1.1
		return errors.New(fmt.Sprintf("there version number should start with a v eg:v1.2.3 %s", version))
	}
	p := strings.Split(version, ".")
	if len(p) >= 2 && len(p) < 4 {
	} else {
		return errors.New(fmt.Sprintf("there version number should match v1.2.3 %s", version))
	}
	return nil
}
