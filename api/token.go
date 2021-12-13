package api

import (
	. "Task3/route"
	"github.com/gin-gonic/gin"
)

func init() {
	Version[V1].Group("/token").GET(tokenGet)
}

func tokenGet(context *gin.Context) {

}
