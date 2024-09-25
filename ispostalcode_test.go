package validatorgo

import "testing"

func TestIsPostalCode(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Valid postal codes with specific locales
		{name: "Valid Postal Code - US", param1: "90210", param2: "US", want: true},
		{name: "Valid Postal Code - Canada", param1: "K1A 0B1", param2: "CA", want: true},
		{name: "Valid Postal Code - UK", param1: "SW1A 1AA", param2: "GB", want: true},
		{name: "Valid Postal Code - Germany", param1: "10115", param2: "DE", want: true},
		{name: "Valid Postal Code - Japan", param1: "100-0001", param2: "JP", want: true},

		// Valid postal codes with "any" locale
		{name: "Valid Postal Code - US (Any Locale)", param1: "90210", param2: "any", want: true},
		{name: "Valid Postal Code - Canada (Any Locale)", param1: "K1A 0B1", param2: "any", want: true},
		{name: "Valid Postal Code - UK (Any Locale)", param1: "SW1A 1AA", param2: "any", want: true},

		// Valid postal codes with no locale
		{name: "Valid Postal Code - Germany (No Locale)", param1: "10115", param2: "", want: true},
		{name: "Valid Postal Code - France (No Locale)", param1: "75001", param2: "", want: true},

		// Invalid postal codes
		{name: "Invalid Postal Code - CN Format in Germany", param1: "902101", param2: "DE", want: false},
		{name: "Invalid Postal Code - Too Short", param1: "12", param2: "any", want: false},
		{name: "Invalid Postal Code - Letters in Number-only Locale", param1: "ABCDE", param2: "DE", want: false},
		{name: "Invalid Postal Code - Wrong Format for Locale", param1: "12345-6789", param2: "GB", want: false},

		// Invalid postal codes with no locale
		{name: "Invalid Postal Code - Gibberish (No Locale)", param1: "XYZ123", param2: "", want: false},
		{name: "Invalid Postal Code - Empty String (No Locale)", param1: "", param2: "", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPostalCode(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
