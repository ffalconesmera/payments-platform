package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContentTypeJsonMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.ContentType()
		if contentType != "application/json" {
			c.Abort()
			c.JSON(http.StatusBadRequest, "Content type is not json")
			return
		}

		c.Next()
	}
}
