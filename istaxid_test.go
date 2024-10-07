package validatorgo

import "testing"

func TestIsTaxID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Empty values
		{name: "Empty params", param1: "", param2: "", want: false},
		{name: "Empty tax ID", param1: "", param2: "en-US", want: false},
		{name: "Empty locale", param1: "123-45-6789", param2: "", want: true},
		{name: "Valid ID Any locale", param1: "123-45-6789", param2: "any", want: true},
		{name: "Invalid ID Any locale", param1: "123-45-6789abc", param2: "any", want: false},

		// Valid tax IDs
		{name: "Valid US SSN", param1: "123-45-6789", param2: "en-US", want: true},
		{name: "Valid Canada SIN", param1: "123456789", param2: "en-CA", want: true},
		{name: "Valid UK UTR", param1: "1234567890", param2: "en-GB", want: true},
		{name: "Valid Spain NIF", param1: "12345678Z", param2: "es-ES", want: true},
		{name: "Valid Germany TIN", param1: "12345678901", param2: "de-DE", want: true},
		{name: "Valid Finland personal identity code", param1: "131052-308T", param2: "fi-FI", want: true},
		{name: "Valid Portugal NIF", param1: "123456789", param2: "pt-PT", want: true},
		{name: "Valid Brazil CPF", param1: "12345678909", param2: "pt-BR", want: true},
		{name: "Valid Italy Codice Fiscale", param1: "RSSMRA85M01H501Z", param2: "it-IT", want: true},

		// Invalid tax IDs
		{name: "Invalid Canada SIN format", param1: "123-456-789", param2: "en-CA", want: false},
		{name: "Invalid UK UTR format", param1: "12345", param2: "en-GB", want: false},
		{name: "Invalid Spain NIF format", param1: "12345678", param2: "es-ES", want: false},
		{name: "Invalid Germany TIN too long", param1: "123456789012", param2: "de-DE", want: false},
		{name: "Invalid Finland personal identity code", param1: "131052308Z", param2: "fi-FI", want: false},
		{name: "Invalid Portugal NIF characters", param1: "12345678A", param2: "pt-PT", want: false},
		{name: "Invalid Italy Codice Fiscale format", param1: "RSSMRA85M01H501", param2: "it-IT", want: false},

		// Edge cases
		{name: "Invalid locale", param1: "123456789", param2: "invalid-locale", want: false},
		{name: "Short tax ID", param1: "12", param2: "en-US", want: false},
		{name: "Long tax ID", param1: "123456789012345", param2: "en-US", want: false},
		{name: "Non-numeric tax ID", param1: "abc123xyz", param2: "en-US", want: false},
		{name: "Special characters tax ID", param1: "123-45!6789", param2: "en-US", want: false},
		{name: "Valid Estonia", param1: "51001091072", param2: "et-EE", want: true},                    // Estonia valid
		{name: "Invalid Estonia too short", param1: "12345", param2: "et-EE", want: false},             // Estonia too short
		{name: "Valid Poland PESEL", param1: "4405140135", param2: "pl-PL", want: true},                // Poland PESEL valid
		{name: "Invalid Poland PESEL characters", param1: "4405140A359", param2: "pl-PL", want: false}, // Poland PESEL with characters
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsTaxID(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
