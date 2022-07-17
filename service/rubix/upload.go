package rubix

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Upload struct {
	Name      string                `json:"name"`
	BuildName string                `json:"build_name"`
	Version   string                `json:"version"`
	File      *multipart.FileHeader `json:"file"`
}

type UploadResponse struct {
	TmpFile      string `json:"tmp_file"`
	UploadedFile string `json:"uploaded_file"`
}

func (inst *App) UploadApp(app *Upload) (*UploadResponse, error) {
	var appName = app.Name
	var appBuildName = app.BuildName
	var version = app.Version
	var file = app.File
	return inst.uploadApp(appName, appBuildName, version, file)
}

// uploadApp
func (inst *App) uploadApp(appName, appBuildName, version string, zip *multipart.FileHeader) (*UploadResponse, error) {
	// make the dirs
	var err error
	if err := inst.MakeTmpDir(); err != nil {
		return nil, err
	}
	var tmpDir string
	if tmpDir, err = inst.MakeTmpDirUpload(); err != nil {
		return nil, err
	}
	log.Infof("upload build to tmp dir:%s", tmpDir)
	log.Infof("app:%s buildName:%s version:%s", appName, appBuildName, version)
	// save app in tmp dir
	zipSource, err := inst.saveUploadedFile(zip, tmpDir)
	if err != nil {
		return nil, err
	}
	return &UploadResponse{
		TmpFile:      tmpDir,
		UploadedFile: zipSource,
	}, err
}

func (inst *App) UploadServiceFile(app *Upload) (*UploadResponse, error) {
	var appName = app.Name
	var appBuildName = app.BuildName
	var version = app.Version
	var file = app.File
	return inst.uploadServiceFile(appName, appBuildName, version, file)
}

// uploadApp
func (inst *App) uploadServiceFile(appName, appBuildName, version string, file *multipart.FileHeader) (*UploadResponse, error) {
	// make the dirs
	var err error
	if filepath.Ext(file.Filename) != ".service" {
		return nil, errors.New(fmt.Sprintf("service file provided:%s, did not have correct file extension must be (.service)", file.Filename))
	}
	if err := inst.MakeTmpDir(); err != nil {
		return nil, err
	}
	var tmpDir string
	if tmpDir, err = inst.MakeTmpDirUpload(); err != nil {
		return nil, err
	}
	log.Infof("upload service to tmp dir:%s", tmpDir)
	log.Infof("app:%s buildName:%s version:%s", appName, appBuildName, version)
	// save app in tmp dir
	zipSource, err := inst.saveUploadedFile(file, tmpDir)
	if err != nil {
		return nil, err
	}
	return &UploadResponse{
		TmpFile:      tmpDir,
		UploadedFile: zipSource,
	}, err
}

func (inst *App) unZip(source, destination string) ([]string, error) {
	return fileutils.New().UnZip(source, destination, os.FileMode(FilePerm))
}

// SaveUploadedFile uploads the form file to specific dst.
// combination's of file name and the destination and will save file as: /data/my-file
// returns the filename and path as a string and any error
func (inst *App) saveUploadedFile(file *multipart.FileHeader, dest string) (destination string, err error) {
	destination = fmt.Sprintf("%s/%s", dest, filepath.Base(file.Filename))
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

func moveFile(sourcePath, destPath string, deleteExisting bool) error {
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
