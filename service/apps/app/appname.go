package app

import (
	"errors"
)

type Name int

//go:generate stringer -type=Name
const (
	FlowFramework Name = iota
	RubixWires
	PiGpio
	Broker
	LoraService
	GoApp
	None
)

func CheckAppName(s string) (checkName Name, appName string, err error) {
	if s == "" {
		return None, "", errors.New("invalid app type selection was EMPTY, try FlowFramework, flow, flow-framework, or ff")
	}
	switch s {
	case FlowFramework.String(), "flow", flow, "ff":
		return FlowFramework, flow, nil
	case RubixWires.String(), "wires", wires:
		return RubixWires, wires, nil
	}
	return None, "", errors.New("invalid app type, try FlowFramework, flow, flow-framework, or ff")
}
