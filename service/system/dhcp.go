package system

import (
	"fmt"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	"strconv"
)

func (inst *System) Exists(iFaceName string) (*Message, error) {
	exists, err := inst.dhcp.Exists(iFaceName)
	if err != nil {
		return nil, err
	}
	return &Message{
		Message: fmt.Sprintf("%s", strconv.FormatBool(exists)),
	}, nil
}

func (inst *System) SetAsAuto(iFaceName string) (*Message, error) {
	exists, err := inst.dhcp.SetAsAuto(iFaceName)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf("was not able :%s to auto", iFaceName)
	if exists {
		msg = fmt.Sprintf("was able to set interface :%s to auto", iFaceName)
	}
	return &Message{
		Message: msg,
	}, nil
}

//SetStaticIP Set a static IP for the specified network interface
func (inst *System) SetStaticIP(body *dhcpd.SetStaticIP) (string, error) {
	return inst.dhcp.SetStaticIP(body)
}
