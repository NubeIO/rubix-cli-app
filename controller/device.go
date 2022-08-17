package controller

import (
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-edge/pkg/model"
	"github.com/gin-gonic/gin"
)

type DeviceProduct struct {
	Device  *model.DeviceInfo  `json:"device"`
	Product *installer.Product `json:"product"`
}

func (inst *Controller) GetDeviceProduct(c *gin.Context) {
	device, err := inst.DB.GetDeviceInfo()
	if err != nil {
		reposeHandler(device, err, c)
		return
	}
	product, err := inst.Rubix.App.GetProduct() // https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7
	if err != nil {
		reposeHandler(device, err, c)
		return
	}
	deviceProduct := &DeviceProduct{
		Device:  device,
		Product: product,
	}
	reposeHandler(deviceProduct, err, c)
}

func (inst *Controller) GetDeviceInfo(c *gin.Context) {
	data, err := inst.DB.GetDeviceInfo()
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) UpdateDeviceInfo(c *gin.Context) {
	var m *model.DeviceInfo
	err = c.ShouldBindJSON(&m)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	data, err := inst.DB.UpdateDeviceInfo(m)
	if err != nil {
		reposeHandler(data, err, c)
		return
	}
	reposeHandler(data, err, c)
}

func (inst *Controller) DropDeviceInfo(c *gin.Context) {
	host, err := inst.DB.DropDeviceInfo()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	reposeHandler(host, err, c)
}
