package main

import (
	config "astraProj/appConfig"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("starting the application")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	routerGroup := router.Group("/api")
	config.InitAppRoute(routerGroup)

	router.Run("localhost:8120")
}
