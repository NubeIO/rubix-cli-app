package controller

import (
	dbase "github.com/NubeIO/edge/database"
	"github.com/gin-gonic/gin"
)

func getAppServiceBody(c *gin.Context) (dto *dbase.SystemCtl) {
	err = c.ShouldBindJSON(&dto)
	return dto
}

func (inst *Controller) AppService(c *gin.Context) {
	body := getAppServiceBody(c)
	data, err := inst.DB.SystemCtlAction(body)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}
