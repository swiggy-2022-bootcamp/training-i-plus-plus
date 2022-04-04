package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"

	"github.com/dhi13man/healthcare-app/bookkeeping_service/models"
	"github.com/dhi13man/healthcare-app/bookkeeping_service/models/dtos"
	"github.com/dhi13man/healthcare-app/bookkeeping_service/repositories"
	"github.com/gin-gonic/gin"
)

var validate = validator.New()

// @Summary      Create a Medicine
// @Description  Create a Medicine in the Database using the data sent by them (REGISTER)
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        medicineDTO  body      models.Medicine  true  "Medicine details to be created"
// @Success      200          {object}  interface{}
// @Failure      400          {object}  dtos.HTTPError
// @Failure      404          {object}  dtos.HTTPError
// @Failure      500          {object}  dtos.HTTPError
// @Router       /users/medicines [post]
func CreateMedicine(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var requestBody models.Medicine
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Error("Binding Request Body Failed: ", err)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err))
	} else if validationErr := validate.Struct(&requestBody); validationErr != nil { // Validation of required fields
		log.Error("Validating Request Body Failed: ", validationErr)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, validationErr))
	} else {
		newMedicine := models.NewMedicine(requestBody.Name, requestBody.Diseases)
		out, err := repositories.CreateMedicine(newMedicine, ctx)
		if err != nil {
			const errMsg string = "Error inserting user."
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		} else {
			c.JSON(http.StatusCreated, out)
		}
	}
}

// @Summary      Return all Medicines in database from repository.
// @Description  Fetches all Medicines in database from repository and return an unfiltered JSON array of them.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Medicine
// @Failure      400          {object}  dtos.HTTPError
// @Failure      404          {object}  dtos.HTTPError
// @Failure      500          {object}  dtos.HTTPError
// @Router       /bookkeeping/medicines/ [get]
func FindAllMedicines(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	medicines, err := repositories.GetMedicines(models.Medicine{}, ctx)

	if err != nil {
		const errMsg string = "Error fetching all medicines."
		log.Error(errMsg, err)
		c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
	} else {
		c.JSON(http.StatusOK, medicines)
	}
}

// @Summary      Get a Medicine by its name.
// @Description  Get a Medicine from the database by its name.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID of the Medicine (currently its name)"
// @Success      200          {object}  models.Medicine
// @Failure      400  {object}  dtos.HTTPError
// @Failure      404  {object}  dtos.HTTPError
// @Failure      500  {object}  dtos.HTTPError
// @Router       /bookkeeping/medicines/{id} [get]
func FindMedicineByName(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	name := c.Param("id")
	if name == "" {
		const errMsg string = "Medicine name is required."
		log.Error(errMsg)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, nil, errMsg))
	} else {
		medicine, err := repositories.GetMedicine(models.Medicine{Name: name}, ctx)
		if err != nil {
			const errMsg string = "Error finding Medicine by name."
			log.Error(errMsg, err)
			c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
			return
		}
		c.JSON(http.StatusOK, medicine)
	}
}

// @Summary      Find all medicines for a Disease identified by its name.
// @Description  Find all medicines for a Disease identified by the disease name from the two databases.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        diseaseName  path      string  true  "ID of the Disease (currently its name)"
// @Success      200  {object}  models.Medicine
// @Failure      400  {object}  dtos.HTTPError
// @Failure      404  {object}  dtos.HTTPError
// @Failure      500  {object}  dtos.HTTPError
// @Router       /bookkeeping/medicines/disease/{diseaseName} [get]
 func FindMedicinesByDiseaseName(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	name := c.Param("diseaseName")
	if name == "" {
		const errMsg string = "Disease name is required."
		log.Error(errMsg)
		c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, nil, errMsg))
		return
	}

	disease, err := repositories.GetDisease(models.Disease{Name: name}, ctx)
	if err != nil {
		const errMsg string = "Error finding Disease by name."
		log.Error(errMsg, err)
		c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
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

// @Summary      Updates Medicines in the Database.
// @Description  Updates the Medicine in the Database using their email.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        medicineDTO  body      models.Medicine  true  "Medicine details to be created"
// @Success      200          {object}  models.Medicine
// @Failure      400          {object}  dtos.HTTPError
// @Failure      404          {object}  dtos.HTTPError
// @Failure      500          {object}  dtos.HTTPError
// @Router       /users/medicines/{name} [put]
func UpdateMedicines(c *gin.Context) {
	 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	 defer cancel()
 
	 var requestBody models.Medicine
	 err := c.BindJSON(&requestBody)
	 if err != nil {
		 log.Error("Binding Request Body Failed: ", err)
		 c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, err))
	 } else if validationErr := validate.Struct(&requestBody); validationErr != nil { // Validation of required fields
		 log.Error("Validating Request Body Failed: ", validationErr)
		 c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, validationErr))
	 } else {
		 medicine := models.NewMedicine(requestBody.Name, requestBody.Diseases)
		 out, err := repositories.UpdateMedicine(medicine, ctx)
		 if err != nil {
			 const errMsg string = "Error updating user."
			 log.Error(errMsg, err)
			 c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
		 } else {
			 c.JSON(http.StatusOK, out)
		 }
	 }
 }

// @Summary      Deletes Medicines in the Database.
// @Description  Deletes the Medicine in the Database using their email.
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        diseaseName  path      string  true  "ID of the Disease (currently its name)"
// @Success      200          {object}  models.Medicine
// @Failure      400          {object}  dtos.HTTPError
// @Failure      404          {object}  dtos.HTTPError
// @Failure      500          {object}  dtos.HTTPError
// @Router       /users/medicines/{name} [delete]
func DeleteMedicines(c *gin.Context) {
	 ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	 defer cancel()
 
	 name := c.Param("name")
	 if name == "" {
		 const errMsg string = "Medicine name is required."
		 log.Error(errMsg)
		 c.JSON(http.StatusBadRequest, dtos.NewError(http.StatusBadRequest, nil, errMsg))
		 return
	 }
	  
	 out, err := repositories.DeleteMedicine(models.Medicine{Name: name}, ctx)
	 if err != nil {
		 const errMsg string = "Error deleting user."
		 log.Error(errMsg, err)
		 c.JSON(http.StatusInternalServerError, dtos.NewError(http.StatusInternalServerError, err, errMsg))
	 } else {
		 c.JSON(http.StatusOK, out)
	 }
 }
