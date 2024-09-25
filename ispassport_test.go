package validatorgo

import "testing"

func TestIsPassportNumber(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Valid passport number (US)
		{name: "Valid US passport number", param1: "123456789", param2: "US", want: true},
		// Invalid passport number (US)
		{name: "Invalid US passport number (letters not allowed)", param1: "A12345678", param2: "US", want: false},
		{name: "Invalid US passport number (too short)", param1: "12345678", param2: "US", want: false},
		{name: "Invalid US passport number (too long)", param1: "1234567890", param2: "US", want: false},

		// Valid passport number (GB)
		{name: "Valid GB passport number", param1: "123456789", param2: "GB", want: true},
		// Invalid passport number (GB)
		{name: "Invalid GB passport number (too short)", param1: "12345678", param2: "GB", want: false},

		// Valid passport number (DE)
		{name: "Valid DE passport number", param1: "C12345678", param2: "DE", want: true},
		// Invalid passport number (DE)
		{name: "Invalid DE passport number (too short)", param1: "C26X0A7", param2: "DE", want: false},

		// Valid passport number (FR)
		{name: "Valid FR passport number", param1: "12AB34567", param2: "FR", want: true},
		// Invalid passport number (FR)
		{name: "Invalid FR passport number (invalid characters)", param1: "12AB@4567", param2: "FR", want: false},

		// Valid passport number (IN)
		{name: "Valid IN passport number", param1: "A1234567", param2: "IN", want: true},
		// Invalid passport number (IN)
		{name: "Invalid IN passport number (too long)", param1: "A12345678", param2: "IN", want: false},
		{name: "Invalid IN passport number (missing letter)", param1: "123456789", param2: "IN", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPassportNumber(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
