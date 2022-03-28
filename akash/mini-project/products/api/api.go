package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products.akash.com/db"
	"products.akash.com/log"
	"products.akash.com/model"
)

func AddProduct(c *gin.Context) {

	productData := model.Product{}
	if err := c.BindJSON(&productData); err != nil {
		panic(err)
	}

	log.Info("Add request for product : ", productData)

	product := db.FindOneWithId(productData.Id)
	if product == nil {
		db.SaveProduct(productData)
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"product added successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"product already exist with this id"}`))
	}
}

func QueryOne(c *gin.Context) {

	id := c.Param("id")
	log.Info("Find product with id : ", id)

	product := db.FindOneWithId(id)
	if product != nil {
		c.JSON(http.StatusOK, product)
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"product not found"}`))
	}
}

func QueryAll(c *gin.Context) {
	products := db.FindAll()
	c.JSON(http.StatusOK, products)
}

func Delete(c *gin.Context) {

	id := c.Param("id")
	log.Info("Delete user with id : ", id)

	if db.DeleteProduct(id) == true {
		c.Data(http.StatusOK, "application/json", []byte(`{"message":"product delete successful"}`))
	} else {
		c.Data(http.StatusUnauthorized, "application/json", []byte(`{"message":"delete failed"}`))
	}

}
