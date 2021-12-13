package api

import (
	. "Task3/route"
	"github.com/gin-gonic/gin"
)

func init() {
	group := Version[V1].Group("/todo")
	group.GET(todoGet)
}

func todoGet(context *gin.Context) {

}

func todoPost(context *gin.Context) {

}

func todoDelete(context *gin.Context) {

}
