package helpers

import (
	"github.com/gin-gonic/gin"
)

// JsonResponse is a singleton for structure http responses
func JsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "ok", "data": data})
}

func JsonFail(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{"status": "fail", "message": message})
}
