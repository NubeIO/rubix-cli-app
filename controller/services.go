package controller

import (
	"github.com/gin-gonic/gin"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
)

func getAppServiceBody(ctx *gin.Context) (dto *dbase.SystemCtl) {
	err = ctx.ShouldBindJSON(&dto)
	return dto
}

func (inst *Controller) AppService(ctx *gin.Context) {
	body := getAppServiceBody(ctx)
	action, err := inst.DB.SystemCtlAction(body)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	reposeHandler(action, err, ctx)
}
