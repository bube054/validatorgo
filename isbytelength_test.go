package validatorgo

import (
	"testing"
)

func TestIsByteLength(t *testing.T) {
	var (
		max1 uint = 8
		max2 uint = 12
	)

	tests := []struct {
		name   string
		param1 string
		param2 *IsByteLengthOpts
		want   bool
	}{
		// Valid vyte length
		{name: "Is greater or equals to min", param1: "We♥Go", param2: &IsByteLengthOpts{Min: 5}, want: true},
		{name: "Is greater or equals to min and less than or equals max", param1: "We♥Go", param2: &IsByteLengthOpts{Min: 5, Max: &max1}, want: true},
		{name: "Valid string with exact min length", param1: "Hello", param2: &IsByteLengthOpts{Min: 5}, want: true},
		{name: "Nil config, valid string comparison", param1: "We♥Go", param2: nil, want: true},

		// Invalid vyte length
		{name: "Is not greater or equals to min", param1: "We♥Go", param2: &IsByteLengthOpts{Min: 8}, want: false},
		{name: "Is not greater or equals to min and less than or equals max", param1: "We♥Go", param2: &IsByteLengthOpts{Min: 8, Max: &max2}, want: false},
		{name: "Invalid string, exceeds max length", param1: "We♥Golang", param2: &IsByteLengthOpts{Min: 1, Max: &max1}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsByteLength(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
