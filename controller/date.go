package controller

import (
	"github.com/NubeIO/lib-date/datelib"
	"github.com/gin-gonic/gin"
	"gthub.com/NubeIO/rubix-cli-app/controller/response"
	"net/http"
)

func (inst *Controller) HostTime(c *gin.Context) {
	data, err := datelib.New(&datelib.Date{}).SystemTime()
	if err != nil {
		response.ReposeHandler(c, http.StatusBadRequest, response.Error, err)
		return
	}
	response.ReposeHandler(c, http.StatusOK, response.Success, data)
}
