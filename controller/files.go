package controller

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

func TimeTrack(start time.Time) (out string) {
	elapsed := time.Since(start)
	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	out = fmt.Sprintf("%s took %s", name, elapsed)
	return out
}

const dirPath = "/home/aidan/testing" //TODO add in config

type UploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}

/*
UploadFile
// curl -X POST http://localhost:1661/api/files/upload?destination=/home/user   -F "file=@/home/user/Downloads/bios-master.zip"   -H "Content-Type: multipart/form-data"
*/
func (inst *Controller) UploadFile(c *gin.Context) {
	now := time.Now()
	destination := c.Query("destination")
	file, err := c.FormFile("file")
	resp := &UploadResponse{}
	if err != nil || file == nil {
		reposeHandler(resp, err, c)
		return
	}
	fileFull := fmt.Sprintf("%s/%s", destination, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, fileFull); err != nil {
		reposeHandler(resp, err, c)
		return
	}
	size, err := fileutils.GetFileSize(fileFull)
	if err != nil {
		reposeHandler(resp, err, c)
	}
	resp = &UploadResponse{
		Destination: file.Filename,
		File:        file.Filename,
		Size:        size.String(),
		UploadTime:  TimeTrack(now),
	}
	reposeHandler(resp, nil, c)

}

/*
DownloadFile
curl -X GET http://localhost:1661/api/files/download/<pathAndFile>
eg curl -X GET http://localhost:1661/api/dirs/Downloads/flow-framework-0.5.4-13ac6506.amd64.zip
*/
func (inst *Controller) DownloadFile(c *gin.Context) {
	inst.readFiles(c, true)
}

/*
ReadDirs
curl -X GET http://localhost:1661/api/files/dirs/<path>

*/
func (inst *Controller) ReadDirs(c *gin.Context) {
	inst.readFiles(c, false)
}

// RenameFile ...
func (inst *Controller) RenameFile(c *gin.Context) {
	existing := c.Query("existing")
	newName := c.Query("new")
	dir, _ := filepath.Split(existing)
	fileFull := fmt.Sprintf("%s/%s", dir, newName)

	if existing == "" || newName == "" {
		reposeHandler(nil, errors.New("path and file name can not be empty"), c)
		return
	}
	if !fileUtils.FileExists(existing) {
		reposeHandler(nil, errors.New("file not found"), c)
		return
	}
	err = fileUtils.Rename(existing, fileFull)
	reposeHandler(Message{Message: "rename ok"}, err, c)
}

func (inst *Controller) MoveFile(c *gin.Context) {
	existing := c.Query("existing")
	destination := c.Query("destination")

	if existing == "" || destination == "" {
		reposeHandler(nil, errors.New("existing and existing name can not be empty"), c)
		return
	}

	err := fileUtils.MoveFile(existing, destination)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "move ok"}, err, c)
}

func (inst *Controller) CopyDir(c *gin.Context) {
	existing := c.Query("existing")
	destination := c.Query("destination")
	exists := fileUtils.DirExists(existing)
	if !exists {
		reposeHandler(nil, errors.New("existing dir not found"), c)
		return
	}

	err := fileUtils.Copy(existing, destination)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "copy ok"}, err, c)
}

/*
DeleteFile
curl -X DELETE http://localhost:1661/api/files/delete/<pathAndFile>
*/
func (inst *Controller) DeleteFile(c *gin.Context) {
	inst.delete(c, false, false)
}

/*
DeleteDir
curl -X DELETE http://localhost:1661/api/files/delete/<pathAndFile>
*/
func (inst *Controller) DeleteDir(c *gin.Context) {
	inst.delete(c, true, false)
}

/*
DeleteDirForce
curl -X DELETE http://localhost:1661/api/files/force/<pathAndFile>
*/
func (inst *Controller) DeleteDirForce(c *gin.Context) {
	inst.delete(c, true, true)
}

func (inst *Controller) delete(c *gin.Context, deleteDir, forceWipeOnDeleteDir bool) {
	localSystemFilePath := concatPath(c.Param("filePath"))

	if !deleteDir { //delete  a file
		if !fileUtils.FileExists(localSystemFilePath) {
			reposeHandler(nil, errors.New("not found"), c)
			return
		}
		reposeHandler(Message{Message: "delete file ok"}, nil, c)
		return
	}

	if deleteDir { //delete  a dir
		if !fileUtils.DirExists(localSystemFilePath) {
			reposeHandler(nil, err, c)
			return
		}
		if forceWipeOnDeleteDir {
			err := fileUtils.RmRF(localSystemFilePath)
			if err != nil {
				reposeHandler(nil, err, c)
				return
			}
		} else {
			err := fileUtils.Rm(localSystemFilePath)
			if err != nil {
				reposeHandler(nil, err, c)
				return
			}
		}
		reposeHandler(Message{Message: "delete directory ok"}, nil, c)
		return
	}
}

func concatPath(localSystemFilePath string) string {
	localSystemFilePath = path.Join(filepath.Dir(dirPath), localSystemFilePath)
	return localSystemFilePath
}

func (inst *Controller) readFiles(c *gin.Context, downloadFile bool) {
	localSystemFilePath := concatPath(c.Param("filePath")) // /api/files/*filePath
	fileInfo, err := os.Stat(localSystemFilePath)
	fileName, _ := filepath.Abs(localSystemFilePath)
	outFileName := fmt.Sprintf("attachment; %s", filepath.Base(fileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			reposeHandler(nil, err, c)
		} else {
			reposeHandler(nil, err, c)
		}
		return
	}
	var dirContent []string
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(localSystemFilePath)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
		for _, file := range files {
			dirContent = append(dirContent, file.Name())
		}
	} else {
		if downloadFile {
			byteFile, err := ioutil.ReadFile(localSystemFilePath)
			if err != nil {
				reposeHandler(nil, err, c)
				return
			}
			c.Header("Content-Disposition", outFileName)
			c.Data(http.StatusOK, "application/octet-stream", byteFile)
		}
	}
	reposeHandler(dirContent, nil, c)
}
