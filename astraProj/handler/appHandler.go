package handler

import (
	"astraProj/service"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveData(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	err := service.SaveData(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data stored successfully"})
}
