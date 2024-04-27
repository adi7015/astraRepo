package config

import (
	"astraProj/handler"

	"github.com/gin-gonic/gin"
)

func InitAppRoute(c *gin.RouterGroup) {

	groupRoute := c.Group("/save")
	groupRoute.POST("/", handler.SaveData)
}
