package controller

import (
	"errors"
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

func (inst *Controller) ListFiles(c *gin.Context) {
	path := c.Query("path")
	fileInfo, err := os.Stat(path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var dirContent []string
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
		for _, file := range files {
			dirContent = append(dirContent, file.Name())
		}
	} else {
		reposeHandler(dirContent, errors.New("it needs to be a directory, found file"), c)
		return
	}

	reposeHandler(dirContent, nil, c)
}

func (inst *Controller) RenameFile(c *gin.Context) {
	oldName := c.Query("old_name_and_path")
	newName := c.Query("new_name_and_path")
	if oldName == "" || newName == "" {
		reposeHandler(nil, errors.New("directory, from and to files name can not be empty"), c)
		return
	}
	err = fileUtils.Rename(oldName, newName)
	reposeHandler(Message{Message: "renaming is successfully done"}, err, c)
}

func (inst *Controller) CopyFile(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if from == "" || to == "" {
		reposeHandler(nil, errors.New("from and to files name can not be empty"), c)
		return
	}
	err = fileUtils.Copy(from, to)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "copying is successfully done"}, err, c)
}

func (inst *Controller) MoveFile(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if from == "" || to == "" {
		reposeHandler(nil, errors.New("from and to files name can not be empty"), c)
		return
	}
	err = fileUtils.MoveFile(from, to)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(Message{Message: "moving is successfully done"}, err, c)
}

func (inst *Controller) DownloadFile(c *gin.Context) {
	path := c.Query("path_and_file")
	fileInfo, err := os.Stat(path)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	var dirContent []string
	if fileInfo.IsDir() {
		reposeHandler(dirContent, errors.New("it needs to be a file, found directory"), c)
		return
	} else {
		byteFile, err := ioutil.ReadFile(path)
		if err != nil {
			reposeHandler(nil, err, c)
			return
		}
		fileName, _ := filepath.Abs(path)
		outFileName := fmt.Sprintf("attachment; filename=%s", filepath.Base(fileName))
		c.Header("Content-Disposition", outFileName)
		c.Data(http.StatusOK, "application/octet-stream", byteFile)
	}
	reposeHandler(dirContent, nil, c)
}

/*
UploadFile
// curl -X POST http://localhost:1661/api/files/upload?to=/data/ -F "file=@/home/user/Downloads/bios-master.zip" -H "Content-Type: multipart/form-data"
*/
func (inst *Controller) UploadFile(c *gin.Context) {
	now := time.Now()
	to := c.Query("to")
	file, err := c.FormFile("file")
	resp := &UploadResponse{}
	if err != nil || file == nil {
		reposeHandler(resp, err, c)
		return
	}
	toFileLocation := fmt.Sprintf("%s/%s", to, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, toFileLocation); err != nil {
		reposeHandler(resp, err, c)
		return
	}
	size, err := fileutils.GetFileSize(toFileLocation)
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

func (inst *Controller) DeleteFile(c *gin.Context) {
	filePath := c.Query("path")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	if !fileUtils.FileExists(filePath) {
		reposeHandler(nil, errors.New("file doesn't exist"), c)
		return
	}
	err = fileUtils.Rm(filePath)
	reposeHandler(Message{Message: "file has been deleted"}, err, c)
}

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

type UploadResponse struct {
	Destination string `json:"destination"`
	File        string `json:"file"`
	Size        string `json:"size"`
	UploadTime  string `json:"upload_time"`
}
