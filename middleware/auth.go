package middleware

import "github.com/gin-gonic/gin"

func Authenticator(context *gin.Context) {
	context.PostForm("api_token")
	param := struct {
		Token string `form:"token"`
	}{}
	err := context.ShouldBindQuery(&param)
	if err != nil {
		return
	}
}
