package controller

import (
	"github.com/gin-gonic/gin"
)

func (inst *Controller) ListFullBackups(c *gin.Context) {
	data, err := inst.Rubix.ListFullBackups()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) ListAppBackupsDirs(c *gin.Context) {
	data, err := inst.Rubix.ListAppBackupsDirs()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}

func (inst *Controller) ListBackupsByApp(c *gin.Context) {
	appName := c.Query("name")
	data, err := inst.Rubix.ListBackupsByApp(appName)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, nil, c)
}
