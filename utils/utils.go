package utils

import (
	"errors"
	"github.com/NubeIO/lib-files/fileutils"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func DeleteDir(source, parentDirectory string, depth int) error {
	dir, _ := os.Open(source)
	obs, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	var errs []error

	for _, obj := range obs {
		fSource := path.Join(source, obj.Name())
		if obj.IsDir() {
			if parentDirectory == "rubix-service/apps/install" &&
				!Contains([]string{"rubix-edge", "rubix-assist"}, obj.Name()) {
				_ = os.RemoveAll(path.Join(parentDirectory, obj.Name()))
			}
			err = DeleteDir(fSource, path.Join(parentDirectory, obj.Name()), depth+1)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	var errString string
	for _, err := range errs {
		errString += err.Error() + "\n"
	}
	if errString != "" {
		return errors.New(errString)
	}
	return nil
}

func CopyDir(source, dest, parentDirectory string, depth int) error {
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
		go func(obj os.FileInfo) {
			defer wg.Done()
			fSource := path.Join(source, obj.Name())
			fDest := path.Join(dest, obj.Name())
			if obj.IsDir() {
				excludesData := []string{"rubix-edge", "rubix-assist", "tmp", "store", "backup", "socat"}
				excludesApps := []string{
					"rubix-service/apps/install/rubix-edge",
					"rubix-service/apps/install/rubix-assist",
				}
				if !((Contains(excludesData, obj.Name()) && depth == 0) ||
					(Contains(excludesApps, path.Join(parentDirectory, obj.Name())) && depth == 3)) {
					err = CopyDir(fSource, fDest, path.Join(parentDirectory, obj.Name()), depth+1)
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
		}(obj)
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
			if !Contains([]string{"nubeio-rubix-edge.service", "nubeio-rubix-assist.service"}, srcFile) {
				err := fileutils.CopyFile(srcFile, path.Join(dest, filepath.Base(srcFile)))
				if err != nil {
					log.Errorf("failed to copy file %s to %s", srcFile, path.Join(dest, filepath.Base(srcFile)))
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
		go func(srcFile string) {
			defer wg.Done()
			err := os.RemoveAll(path.Join(dest, filepath.Base(srcFile)))
			if err != nil {
				log.Errorf("failed to remove file %s", os.RemoveAll(path.Join(dest, filepath.Base(srcFile))))
			}
		}(srcFile)
	}
	wg.Wait()
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
