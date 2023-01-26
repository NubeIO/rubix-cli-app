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
	responseHandler(logStreams, nil, c)
}

func (inst *Controller) GetStreamLog(c *gin.Context) {
	u := c.Param("uuid")
	logStream := streamlog.GetStreamLog(u)
	if logStream == nil {
		responseHandler(nil, errors.New("log not found"), c)
		return
	}
	responseHandler(logStream, nil, c)
}

func (inst *Controller) CreateStreamLog(c *gin.Context) {
	body, err := getBodyLogStreamCreate(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	uuid, err := streamlog.CreateStreamLog(body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(map[string]interface{}{"UUID": uuid}, err, c)
}

func (inst *Controller) CreateLogAndReturn(c *gin.Context) {
	body, err := getBodyLogStreamCreate(c)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	logStream, err := streamlog.CreateLogAndReturn(body)
	if err != nil {
		responseHandler(nil, err, c)
		return
	}
	responseHandler(logStream, nil, c)
}

func (inst *Controller) DeleteStreamLog(c *gin.Context) {
	u := c.Param("uuid")

	deleted := streamlog.DeleteStreamLog(u)
	if !deleted {
		responseHandler(nil, errors.New("log not found"), c)
		return
	}
	responseHandler(deleted, nil, c)
}

func (inst *Controller) DeleteStreamLogs(c *gin.Context) {
	streamlog.DeleteStreamLogs()
	responseHandler(true, nil, c)
}
