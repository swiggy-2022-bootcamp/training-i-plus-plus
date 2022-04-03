package controllers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/dhi13man/healthcare-app/bookkeeping_service/models"
	"github.com/dhi13man/healthcare-app/bookkeeping_service/repositories"
	"github.com/gin-gonic/gin"
)

// Return all Medicines in database from repository.
func FindAllMedicines(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	medicines, err := repositories.GetMedicines(models.Medicine{}, ctx)

	if err != nil {
		const errMsg string = "Error fetching all medicines."
		log.Error(errMsg, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errMsg, "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, medicines)
	}
}

// Get a Medicine by its name.
func FindMedicineByName(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	name := c.Param("id")
	if name == "" {
		const errMsg string = "Medicine name is required."
		log.Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	medicine, err := repositories.GetMedicine(models.Medicine{Name: name}, ctx)
	if err != nil {
		const errMsg string = "Error finding Medicine by name."
		log.Error(errMsg, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errMsg, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medicine)
}

// Find all medicines for a Disease identified by its name.
 func FindMedicinesByDiseaseName(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	name := c.Param("diseaseName")
	if name == "" {
		const errMsg string = "Disease name is required."
		log.Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	disease, err := repositories.GetDisease(models.Disease{Name: name}, ctx)
	if err != nil {
		const errMsg string = "Error finding Disease by name."
		log.Error(errMsg, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errMsg, "error": err.Error()})
		return
	}

	medicines := make([]models.Medicine, len(disease.Medicines))
	for i, v := range disease.Medicines {
		medicines[i], err = repositories.GetMedicine(models.Medicine{Name: v}, ctx)
		if err != nil {
			var errMsg string = "Error finding Medicine by name: " + v
			log.Error(errMsg, err)
		}
	}
	c.JSON(http.StatusOK, medicines)
 }
