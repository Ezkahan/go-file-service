package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func JSONSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}
