package controller

import (
	"github.com/gin-gonic/gin"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
)

func (inst *Controller) InstallApp(c *gin.Context) {
	var m *dbase.App
	err = c.ShouldBindJSON(&m)
	data, err := inst.DB.InstallApp(m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}
