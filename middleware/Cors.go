package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() func(context *gin.Context) {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == http.MethodOptions {
			context.AbortWithStatus(http.StatusNoContent)
		}
		fmt.Println(2)
		context.Next()
	}
}
