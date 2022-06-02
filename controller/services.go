package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
	"net/http"
)

func getAppServiceBody(c *gin.Context) (dto *dbase.SystemCtl) {
	err = c.ShouldBindJSON(&dto)
	return dto
}

func (inst *Controller) AppService(c *gin.Context) {
	body := getAppServiceBody(c)
	data, err := inst.DB.SystemCtlAction(body)
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}
