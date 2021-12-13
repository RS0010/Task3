package middleware

import (
	"Task3/Schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authnHeader struct {
	Token string `form:"token" binding:"required"`
}

func Authenticator(context *gin.Context) {
	var header authnHeader
	var reply Schemas.Reply

	if err := context.ShouldBindHeader(&header); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Data = err.Error()
		context.JSON(reply.Status, reply)
		context.Abort()
	}
	context.Next()
}
