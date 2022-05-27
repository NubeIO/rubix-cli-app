package remote

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/str"
	"gthub.com/NubeIO/rubix-cli-app/service/remote/command"
	"strings"
)

type Node struct {
	IsInstalled      bool   `json:"is_installed"`
	InstalledVersion string `json:"installed_version"`
}

func (inst *Admin) NodeGetVersion() (res *command.Response, node *Node) {
	cmd := "nodejs -v"
	inst.CMD.Commands = command.Builder(cmd)
	res = inst.CMD.RunCommand()
	if res.Err != nil {
		return
	}
	cmdOut := res.Out
	if strings.Contains(cmdOut, "v") {
		node.InstalledVersion = str.RemoveNewLine(cmdOut)
		node.IsInstalled = true
		return res, node
	} else {
		node.InstalledVersion = cmdOut
		node.IsInstalled = false
		return res, node
	}
}

type NodeJSInstall struct {
	AlreadyInstalled bool   `json:"already_installed"`
	InstalledOk      bool   `json:"installed_ok"`
	TextOut          string `json:"text_out"`
}
