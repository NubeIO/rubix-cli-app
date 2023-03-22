package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/pkg/interfaces"
	"github.com/NubeIO/rubix-edge/utils"
	"github.com/gin-gonic/gin"
)

func getBodyRestartJob(c *gin.Context) (dto *interfaces.RestartJob, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetRestartJob(c *gin.Context) {
	restartJobs := utils.GetRestartJobs()
	if restartJobs == nil {
		responseHandler([]interfaces.RestartJob{}, nil, c)
		return
	}
	responseHandler(restartJobs, nil, c)
}

func (inst *Controller) UpdateRestartJob(c *gin.Context) {
	body, err := getBodyRestartJob(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	err = utils.ValidateCornExpression(body.Expression)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	update := false
	restartJobs := utils.GetRestartJobs()
	for i, restartJob := range restartJobs {
		if restartJob.Unit == body.Unit {
			restartJobs[i] = body
			_ = inst.Scheduler.RemoveByTag(restartJob.Unit)
			update = true
			break
		}
	}
	if !update {
		restartJobs = append(restartJobs, body)
	}
	err = utils.SaveRestartJobs(restartJobs, inst.FileMode)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	_, err = inst.Scheduler.Cron(body.Expression).Tag(body.Unit).Do(func() {
		_ = inst.SystemCtl.Restart(body.Unit)
	})
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(body, nil, c)
}

func (inst *Controller) DeleteRestartJob(c *gin.Context) {
	unit := c.Param("unit")
	restartJobs := utils.GetRestartJobs()
	deleted := false
	for i, restartJob := range restartJobs {
		if restartJob.Unit == unit {
			err := inst.Scheduler.RemoveByTag(restartJob.Unit)
			if err != nil {
				responseHandler(nil, err, c)
				return
			}
			restartJobs = append(restartJobs[:i], restartJobs[i+1:]...)
			deleted = true
			break
		}
	}
	if !deleted {
		err := errors.New("unit not found")
		responseHandler(nil, err, c)
		return
	}
	err := utils.SaveRestartJobs(restartJobs, inst.FileMode)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(model.Message{Message: fmt.Sprintf("deleted %s restart job successfully", unit)}, nil, c)
}
