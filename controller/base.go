package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
)

type Controller struct {
	DB *dbase.DB
	WS *melody.Melody //web socket
}

type WsMsg struct {
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
	IsError bool        `json:"is_error"`
}

var err error

func bodyAsJSON(ctx *gin.Context) (interface{}, error) {
	var body interface{} //get the body and put it into an interface
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func useHostNameOrID(ctx *gin.Context) (idName string, useID bool) {
	hostID := resolveHeaderHostID(ctx)
	hostName := resolveHeaderHostName(ctx)
	if hostID != "" {
		return hostID, true
	} else if hostName != "" {
		return hostName, false
	} else {
		return "", false
	}
}

func resolveHeaderHostID(ctx *gin.Context) string {
	return ctx.GetHeader("host_uuid")
}

func resolveHeaderHostName(ctx *gin.Context) string {
	return ctx.GetHeader("host_name")
}

func resolveHeaderGitToken(ctx *gin.Context) string {
	return ctx.GetHeader("git_token")
}

func reposeWithCode(code int, body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(code, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(code, body)
			} else {
				ctx.JSON(code, Message{Message: err.Error()})
			}

		}
	} else {
		ctx.JSON(code, body)
	}
}

type Response struct {
	StatusCode   int         `json:"status_code"`
	ErrorMessage string      `json:"error_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(404, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(404, body)
			} else {
				ctx.JSON(404, Message{Message: err.Error()})
			}
		}
	} else {
		ctx.JSON(200, body)
	}
}

type Message struct {
	Message string `json:"message"`
}

func reposeMessage(code int, body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(code, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(code, body)
			} else {
				ctx.JSON(code, Message{Message: err.Error()})
			}

		}
	} else {
		ctx.JSON(code, body)
	}
}
