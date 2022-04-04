package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"product/domain"
	"product/utils/errs"
	"product/utils/logger"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	service domain.ProductService
}

type ProductDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type ProductResponseDTO struct {
	Message string `json:"message"`
}

// @Schemes
// @Description Fetches product details by id
// @Tags products
// @Param        productId   path      string  true  "Id"
// @Produce json
// @Success 200 {object} domain.Product
// @Failure      403  {object} errs.AppError
// @Router /products/{productId} [get]
func (uh *ProductHandlers) GetProductById(c *gin.Context) {

	productId, ok := c.Params.Get("productId")

	if !ok {
		logger.Error("Product Id not present in request params")
		err := errs.NewValidationError("Product Id not present in request params")
		c.JSON(err.Code, err.AsMessage())

	} else {
		product, err := uh.service.FindById(productId)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(product)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

// @Schemes
// @Description Creates a product
// @Tags products
// @Produce json
// @Accept json
// @Param        product  body      ProductDTO  true  "Product Creation"
// @Success 201 {object} domain.Product
// @Router /products/ [post]
func (uh *ProductHandlers) Register(c *gin.Context) {

	var newProduct domain.Product
	err := c.Bind(&newProduct)
	fmt.Println(newProduct, err)

	if err != nil {
		logger.Error("Invalid request body")
		err := errs.NewValidationError("Invalid request body")
		c.JSON(err.Code, err.AsMessage())

	} else {
		regProduct, err := uh.service.Register(newProduct)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(regProduct)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusCreated, "application/json", data)
		}
	}
}

// @Schemes
// @Description Updates product by productId
// @Tags products
// @Param        productId   path      string  true  "Id"
// @Param        product details   body      ProductDTO true  "Product details"
// @Produce json
// @Success 200 {object} domain.Product
// @Failure      500  {object} errs.AppError
// @Router /products/{productId} [put]
func (uh *ProductHandlers) UpdateProduct(c *gin.Context) {

	productId, ok := c.Params.Get("productId")

	if !ok {
		logger.Error("Product Id not present in request params")
		err := errs.NewValidationError("Product Id not present in request params")
		c.JSON(err.Code, err.AsMessage())
	} else {
		var updatedProduct domain.Product
		err := c.Bind(&updatedProduct)

		if err != nil {
			logger.Error("Invalid request body")
			err := errs.NewValidationError("Invalid request body")
			c.JSON(err.Code, err.AsMessage())

		} else {
			product, err := uh.service.Update(productId, updatedProduct)
			if err != nil {
				c.JSON(err.Code, err.AsMessage())
			} else {

				data, err := json.Marshal(product)
				if err != nil {
					err1 := errs.NewUnexpectedError("Unexpected error")
					c.JSON(err1.Code, err1.AsMessage())
				}
				c.Data(http.StatusCreated, "application/json", data)
			}
		}
	}
}

// @Schemes
// @Description Deletes product by id
// @Tags products
// @Param        productId   path      string  true  "Product Id"
// @Produce json
// @Success 200 {object} ProductDTO
// @Failure      500  {object} errs.AppError
// @Router /products/{productId} [delete]
func (uh *ProductHandlers) DeleteProduct(c *gin.Context) {

	productId, ok := c.Params.Get("productId")

	if !ok {
		logger.Error("Product Id not present in request params")
		err := errs.NewValidationError("Product Id not present in request params")
		c.JSON(err.Code, err.AsMessage())

	} else {
		product, err := uh.service.DeleteById(productId)
		if err != nil {
			c.JSON(err.Code, err.AsMessage())
		} else {

			data, err := json.Marshal(product)
			if err != nil {
				err1 := errs.NewUnexpectedError("Unexpected error")
				c.JSON(err1.Code, err1.AsMessage())
			}
			c.Data(http.StatusOK, "application/json", data)
		}
	}
}

func (uh *ProductHandlers) HelloWorldHandler(c *gin.Context) {

	token := "Hello Product World!"
	data, _ := json.Marshal(token)
	c.Data(http.StatusOK, "application/json", data)
}
