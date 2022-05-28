package app

import "errors"

type Name int

//go:generate stringer -type=Name
const (
	flowFramework Name = iota
	rubixWires
	piGpio
	broker
	loraService
)

func CheckAppName(s string) error {
	switch s {
	case flowFramework.String():
		return nil
	case rubixWires.String():
		return nil
	case piGpio.String():
		return nil
	case broker.String():
		return nil
	case loraService.String():
		return nil
	}
	return errors.New("invalid app type, try rubixWires")

}
