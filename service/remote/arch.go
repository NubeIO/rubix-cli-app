package remote

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"gthub.com/NubeIO/rubix-cli-app/service/remote/command"
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
	Err          error
}

//DetectArch can detect hardware type is in ARM or AMD and also if hardware is for example a Raspberry PI
func (inst *Admin) DetectArch() (res *command.Response, arch *Arch) {
	arch = &Arch{}
	cmd := "tr '\\0' '\\n' </proc/device-tree/model;arch &&  dpkg --print-architecture"
	inst.CMD.Commands = command.Builder(cmd)
	res = inst.CMD.RunCommand()
	cmdOut := res.Out
	err := res.Err
	if err != nil || strings.Contains(cmdOut, " No such file or directory") {
		cmd = "dpkg --print-architecture"
		inst.CMD.Commands = command.Builder(cmd)
		res = inst.CMD.RunCommand()
		arch.ArchModel = res.Out
		if err != nil {
			return res, arch
		}
	}
	cmdOut = str.RemoveNewLine(cmdOut)
	if strings.Contains(cmdOut, "Raspberry Pi") {
		arch.IsRaspberry = true
		arch.IsArm = true
		return res, arch
	} else if strings.Contains(cmdOut, "BeagleBone Black") {
		arch.ArchModel = "BeagleBone Black"
		arch.IsBeagleBone = true
		arch.IsArm = true
		return res, arch
	} else if strings.Contains(cmdOut, "amd64") {
		arch.ArchModel = cmdOut
		arch.IsAMD64 = true
		return res, arch
	} else if strings.Contains(cmdOut, "amd32") {
		arch.ArchModel = cmdOut
		arch.IsAMD32 = true
		return res, arch
	} else if strings.Contains(cmdOut, "armhf") {
		arch.ArchModel = cmdOut
		arch.IsARMf = true
		arch.IsArm = true
		return res, arch
	} else if strings.Contains(cmdOut, "armv7l") {
		arch.ArchModel = cmdOut
		arch.IsArmv7l = true
		arch.IsArm = true
		return res, arch
	}
	return res, arch
}
