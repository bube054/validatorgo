package validatorgo

import (
	"testing"
)

func TestIsNumericWithNoSymbolsAndNoLocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsNumericOpts
		want   bool
	}{
		// Valid default Configuration NoSymbols is false
		{name: "Valid default config", param1: "12345", param2: IsNumericOpts{}, want: true},
		{name: "Valid default config", param1: "+00123", param2: IsNumericOpts{}, want: true},
		{name: "Valid default config", param1: "9876543210", param2: IsNumericOpts{}, want: true},
		{name: "Valid default config", param1: "-123.456", param2: IsNumericOpts{}, want: true},
		// Invalid default Configuration NoSymbols is false
		{name: "Invalid default config, multiple periods", param1: "12.34.56", param2: IsNumericOpts{}, want: false},
		{name: "Invalid default config, contains letters", param1: "123abc", param2: IsNumericOpts{}, want: false},
		{name: "Invalid default config, contains comma", param1: "123,456", param2: IsNumericOpts{}, want: false},
		{name: "Invalid default config, empty space", param1: " ", param2: IsNumericOpts{}, want: false},

		// Valid Configuration NoSymbols is true
		{name: "Valid config NoSymbol true", param1: "12345", param2: IsNumericOpts{NoSymbols: true}, want: true},
		{name: "Valid config NoSymbol true", param1: "00123", param2: IsNumericOpts{NoSymbols: true}, want: true},
		{name: "Valid config NoSymbol true", param1: "0", param2: IsNumericOpts{NoSymbols: true}, want: true},
		// Invalid Configuration NoSymbols is true
		{name: "Invalid config, decimal point is considered a symbol", param1: "123.456", param2: IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, minus is considered a symbol", param1: "-123", param2: IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, plus is considered a symbol", param1: "+123", param2: IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, comma point is considered a symbol", param1: "12,345", param2: IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, contains letters", param1: "12abc", param2: IsNumericOpts{NoSymbols: true}, want: false},
		{name: "Invalid config, no spaces", param1: " 123 ", param2: IsNumericOpts{NoSymbols: true}, want: false},

		// Valid Configuration Locale is present (French locale where commas are used as decimal separators)
		{name: "Valid config Locale french, greater than 1000th(s) with decimal", param1: "+1.234.567,89", param2: IsNumericOpts{Locale: "fr-FR"}, want: true},
		{name: "Valid config Locale french, less than 1000th(s) with decimal", param1: "567,89", param2: IsNumericOpts{Locale: "fr-FR"}, want: true},
		{name: "Valid config Locale french, less than 100th(s) with decimal", param1: "-0,345", param2: IsNumericOpts{Locale: "fr-FR"}, want: true},
		{name: "Invalid config Locale french, less than 100th(s) with no decimal", param1: "12.345", param2: IsNumericOpts{Locale: "fr-FR"}, want: true},
		// Invalid Configuration Locale is present (French locale where commas are used as decimal separators)
		{name: "Invalid config Locale french, contains letter", param1: "123abc", param2: IsNumericOpts{Locale: "fr-FR"}, want: false},
		{name: "Invalid config Locale french, empty space", param1: " ", param2: IsNumericOpts{Locale: "fr-FR"}, want: false},

		// Valid Configuration with  Locale is present (French locale where commas are used as decimal separators)
		{name: "Valid config Locale french, No plus", param1: "1.234.567,89", param2: IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: true},
		{name: "Valid config Locale french, No minus", param1: "1.234.567,89", param2: IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: true},
		// Invalid Configuration Locale is present (French locale where commas are used as decimal separators)
		{name: "Valid config Locale french, No plus", param1: "+1.234.567,89", param2: IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: false},
		{name: "Valid config Locale french, No minus", param1: "-1.234.567,89", param2: IsNumericOpts{Locale: "fr-FR", NoSymbols: true}, want: false},

		// {
		// 	name:   "Just a plus",
		// 	param1: "+",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   false,
		// },
		// {
		// 	name:   "Just a minus",
		// 	param1: "-",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   false,
		// },
		// {
		// 	name:   "Just a period",
		// 	param1: ".",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   false,
		// },
		// {
		// 	name:   "A number with no symbols (+ or - or . or ,)",
		// 	param1: "12345",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   true,
		// },
		// {
		// 	name:   "A number with a period",
		// 	param1: "123.45",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   false,
		// },
		// {
		// 	name:   "A number with a comma",
		// 	param1: "123,45",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   false,
		// },
		// {
		// 	name:   "A number with a plus",
		// 	param1: "+12345",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   false,
		// },
		// {
		// 	name:   "A number with a a minus",
		// 	param1: "-12345",
		// 	param2: IsNumericOpts{NoSymbols: true, Locale: ""},
		// 	want:   false,
		// },
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

// func TestIsNumericWithSymbolsAndNoLocale(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		param1 string
// 		param2 IsNumericOpts
// 		want   bool
// 	}{
// 		{
// 			name:   "Just a plus",
// 			param1: "+",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   false,
// 		},
// 		{
// 			name:   "Just a minus",
// 			param1: "-",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   false,
// 		},
// 		{
// 			name:   "Just a period",
// 			param1: ".",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   false,
// 		},
// 		{
// 			name:   "A number with no symbols (+ or - or . or ,)",
// 			param1: "12345",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   true,
// 		},
// 		{
// 			name:   "A number with a period",
// 			param1: "123.45",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   true,
// 		},
// 		{
// 			name:   "A number with a comma",
// 			param1: "123,45",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   false,
// 		},
// 		{
// 			name:   "A number with a plus",
// 			param1: "+12345",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   true,
// 		},
// 		{
// 			name:   "A number with a a minus",
// 			param1: "-12345",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   true,
// 		},
// 		{
// 			name:   "A number starting with a plus and a decimal end",
// 			param1: "+12345.00",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   true,
// 		},
// 		{
// 			name:   "A number starting with a minus and a decimal end",
// 			param1: "-12345.78",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   true,
// 		},
// 		{
// 			name:   "A number starting with a minus and ending with just a decimal point .",
// 			param1: "-12345.",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: ""},
// 			want:   false,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			result := IsNumeric(test.param1, test.param2)

// 			if result != test.want {
// 				t.Errorf("got `%t`, wanted `%t`", result, test.want)
// 			}
// 		})
// 	}
// }

// func TestIsNumericWithLocales(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		param1 string
// 		param2 IsNumericOpts
// 		want   bool
// 	}{
// 		{
// 			name:   "Just a plus",
// 			param1: "+",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: "en-US"},
// 			want:   false,
// 		},
// 		{
// 			name:   "Just a minus",
// 			param1: "-",
// 			param2: IsNumericOpts{NoSymbols: true, Locale: "en-US"},
// 			want:   false,
// 		},
// 		{
// 			name:   "Just a period",
// 			param1: ".",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: "en-US"},
// 			want:   false,
// 		},
// 		{
// 			name:   "Has 3 digits or less",
// 			param1: "567",
// 			param2: IsNumericOpts{NoSymbols: true, Locale: "en-US"},
// 			want:   true,
// 		},
// 		{
// 			name:   "Has 6 digits or less",
// 			param1: "234,567",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: "en-US"},
// 			want:   true,
// 		},
// 		{
// 			name:   "Has 9 digits or less",
// 			param1: "1,234,567",
// 			param2: IsNumericOpts{NoSymbols: true, Locale: "en-US"},
// 			want:   true,
// 		},
// 		{
// 			name:   "Has 3 digits or less with decimal point",
// 			param1: "567.89",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: "en-US"},
// 			want:   true,
// 		},
// 		{
// 			name:   "Has 6 digits or less with decimal point",
// 			param1: "234,567.89",
// 			param2: IsNumericOpts{NoSymbols: true, Locale: "en-US"},
// 			want:   true,
// 		},
// 		{
// 			name:   "Has 9 digits or less with decimal point",
// 			param1: "1,234,567.89",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: "en-US"},
// 			want:   true,
// 		},
// 		{
// 			name:   "Has 3 digits or less with decimal point and plus sign",
// 			param1: "+567.89",
// 			param2: IsNumericOpts{NoSymbols: true, Locale: "en-US"},
// 			want:   true,
// 		},
// 		{
// 			name:   "Has 6 digits or less with decimal point and minus sign",
// 			param1: "-234,567",
// 			param2: IsNumericOpts{NoSymbols: false, Locale: "en-US"},
// 			want:   true,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			result := IsNumeric(test.param1, test.param2)

// 			if result != test.want {
// 				t.Errorf("got `%t`, wanted `%t`", result, test.want)
// 			}
// 		})
// 	}
// }
