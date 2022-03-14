package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func FormatLogger(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("%s - [%s] \"%s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
