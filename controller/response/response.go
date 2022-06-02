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
func ReposeHandler(ctx *gin.Context, httpCode, responseCode int, data interface{}) {
	response := Response{
		Code:    responseCode,
		Message: GetMsg(responseCode),
		Data:    data,
	}
	ctx.JSON(httpCode, response)
	if mode := gin.Mode(); mode == gin.DebugMode {
		switch data.(type) {
		case error:
			log.Panicf("[error] %+v", data)
		}
	}
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if body != nil {
			j, _ := json.Marshal(body)
			if string(j) == "null" {
				ctx.JSON(404, Message{Message: err.Error()})
				return
			}
			ctx.JSON(404, body)
		} else {
			ctx.JSON(404, Message{Message: err.Error()})
		}
	} else {
		ctx.JSON(200, body)
	}
}

type Message struct {
	Message string `json:"message"`
}
