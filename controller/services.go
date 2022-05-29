package controller

import (
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/service/apps"
)

func getAppServiceBody(ctx *gin.Context) (dto *apps.AppService) {
	err = ctx.ShouldBindJSON(&dto)
	return dto
}

func (inst *Controller) AppService(ctx *gin.Context) {
	//body := getAppServiceBody(ctx)
	//action, err := inst.Apps.Action(body)
	//if err != nil {
	//	reposeHandler(nil, err, ctx)
	//	return
	//}
	//reposeHandler(action, err, ctx)
}
