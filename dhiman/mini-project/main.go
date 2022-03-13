package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/", func(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  router.Run()
}
