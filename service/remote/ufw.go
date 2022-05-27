package remote

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gthub.com/NubeIO/rubix-cli-app/service/remote/command"
	"strings"
)

type UFW struct {
	PortsCurrentState map[string]map[string]bool
	//Host              ssh.Host
}

type UWFInstall struct {
	AlreadyInstalled bool   `json:"already_installed"`
	InstalledOk      bool   `json:"installed_ok"`
	TextOut          string `json:"text_out"`
}

func logErr(msg string, err error) error {
	errMsg := fmt.Sprintf("%s", msg)
	if err != nil {
		errMsg = fmt.Sprintf("%s err: %f", msg, err)
	}
	log.Errorf(errMsg)
	return errors.New(errMsg)
}

//UWFStatus check status and also if ufw is installed
func (inst *Admin) UWFStatus() (res *command.Response, isActive bool) {
	cmd := "sudo ufw status"
	inst.CMD.Commands = command.Builder(cmd)
	res = inst.CMD.RunCommand()
	if res.Err != nil {
		log.Error()
		res.Err = logErr("ufw: FirewallPort Error: ", res.Err)
		return
	}
	if strings.Contains(res.Out, "active") {
		isActive = true
		return
	} else if strings.Contains(res.Out, "inactive") {
		isActive = false
		return
	} else {
		res.Err = logErr("ufw: failed to get any valid response to check firewall", nil)
		return
	}
}

//func (inst *Admin) UWFInstall() (UWFInstall UWFInstall, err error) {
//	log.Info("ufw: run install command")
//	cmd := "sudo apt install ufw -y"
//	inst.Host.CommandOpts.CMD = cmd
//	out, ok, err := inst.Host.RunCommand()
//	if err != nil {
//		log.Error("ufw: UWFInstall Error: ", err)
//		UWFInstall.TextOut = out
//		UWFInstall.InstalledOk = ok
//		return UWFInstall, err
//	}
//	log.Info("ufw: run install command")
//	_, installed, err := inst.UWFStatus()
//	if err != nil {
//		log.Error("ufw: UWFInstall Error: ", err)
//		UWFInstall.TextOut = ""
//		UWFInstall.InstalledOk = installed
//		return UWFInstall, err
//	}
//	UWFInstall.TextOut = ""
//	UWFInstall.InstalledOk = installed
//	return UWFInstall, err
//}
//
//func (ufw *UFW) UWFReset() (ok bool, err error) {
//	cmd := "echo \"y\" | sudo ufw reset"
//	ufw.Host.CommandOpts.CMD = cmd
//	_, ok, err = ufw.Host.RunCommand()
//	if err != nil {
//		log.Error("ufw: Enable Error: ", err)
//		return ok, err
//	}
//	return ok, err
//}
//
//func (ufw *UFW) UWFEnable() (ok bool, err error) {
//	//first make sure port 22 is open, this is to make sure we don't lock ourselves out
//	cmd := "sudo ufw allow 22"
//	ufw.Host.CommandOpts.CMD = cmd
//	_, ok, err = ufw.Host.RunCommand()
//	if err != nil {
//		log.Error("ufw: failed to open port 22: ", err)
//		return ok, errors.New("ufw: failed to open port 22")
//	}
//	cmd = "echo \"yes\" | sudo ufw enable"
//	ufw.Host.CommandOpts.CMD = cmd
//	_, ok, err = ufw.Host.RunCommand()
//	if err != nil {
//		log.Error("ufw: Enable Error: ", err)
//		return ok, err
//	}
//	if !ok {
//		log.Error("ufw: run-command returned false: ", err)
//		return ok, err
//	}
//	return ok, err
//}
//
//func (ufw *UFW) UWFDisable() (ok bool, err error) {
//	cmd := "echo \"yes\" | sudo ufw disable"
//	if ufw.Host.CommandOpts.Sudo {
//		cmd = "sudo ufw disable"
//	}
//	ufw.Host.CommandOpts.CMD = cmd
//	_, ok, err = ufw.Host.RunCommand()
//	if err != nil {
//		log.Error("ufw: Disable Error: ", err)
//		return ok, err
//	}
//	return ok, err
//}
//
////UWFPort allow or deny a port, default is allow
//func (ufw *UFW) UWFPort(port int, deny bool) (ok bool, err error) {
//	cmd := fmt.Sprintf("sudo ufw allow %d", port)
//	if deny {
//		if port == 22 {
//			log.Error("ufw: port 22 must be kept open")
//			return ok, errors.New("ufw: failed to open port 22, port 22 must be kept open")
//		}
//		cmd = fmt.Sprintf("sudo ufw deny %d", port)
//	}
//	ufw.Host.CommandOpts.CMD = cmd
//	_, ok, err = ufw.Host.RunCommand()
//	if err != nil {
//		log.Error("ufw: UWFPort Error: ", err)
//		return ok, err
//	}
//	return ok, err
//}
//
////UWFDefaultPorts nube-io default ports
//func (ufw *UFW) UWFDefaultPorts() (ok bool, err error) {
//	cmd := fmt.Sprintf("sudo ufw allow 22 && sudo ufw allow 1313 && sudo ufw allow 1414 && sudo ufw allow 1616 && sudo ufw allow 1615")
//	ufw.Host.CommandOpts.CMD = cmd
//	_, ok, err = ufw.Host.RunCommand()
//	if err != nil {
//		log.Error("ufw: UWFPort Error: ", err)
//		return ok, err
//	}
//	return ok, err
//}
//
//
//}
//
//func (ufw *UFW) UFWLoadProfile(asSudo bool) (*UFW, error) {
//	ufw.PortsCurrentState = map[string]map[string]bool{}
//	cmd := "ufw status | grep ALLOW"
//	if asSudo {
//		cmd = "sudo ufw status | grep ALLOW"
//	}
//	ufw.Host.CommandOpts.CMD = cmd
//	output, _, err := ufw.Host.RunCommand()
//	if err != nil {
//		return nil, err
//	}
//	lines := strings.Split(output, "\n")
//	for _, line := range lines {
//		if line != "" {
//			if strings.Contains(strings.ToLower(line), "reject") == true {
//				continue
//			}
//			for cc := 20; cc > 0; cc-- {
//				replace := ""
//				for ttt := 0; ttt < cc; ttt++ {
//					replace += " "
//				}
//				line = strings.Replace(line, replace, " ", -1)
//			}
//			tokens := strings.Split(line, " ")
//			address := tokens[2]
//			tokens1 := strings.Split(tokens[0], "/")
//			protocol := tokens1[0]
//			port := tokens1[0]
//			if address != "" && protocol != "" && port != "" {
//				_, ok := ufw.PortsCurrentState[address]
//				if ok == false {
//					ufw.PortsCurrentState[address] = map[string]bool{}
//				}
//				ufw.PortsCurrentState[address][protocol+":"+port] = true
//			}
//		}
//	}
//	return ufw, nil
//}
