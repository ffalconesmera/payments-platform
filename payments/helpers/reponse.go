package helpers

import (
	"github.com/gin-gonic/gin"
)

type jsonResponse struct{}

var jsonRes *jsonResponse

func JsonResponse() *jsonResponse {
	if jsonRes == nil {
		jsonRes = &jsonResponse{}
	}

	return jsonRes
}

func (r *jsonResponse) JsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "ok", "data": data})
}

func (r *jsonResponse) JsonFail(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{"status": "fail", "message": message})
}
