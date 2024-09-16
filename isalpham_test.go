package validatorgo

import (
	"testing"
)

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsAlphaOpts
		want   bool
	}{
		// Valid alpha
		{name: "Basic alpha check", param1: "hello", param2: IsAlphaOpts{}, want: true},
		{name: "Uppercase alpha check", param1: "HELLO", param2: IsAlphaOpts{}, want: true},
		{name: "Mixed case alpha check", param1: "HelloWorld", param2: IsAlphaOpts{}, want: true},
		{name: "German locale alpha check", param1: "äöüß", param2: IsAlphaOpts{Locale: "de-DE"}, want: true},
		{name: "German locale with uppercase", param1: "Schön", param2: IsAlphaOpts{Locale: "de-DE"}, want: true},
		{name: "Spanish locale alpha check", param1: "ÁÉÍÓÚáéíóú", param2: IsAlphaOpts{Locale: "es-ES"}, want: true},
		{name: "Spanish locale with ñ", param1: "niño", param2: IsAlphaOpts{Locale: "es-ES"}, want: true},
		{name: "French locale alpha check", param1: "çàèéêô", param2: IsAlphaOpts{Locale: "fr-FR"}, want: true},
		{name: "French locale with mixed case", param1: "élève", param2: IsAlphaOpts{Locale: "fr-FR"}, want: true},
		{name: "Ignore hyphen option", param1: "hello-world", param2: IsAlphaOpts{Ignore: "-"}, want: true},
		{name: "Ignore digits option", param1: "Hello123", param2: IsAlphaOpts{Ignore: "123"}, want: true},
		// Invalid alpha
		{name: "Invalid with digits", param1: "hello123", param2: IsAlphaOpts{}, want: false},
		{name: "Invalid with special character", param1: "hello!", param2: IsAlphaOpts{}, want: false},
		{name: "Only digits", param1: "123", param2: IsAlphaOpts{}, want: false},
		{name: "German with digits", param1: "äöüß123", param2: IsAlphaOpts{Locale: "de-DE"}, want: false},
		{name: "German with special character", param1: "Schön!", param2: IsAlphaOpts{Locale: "de-DE"}, want: false},
		{name: "Spanish with digits", param1: "ÁÉÍÓÚ123", param2: IsAlphaOpts{Locale: "es-ES"}, want: false},
		{name: "Spanish with special character", param1: "niño!", param2: IsAlphaOpts{Locale: "es-ES"}, want: false},
		{name: "French with digits", param1: "çàèéêô123", param2: IsAlphaOpts{Locale: "fr-FR"}, want: false},
		{name: "French with special character", param1: "élève!", param2: IsAlphaOpts{Locale: "fr-FR"}, want: false},
		{name: "Ignore option fails", param1: "hello-world", param2: IsAlphaOpts{Ignore: "!"}, want: false},
		{name: "Ignore incorrect digits", param1: "Hello123", param2: IsAlphaOpts{Ignore: "456"}, want: false},
		{name: "Invalid locale", param1: "hello", param2: IsAlphaOpts{Locale: "invalid-locale"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAlpha(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}