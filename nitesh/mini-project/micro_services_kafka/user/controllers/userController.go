package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func BookTicket() gin.HandlerFunc {
	return func(g *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		var search struct {
			TrainID string
			Date    string
		}

		if err := g.BindJSON(&search); err != nil {
			g.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		g.JSON(http.StatusOK, gin.H{"msg": "done"})
	}
}
