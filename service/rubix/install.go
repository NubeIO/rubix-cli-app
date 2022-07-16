package rubix

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// InstallApp make all the required dirs
func (inst *App) InstallApp(appName, appBuildName, version string, zip *multipart.FileHeader, localZip string) error {
	// make the dirs
	err := inst.DirsInstallApp(appName, appBuildName, version)
	if err != nil {
		return err
	}
	log.Infof("made all dirs for app:%s,  buildName:%s, version:%s", appName, appBuildName, version)
	var zipSource string
	if localZip != "" { //
		source := fmt.Sprintf("%s/%s", HostDownloadPath, localZip)
		dest := fmt.Sprintf("%s/%s", TmpDir, localZip)
		zipSource = dest
		err := MoveFile(source, dest, false)
		if err != nil {
			log.Errorf("move zip:%s: err:%s", dest, err.Error())
			return err
		}
	} else {
		// save app in tmp dir
		zipSource, err = inst.saveUploadedFile(zip, TmpDir)
		if err != nil {
			return err
		}
	}

	// unzip the build to the app dir  /data/rubix-service/install/wires-build
	zipDest := inst.getAppInstallPathAndVersion(appBuildName, version)
	_, err = inst.unZip(zipSource, zipDest)
	if err != nil {
		return err
	}
	return nil
}

func (inst *App) unZip(source, destination string) ([]string, error) {
	return fileutils.New().UnZip(source, destination, os.FileMode(FilePerm))
}

// SaveUploadedFile uploads the form file to specific dst.
// combination's of file name and the destination and will save file as: /data/my-file
// returns the filename and path as a string and any error
func (inst *App) saveUploadedFile(file *multipart.FileHeader, destination string) (string, error) {
	destination = fmt.Sprintf("%s/%s", destination, filepath.Base(file.Filename))
	src, err := file.Open()
	if err != nil {
		return destination, err
	}
	defer src.Close()

	out, err := os.Create(destination)
	if err != nil {
		return destination, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return destination, err
}

func MoveFile(sourcePath, destPath string, deleteExisting bool) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	if deleteExisting {
		err = os.Remove(sourcePath)
		if err != nil {
			return fmt.Errorf("failed removing original file: %s", err)
		}
	}
	return nil
}
