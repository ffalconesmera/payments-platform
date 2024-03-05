package middleware

import (
	"log"
	"net/http"

	"github.com/ffalconesmera/payments-platform/payments/helpers"
	"github.com/gin-gonic/gin"
)

func JWTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			log.Println("Missing authorization header")
			c.Abort()
			c.JSON(http.StatusBadRequest, "Missing authorization header")
			return
		}

		tokenString = tokenString[len("Bearer "):]

		ok, err := helpers.CustomHash().CheckJWToken(tokenString)
		if !ok {
			c.Abort()
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.Next()
	}
}
