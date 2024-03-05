package helpers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CustomLog is a singleton for loggin

// Generate and return and uuid for identify context
func GetRequestContextId(k string) string {
	return fmt.Sprintf("%s_%s", k, NewUUIDString())
}

// Print by console a info log
func PrintInfo(ctx context.Context, c *gin.Context, v any) {
	logInfo(ctx, c, v)
}

// Print by console a error log
func PrintError(ctx context.Context, c *gin.Context, v any, shutdown bool) {
	logError(ctx, c, v)
	if shutdown {
		panic("shutdown..!")
	}
}

func logError(ctx context.Context, c *gin.Context, v any) {
	// Processing request
	c.Next()

	log.WithFields(log.Fields{
		"PID":       ctx.Value("REQUEST_ID"),
		"METHOD":    c.Request.Method,
		"URI":       c.Request.RequestURI,
		"STATUS":    c.Writer.Status(),
		"CLIENT_IP": c.ClientIP(),
		"PAYLOAD":   c.Request.Body,
		"QUERY":     c.Request.URL.Query(),
	}).Error(v)
}

func logInfo(ctx context.Context, c *gin.Context, v any) {
	// Processing request
	c.Next()

	log.WithFields(log.Fields{
		"PID":       ctx.Value("REQUEST_ID"),
		"METHOD":    c.Request.Method,
		"URI":       c.Request.RequestURI,
		"STATUS":    c.Writer.Status(),
		"CLIENT_IP": c.ClientIP(),
		"PAYLOAD":   c.Request.Body,
		"QUERY":     c.Request.URL.Query(),
	}).Info(v)
}
