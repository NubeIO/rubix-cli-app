package httpresp

import "net/http"

const (
	Success        = http.StatusOK
	InvalidParams  = http.StatusBadRequest
	Error          = http.StatusInternalServerError
	StatusNotFound = http.StatusNotFound
)

type ErrorCode struct {
	code int
}

func (e *ErrorCode) Error() string {
	return GetMsg(e.code)
}

func (e *ErrorCode) Code() int {
	return e.code
}

func NewError(code int) error {
	return &ErrorCode{code}
}
