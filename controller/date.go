package controller

import (
	"github.com/NubeIO/lib-date/datelib"
	"github.com/gin-gonic/gin"
)

func (inst *Controller) HostTime(c *gin.Context) {
	data := datelib.New(&datelib.Date{}).SystemTime()
	reposeHandler(data, err, c)
}
