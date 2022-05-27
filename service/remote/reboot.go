package remote

import (
	"gthub.com/NubeIO/rubix-cli-app/service/remote/command"
)

func (inst *Admin) HostReboot() (res *command.Response) {
	cmd := "sudo shutdown -r now"
	inst.CMD.Commands = command.Builder(cmd)
	res = inst.CMD.RunCommand()
	return
}
