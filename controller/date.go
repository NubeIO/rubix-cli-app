package controller

import (
	"github.com/NubeIO/lib-date/datelib"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) HostTime(c *gin.Context) {

	data, err := datelib.New(&datelib.Date{}).SystemTime()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(data, err, c)
}
