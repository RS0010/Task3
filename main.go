package main

import (
	_ "Task3/controller"
	"Task3/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	c := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "userid", "usercode", "token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	})
	router.Use(gin.Logger(), gin.Recovery(), c)
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
