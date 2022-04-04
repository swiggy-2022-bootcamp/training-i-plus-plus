package helpers

import (
	"TrainService/models"
	"testing"

	"github.com/go-playground/validator/v10"
)

var avalidate = validator.New()

func TestCreateTrain(t *testing.T) {

	cases := []struct {
		name          string
		args          models.Train
		expectedError bool
	}{
		{
			name: "Create train",
			args: models.Train{
				Source:      "Bangalore",
				Destination: "Chennai",
			},
			expectedError: false,
		},
		{
			name: "Create train with empty source",
			args: models.Train{
				Source:      "",
				Destination: "Chennai",
			},
			expectedError: true,
		},
		{
			name: "Create train with empty destination",
			args: models.Train{
				Source:      "Bangalore",
				Destination: "",
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
