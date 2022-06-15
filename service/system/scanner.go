package system

import (
	"github.com/NubeIO/lib-networking/scanner"
)

type Scanner struct {
	Count int
	Iface string
	Ip    string
	Ports []string
}

func RunScanner(params *Scanner) (*scanner.Hosts, error) {
	iface := ""
	ip := ""
	count := 254
	ports := []string{"22", "1414", "1883", "1660", "502", "1313", "1616"}
	if params != nil {
		iface = params.Iface
		ip = params.Ip
		if params.Count > 0 {
			count = params.Count
		}
		if len(params.Ports) > 0 {
			ports = params.Ports
		}

	}
	scan := scanner.New()
	address, err := scan.ResoleAddress(ip, count, iface)
	if err != nil {
		return nil, err
	}
	host := scan.IPScanner(address, ports, true)
	return host, nil
}
