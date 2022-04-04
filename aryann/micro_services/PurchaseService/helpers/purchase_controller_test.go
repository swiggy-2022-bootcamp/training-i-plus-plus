package helpers

import (
	"PurchaseService/models"
	"testing"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var avalidate = validator.New()

func TestCreatePurchased(t *testing.T) {

	cases := []struct {
		name          string
		args          models.Purchased
		expectedError bool
	}{
		{
			name: "Create Purchased",
			args: models.Purchased{
				Train_id:       primitive.NewObjectID(),
				Departure:      "Mumbai",
				Arrival:        "Bangalore",
				Departure_time: "2020-01-01T00:00:00Z",
				Arrival_time:   "2020-01-01T00:00:00Z",
				Cost:           100,
				User_id:        primitive.NewObjectID(),
			},
			expectedError: false,
		},
		{
			name: "Create Purchased with empty Arrival",
			args: models.Purchased{
				Train_id:       primitive.NewObjectID(),
				Departure:      "Kuala Lumpur",
				Arrival:        "",
				Departure_time: "2020-01-01T00:00:00Z",
				Arrival_time:   "2020-02-01T00:00:00Z",
				Cost:           10,
				User_id:        primitive.NewObjectID(),
			},
			expectedError: true,
		},
		{
			name: "Create Purchased with empty Arrival time",
			args: models.Purchased{
				Train_id:       primitive.NewObjectID(),
				Departure:      "Delhi",
				Arrival:        "Pune",
				Departure_time: "2020-01-01T00:00:00Z",
				Arrival_time:   "",
				Cost:           20,
				User_id:        primitive.NewObjectID(),
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
