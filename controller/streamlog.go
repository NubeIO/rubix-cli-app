package controller

import (
	"errors"
	"github.com/NubeIO/rubix-edge/pkg/streamlog"
	"github.com/gin-gonic/gin"
)

func getBodyLogStreamCreate(c *gin.Context) (dto *streamlog.Log, err error) {
	err = c.ShouldBindJSON(&dto)
	return dto, err
}

func (inst *Controller) GetStreamLogs(c *gin.Context) {
	logStreams := streamlog.GetStreamsLogs()
	reposeHandler(logStreams, nil, c)
}

func (inst *Controller) GetStreamLog(c *gin.Context) {
	u := c.Param("uuid")
	logStream := streamlog.GetStreamLog(u)
	if logStream == nil {
		reposeHandler(nil, errors.New("log not found"), c)
		return
	}
	reposeHandler(logStream, nil, c)
}

func (inst *Controller) CreateStreamLog(c *gin.Context) {
	body, err := getBodyLogStreamCreate(c)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	uuid := streamlog.CreateStreamLog(body)
	reposeHandler(map[string]interface{}{"UUID": uuid}, err, c)
}

func (inst *Controller) DeleteStreamLog(c *gin.Context) {
	u := c.Param("uuid")
	deleted := streamlog.DeleteStreamLog(u)
	if !deleted {
		reposeHandler(nil, errors.New("log not found"), c)
		return
	}
	reposeHandler(deleted, nil, c)
}

func (inst *Controller) DeleteStreamLogs(c *gin.Context) {
	streamlog.DeleteStreamLogs()
	reposeHandler(true, nil, c)
}
