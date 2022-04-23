package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//UseLogger ..
func UseLogger(formatter func(params gin.LogFormatterParams) string) gin.HandlerFunc {
	loggerConfig := LoggerConfigFunc(formatter)
	return gin.LoggerWithConfig(loggerConfig)
}

//LoggerConfigFunc ..
func LoggerConfigFunc(formatter func(gin.LogFormatterParams) string) gin.LoggerConfig {
	return gin.LoggerConfig{
		Formatter: formatter,
	}
}

//DefaultLoggerFormatter ..
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
