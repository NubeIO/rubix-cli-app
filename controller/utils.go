package controller

import (
	"errors"
	"github.com/NubeIO/lib-files/fileutils"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func toBool(value string) (bool, error) {
	if value == "" {
		return false, nil
	} else {
		c, err := bools.Boolean(value)
		return c, err
	}
}

func resolve(name string) string {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return ""
	}
	return filepath.FromSlash(fileutils.SlashClean(name))
}

func copyFolder(source string, dest string) error {
	if source = resolve(source); source == "" {
		return os.ErrNotExist
	}
	if dest = resolve(dest); dest == "" {
		return os.ErrNotExist
	}
	srcinfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.MkdirAll(dest, srcinfo.Mode())
	if err != nil {
		return err
	}
	dir, _ := os.Open(source)
	obs, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	var errs []error
	for _, obj := range obs {
		wg.Add(1)
		obj := obj
		go func() {
			defer wg.Done()
			fsource := source + "/" + obj.Name()
			fdest := dest + "/" + obj.Name()

			if obj.IsDir() {
				if obj.Name() != "rubix-edge" {
					err = copyFolder(fsource, fdest)
					if err != nil {
						errs = append(errs, err)
					}
				}
			} else {
				err = fileutils.CopyFile(fsource, fdest)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}()
	}
	wg.Wait()
	var errString string
	for _, err := range errs {
		errString += err.Error() + "\n"
	}
	if errString != "" {
		return errors.New(errString)
	}
	return nil
}

func copyFiles(srcFiles []string, dest string) {
	var wg sync.WaitGroup
	for _, srcFile := range srcFiles {
		wg.Add(1)
		srcFile := srcFile
		go func() {
			defer wg.Done()
			if !strings.Contains(srcFile, "rubix-edge") {
				err := fileutils.CopyFile(srcFile, path.Join(dest, filepath.Base(srcFile)))
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
			}
		}()
	}
	wg.Wait()
}

func checkSnapshotSize(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.Size()/1024/1024/1024 > 1 {
		return errors.New("maximum response size reached")
	}
	return nil
}
