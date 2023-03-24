package controller

import (
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (inst *Controller) GetSystem(c *gin.Context) {
	data, err := inst.System.GetSystem()
	responseHandler(data, err, c)
}

func (inst *Controller) GetMemoryUsage(c *gin.Context) {
	data, err := inst.System.GetMemoryUsage()
	responseHandler(data, err, c)
}

func (inst *Controller) GetMemory(c *gin.Context) {
	data, err := inst.System.GetMemory()
	responseHandler(data, err, c)
}

func (inst *Controller) GetTopProcesses(c *gin.Context) {
	count, err := strconv.Atoi(c.Query("count"))
	m := system.TopProcesses{
		Count: count,
		Sort:  c.Query("sort"),
	}
	data, err := inst.System.GetTopProcesses(m)
	responseHandler(data, err, c)
}

func (inst *Controller) GetSwap(c *gin.Context) {
	data, err := inst.System.GetSwap()
	responseHandler(data, err, c)
}

func (inst *Controller) DiscUsage(c *gin.Context) {
	data, err := inst.System.DiscUsage()
	responseHandler(data, err, c)
}

func (inst *Controller) DiscUsagePretty(c *gin.Context) {
	data, err := inst.System.DiscUsagePretty()
	responseHandler(data, err, c)
}
