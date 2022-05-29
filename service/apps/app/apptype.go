package app

import "errors"

type Type int

//go:generate stringer -type=Type
const (
	Python Type = iota
	Go
	Node
	NoApp
)

func CheckType(s string) (appType Type, appName string, err error) {
	if s == "" {
		return NoApp, "", errors.New("invalid app type selection was EMPTY, Python, Go ")
	}
	switch s {
	case Python.String():
		return Python, Python.String(), nil
	case Go.String():
		return Go, Go.String(), nil
	case Node.String():
		return Node, Node.String(), nil
	}
	return NoApp, "", errors.New("invalid app type selection was EMPTY, Python, Go ")
}
