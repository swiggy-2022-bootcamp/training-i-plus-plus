package utils_test

import (
	"paymentService/utils"
	"testing"
)

func TestISODateToTime(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"2022-03-28T00:00:00.000Z", "2022-03-28 00:00:00 +0000 UTC"},
		{"2022-03-29T00:00:00.000Z", "2022-03-29 00:00:00 +0000 UTC"},
		{"2022-03-30T00:00:00.000Z", "2022-03-30 00:00:00 +0000 UTC"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output, err := utils.ISODateToTime(test.input)
			if err != nil {
				t.Errorf("ISODateToTime(%s) returned error: %s", test.input, err)
			}
			if output.String() != test.output {
				t.Errorf("ISODateToTime(%s) returned %s, want %s", test.input, output, test.output)
			}
		})
	}

}
