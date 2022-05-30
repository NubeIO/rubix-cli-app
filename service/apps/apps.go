package apps

import (
	"context"
	"errors"
	"github.com/NubeIO/git/pkg/git"
)

var err error

type Apps struct {
	Token   string `json:"token"`   // git token
	Version string `json:"version"` //version to install
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

func New(inst *Apps, rubixApp string) (*Apps, error) {
	if inst == nil {
		return nil, errors.New("type apps must not be nil")
	}
	if rubixApp == "" {
		return nil, errors.New("no app was passed in, try ff, flow or flow-framework")
	}
	if inst.Perm == 0 {
		inst.Perm = 0700
	}
	if inst.Token == "" {
		return nil, errors.New("git token can not be empty")
	}

	opts := &git.AssetOptions{
		Owner: inst.App.Owner,
		Repo:  inst.App.Repo,
		Tag:   inst.Version,
		Arch:  inst.App.Arch,
	}
	ctx := context.Background()
	gitClient = git.NewClient(inst.Token, opts, ctx)
	return inst, err
}
