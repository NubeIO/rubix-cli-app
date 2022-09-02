package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-networking/networking"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-edge/pkg/model"
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
	err = c.ShouldBindJSON(&body)
	if body != nil || err != nil {
		reposeHandler(nil, errors.New("ping body can not be empty"), c)
		return
	}
	ping := Ping(body.Ip, body.Port, body.TimeOut)
	reposeHandler(ping, nil, c)
}

type DeviceProduct struct {
	Device     *model.DeviceInfo  `json:"device"`
	Product    *installer.Product `json:"product"`
	Networking []networking.NetworkInterfaces
}

func (inst *Controller) GetDeviceProduct(c *gin.Context) {
	device, err := inst.DB.GetDeviceInfo()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	product, err := inst.EdgeApp.App.GetProduct() // https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	networks, err := nets.GetNetworks()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	deviceProduct := &DeviceProduct{
		Device:     device,
		Product:    product,
		Networking: networks,
	}
	reposeHandler(deviceProduct, err, c)
}
