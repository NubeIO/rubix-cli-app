package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/files"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

type ZipParams struct {
	Source             string `json:"source"`
	Destination        string `json:"destination"`
	Perm               int    `json:"perm"`
	OverrideIfExisting bool   `json:"override_if_existing"`
}

func getZipBody(c *gin.Context) (dto *ZipParams, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) Unzip(c *gin.Context) {
	body, err := getZipBody(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	pathToZip := body.Source
	destination := body.Destination
	if body.Source == "" {
		reposeHandler(nil, errors.New("zip source can not be empty, try /data/zip.zip"), c)
		return
	}
	if body.Destination == "" {
		reposeHandler(nil, errors.New("zip destination can not be empty, try /data/unzip-test"), c)
		return
	}
	zip, err := fileUtils.UnZip(pathToZip, destination, os.FileMode(body.Perm))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: fmt.Sprintf("unzipped successfully, files count: %d", len(zip))}, err, c)
}

func (inst *Controller) ZipDir(c *gin.Context) {
	body, err := getZipBody(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	pathToZip := body.Source
	destination := body.Destination
	if body.Source == "" {
		reposeHandler(nil, errors.New("zip source can not be empty, try /data/flow-framework"), c)
		return
	}
	if body.Destination == "" {
		reposeHandler(nil, errors.New("zip destination can not be empty, try /home/test/flow-framework.zip"), c)
		return
	}

	exists := fileUtils.DirExists(pathToZip)
	if !exists {
		reposeHandler(nil, errors.New("dir to zip not found"), c)
		return
	}
	files.MakeDirectoryIfNotExists(path.Dir(destination))
	err = fileUtils.RecursiveZip(pathToZip, destination)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: fmt.Sprintf("zip file is created on: %s", destination)}, nil, c)
}
