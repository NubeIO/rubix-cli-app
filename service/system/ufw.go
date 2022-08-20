package system

import (
	"errors"
	"fmt"
	pprint "github.com/NubeIO/rubix-edge/pkg/helpers/print"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strconv"
	"strings"
)

type UWFInstall struct {
	AlreadyInstalled bool   `json:"already_installed"`
	InstalledOk      bool   `json:"installed_ok"`
	TextOut          string `json:"text_out"`
}

//UWFActive check status and also if ufw is installed
func (inst *System) UWFActive() (bool, error) {
	cmd := exec.Command("sh", "-c", "sudo ufw status")
	output, err := cmd.Output()
	res := cleanCommand(string(output), cmd, err, debug)
	if strings.Contains(res, "Status: active") {
		return true, nil
	} else if strings.Contains(res, "Status: inactive") {
		return false, err
	} else {
		return false, errors.New("failed to check ufw status")
	}
}

//UWFStatus check status and also if ufw is installed
func (inst *System) UWFStatus() (*Message, error) {
	active, err := inst.UWFActive()
	if err != nil {
		return nil, err
	}
	msg := "firewall is disabled"
	if active {
		msg = "firewall is enabled"
	}
	return &Message{
		Message: msg,
	}, err
}

//UWFStatusList check status and also if ufw is installed
func (inst *System) UWFStatusList() ([]UFWStatus, error) {
	cmd := exec.Command("sh", "-c", "sudo ufw status")
	output, err := cmd.Output()
	res := cleanCommand(string(output), cmd, err, debug)
	if res != "" {
		list := ufwStats(res)
		if debug {
			pprint.PrintJOSN(list)
		}
		return list, nil
	} else {
		return nil, errors.New("failed to get ufw rule list")
	}
}

func (inst *System) UWFOpenPort(port int) (*Message, error) {
	return inst.ufwPort(port, false)
}

func (inst *System) UWFClosePort(port int) (*Message, error) {
	return inst.ufwPort(port, true)
}

func (inst *System) ufwPort(port int, deny bool) (*Message, error) {
	if port == 0 {
		return nil, errors.New("port must not be 0")
	}
	cmdStr := fmt.Sprintf("sudo ufw allow %d", port)
	if deny {
		if port == 22 {
			log.Error("ufw: port 22 must be kept open")
			return nil, errors.New("ufw: failed to close port 22, port 22 must be kept open")
		}
		cmdStr = fmt.Sprintf("sudo ufw deny %d", port)
	}
	cmd := exec.Command("sh", "-c", cmdStr)
	output, err := cmd.Output()
	res := cleanCommand(string(output), cmd, err, debug)
	return &Message{
		Message: res,
	}, err
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

func (inst *System) UWFDisable() (*Message, error) {
	return inst.ufwEnable(true)
}

func (inst *System) ufwEnable(disable bool) (*Message, error) {
	cmdString := "sudo ufw enable"
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

type UFWStatus struct {
	Port   int    `json:"port"`
	Status string `json:"status"`
}

func ufwStats(output string) []UFWStatus {
	var ufwStatus []UFWStatus
	var portStatus UFWStatus
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line != "" {
			if strings.Contains(strings.ToLower(line), "reject") == true {
				continue
			}
			for cc := 20; cc > 0; cc-- {
				replace := ""
				for ttt := 0; ttt < cc; ttt++ {
					replace += " "
				}
				line = strings.Replace(line, replace, " ", -1)
			}
			list := strings.Split(line, " ")
			var port int
			var status string
			for i, s := range list {
				if i == 0 { // get port
					intVar, err := strconv.Atoi(s)
					if err == nil {
						port = intVar
					}
				}
				if i == 2 { // get status
					if strings.Contains(s, "DENY") {
						status = "closed"
					}
					if strings.Contains(s, "ALLOW") {
						status = "open"
					}
				}
			}
			if port != 0 && status != "" {
				portStatus.Port = port
				portStatus.Status = status
				ufwStatus = append(ufwStatus, portStatus)
			}
		}
	}
	return ufwStatus

}
