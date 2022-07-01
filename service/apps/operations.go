package apps

import "errors"

func (inst *Apps) SystemCtlAction(action string, timeout int) (*SystemResponse, error) {
	actionResp := &SystemResponse{}
	switch action {
	case start.String():
		return inst.Start(timeout)
	case stop.String():
		return inst.Stop(timeout)
	case enable.String():
		return inst.Enable(timeout)
	case disable.String():
		return inst.Disable(timeout)
	}
	return actionResp, errors.New("no valid action found try, start, stop, enable or disable")
}

func (inst *Apps) SystemCtlStatus(action string, timeout int) (*SystemResponseChecks, error) {
	actionResp := &SystemResponseChecks{}
	switch action {
	case isRunning.String():
		return inst.IsRunning(timeout)
	case isInstalled.String():
		return inst.IsInstalled(timeout)
	case isEnabled.String():
		return inst.IsEnabled(timeout)
	case isActive.String():
		return inst.IsActive(timeout)
	case isFailed.String():
		return inst.IsFailed(timeout)
	}
	return actionResp, errors.New("no valid action found try, isRunning, isInstalled, isEnabled, isActive or isFailed")
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
	// for _, app := range mass.Apps {
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
	//	httpresp = append(httpresp, res)
	// }
	return response, nil
}
