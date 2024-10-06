package validatorgo

import "testing"

func TestIsIBAN(t *testing.T) {
	tests := []struct {
		name   string
		param1 string // IBAN string
		param2 string // Country code
		want   bool   // Expected result
	}{
		// Valid IBANs
		{name: "Valid IBAN for Germany", param1: "DE89370400440532013000", param2: "DE", want: true},
		{name: "Valid IBAN for France", param1: "FR7630006000011234567890189", param2: "FR", want: true},
		{name: "Valid IBAN for Spain", param1: "ES9121000418450200051332", param2: "ES", want: true},
		{name: "Valid IBAN for United Kingdom", param1: "GB33BUKB20201555555555", param2: "GB", want: true},
		{name: "Valid IBAN for Italy", param1: "IT60X0542811101000000123456", param2: "IT", want: true},
		{name: "Valid IBAN for Switzerland", param1: "CH9300762011623852957", param2: "CH", want: true},

		// Invalid IBANs
		{name: "Invalid IBAN for Germany (too short)", param1: "DE8937040044053201300", param2: "DE", want: false},
		{name: "Invalid IBAN for France (wrong format)", param1: "FR7630006000A11234567890189", param2: "FR", want: false},
		{name: "Invalid IBAN for Spain (invalid character)", param1: "ES91210004184502000513A32", param2: "ES", want: false},
		{name: "Invalid IBAN for United Kingdom (wrong country)", param1: "GB33BUKB20201555555555", param2: "DE", want: false},
		{name: "Invalid IBAN for Italy (incorrect length)", param1: "IT60X05428111010000001234", param2: "IT", want: false},
		{name: "Invalid IBAN for unknown country code", param1: "CH9300762011623852957", param2: "XX", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIBAN(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
