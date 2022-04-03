package middlewares

import (
	"bytes"
	"fmt"
	"paymentService/services"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware(topic string) gin.HandlerFunc {
	l := services.NewLoggerService(topic)
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		l.Log(fmt.Sprintf("%s %s %s %d %s (request) %+v (response) %+v", clientIP, method, path, statusCode, latency, c.Request, blw.body.String()))
	}
}
