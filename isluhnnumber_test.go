package validatorgo

import "testing"

func TestIsLuhnNumber(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Luhn number
		{name: "Valid Visa card number", param1: "4532015112830366", want: true},
		{name: "Valid Discover card number", param1: "6011514433546201", want: true},
		{name: "Valid Visa card number", param1: "4485275742308327", want: true},
		{name: "Valid IMEI number", param1: "490154203237518", want: true},
		{name: "A valid Luhn number", param1: "79927398713", want: true},
		{name: "A valid Luhn number", param1: "1234567812345670", want: true},
		// Invalid Luhn number
		{name: "Invalid due to incorrect checksum", param1: "4532015112830367", want: false},
		{name: "Invalid Discover card number", param1: "6011514433546202", want: false},
		{name: "Invalid due to incorrect checksum", param1: "490154203237519", want: false},
		{name: "Invalid Luhn number", param1: "79927398714", want: false},
		{name: "Invalid due to incorrect checksum", param1: "1234567812345671", want: false},
		{name: "Invalid due to incorrect non number included", param1: "12345abc12345671", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsLuhnNumber(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
