package system

import (
	"errors"
	"github.com/NubeIO/lib-ufw/ufw"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type UFWBody struct {
	Port int `json:"port"`
}

//UWFActive check status and also if ufw is installed
func (inst *System) UWFActive() (bool, error) {
	return inst.ufw.UWFActive()
}

//UWFStatus check status and also if ufw is installed
func (inst *System) UWFStatus() (*ufw.Message, error) {
	return inst.ufw.UWFStatus()
}

//UWFStatusList check status and also if ufw is installed
func (inst *System) UWFStatusList() ([]ufw.UFWStatus, error) {
	return inst.ufw.UWFStatusList()
}

func (inst *System) UWFOpenPort(body UFWBody) (*ufw.Message, error) {
	return inst.ufw.UWFOpenPort(body.Port)
}

func (inst *System) UWFClosePort(body UFWBody) (*ufw.Message, error) {
	return inst.ufw.UWFOpenPort(body.Port)
}

func (inst *System) ufwEnable(disable bool) (*Message, error) {
	cmdString := "echo y | sudo ufw enable"
	if disable {
		cmdString = "sudo ufw disable"
	}
	cmd := exec.Command("sh", "-c", cmdString)
	output, err := cmd.Output()
	res := cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		return &Message{Message: res}, err
	}
	return &Message{Message: res}, nil

}

func (inst *System) UWFEnable() (*Message, error) {
	//first make sure port 22 is open, this is to make sure we don't lock ourselves out
	cmdString := "sudo ufw allow 22"
	cmd := exec.Command("sh", "-c", cmdString)
	output, err := cmd.Output()
	res := cleanCommand(string(output), cmd, err, debug)
	if err != nil {
		log.Error("ufw: failed to open port 22: ", err)
		return &Message{Message: res}, errors.New("ufw: failed to open port 22")
	}
	return inst.ufwEnable(false)
}

//
//func (inst *System) UWFEnable() (*ufw.Message, error) {
//	return inst.ufw.UWFEnable()
//}

func (inst *System) UWFDisable() (*ufw.Message, error) {
	return inst.ufw.UWFDisable()
}
