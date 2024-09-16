package validatorgo

import "testing"

func TestIsDecimal(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsDecimalOpts
		want   bool
	}{
		// Basic valid decimal numbers without forcing a decimal point
		{name: "Valid integer without forcing decimal", param1: "123", param2: IsDecimalOpts{ForceDecimal: false} , want: true},
		{name: "Valid decimal number without forcing decimal", param1: "123.45", param2: IsDecimalOpts{ForceDecimal: false}, want: true},

		// Force decimal point presence
		{name: "Valid decimal number with forced decimal", param1: "123.45", param2: IsDecimalOpts{ForceDecimal: true}, want: true},
		{name: "Integer but decimal point required", param1: "123", param2: IsDecimalOpts{ForceDecimal: true}, want: false},

		// Testing with different decimal digit ranges
		{name: "Decimal number with exact digits", param1: "123.45", param2: IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Decimal number with fewer digits", param1: "123.4", param2: IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Decimal number with more digits", param1: "123.456", param2: IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Decimal number with valid range of digits", param1: "123.456", param2: IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(5)}}, want: true},

		// Testing locale-specific decimal formats
		{name: "Valid decimal with French locale", param1: "123,45", param2: IsDecimalOpts{ForceDecimal: true, Locale: "fr-FR"}, want: true},
		{name: "Invalid decimal for French locale", param1: "123.45", param2: IsDecimalOpts{ForceDecimal: true, Locale: "fr-FR"}, want: false},
		{name: "Valid decimal with German locale", param1: "123,45", param2: IsDecimalOpts{ForceDecimal: true, Locale: "de-DE"}, want: true},

		// Invalid cases
		{name: "Non-numeric string", param1: "abc", param2: IsDecimalOpts{ForceDecimal: false}, want: false},
		{name: "Invalid decimal number", param1: "123..45", param2: IsDecimalOpts{ForceDecimal: true}, want: false},
		{name: "Negative decimal number", param1: "-123.45", param2: IsDecimalOpts{ForceDecimal: true}, want: true},
		{name: "Positive number with locale", param1: "123.45", param2: IsDecimalOpts{Locale: "en-US"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsDecimal(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
