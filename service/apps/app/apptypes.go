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
	None
)

func (inst *Service) checkAppName(s string) (Name, error) {
	if s == "" {
		return None, errors.New("invalid app type selection was EMPTY, try FlowFramework, flow, flow-framework, or ff")
	}
	switch s {
	case FlowFramework.String(), "flow", flow, "ff":
		return FlowFramework, nil
	case RubixWires.String(), "wires", wires:
		return RubixWires, nil
	}
	return None, errors.New("invalid app type, try FlowFramework, flow, flow-framework, or ff")
}
