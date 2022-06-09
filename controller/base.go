package controller

import (
	"encoding/json"
	dbase "github.com/NubeIO/edge/database"
	"github.com/NubeIO/edge/service/apps/installer"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type Controller struct {
	DB        *dbase.DB
	WS        *melody.Melody //web socket
	Installer *installer.Installer
}

var fileUtils = fileutils.New()

type WsMsg struct {
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
	IsError bool        `json:"is_error"`
}

var err error

func bodyAsJSON(c *gin.Context) (interface{}, error) {
	var body interface{} //get the body and put it into an interface
	err = c.ShouldBindJSON(&body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func resolveHeaderHostID(c *gin.Context) string {
	return c.GetHeader("host_uuid")
}

func resolveHeaderHostName(c *gin.Context) string {
	return c.GetHeader("host_name")
}

func resolveHeaderGitToken(c *gin.Context) string {
	return c.GetHeader("git_token")
}

func reposeWithCode(code int, body interface{}, err error, c *gin.Context) {
	if err != nil {
		if err == nil {
			c.JSON(code, Message{Message: "unknown error"})
		} else {
			if body != nil {
				c.JSON(code, body)
			} else {
				c.JSON(code, Message{Message: err.Error()})
			}

		}
	} else {
		c.JSON(code, body)
	}
}

type Response struct {
	StatusCode   int         `json:"status_code"`
	ErrorMessage string      `json:"error_message"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func reposeHandler(body interface{}, err error, c *gin.Context) {
	if err != nil {
		if body != nil {
			j, _ := json.Marshal(body)
			if string(j) == "null" {
				c.JSON(404, Message{Message: err.Error()})
				return
			}
			c.JSON(404, body)
		} else {
			c.JSON(404, Message{Message: err.Error()})
		}
	} else {
		c.JSON(200, body)
	}
}

type Message struct {
	Message string `json:"message"`
}
