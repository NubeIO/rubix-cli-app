package response

var MsgFlags = map[int]string{
	Success:        "Ok",
	InvalidParams:  "Invalid params error",
	Error:          "Fail",
	StatusNotFound: "not found",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
