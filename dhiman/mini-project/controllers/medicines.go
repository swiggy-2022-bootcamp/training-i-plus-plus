package controllers

import (
	"net/http"

	"github.com/dhi13man/healthcare-app/models"
	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
  func FindMedicines(c *gin.Context) {
	var books []models.Medicine
	models.DB.Find(&books)
  
	c.JSON(http.StatusOK, gin.H{"data": books})
  }