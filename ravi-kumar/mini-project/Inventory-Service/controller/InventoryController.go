package controller

import (
	repository "Inventory-Service/Repository"
	errors "Inventory-Service/errors"
	mockdata "Inventory-Service/model"
	service "Inventory-Service/service"
	"encoding/json"
	"strconv"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var inventoryService service.IInventoryService

func init() {
	inventoryService = service.InitInventoryService(&repository.MongoDAO{})
}

func CreateProduct(c *gin.Context) {
	//access: Seller
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))

	if mockdata.Role(acessorUserRole) != mockdata.Seller {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var newProduct mockdata.Product
	json.NewDecoder(c.Request.Body).Decode(&newProduct)

	result := inventoryService.CreateProduct(newProduct)
	c.JSON(http.StatusOK, result)
}

func GetCatalog(c *gin.Context) {
	allProducts := inventoryService.GetCatalog()
	c.JSON(http.StatusOK, allProducts)
}

func UpdateProductQuantity(c *gin.Context) {
	//access: Admin and Seller
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))
	if !(mockdata.Role(acessorUserRole) == mockdata.Seller || mockdata.Role(acessorUserRole) == mockdata.Admin) {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var productId string = c.Param("productId")
	updateCount, err := strconv.Atoi(c.Param("updateCount"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "update count should be a valid integer")
		return
	}

	quantityAfterUpdation, error := inventoryService.UpdateProductQuantity(productId, updateCount)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			c.JSON(productError.Status, productError.ErrorMessage)
			return
		} else {
			fmt.Println("productError casting error in UpdateProductQuantity")
			return
		}
	}
	c.JSON(http.StatusOK, quantityAfterUpdation)
}

func GetProductById(c *gin.Context) {
	var productId string = c.Param("productId")
	productRetrieved, error := inventoryService.GetProductById(productId)

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
	//access: Seller
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))

	if mockdata.Role(acessorUserRole) != mockdata.Seller {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var updatedProduct mockdata.Product
	unmarshalErr := json.NewDecoder(c.Request.Body).Decode(&updatedProduct)
	if unmarshalErr != nil {
		c.JSON(errors.UnmarshallError().Status, errors.UnmarshallError().ErrorMessage)
		return
	}

	var productId string = c.Param("productId")
	productRetrieved, error := inventoryService.UpdateProductById(productId, updatedProduct)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			c.JSON(productError.Status, productError.ErrorMessage)
			return
		} else {
			fmt.Println("productError casting error in UpdateProductById")
			return
		}
	}
	c.JSON(http.StatusOK, productRetrieved)
}

func DeleteProductbyId(c *gin.Context) {
	//access: Seller
	acessorUserRole, _ := strconv.Atoi(c.Param("acessorUserRole"))

	if mockdata.Role(acessorUserRole) != mockdata.Seller {
		c.JSON(http.StatusUnauthorized, errors.AccessDenied())
		return
	}

	var productId string = c.Param("productId")
	successMessage, error := inventoryService.DeleteProductbyId(productId)

	if error != nil {
		productError, ok := error.(*errors.ProductError)
		if ok {
			c.JSON(productError.Status, productError.ErrorMessage)
			return
		} else {
			fmt.Println("productError casting error in DeleteProductbyId")
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
