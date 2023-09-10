package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
	})
}

func ResponseData(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": statusCode,
		"data":   data,
	})
}
