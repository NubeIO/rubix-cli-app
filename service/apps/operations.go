package apps

import "errors"

type AppService struct {
	AppName string `json:"service"`
	Action  string `json:"action"`
	Timeout int    `json:"timeout"`
}

func (inst *Apps) Action(appService *AppService) (*Response, error) {
	actionResp := &Response{}
	if appService == nil {
		return nil, errors.New("action must not be nil")
	}
	app := appService.AppName
	actionType := appService.Action
	timeout := appService.Timeout
	err := CheckAction(actionType)
	if err != nil {
		return nil, err
	}
	switch actionType {
	case start.String():
		actionResp = inst.Start(app, timeout)
	case stop.String():
		actionResp = inst.Stop(app, timeout)
	case status.String():
		actionResp = inst.Status(app, timeout)
	case enable.String():
		actionResp = inst.Enable(app, timeout)
	case disable.String():
		actionResp = inst.Disable(app, timeout)
	case isRunning.String():
		actionResp = inst.IsRunning(app, timeout)
	case isInstalled.String():
		actionResp = inst.IsInstalled(app, timeout)
	case isEnabled.String():
		actionResp = inst.IsEnabled(app, timeout)
	case isActive.String():
		actionResp = inst.IsActive(app, timeout)
	case isFailed.String():
		actionResp = inst.Status(app, timeout)
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
	for _, app := range mass.Apps {
		appService := &AppService{}
		actionType := appService.Action
		actionResp, err := inst.Action(appService)
		if err != nil {
			return nil, err
		}
		res := massResponse{
			AppName: app,
			Action:  actionType,
			Ok:      actionResp.Ok,
			Message: actionResp.Message,
			Err:     actionResp.Err,
		}
		response = append(response, res)
	}
	return response, nil
}
