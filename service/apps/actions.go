package apps

import "errors"

type ActionType int

//go:generate stringer -type=ActionType
const (
	start ActionType = iota
	stop
	status
	enable
	disable
)

func CheckAction(s string) error {
	switch s {
	case start.String():
		return nil
	case stop.String():
		return nil
	case status.String():
		return nil
	case enable.String():
		return nil
	case disable.String():
		return nil
	}
	return errors.New("invalid action type, try start")

}
