package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/gin-gonic/gin"
)

// AddUploadApp
// upload the build
func (inst *Controller) AddUploadApp(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:      c.Query("name"),
		BuildName: c.Query("buildName"),
		Version:   c.Query("version"),
		File:      file,
	}
	data, err := inst.Rubix.AddUploadEdgeApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// UploadService
// upload the service file
func (inst *Controller) UploadService(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &installer.Upload{
		Name:      c.Query("name"),
		BuildName: c.Query("buildName"),
		Version:   c.Query("version"),
		File:      file,
	}
	data, err := inst.Rubix.UploadServiceFile(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) InstallService(c *gin.Context) {
	var m *installer.Install
	err = c.ShouldBindJSON(&m)
	data, err := inst.Rubix.InstallService(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
