package api

import (
	"github.com/gin-gonic/gin"
)

func init() {
	todoGroup := version[v1].addGroup(group{
		path:       "/todo",
		apis:       nil,
		middleware: nil,
	})

	version[v1].addApi(api{
		path:   "/todo",
		POST:   todoGet,
		GET:    todoGet,
		DELETE: nil,
	})
}

func todoGet(context *gin.Context) {

}

func todoPost(context *gin.Context) {

}

func todoDelete(context *gin.Context) {

}
