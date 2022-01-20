package Authenticator

import (
	"Task3/database"
	"Task3/schemas"
	"Task3/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authnHeader struct {
	Token string `form:"token" binding:"required"`
}

func Authenticator(type_ int) func(context *gin.Context) {
	return func(context *gin.Context) {
		var header authnHeader
		var reply schemas.Reply

		if err := context.ShouldBindHeader(&header); err != nil {
			reply.Status = http.StatusBadRequest
			reply.Error = err.Error()
			context.JSON(reply.Status, reply)
			context.Abort()
			return
		}

		id, t, err := tools.JWTVerify(header.Token)
		if err != nil {
			reply.Status = http.StatusUnauthorized
			reply.Error = err.Error()
			context.JSON(reply.Status, reply)
			context.Abort()
			return
		} else if t != type_ {
			reply.Status = http.StatusUnauthorized
			reply.Error = "Incorrect token type"
			context.JSON(reply.Status, reply)
			context.Abort()
			return
		}

		user, _ := database.UserGet(id)
		context.Set("user", user)
		context.Next()
	}
}
