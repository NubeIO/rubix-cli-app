package remote

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/validation"
	"gthub.com/NubeIO/rubix-cli-app/service/remote/command"
	"strings"
)

type EdgeNetworking struct {
	IPAddress  string `json:"ip_address" post:"true"`
	SubnetMask string `json:"subnet_mask" post:"true"`
	Gateway    string `json:"gateway" post:"true"`
	SetDHCP    bool   `json:"set_dhcp" post:"true"`
	Password   string `json:"password"`
}

func (inst *Admin) EdgeSetIP(net *EdgeNetworking) (ok bool, err error) {
	if net == nil {
		return false, errors.New("no values where valid")
	}
	arch, err := inst.DetectModel()
	if arch.IsBeagleBone {
		return false, errors.New("error incorrect arch type")
	}
	iface, err := inst.edge28Iface()
	if err != nil || iface == "" {
		return false, errors.New("error on get network interface name")
	}
	cmd := ""
	if !net.SetDHCP {
		_, err = validation.IsIPAddr(net.IPAddress)
		if err != nil {
			return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an IPAddress", net.IPAddress))
		}
		_, err = validation.IsIPAddr(net.SubnetMask)
		if err != nil {
			return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an SubnetMask", net.SubnetMask))
		}
		_, err = validation.IsIPAddr(net.Gateway)
		if err != nil {
			return false, errors.New(fmt.Sprintf(" %s couldn't be parsed as an Gateway", net.Gateway))
		}
		cmd = fmt.Sprintf("echo N00B2828 | sudo connmanctl config %s --ipv4 manual %s %s %s", iface, net.IPAddress, net.SubnetMask, net.Gateway)
	} else {
		cmd = fmt.Sprintf("sudo connmanctl config %s --ipv4 dhcp", iface)
	}
	inst.CMD.Commands = command.Builder(cmd)
	res := inst.CMD.RunCommand()
	if res.Err == nil {
		ok = true
	}
	return

}

func (inst *Admin) edge28Iface() (interfaceName string, err error) {

	inst.CMD.Commands = command.Builder("connmanctl services")
	res := inst.CMD.RunCommand()
	if err != nil {
		return "", errors.New("failed to get interface")
	} else {
		if strings.Contains(res.Out, "*AO") {
			interfaceName = strings.ReplaceAll(res.Out, "*AO Wired", "")
			return str.StandardizeSpaces(interfaceName), nil
		}
		if strings.Contains(res.Out, "*AR") {
			interfaceName = strings.ReplaceAll(res.Out, "*AR Wired", "")
			return str.StandardizeSpaces(interfaceName), nil
		} else {
			return "", errors.New("failed to parse interface")
		}

	}
}
