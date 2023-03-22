package controller

import (
	"errors"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/pkg/interfaces"
	"github.com/NubeIO/rubix-edge/utils"
	"github.com/gin-gonic/gin"
)

var rebootHostJobTag = "reboot.host"

func getBodyRebootJob(c *gin.Context) (dto *interfaces.RebootJob, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetRebootHostJob(c *gin.Context) {
	rebootJob := utils.GetRebootJob()
	if rebootJob == nil {
		responseHandler(nil, errors.New("reboot job not found"), c)
		return
	}
	responseHandler(rebootJob, nil, c)
}

func (inst *Controller) UpdateRebootHostJob(c *gin.Context) {
	body, err := getBodyRebootJob(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	err = utils.ValidateCornExpression(body.Expression)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	body.Tag = rebootHostJobTag
	err = utils.SaveRebootJob(body, inst.FileMode)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	_, err = inst.Scheduler.Cron(body.Expression).Tag(body.Tag).Do(func() {
		_, _ = inst.System.RebootHost()
	})
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(body, nil, c)
}

func (inst *Controller) DeleteRebootHostJob(c *gin.Context) {
	err := utils.SaveRebootJob(nil, inst.FileMode)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	err = inst.Scheduler.RemoveByTag(rebootHostJobTag)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(model.Message{Message: "deleted system reboot job successfully"}, nil, c)
}
