package helpers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CustomLog is a singleton for loggin
type customLog struct{}

var clog *customLog

func CustomLog() *customLog {
	if clog == nil {
		clog = &customLog{}
	}

	return clog
}

// Generate and return and uuid for identify context
func (l *customLog) GetRequestContextId(k string) string {
	return fmt.Sprintf("%s_%s", k, CustomHash().NewUUIDString())
}

// Print by console a info log
func (l *customLog) PrintInfo(ctx context.Context, c *gin.Context, v any) {
	l.logInfo(ctx, c, v)
}

// Print by console a error log
func (l *customLog) PrintError(ctx context.Context, c *gin.Context, v any, shutdown bool) {
	l.logError(ctx, c, v)
	if shutdown {
		panic("shutdown..!")
	}
}

func (l *customLog) logError(ctx context.Context, c *gin.Context, v any) {
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

func (l *customLog) logInfo(ctx context.Context, c *gin.Context, v any) {
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
