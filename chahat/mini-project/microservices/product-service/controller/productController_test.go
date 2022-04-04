package productcontroller


import (
	"bhatiachahat/product-service/model"
	"testing"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var avalidate = validator.New()

func TestCreateProduct(t *testing.T) {

	cases := []struct {
		name          string
		args          model.Product
		expectedError bool
	}{
		{
			name: "Create Product",
			args: models.Product {
				
				"title": "Solid Gold Petite Micropave",
				"category": "Jewelery",
				"price": 30000,
				"seller": "White Gold Plated Princess",
				
				"ratings": 3,
				"image_url": "https://fakestoreapi.com/img/61sbMiUnoGL._AC_UL640_QL65_ML3_.jpg"
			}},
			expectedError: false,
		},
		{
			name: "Create Ticket with empty title",
			args: models.Product {
				
				"title": "",
				"category": "Jewelery",
				"price": 30000,
				"seller": "White Gold Plated Princess",
				
				"ratings": 3,
				"image_url": "https://fakestoreapi.com/img/61sbMiUnoGL._AC_UL640_QL65_ML3_.jpg"
			},
			expectedError: true,
		},
		{
			name: "Create Ticket with empty category",
			args: models.Product {
				
				"title": "Solid Gold Petite Micropave",
				"category": "",
				"price": 30000,
				"seller": "White Gold Plated Princess",
				
				"ratings": 3,
				"image_url": "https://fakestoreapi.com/img/61sbMiUnoGL._AC_UL640_QL65_ML3_.jpg"
			},
			expectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			if validationErr := avalidate.Struct(&c.args); validationErr != nil && !c.expectedError {
				t.Errorf("%s: %s", c.name, validationErr.Error())
			}

		})
	}
}