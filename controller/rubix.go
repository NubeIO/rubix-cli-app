package controller

import (
	"github.com/NubeIO/rubix-edge/service/rubix"
	"github.com/gin-gonic/gin"
)

// UploadApp
// upload the build
func (inst *Controller) UploadApp(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	m := &rubix.Upload{
		Name:      c.Query("name"),
		BuildName: c.Query("buildName"),
		Version:   c.Query("version"),
		File:      file,
	}
	data, err := inst.Rubix.UploadApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

// InstallApp
// make all the dirs and install the uploaded build
func (inst *Controller) InstallApp(c *gin.Context) {
	var m *rubix.Install
	err = c.ShouldBindJSON(&m)
	data, err := inst.Rubix.InstallApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) InstallAppService(c *gin.Context) {
	var m *rubix.Install
	err = c.ShouldBindJSON(&m)
	data, err := inst.Rubix.InstallApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
