package controller

import (
	"Task3/database"
	"Task3/middleware/Authenticator"
	. "Task3/route"
	"Task3/schemas"
	"Task3/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	Version[V1].Group("/token").GET(tokenGet)
	Version[V1].Group("/token/new").Use(Authenticator.Authenticator(tools.RefreshToken)).GET(tokenNew)
}

func tokenGet(context *gin.Context) {
	var token schemas.LoginToken
	var reply schemas.Reply

	if err := context.ShouldBindHeader(&token); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Error = err.Error()
		context.JSON(reply.Status, reply)
		return
	}

	user, ok := database.UserGet(token.Userid)
	if !ok {
		reply.Status = http.StatusUnauthorized
		reply.Error = "Unknown user id: " + strconv.Itoa(token.Userid)
		context.JSON(reply.Status, reply)
		return
	}

	if !tools.HashCompare(token.Passcode, user.Code) {
		reply.Status = http.StatusUnauthorized
		reply.Error = "Wrong passcode"
		context.JSON(reply.Status, reply)
		return
	}

	reply.Status = http.StatusOK
	reply.Data = schemas.TokenGetting{
		AuthToken:    tools.JWTGenerate(user.Id, "tokenGet", tools.AuthenticationToken),
		RefreshToken: tools.JWTGenerate(user.Id, "tokenGet", tools.RefreshToken),
	}
	context.JSON(reply.Status, reply)
}

func tokenNew(context *gin.Context) {
	user := context.MustGet("user").(database.User)

	reply := schemas.Reply{
		Status: http.StatusOK,
		Data: schemas.TokenGetting{
			AuthToken:    tools.JWTGenerate(user.Id, "tokenGet", tools.AuthenticationToken),
			RefreshToken: tools.JWTGenerate(user.Id, "tokenGet", tools.RefreshToken),
		},
	}
	context.JSON(200, reply)
}
