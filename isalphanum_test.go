package validatorgo

import (
	"testing"
)

func TestIsAlphanumeric(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsAlphanumericOpts
		want   bool
	}{
		{name: "Basic alphanumeric check", param1: "hello123", param2: IsAlphanumericOpts{}, want: true},
		{name: "Uppercase alphanumeric check", param1: "HELLO123", param2: IsAlphanumericOpts{}, want: true},
		{name: "Mixed case alphanumeric check", param1: "Hello123World", param2: IsAlphanumericOpts{}, want: true},
		{name: "German locale alphanumeric check", param1: "äöüß123", param2: IsAlphanumericOpts{Locale: "de-DE"}, want: true},
		{name: "German locale with uppercase", param1: "Schön123", param2: IsAlphanumericOpts{Locale: "de-DE"}, want: true},
		{name: "Spanish locale alphanumeric check", param1: "ÁÉÍÓÚáéíóú123", param2: IsAlphanumericOpts{Locale: "es-ES"}, want: true},
		{name: "Spanish locale with ñ", param1: "niño123", param2: IsAlphanumericOpts{Locale: "es-ES"}, want: true},
		{name: "French locale alphanumeric check", param1: "çàèéêô123", param2: IsAlphanumericOpts{Locale: "fr-FR"}, want: true},
		{name: "French locale with mixed case", param1: "élève123", param2: IsAlphanumericOpts{Locale: "fr-FR"}, want: true},
		{name: "Ignore hyphen option", param1: "hello-123", param2: IsAlphanumericOpts{Ignore: "-"}, want: true},
		{name: "Ignore space option", param1: "Hello 123", param2: IsAlphanumericOpts{Ignore: " "}, want: true},
		{name: "Invalid with special character", param1: "hello!", param2: IsAlphanumericOpts{}, want: false},
		{name: "Only special characters", param1: "!!!", param2: IsAlphanumericOpts{}, want: false},
		{name: "German with special character", param1: "äöüß!123", param2: IsAlphanumericOpts{Locale: "de-DE"}, want: false},
		{name: "Spanish with special character", param1: "ÁÉÍÓÚ!123", param2: IsAlphanumericOpts{Locale: "es-ES"}, want: false},
		{name: "French with special character", param1: "çàèéêô!123", param2: IsAlphanumericOpts{Locale: "fr-FR"}, want: false},
		{name: "Ignore option fails", param1: "hello-123", param2: IsAlphanumericOpts{Ignore: "!"}, want: false},
		{name: "Ignore incorrect characters", param1: "Hello123!", param2: IsAlphanumericOpts{Ignore: "-"}, want: false},
		{name: "Invalid locale", param1: "hello123", param2: IsAlphanumericOpts{Locale: "invalid-locale"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAlphanumeric(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}