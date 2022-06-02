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

const dirPath = "/home/aidan/" //TODO add in config

func ConcatPath(localSystemFilePath string) string {
	localSystemFilePath = path.Join(filepath.Dir(dirPath), localSystemFilePath)
	return localSystemFilePath
}

/*
UploadFile
//curl -X POST http://localhost:8090/api/files/upload/code/go   -F "file=@/home/user/Downloads/bios-master.zip"   -H "Content-Type: multipart/form-data"
*/
func (inst *Controller) UploadFile(ctx *gin.Context) {
	localSystemFilePath := ConcatPath(ctx.Param("filePath"))
	file, err := ctx.FormFile("file")
	if err != nil || file == nil {
		response.ReposeHandler(ctx, http.StatusOK, response.Error, err)
		return
	}
	fileFull := fmt.Sprintf("%s/%s", localSystemFilePath, filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, fileFull); err != nil {
		response.ReposeHandler(ctx, http.StatusOK, response.Error, err)
		return
	}
	response.ReposeHandler(ctx, http.StatusOK, response.Success, gin.H{"uploaded": fileFull})
}

/*
DownloadFile
curl -X POST http://localhost:8080/api/files/download/<pathAndFile>
eg curl -X POST http://localhost:8090/api/files/download/code/go/nube/rubix-cli-app/docs/api.md
*/
func (inst *Controller) DownloadFile(ctx *gin.Context) {
	inst.readFiles(ctx, true)
}

/*
ReadDirs
curl -X GET http://localhost:8080/api/files/dirs/<path>

*/
func (inst *Controller) ReadDirs(ctx *gin.Context) {
	inst.readFiles(ctx, false)
}

/*
DeleteFile
curl -X DELETE http://localhost:8080/api/files/delete/<pathAndFile>
*/
func (inst *Controller) DeleteFile(ctx *gin.Context) {
	localSystemFilePath := ConcatPath(ctx.Param("filePath"))
	if _, err := os.Stat(localSystemFilePath); !errors.Is(err, os.ErrNotExist) {
		response.ReposeHandler(ctx, http.StatusOK, response.Error, err)
		return
	}
	if err := os.Remove(localSystemFilePath); err != nil {
		response.ReposeHandler(ctx, http.StatusOK, response.Error, err)
		return
	}
	response.ReposeHandler(ctx, http.StatusOK, response.Success, gin.H{"path": localSystemFilePath})
}

func (inst *Controller) readFiles(ctx *gin.Context, downloadFile bool) {
	localSystemFilePath := ConcatPath(ctx.Param("filePath")) // /api/files/*filePath
	fileInfo, err := os.Stat(localSystemFilePath)
	fileName, _ := filepath.Abs(localSystemFilePath)
	outFileName := fmt.Sprintf("attachment; %s", filepath.Base(fileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			response.ReposeHandler(ctx, http.StatusOK, response.StatusNotFound, err)
		} else {
			response.ReposeHandler(ctx, http.StatusOK, response.Error, err)
		}
		return
	}
	var dirContent []string
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(localSystemFilePath)
		if err != nil {
			response.ReposeHandler(ctx, http.StatusOK, response.Error, err)
			return
		}
		for _, file := range files {
			dirContent = append(dirContent, file.Name())
		}
	} else {
		if downloadFile {
			byteFile, err := ioutil.ReadFile(localSystemFilePath)
			if err != nil {
				response.ReposeHandler(ctx, http.StatusOK, response.Error, err)
				return
			}
			ctx.Header("Content-Disposition", outFileName)
			ctx.Data(http.StatusOK, "application/octet-stream", byteFile)
		}
	}
	response.ReposeHandler(ctx, http.StatusOK, response.Success, gin.H{"path": dirContent})
}
