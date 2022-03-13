package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/configs"
	"github.com/dhi13man/healthcare-app/models"
	"github.com/gin-gonic/gin"
)

// Get all Medicines
func FindMedicines(c *gin.Context) []models.Medicine {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := configs.MedicinesCollection.Find(ctx, models.Medicine{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error inserting user", "error": err.Error()})
		return []models.Medicine{}
	}

	var books []models.Medicine
	cursor.Decode(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
	return books
}
