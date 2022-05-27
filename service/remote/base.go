package remote

import (
	"gthub.com/NubeIO/rubix-cli-app/service/remote/command"
)

type Response struct {
	Ok  bool
	Out string
	Err error
}

type Admin struct {
	CMD *command.Command
}

func New(admin *Admin) *Admin {
	opts := &command.Command{}
	admin.CMD = opts
	return admin
}

func (inst *Admin) Uptime() (res *command.Response) {
	cmd := "uptime"
	inst.CMD.Commands = command.Builder(cmd)
	res = inst.CMD.RunCommand()
	return
}
