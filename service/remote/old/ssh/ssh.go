package ssh

//type Host struct {
//	Host   *model.Host
//	CMD    *command.Opts
//	IsSudo bool
//}
//
////RunCommand will run a local or remote command, if CommandOpts.Sudo is true then a sudo is added to the existing command (cmd = "sudo " + CommandOpts.CMD)
//func (h *Host) RunCommand() (res *command.Response) {
//	var err error
//	cmd := h.CMD
//	res = &command.Response{}
//	if nils.BoolIsNil(h.Host.IsLocalhost) {
//		res = command.Run(cmd)
//		cmdOut := res.Out
//		err = res.Err
//		if err != nil {
//			res.Err = err
//			return res
//		}
//		res.Ok = true
//		res.Out = cmdOut
//		return res
//	}
//	return
//}
