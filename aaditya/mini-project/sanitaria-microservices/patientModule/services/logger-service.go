package services

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

func UseLogger(formatter func(params gin.LogFormatterParams) string) gin.HandlerFunc {
	loggerConfig := NewLoggerConfig(formatter)
	return gin.LoggerWithConfig(loggerConfig)
}

func NewLoggerConfig(formatter func(gin.LogFormatterParams) string) gin.LoggerConfig {
	return gin.LoggerConfig{
		Formatter: formatter,
	}
}

func DefaultLoggerFormatter(param gin.LogFormatterParams) string {

	return fmt.Sprintf("%s - [%d] %s %s %s %s %s %s %s\n",
		param.Method,
		param.StatusCode,
		param.ClientIP,
		param.Request.Proto,
		param.Request.UserAgent(),
		param.Path,
		param.TimeStamp.Format(time.RFC3339),
		param.Latency,
		param.ErrorMessage,
	)
}
