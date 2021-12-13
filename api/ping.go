package api

import (
	"Task3/Schemas"
	"Task3/middleware"
	. "Task3/route"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	Version[V1].Group("/ping").Use(middleware.Authenticator).GET(pongGet)
}

func pongGet(context *gin.Context) {
	reply := Schemas.Reply{
		Status: http.StatusOK,
		Data:   "和",
	}
	context.JSON(200, reply)
}
