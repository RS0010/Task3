package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func ShowRequest() func(context *gin.Context) {
	return func(context *gin.Context) {
		fmt.Println(context.Request.Method, context.Request.URL, context.Request.RequestURI)
		for k, v := range context.Request.Header {
			fmt.Println(k, v)
		}
		all, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(all))
		context.Next()
	}
}
