package validatorgo

import "testing"

func TestIsLicensePlate(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Valid examples
		{name: "Valid Czech plate", param1: "AA12345", param2: "cs-CZ", want: true},
		{name: "Valid German plate", param1: "B-MA-1234", param2: "de-DE", want: true},
		{name: "Valid Liechtenstein plate", param1: "12345", param2: "de-LI", want: true},
		{name: "Valid Indian plate", param1: "DL12AB1234", param2: "en-IN", want: true},
		{name: "Valid Singapore plate", param1: "SGP1234A", param2: "en-SG", want: true},
		{name: "Valid Pakistan plate", param1: "ABC-1234", param2: "en-PK", want: true},
		{name: "Valid Argentine plate", param1: "AB123CD", param2: "es-AR", want: true},
		{name: "Valid Hungarian plate", param1: "ABC-123", param2: "hu-HU", want: true},
		{name: "Valid Brazilian plate", param1: "ABC-1234", param2: "pt-BR", want: true},
		{name: "Valid Portuguese plate", param1: "12-34-AB", param2: "pt-PT", want: true},
		{name: "Valid Albanian plate", param1: "AB-123-CD", param2: "sq-AL", want: true},
		{name: "Valid Swedish plate", param1: "ABC123", param2: "sv-SE", want: true},
		{name: "Valid Swedish plate, no locale provided", param1: "ABC123", param2: "", want: true},
		{name: "Valid Swedish plate, any locale provided", param1: "ABC123", param2: "any", want: true},

		// Invalid examples
		{name: "Invalid Czech plate", param1: "12A3456", param2: "cs-CZ", want: false},
		{name: "Invalid German plate", param1: "B-12345", param2: "de-DE", want: false},
		{name: "Invalid Liechtenstein plate", param1: "ABCDE", param2: "de-LI", want: false},
		{name: "Invalid Indian plate", param1: "DL1234AB12", param2: "en-IN", want: false},
		{name: "Invalid Singapore plate", param1: "S12345GP", param2: "en-SG", want: false},
		{name: "Invalid Pakistan plate", param1: "1234-ABC", param2: "en-PK", want: false},
		{name: "Invalid Argentine plate", param1: "123ABCD", param2: "es-AR", want: false},
		{name: "Invalid Hungarian plate", param1: "AB-1234", param2: "hu-HU", want: false},
		{name: "Invalid Brazilian plate", param1: "AB-12345", param2: "pt-BR", want: false},
		{name: "Invalid Portuguese plate", param1: "12-345-AB", param2: "pt-PT", want: false},
		{name: "Invalid Albanian plate", param1: "AB1234CD", param2: "sq-AL", want: false},
		{name: "Invalid Swedish plate", param1: "ABC-123", param2: "sv-SE", want: false},
		{name: "Valid Swedish plate, unrecognized locale provided", param1: "ABC-123", param2: "ta-XY", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsLicensePlate(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
