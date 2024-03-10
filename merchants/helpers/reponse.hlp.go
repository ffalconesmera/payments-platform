package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JsonResponse is a singleton for structure http responses
func jsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "ok", "data": data})
}

func jsonFail(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": message})
}

func ResponseJson(c *gin.Context, data interface{}, err error) {
	if err != nil {
		PrintError(c, fmt.Sprintf("something happened: %s", err.Error()), false)
		jsonFail(c, fmt.Sprintf("something happened: %s", err.Error()))
		return
	}

	jsonSuccess(c, data)
}
