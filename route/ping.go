package route

import (
	"Task3/Schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	version[v1].addApi(api{
		path: "/ping",
		GET:  pongGet,
	})
}

func pongGet(context *gin.Context) {
	reply := Schemas.Reply{
		Status: http.StatusOK,
		Data: struct {
			A int `json:"a"`
		}{A: 2},
	}
	context.JSON(200, reply)
}
