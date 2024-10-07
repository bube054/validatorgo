package validatorgo

import "testing"

func TestIsVAT(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		{name: "Valid VAT Germany", param1: "DE123456789", param2: "DE", want: true},
		{name: "Invalid VAT Germany", param1: "DE12345678", param2: "DE", want: false},
		{name: "Valid VAT Italy", param1: "IT12345678901", param2: "IT", want: true},
		{name: "Invalid VAT Italy", param1: "IT1234567890", param2: "IT", want: false},
		{name: "Valid VAT Netherlands", param1: "NL123456789B01", param2: "NL", want: true},
		{name: "Invalid VAT Netherlands", param1: "NL123456789B", param2: "NL", want: false},
		{name: "Valid VAT Austria", param1: "U12345678", param2: "AT", want: true},
		{name: "Invalid VAT Austria", param1: "U1234567", param2: "AT", want: false},
		{name: "Valid VAT with empty country", param1: "DE123456789", param2: "", want: true},
		{name: "Invalid VAT with empty country", param1: "XYZ123456789", param2: "", want: false},

		{name: "Valid VAT with invalid country", param1: "DE123456789", param2: "XYX", want: false},
		{name: "Valid VAT with no country", param1: "DE123456789", param2: "", want: true},
		{name: "Valid VAT with any country", param1: "DE123456789", param2: "any", want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsVAT(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
