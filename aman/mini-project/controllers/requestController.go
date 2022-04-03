package controllers

import (
	"aman-swiggy-mini-project/kafka"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		consumedRequests := kafka.Consume(ctx)
		c.JSON(http.StatusOK, gin.H{"Requests": consumedRequests})
	}
}

func PostRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		var bMap map[string]string

		body := c.Request.Body
		b, err := io.ReadAll(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		json.Unmarshal([]byte(b), &bMap)
		requestString := bMap["request"]
		fmt.Println(requestString)
		kafka.Produce(ctx, requestString)
		c.JSON(http.StatusOK, gin.H{"Request Submitted": requestString})
	}
}
