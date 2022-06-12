package apps

import (
	"errors"
	"github.com/NubeIO/git/pkg/git"
)

var err error

type Apps struct {
	Token   string `json:"token"`   // git token
	Version string `json:"version"` // version to install
	Perm    int
	App     *Store
}

const (
	Owner     = "NubeIO"
	User      = "root"
	ArchAmd64 = "amd64"
	ArchArm7  = "armv7"
)

var gitClient *git.Client

const Permission = 0700

func New(inst *Apps) (*Apps, error) {
	if inst == nil {
		return nil, errors.New("type apps must not be nil")
	}
	if inst.Perm == 0 {
		inst.Perm = Permission
	}
	return inst, err
}
