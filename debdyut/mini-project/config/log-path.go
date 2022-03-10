package config

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupLogPath() {
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
