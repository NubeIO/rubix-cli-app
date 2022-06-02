package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const dirPath = "/home/aidan/testing/" //TODO add in config

/*
UploadFile
//curl -X POST http://localhost:8090/api/files/upload/code/go   -F "file=@/home/user/Downloads/bios-master.zip"   -H "Content-Type: multipart/form-data"
*/
func (inst *Controller) UploadFile(c *gin.Context) {
	localSystemFilePath := ConcatPath(c.Param("filePath"))
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		response.ReposeHandler(c, http.StatusOK, response.Error, err)
		return
	}
	fileFull := fmt.Sprintf("%s/%s", localSystemFilePath, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, fileFull); err != nil {
		response.ReposeHandler(c, http.StatusOK, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, gin.H{"uploaded": fileFull})
}

/*
DownloadFile
curl -X POST http://localhost:8090/api/files/download/<pathAndFile>
eg curl -X POST http://localhost:8090/api/files/download/code/go/nube/rubix-cli-app/docs/api.md
*/
func (inst *Controller) DownloadFile(c *gin.Context) {
	inst.readFiles(c, true)
}

/*
ReadDirs
curl -X GET http://localhost:8090/api/files/dirs/<path>

*/
func (inst *Controller) ReadDirs(c *gin.Context) {
	inst.readFiles(c, false)
}

/*
DeleteFile
curl -X DELETE http://localhost:8090/api/files/delete/<pathAndFile>
*/
func (inst *Controller) DeleteFile(c *gin.Context) {
	inst.delete(c, false, false)
}

/*
DeleteDir
curl -X DELETE http://localhost:8090/api/files/delete/<pathAndFile>
*/
func (inst *Controller) DeleteDir(c *gin.Context) {
	inst.delete(c, true, false)
}

/*
DeleteDirForce
curl -X DELETE http://localhost:8090/api/files/delete/<pathAndFile>
*/
func (inst *Controller) DeleteDirForce(c *gin.Context) {
	inst.delete(c, true, true)
}

func (inst *Controller) delete(c *gin.Context, deleteDir, forceWipeOnDeleteDir bool) {
	localSystemFilePath := ConcatPath(c.Param("filePath"))

	if !deleteDir { //delete  a file
		if !fileUtils.FileExists(localSystemFilePath) {
			response.ReposeHandler(c, http.StatusOK, response.Error, err)
			return
		}
		response.ReposeHandler(c, http.StatusOK, response.Success, gin.H{"path": localSystemFilePath})
		return
	}

	if deleteDir { //delete  a dir
		if !fileUtils.DirExists(localSystemFilePath) {
			response.ReposeHandler(c, http.StatusOK, response.Error, err)
			return
		}
		if forceWipeOnDeleteDir {
			err := fileUtils.RmRF(localSystemFilePath)
			if err != nil {
				response.ReposeHandler(c, http.StatusOK, response.Error, err)
				return
			}
		} else {
			err := fileUtils.Rm(localSystemFilePath)
			if err != nil {
				response.ReposeHandler(c, http.StatusOK, response.Error, err)
				return
			}
		}
		response.ReposeHandler(c, http.StatusOK, response.Success, gin.H{"path": localSystemFilePath})
		return
	}
}

func ConcatPath(localSystemFilePath string) string {
	localSystemFilePath = path.Join(filepath.Dir(dirPath), localSystemFilePath)
	return localSystemFilePath
}

func (inst *Controller) readFiles(c *gin.Context, downloadFile bool) {
	localSystemFilePath := ConcatPath(c.Param("filePath")) // /api/files/*filePath
	fileInfo, err := os.Stat(localSystemFilePath)
	fileName, _ := filepath.Abs(localSystemFilePath)
	outFileName := fmt.Sprintf("attachment; %s", filepath.Base(fileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			response.ReposeHandler(c, http.StatusOK, response.StatusNotFound, err)
		} else {
			response.ReposeHandler(c, http.StatusOK, response.Error, err)
		}
		return
	}
	var dirContent []string
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(localSystemFilePath)
		if err != nil {
			response.ReposeHandler(c, http.StatusOK, response.Error, err)
			return
		}
		for _, file := range files {
			dirContent = append(dirContent, file.Name())
		}
	} else {
		if downloadFile {
			byteFile, err := ioutil.ReadFile(localSystemFilePath)
			if err != nil {
				response.ReposeHandler(c, http.StatusOK, response.Error, err)
				return
			}
			c.Header("Content-Disposition", outFileName)
			c.Data(http.StatusOK, "application/octet-stream", byteFile)
		}
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, gin.H{"path": dirContent})
}
