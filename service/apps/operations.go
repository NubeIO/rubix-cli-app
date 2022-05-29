package apps

func (inst *Apps) SystemCtlAction(action string, timeout int) (*Response, error) {
	actionResp := &Response{}

	err := CheckAction(action)
	if err != nil {
		return nil, err
	}
	switch action {
	case start.String():
		actionResp = inst.Start(timeout)
	case stop.String():
		actionResp = inst.Stop(timeout)
	case status.String():
		actionResp = inst.Status(timeout)
	case enable.String():
		actionResp = inst.Enable(timeout)
	case disable.String():
		actionResp = inst.Disable(timeout)
	case isRunning.String():
		actionResp, _ = inst.IsRunning(timeout)
	case isInstalled.String():
		actionResp, _ = inst.IsInstalled(timeout)
	case isEnabled.String():
		actionResp, _ = inst.IsEnabled(timeout)
	case isActive.String():
		actionResp, _ = inst.IsActive(timeout)
	case isFailed.String():
		actionResp = inst.Status(timeout)
	}
	return actionResp, nil
}

type Mass struct {
	Apps    []string
	Action  string
	Timeout int `json:"timeout"`
}

type massResponse struct {
	AppName string `json:"app_name"`
	Action  string `json:"action"`
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func (inst *Apps) Mass(mass *Mass) ([]massResponse, error) {
	var response []massResponse
	//for _, app := range mass.Apps {
	//	appService := &AppService{}
	//	actionType := appService.Action
	//	actionResp, err := inst.Action(appService)
	//	if err != nil {
	//		return nil, err
	//	}
	//	res := massResponse{
	//		AppName: app,
	//		Action:  actionType,
	//		Ok:      actionResp.Ok,
	//		Message: actionResp.Message,
	//		Err:     actionResp.Err,
	//	}
	//	response = append(response, res)
	//}
	return response, nil
}
