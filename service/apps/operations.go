package apps

func (inst *Apps) Action(app, action string, timeout int) (*Response, error) {
	actionResp := &Response{}
	err := CheckAction(action)
	if err != nil {
		return nil, err
	}
	switch action {
	case start.String():
		actionResp = inst.Start(app, timeout)
	case stop.String():
	case status.String():
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
		action := mass.Action
		actionResp, err := inst.Action(app, action, mass.Timeout)
		if err != nil {
			return nil, err
		}
		res := massResponse{
			AppName: app,
			Action:  action,
			Ok:      actionResp.Ok,
			Message: actionResp.Message,
			Err:     actionResp.Err,
		}
		response = append(response, res)
	}
	return response, nil
}
