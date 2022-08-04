package apps

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func userHomeDir() string {
	homeDir, _ := fileutils.Dir()
	return homeDir
}

func empty(name string) error {
	if name == "" {
		return errors.New("can not be empty")
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
	path = filePath(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
}

// filePath make the file path work for unix or windows
func filePath(path string, debug ...bool) string {
	updated := filepath.FromSlash(path)
	if len(debug) > 0 {
		if debug[0] {
			log.Infof("existing-path: %s", path)
			log.Infof("updated-path: %s", updated)
		}
	}
	return filepath.FromSlash(updated)
}

func checkVersion(version string) error {
	if version[0:1] != "v" { // make sure have a v at the start v0.1.1
		return errors.New(fmt.Sprintf("incorrect provided:%s version number try: v1.2.3", version))
	}
	p := strings.Split(version, ".")
	if len(p) >= 2 && len(p) < 4 {
	} else {
		return errors.New(fmt.Sprintf("incorrect lenght provided:%s version number try: v1.2.3", version))
	}
	return nil
}
