package helpers

import (
	"TicketService/models"
	"testing"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var avalidate = validator.New()

func TestCreateTicket(t *testing.T) {

	cases := []struct {
		name          string
		args          models.Ticket
		expectedError bool
	}{
		{
			name: "Create Ticket",
			args: models.Ticket{
				Train_id:       primitive.NewObjectID(),
				Capacity:       10,
				Cost:           100,
				Departure_time: "2020-01-01T00:00:00Z",
				Arrival_time:   "2020-01-01T00:00:00Z",
			},
			expectedError: false,
		},
		{
			name: "Create Ticket with empty Departure time",
			args: models.Ticket{
				Train_id:       primitive.NewObjectID(),
				Capacity:       10,
				Cost:           100,
				Departure_time: "",
				Arrival_time:   "2020-01-01T00:00:00Z",
			},
			expectedError: true,
		},
		{
			name: "Create Ticket with empty Arrival time",
			args: models.Ticket{
				Train_id:       primitive.NewObjectID(),
				Capacity:       10,
				Cost:           100,
				Departure_time: "2020-01-01T00:00:00Z",
				Arrival_time:   "",
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
