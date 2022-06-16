package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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
	pathToZip := body.Source
	destination := body.Destination
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if body.Source == "" {
		reposeHandler(nil, errors.New("zip source can not be empty, try /home/user/zip.zip"), c)
		return
	}
	if body.Destination == "" {
		reposeHandler(nil, errors.New("zip destination can not be empty, try /home/user"), c)
		return
	}
	zip, err := fileUtils.UnZip(pathToZip, destination, os.FileMode(body.Perm))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: fmt.Sprintf("ok amout of files unziped %d", len(zip))}, err, c)
}

func (inst *Controller) ZipDir(c *gin.Context) {
	body, err := getZipBody(c)
	pathToZip := body.Source
	destination := body.Destination
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if body.Source == "" {
		reposeHandler(nil, errors.New("zip source can not be empty, try try /home/user/test"), c)
		return
	}
	if body.Destination == "" {
		reposeHandler(nil, errors.New("zip destination can not be empty, try try /home/user/test.zip"), c)
		return
	}

	exists := fileUtils.DirExists(pathToZip)
	if !exists {
		reposeHandler(nil, errors.New("dir to zip not found"), c)
		return
	}
	err = fileUtils.RecursiveZip(pathToZip, destination)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: fmt.Sprintf("new zip:%s", destination)}, nil, c)

}
