package apps

func (inst *Apps) SystemCtlAction(action string, timeout int) (interface{}, error) {

	return nil, nil
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
