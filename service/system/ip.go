package system

import (
	"errors"
	"fmt"
	"github.com/NubeIO/lib-command/command"
	"github.com/NubeIO/lib-command/unixcmd"
	"github.com/NubeIO/lib-dhcpd/dhcpd"
	address "github.com/NubeIO/lib-networking/ip"
	"github.com/NubeIO/lib-networking/networking"
)

type IP struct {
	Interface        string `json:"interface"`
	ConfirmInterface bool   `json:"confirm_interface"` //check if the interface exists
	IPAddress        string `json:"ip_address"`        //192.168.15.10
	Netmask          string `json:"netmask"`           //255.255.255.0
	Gateway          string `json:"gateway"`           //192.168.15.1
	DHCP             bool   `json:"-"`
}

func NewIP(newIp *IP) *IP {
	return newIp
}

var nets = networking.New()
var cmd = unixcmd.New(&command.Command{})
var isTypeEdge bool
var isTypeRc bool

func (ip *IP) SetDHCP() (ok bool, err error) {
	return ip.update()
}

func (ip *IP) SetStaticIP() (ok bool, err error) {
	return ip.update()
}

func (ip *IP) update() (ok bool, err error) {
	if ip == nil {
		return false, errors.New("ip struct is nil")
	}
	err = ip.checks()
	if err != nil {
		return false, err
	}
	res, err := cmd.DetectNubeProduct()
	if err != nil {
		return false, err
	}
	if res.IsEdge {
		_, err := ip.updateEdge()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if res.IsRC {
		_, err := ip.updateRC()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New("arch type is not a number product")
}

func (ip *IP) checks() (err error) {
	if !isTypeEdge {
		if ip.ConfirmInterface {
			_, err = nets.CheckInterfacesName(ip.Interface)
			if err != nil {
				return err
			}
		}
	}
	if !ip.DHCP {
		ips := address.New()
		err = ips.IsIPAddrErr(ip.IPAddress)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ip *IP) updateEdge() (bool, error) {
	nets := dhcpd.New()
	fmt.Println("SET IP not finished")
	fmt.Println(nets)
	return false, errors.New("code not finished")

}

func (ip *IP) updateRC() (bool, error) {
	nets := dhcpd.New()
	fmt.Println("SET IP not finished")
	fmt.Println(nets)
	return false, errors.New("code not finished")

}
