package main

import (
	"Task3/route"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	err := router.SetTrustedProxies([]string{})
	if err != nil {
		panic(err)
	}
	route.Init(router)
	err = router.Run(":80")
	if err != nil {
		panic(err)
	}
}
