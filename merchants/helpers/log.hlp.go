package helpers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CustomLog is a singleton for logging

// Print by console a error log
func PrintError(c *gin.Context, v any, shutdown bool) {
	logError(c, v)
	if shutdown {
		panic("shutdown..!")
	}
}

func logError(c *gin.Context, v any) {
	// Processing request
	c.Next()

	log.WithFields(log.Fields{
		"PID":       c.Param("REQUEST_ID"),
		"METHOD":    c.Request.Method,
		"URI":       c.Request.RequestURI,
		"STATUS":    c.Writer.Status(),
		"CLIENT_IP": c.ClientIP(),
		"PAYLOAD":   c.Request.Body,
		"QUERY":     c.Request.URL.Query(),
	}).Error(v)
}
