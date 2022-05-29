package remote

import (
	"errors"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"gthub.com/NubeIO/rubix-cli-app/service/remote/command"
	"runtime"
	"strings"
)

type Arch struct {
	ArchModel    string `json:"arch_model"`
	IsBeagleBone bool   `json:"is_beagle_bone,omitempty"`
	IsRaspberry  bool   `json:"is_raspberry,omitempty"`
	IsArm        bool   `json:"is_arm,omitempty"`
	IsAMD64      bool   `json:"is_amd64,omitempty"`
	IsAMD32      bool   `json:"is_amd32,omitempty"`
	IsARMf       bool   `json:"is_armf,omitempty"`
	IsArmv7l     bool   `json:"is_armv7l,omitempty"`
	IsLinux      bool   `json:"is_linux"`
	Err          error
}

func (inst *Admin) ArchIsLinux() bool {
	s := runtime.GOOS
	fmt.Println(s)
	switch s {
	case "linux":
		return true
	}
	return false
}

//DetectArch can detect hardware type is in ARM or AMD
func (inst *Admin) DetectArch() (arch *Arch, err error) {
	arch = &Arch{}
	inst.CMD.Commands = command.Builder("dpkg", "--print-architecture")
	res := inst.CMD.RunCommand()
	cmdOut := res.Out
	cmdOut = str.RemoveNewLine(cmdOut)
	if strings.Contains(cmdOut, "amd64") {
		arch.ArchModel = cmdOut
		arch.IsAMD64 = true
		return arch, nil
	} else if strings.Contains(cmdOut, "amd32") {
		arch.ArchModel = cmdOut
		arch.IsAMD32 = true
		return arch, nil
	} else if strings.Contains(cmdOut, "armhf") {
		arch.ArchModel = cmdOut
		arch.IsARMf = true
		arch.IsArm = true
		return arch, nil
	} else if strings.Contains(cmdOut, "armv7l") {
		arch.ArchModel = cmdOut
		arch.IsArmv7l = true
		arch.IsArm = true
		return arch, nil
	}
	return arch, errors.New("could not find correct arch type")
}

//DetectNubeProduct can detect hardware type is in ARM or AMD and also if hardware is for example a Raspberry PI
func (inst *Admin) DetectNubeProduct() (isRc, isEdge bool) {
	inst.CMD.Commands = command.Builder("cat", "/proc/device-tree/model")
	res := inst.CMD.RunCommand()
	cmdOut := res.Out
	cmdOut = str.RemoveNewLine(cmdOut)
	if strings.Contains(cmdOut, "Raspberry Pi") {
		isRc = true
		return
	} else if strings.Contains(cmdOut, "BeagleBone Black") {
		isEdge = true
		return
	}
	return
}

func (inst *Admin) CheckEdge28() error {
	_, isEdge := inst.DetectNubeProduct()
	if isEdge {
	} else {
		return errors.New("the host product is not type edge-28")
	}
	return nil

}
