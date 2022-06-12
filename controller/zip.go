package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type UnzipParams struct {
	Source             string `json:"source"`
	Destination        string `json:"destination"`
	Perm               int    `json:"perm"`
	OverrideIfExisting bool   `json:"override_if_existing"`
}

func getZipBody(c *gin.Context) (dto *UnzipParams, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

/*
Unzip
curl -X DELETE http://localhost:8090/api/files/delete/<pathAndFile>
*/
func (inst *Controller) Unzip(c *gin.Context) {
	body, err := getZipBody(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if body.Source == "" {
		reposeHandler(nil, errors.New(fmt.Sprintf("zip sorce can not be empty, try /home/user/zip.zip")), c)
		return
	}
	if body.Destination == "" {
		reposeHandler(nil, errors.New(fmt.Sprintf("zip destination can not be empty, try /home/user")), c)
		return
	}
	zip, err := fileUtils.UnZip(body.Source, body.Destination, os.FileMode(body.Perm))
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{fmt.Sprintf("amout of files unziped %d", len(zip))}, err, c)
}
