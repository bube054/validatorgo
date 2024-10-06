package validatorgo

import (
	"testing"
)

func TestIsNumericWithNoSymbolsAndNoLocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsNumericOpts
		want   bool
	}{
		// en-AU: Locale 0 (1,234,567.89)
		{name: "Valid config Locale en-AU, NoSymbols true", param1: "1,234,567.89", param2: &IsNumericOpts{Locale: "en-AU", NoSymbols: true}, want: true},
		{name: "Valid config Locale en-AU, NoSymbols false", param1: "+1,234,567.89", param2: &IsNumericOpts{Locale: "en-AU", NoSymbols: false}, want: true},
		{name: "Invalid config Locale en-AU, NoSymbols true", param1: "1.234.567,89", param2: &IsNumericOpts{Locale: "en-AU", NoSymbols: true}, want: false},
		{name: "Invalid config Locale en-AU, NoSymbols false", param1: "+1.234.567,89", param2: &IsNumericOpts{Locale: "en-AU", NoSymbols: false}, want: false},

		// eo: Locale 1 (1234567.89)
		{name: "Valid config Locale eo, NoSymbols true", param1: "1234567.89", param2: &IsNumericOpts{Locale: "eo", NoSymbols: true}, want: true},
		{name: "Valid config Locale eo, NoSymbols false", param1: "+1234567.89", param2: &IsNumericOpts{Locale: "eo", NoSymbols: false}, want: true},
		{name: "Invalid config Locale eo, NoSymbols true", param1: "1,234,567.89", param2: &IsNumericOpts{Locale: "eo", NoSymbols: true}, want: false},
		{name: "Invalid config Locale eo, NoSymbols false", param1: "+1,234,567.89", param2: &IsNumericOpts{Locale: "eo", NoSymbols: false}, want: false},

		// bg-BG: Locale 3 (1.234.567,89)
		{name: "Valid config Locale bg-BG, NoSymbols true", param1: "1.234.567,89", param2: &IsNumericOpts{Locale: "bg-BG", NoSymbols: true}, want: true},
		{name: "Valid config Locale bg-BG, NoSymbols false", param1: "+1.234.567,89", param2: &IsNumericOpts{Locale: "bg-BG", NoSymbols: false}, want: true},
		{name: "Invalid config Locale bg-BG, NoSymbols true", param1: "1,234,567.89", param2: &IsNumericOpts{Locale: "bg-BG", NoSymbols: true}, want: false},
		{name: "Invalid config Locale bg-BG, NoSymbols false", param1: "+1,234,567.89", param2: &IsNumericOpts{Locale: "bg-BG", NoSymbols: false}, want: false},

		// en-IN: Locale 5 (12,34,567.89)
		{name: "Valid config Locale en-IN, NoSymbols true", param1: "12,34,567.89", param2: &IsNumericOpts{Locale: "en-IN", NoSymbols: true}, want: true},
		{name: "Valid config Locale en-IN, NoSymbols false", param1: "+12,34,567.89", param2: &IsNumericOpts{Locale: "en-IN", NoSymbols: false}, want: true},
		{name: "Invalid config Locale en-IN, NoSymbols true", param1: "1,234,567.89", param2: &IsNumericOpts{Locale: "en-IN", NoSymbols: true}, want: false},
		{name: "Invalid config Locale en-IN, NoSymbols false", param1: "+1,234,567.89", param2: &IsNumericOpts{Locale: "en-IN", NoSymbols: false}, want: false},

		// ar-AE: Locale 8 (١٬٢٣٤٬٥٦٧٫٨٩)
		{name: "Valid config Locale ar-AE, NoSymbols true", param1: "١٬٢٣٤٬٥٦٧٫٨٩", param2: &IsNumericOpts{Locale: "ar-AE", NoSymbols: true}, want: true},
		{name: "Valid config Locale ar-AE, NoSymbols false", param1: "+١٬٢٣٤٬٥٦٧٫٨٩", param2: &IsNumericOpts{Locale: "ar-AE", NoSymbols: false}, want: true},

		{name: "Invalid config Locale ar-AE, NoSymbols true", param1: "1,234,567.89", param2: &IsNumericOpts{Locale: "ar-AE", NoSymbols: true}, want: false},
		{name: "Invalid config Locale ar-AE, NoSymbols false", param1: "+1,234,567.89", param2: &IsNumericOpts{Locale: "ar-AE", NoSymbols: false}, want: false},

		// Valid default Configuration NoSymbols is false
		{name: "Valid default config", param1: "+12,345", param2: nil, want: true},
		{name: "Valid default config", param1: "+00123", param2: &IsNumericOpts{}, want: true},
		{name: "Valid default config", param1: "9876543210", param2: &IsNumericOpts{}, want: true},
		{name: "Valid default config", param1: "-123.456", param2: &IsNumericOpts{}, want: true},
		// Invalid default Configuration NoSymbols is false
		{name: "Invalid default config, multiple periods", param1: "12.34.56", param2: &IsNumericOpts{}, want: false},
		{name: "Invalid default config, contains letters", param1: "123abc", param2: &IsNumericOpts{}, want: false},
		{name: "Invalid default config, contains comma", param1: "123,456", param2: &IsNumericOpts{}, want: false},
		{name: "Invalid default config, empty space", param1: " ", param2: &IsNumericOpts{}, want: false},

		// Valid Configuration NoSymbols is true
		{name: "Valid config NoSymbol true", param1: "12345", param2: &IsNumericOpts{NoSymbols: true}, want: true},
		{name: "Valid config NoSymbol true", param1: "00123", param2: &IsNumericOpts{NoSymbols: true}, want: true},
		{name: "Valid config NoSymbol true", param1: "0", param2: &IsNumericOpts{NoSymbols: true}, want: true},
		// Invalid Configuration NoSymbols is true
		{name: "Invalid config, decimal point is considered a symbol", param1: "123.456", param2: &IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, minus is considered a symbol", param1: "-123", param2: &IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, plus is considered a symbol", param1: "+123", param2: &IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, comma point is considered a symbol", param1: "12,345", param2: &IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, contains letters", param1: "12abc", param2: &IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, no spaces", param1: " 123 ", param2: &IsNumericOpts{NoSymbols: true}, want: false},

		// Valid Configuration Locale is present (French locale where commas are used as decimal separators)
		{name: "Valid config Locale french, greater than 1000th(s) with decimal", param1: "+1.234.567,89", param2: &IsNumericOpts{Locale: "fr-FR"}, want: true},
		{name: "Valid config Locale french, less than 1000th(s) with decimal", param1: "567,89", param2: &IsNumericOpts{Locale: "fr-FR"}, want: true},
		{name: "Valid config Locale french, less than 100th(s) with decimal", param1: "-0,345", param2: &IsNumericOpts{Locale: "fr-FR"}, want: true},
		{name: "Invalid config Locale french, less than 100th(s) with no decimal", param1: "12.345", param2: &IsNumericOpts{Locale: "fr-FR"}, want: true},
		// Invalid Configuration Locale is present (French locale where commas are used as decimal separators)
		{name: "Invalid config Locale french, contains letter", param1: "123abc", param2: &IsNumericOpts{Locale: "fr-FR"}, want: false},
		{name: "Invalid config Locale french, empty space", param1: " ", param2: &IsNumericOpts{Locale: "fr-FR"}, want: false},

		// Valid Configuration with  Locale is present (French locale where commas are used as decimal separators)
		{name: "Valid config Locale french, No plus", param1: "1.234.567,89", param2: &IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: true},
		{name: "Valid config Locale french, No minus", param1: "1.234.567,89", param2: &IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: true},
		// Invalid Configuration Locale is present (French locale where commas are used as decimal separators)
		{name: "Valid config Locale french, No plus", param1: "+1.234.567,89", param2: &IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: false},
		{name: "Valid config Locale french, No minus", param1: "-1.234.567,89", param2: &IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsNumeric(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
