package middlewares
import "github.com/gin-gonic/gin"
import (
	"fmt"
	"time"
)
func Logger() gin.HandlerFunc{
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string{
		return fmt.Sprintf("%s - [%s] %s %s %s %d  %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}