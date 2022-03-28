package controller

import (
	errors "Inventory-Service/errors"
	service "Inventory-Service/service"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	result := service.CreateProduct(&c.Request.Body)
	c.JSON(http.StatusOK, result)
}

func GetCatalog(c *gin.Context) {
	allProducts := service.GetCatalog()
	c.JSON(http.StatusOK, allProducts)
}

func GetProductById(c *gin.Context) {
	var productId string = c.Param("productId")
	productRetrieved, error := service.GetProductById(productId)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			c.JSON(productError.Status, productError.ErrorMessage)
			return
		} else {
			fmt.Println("productError casting error in GetProductById")
			return
		}
	}
	c.JSON(http.StatusOK, productRetrieved)
}

func UpdateProductById(c *gin.Context) {
	var productId string = c.Param("productId")
	productRetrieved, error := service.UpdateProductById(productId, &c.Request.Body)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			c.JSON(productError.Status, productError.ErrorMessage)
			return
		} else {
			fmt.Println("productError casting error in GetProductById")
			return
		}
	}
	c.JSON(http.StatusOK, productRetrieved)
}

func DeleteProductbyId(c *gin.Context) {
	var productId string = c.Param("productId")
	successMessage, error := service.DeleteProductbyId(productId)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			c.JSON(productError.Status, productError.ErrorMessage)
			return
		} else {
			fmt.Println("productError casting error in GetProductById")
			return
		}
	}
	c.JSON(http.StatusOK, *successMessage)
}

//decomissioned code base

// func GetCatalog(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Add("Content-type", "application/json")
// 	res.WriteHeader(http.StatusOK)
// 	json.NewEncoder(res).Encode(mockdata.GetProductCatalog())
// }

// func GetProductById(res http.ResponseWriter, req *http.Request) {
// 	//storing route variables for a request
// 	variables := mux.Vars(req)
// 	productId := variables["productId"]

// 	for _, productDetail := range mockdata.GetProductCatalog() {
// 		productIdInt, _ := strconv.Atoi(productId)
// 		if productDetail.Id == productIdInt {
// 			res.Header().Add("Content-type", "application/json")
// 			res.WriteHeader(http.StatusOK)
// 			json.NewEncoder(res).Encode(productDetail)
// 			return
// 		}
// 	}

// 	res.Header().Add("Content-type", "application/json")
// 	res.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(res).Encode("product with given id not found")
// }

// func DeleteProductbyId(res http.ResponseWriter, req *http.Request) {
// 	variables := mux.Vars(req)
// 	productId := variables["productId"]

// 	for index, productDetail := range mockdata.GetProductCatalog() {
// 		productIdInt, _ := strconv.Atoi(productId)
// 		if productDetail.Id == productIdInt {
// 			//to mockdata.Catalog[:index], append everything beyond index (excluding index ofcourse).
// 			//since append() takes "elements" to append as 2nd param, use "..." to lay out elements of slice as independent elements
// 			mockdata.Catalog = append(mockdata.Catalog[:index], mockdata.Catalog[index+1:]...)

// 			res.Header().Add("Content-type", "application/json")
// 			res.WriteHeader(http.StatusOK)
// 			json.NewEncoder(res).Encode("product deleted")
// 			return
// 		}
// 	}

// 	res.Header().Add("Content-type", "application/json")
// 	res.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(res).Encode("product with given id not found")
// }

// func PrintCatalog() {
// 	catalog := mockdata.GetProductCatalog()
// 	for _, product := range catalog {
// 		PrintProduct(&product)
// 	}
// }

// func PrintProduct(product *mockdata.Product) {
// 	fmt.Println("\nname: ", product.Name,
// 		"\nprice: ", product.Price,
// 		"\ndescription: ", product.Description,
// 		"\nseller: ", product.Seller,
// 		"\nrating: ", product.Rating,
// 		"\nreview: ", strings.Join(product.Review, ", "))
// }
