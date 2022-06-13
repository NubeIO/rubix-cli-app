package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/edge/controller/httpresp"
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

/*
UploadFile
// curl -X POST http://localhost:1661/api/files/upload?destination=/home/user   -F "file=@/home/user/Downloads/bios-master.zip"   -H "Content-Type: multipart/form-data"
*/
func (inst *Controller) UploadFile(c *gin.Context) {
	now := time.Now()
	destination := c.Query("destination")
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
		return
	}
	fileFull := fmt.Sprintf("%s/%s", destination, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, fileFull); err != nil {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
		return
	}
	size, err := fileutils.GetFileSize(fileFull)
	if err != nil {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
	}
	httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{
		"file":        file.Filename,
		"destination": destination,
		"upload_time": TimeTrack(now),
		"size":        size.String(),
	})
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
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error": "path and file name can not be empty",
			"path":  existing,
			"new":   fileFull,
		})
		return
	}
	if !fileUtils.FileExists(existing) {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error": "file not found",
			"path":  existing,
			"new":   fileFull,
		})
		return
	}
	err = fileUtils.Rename(existing, fileFull)
	httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{
		"path": existing,
		"new":  fileFull,
	})
}

func (inst *Controller) MoveFile(c *gin.Context) {
	existing := c.Query("existing")
	destination := c.Query("destination")

	if existing == "" || destination == "" {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       "existing and existing name can not be empty",
			"existing":    existing,
			"destination": destination,
		})
		return
	}

	err := fileUtils.MoveFile(existing, destination)
	if err != nil {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       err.Error(),
			"existing":    existing,
			"destination": destination,
		})
		return
	}
	httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{
		"message":     "moved ok",
		"existing":    existing,
		"destination": destination,
	})
}

func (inst *Controller) CopyDir(c *gin.Context) {
	existing := c.Query("existing")
	destination := c.Query("destination")
	exists := fileUtils.DirExists(existing)
	if !exists {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       "existing dir not found",
			"existing":    existing,
			"destination": destination,
		})
		return
	}

	err := fileUtils.Copy(existing, destination)
	if err != nil {
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, gin.H{
			"error":       err.Error(),
			"existing":    existing,
			"destination": destination,
		})
		return
	}
	httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{
		"message":     "moved ok",
		"existing":    existing,
		"destination": destination,
	})
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
			httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
			return
		}
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{"path": localSystemFilePath})
		return
	}

	if deleteDir { //delete  a dir
		if !fileUtils.DirExists(localSystemFilePath) {
			httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
			return
		}
		if forceWipeOnDeleteDir {
			err := fileUtils.RmRF(localSystemFilePath)
			if err != nil {
				httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
				return
			}
		} else {
			err := fileUtils.Rm(localSystemFilePath)
			if err != nil {
				httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
				return
			}
		}
		httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{"path": localSystemFilePath})
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
			httpresp.ReposeHandler(c, http.StatusOK, httpresp.StatusNotFound, err)
		} else {
			httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
		}
		return
	}
	var dirContent []string
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(localSystemFilePath)
		if err != nil {
			httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
			return
		}
		for _, file := range files {
			fmt.Println(file)
			dirContent = append(dirContent, file.Name())
		}
	} else {
		if downloadFile {
			byteFile, err := ioutil.ReadFile(localSystemFilePath)
			if err != nil {
				httpresp.ReposeHandler(c, http.StatusOK, httpresp.Error, err)
				return
			}
			c.Header("Content-Disposition", outFileName)
			c.Data(http.StatusOK, "application/octet-stream", byteFile)
		}
	}
	httpresp.ReposeHandler(c, http.StatusOK, httpresp.Success, gin.H{"path": dirContent})
}
