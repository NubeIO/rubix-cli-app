package rubix

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func (inst *Rubix) MakeAllDirs() error {
	err := inst.MakeDataDir()
	if err != nil {
		return err
	}
	err = inst.MakeTmpDir()
	if err != nil {
		return err
	}
	err = inst.MakeInstallDir()
	if err != nil {
		return err
	}
	return nil
}

//MakeTmpDir  => /data/tmp
func (inst *Rubix) MakeTmpDir() error {
	if err := checkDir(inst.DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", inst.DataDir))
	}
	return makeDirectoryIfNotExists(TmpDir, os.FileMode(FilePerm))
}

//MakeAppInstallDir  => /data/rubix-service/apps/install/flow-framework
func (inst *Rubix) MakeAppInstallDir(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkDir(inst.DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", inst.DataDir))
	}
	if err := checkDir(AppsInstallDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", AppsInstallDir))
	}
	return makeDirectoryIfNotExists(fmt.Sprintf("%s/%s", AppsInstallDir, appName), os.FileMode(FilePerm))
}

//MakeAppVersionDir  => /data/rubix-service/apps/install/flow-framework/v0.0.1
func (inst *Rubix) MakeAppVersionDir(appName, version string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkVersion(version); err != nil {
		return err
	}
	if err := checkDir(inst.DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", inst.DataDir))
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

//MakeAppDir  => /data/flow-framework
func (inst *Rubix) MakeAppDir(appName string) error {
	if err := emptyPath(appName); err != nil {
		return err
	}
	if err := checkDir(inst.DataDir); err != nil {
		return errors.New(fmt.Sprintf("dir not exists %s", inst.DataDir))
	}
	return makeDirectoryIfNotExists(fmt.Sprintf("%s/%s", inst.DataDir, appName), os.FileMode(FilePerm))
}

//MakeInstallDir  => /data/rubix-service/install
func (inst *Rubix) MakeInstallDir() error {
	if AppsInstallDir == "" {
		return errors.New("path can not be empty")
	}
	return os.MkdirAll(AppsInstallDir, os.FileMode(FilePerm))
}

//MakeDataDir  => /data
func (inst *Rubix) MakeDataDir() error {
	if inst.DataDir == "" {
		return errors.New("path can not be empty")
	}
	return makeDirectoryIfNotExists(inst.DataDir, os.FileMode(FilePerm))
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
