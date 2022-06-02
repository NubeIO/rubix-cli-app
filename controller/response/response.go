package response

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

// Response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ReposeHandler response request
func ReposeHandler(c *gin.Context, httpCode, responseCode int, data interface{}) {
	response := Response{
		Code:    responseCode,
		Message: GetMsg(responseCode),
		Data:    data,
	}
	c.JSON(httpCode, response)
	if mode := gin.Mode(); mode == gin.DebugMode {
		switch data.(type) {
		case error:
			log.Panicf("[error] %+v", data)
		}
	}
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
