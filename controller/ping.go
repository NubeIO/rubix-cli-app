package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"time"
)

type PingBody struct {
	Ip      string        `json:"ip"`
	Port    int           `json:"port"`
	TimeOut time.Duration `json:"time_out"`
}

// Ping ping from the edge device
func Ping(ip string, port int, timeOut time.Duration) bool {
	if timeOut == 0 {
		timeOut = 1000 * time.Millisecond
	}
	ip_ := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", ip_, timeOut)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}

func (inst *Controller) Ping(c *gin.Context) {
	var body *PingBody
	err := c.ShouldBindJSON(&body)
	if body != nil || err != nil {
		responseHandler(nil, errors.New("ping body can not be empty"), c)
		return
	}
	ping := Ping(body.Ip, body.Port, body.TimeOut)
	responseHandler(ping, nil, c)
}
