package app

//CheckEdge28 if user try's to install an app on the incorrect host the reject the installation
func (inst *Service) CheckEdge28() (err error) {
	run := inst.initArchCheck()
	return run.CheckEdge28()
}
