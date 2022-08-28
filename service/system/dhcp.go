package system

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
)

func (inst *System) DHCPPortExists(body NetworkingBody) (bool, error) {
	return inst.dhcp.Exists(body.PortName)
}

func (inst *System) DHCPSetAsAuto(body NetworkingBody) (*Message, error) {
	ok, err := inst.dhcp.SetAsAuto(body.PortName)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("was not able :%s to auto", body.PortName)
	if ok {
		msg = fmt.Sprintf("was able to set interface :%s to auto", body.PortName)
	} else {
		return nil, errors.New(fmt.Sprintf("was not able :%s to auto", body.PortName))
	}
	return &Message{
		Message: msg,
	}, nil
}

func (inst *System) DHCPSetStaticIP(body *dhcpd.SetStaticIP) (string, error) {
	return inst.dhcp.SetStaticIP(body)
}
