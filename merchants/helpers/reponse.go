package helpers

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse interface {
	JsonSuccess(data interface{})
	JsonFail(c *gin.Context, message string)
}

type jsonResponse struct{}

func NewJsonResponse() *jsonResponse {
	return &jsonResponse{}
}

func (r *jsonResponse) JsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "ok", "data": data})
}

func (r *jsonResponse) JsonFail(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{"status": "fail", "message": message})
}
