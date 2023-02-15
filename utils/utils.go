package utils

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

func ToBool(value string) (bool, error) {
	if value == "" {
		return false, nil
	} else {
		c, err := bools.Boolean(value)
		return c, err
	}
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func CopyDir(source string, dest string) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.MkdirAll(dest, srcInfo.Mode())
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
			fSource := path.Join(source, obj.Name())
			fDest := path.Join(dest, obj.Name())
			if obj.IsDir() {
				if obj.Name() != "rubix-edge" && obj.Name() != "tmp" && obj.Name() != "store" && obj.Name() != "backup" {
					err = CopyDir(fSource, fDest)
					if err != nil {
						errs = append(errs, err)
					}
				}
			} else {
				err = fileutils.CopyFile(fSource, fDest)
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

func CopyFiles(srcFiles []string, dest string) {
	var wg sync.WaitGroup
	for _, srcFile := range srcFiles {
		wg.Add(1)
		go func(srcFile string) {
			defer wg.Done()
			if srcFile != "nubeio-rubix-edge.service" {
				err := fileutils.CopyFile(srcFile, path.Join(dest, filepath.Base(srcFile)))
				if err != nil {
					log.Errorf("err: %s", err.Error())
				}
			}
		}(srcFile)
	}
	wg.Wait()
}

func DeleteFiles(srcFiles []string, dest string) {
	var wg sync.WaitGroup
	for _, srcFile := range srcFiles {
		wg.Add(1)
		srcFile := srcFile
		go func() {
			defer wg.Done()
			err := os.RemoveAll(path.Join(dest, filepath.Base(srcFile)))
			if err != nil {
				log.Errorf("err: %s", err.Error())
			}
		}()
	}
	wg.Wait()
}
