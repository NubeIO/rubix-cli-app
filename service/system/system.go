package system

import (
	"errors"
)

type Apps struct {
	Token       string `json:"token"`   // git token
	Version     string `json:"version"` // version to install
	Perm        int
	ServiceName string
}

var err error

func New(inst *Apps) (*Apps, error) {
	if inst == nil {
		return nil, errors.New("type apps must not be nil")
	}
	return inst, err
}

const Permission = 0700
