package controller

import (
	"fmt"
	"github.com/NubeIO/edge/controller/httpresp"
	"github.com/gin-gonic/gin"
	"net/http"
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
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       "zip source can not be empty, try /home/user/zip.zip",
			"existing":    pathToZip,
			"destination": destination,
		})
		return
	}
	if body.Destination == "" {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       "zip destination can not be empty, try /home/user",
			"existing":    pathToZip,
			"destination": destination,
		})
		return
	}
	zip, err := fileUtils.UnZip(body.Source, body.Destination, os.FileMode(body.Perm))
	if err != nil {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       err.Error(),
			"existing":    pathToZip,
			"destination": destination,
		})
		return
	}
	httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{
		"message":     fmt.Sprintf("ok amout of files unziped %d", len(zip)),
		"existing":    pathToZip,
		"destination": destination,
	})
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
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       "zip source can not be empty, try /home/user/test",
			"existing":    pathToZip,
			"destination": destination,
		})
		return
	}
	if body.Destination == "" {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       "zip destination can not be empty, try /home/user/test.zip",
			"existing":    pathToZip,
			"destination": destination,
		})
		return
	}

	exists := fileUtils.DirExists(pathToZip)
	if !exists {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       "existing dir not found",
			"existing":    pathToZip,
			"destination": destination,
		})
		return
	}
	err = fileUtils.RecursiveZip(pathToZip, destination)
	if err != nil {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       err.Error(),
			"existing":    pathToZip,
			"destination": destination,
		})
		return
	}
	httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{
		"message":     "zipped ok",
		"existing":    pathToZip,
		"destination": destination,
	})
}
