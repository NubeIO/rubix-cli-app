package controller

import (
	"errors"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/system/files"
	"github.com/gin-gonic/gin"
)

type DirsParams struct {
	Path string `json:"path"`
	From string `json:"from"`
	To   string `json:"to"`
}

func getDirsBody(c *gin.Context) (dto *DirsParams, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) CreateDir(c *gin.Context) {
	body, err := getDirsBody(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if body.Path == "" {
		reposeHandler(nil, errors.New("path can not be empty"), c)
		return
	}
	err = files.MakeDirectoryIfNotExists(body.Path)
	reposeHandler(Message{Message: "directory creation is successfully executed"}, err, c)
}

func (inst *Controller) CopyDir(c *gin.Context) {
	body, err := getDirsBody(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if body.From == "" || body.To == "" {
		reposeHandler(nil, errors.New("from and to directories name can not be empty"), c)
		return
	}
	exists := fileUtils.DirExists(body.From)
	if !exists {
		reposeHandler(nil, errors.New("from dir not found"), c)
		return
	}

	err = fileUtils.Copy(body.From, body.To)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "copying directory is successfully executed"}, err, c)
}

func (inst *Controller) DeleteDir(c *gin.Context) {
	body, err := getDirsBody(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	force := c.Query("force") == "true"
	if body.Path == "" {
		reposeHandler(nil, errors.New("path can not be empty"), c)
		return
	}
	if !fileUtils.DirExists(body.Path) {
		reposeHandler(nil, err, c)
		return
	}
	if force {
		err := fileUtils.RmRF(body.Path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
	} else {
		err := fileUtils.Rm(body.Path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
	}
	reposeHandler(Message{Message: "deletion of directory is successfully executed"}, nil, c)
	return
}
