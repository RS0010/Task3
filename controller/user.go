package controller

import (
	"Task3/database"
	. "Task3/route"
	"Task3/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	group := Version[V1].Group("/user")
	group.POST(userPost)
	//group.OPTIONS(userOptions)
}

func userOptions(context *gin.Context) {

}

func userPost(context *gin.Context) {
	var user schemas.User
	var reply schemas.Reply

	if err := context.ShouldBindJSON(&user); err != nil {
		reply.Status = http.StatusBadRequest
		reply.Error = err.Error()
		context.JSON(reply.Status, reply)
		return
	}

	u := database.UserAdd(user)
	reply.Status = http.StatusOK
	reply.Data = schemas.UserGetting{
		Userid:   u.Id,
		Username: u.Name,
	}
	context.JSON(reply.Status, reply)
}
