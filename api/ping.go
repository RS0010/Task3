package api

import (
	"Task3/Schemas"
	. "Task3/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Version[V1].Group("/ping").GET(pongGet)
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
