package validatorgo

import (
	"fmt"
	"regexp"
)

// DecimalDigits is used to configure IsDecimalOpts
type DecimalDigits struct {
	Min uint
	Max uint
}

// IsDecimalOpts is used to configure IsDecimal
type IsDecimalOpts struct {
	ForceDecimal bool
	DecimalDigits
	Locale string
}

// decimalFormats is the set of decimal formats and their validating functions
var decimalFormats = map[string]func(IsDecimalOpts) *regexp.Regexp{
	"dot_decimal_comma_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 1,234.56
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d{1,3}(,\d{3})*)?(\.\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"comma_decimal_dot_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 1.234,56
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d{1,3}(\.\d{3})*)?(\,\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"comma_decimal_space_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 1 234,56
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d{1,3}( \d{3})*)?(\,\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"dot_decimal_space_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 1 234.56
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d{1,3}( \d{3})*)?(\.\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"dot_decimal_apostrophe_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 1'234.56
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d{1,3}('\d{3})*)?(\.\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"comma_decimal_apostrophe_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 1'234,56
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d{1,3}('\d{3})*)?(\,\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"dot_decimal_no_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 123456.78
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d*)?(\.\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"comma_decimal_no_thousands": func(ido IsDecimalOpts) *regexp.Regexp { // 123456,78
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d*)?(\,\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"indian_numbering_system": func(ido IsDecimalOpts) *regexp.Regexp { // 1,23,456.78
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?(\d{1,2}(,\d{2})*(,\d{3}))?(\.\d{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"arabic_numerals_dot_decimal": func(ido IsDecimalOpts) *regexp.Regexp { // ١٫٢٣٤,٥٦
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?[٠١٢٣٤٥٦٧٨٩]*(\.[٠١٢٣٤٥٦٧٨٩]{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
	"arabic_numerals_comma_decimal": func(ido IsDecimalOpts) *regexp.Regexp { // ١٬٢٣٤٫٥٦
		var q = "?"
		if ido.ForceDecimal {
			q = ""
		}
		return regexp.MustCompile(fmt.Sprintf(`^[+-]?[٠١٢٣٤٥٦٧٨٩]*(٬[٠١٢٣٤٥٦٧٨٩]{%d,%d})%s$`, ido.DecimalDigits.Min, ido.DecimalDigits.Max, q))
	},
}

// localeDecimalFormats is the set of Locales and their decimal formats
var localeDecimalFormats = map[string]string{
	"ar":          "arabic_numerals_dot_decimal",
	"ar-AE":       "arabic_numerals_dot_decimal",
	"ar-BH":       "arabic_numerals_dot_decimal",
	"ar-DZ":       "arabic_numerals_dot_decimal",
	"ar-EG":       "arabic_numerals_dot_decimal",
	"ar-IQ":       "arabic_numerals_dot_decimal",
	"ar-JO":       "arabic_numerals_dot_decimal",
	"ar-KW":       "arabic_numerals_dot_decimal",
	"ar-LB":       "arabic_numerals_dot_decimal",
	"ar-LY":       "arabic_numerals_dot_decimal",
	"ar-MA":       "arabic_numerals_dot_decimal",
	"ar-QA":       "arabic_numerals_dot_decimal",
	"ar-QM":       "arabic_numerals_dot_decimal",
	"ar-SA":       "arabic_numerals_dot_decimal",
	"ar-SD":       "arabic_numerals_dot_decimal",
	"ar-SY":       "arabic_numerals_dot_decimal",
	"ar-TN":       "arabic_numerals_dot_decimal",
	"ar-YE":       "arabic_numerals_dot_decimal",
	"bg-BG":       "comma_decimal_space_thousands",
	"cs-CZ":       "comma_decimal_dot_thousands",
	"da-DK":       "comma_decimal_dot_thousands",
	"de-DE":       "comma_decimal_dot_thousands",
	"el-GR":       "comma_decimal_dot_thousands",
	"en-AU":       "dot_decimal_comma_thousands",
	"en-GB":       "dot_decimal_comma_thousands",
	"en-HK":       "dot_decimal_comma_thousands",
	"en-IN":       "indian_numbering_system",
	"en-NZ":       "dot_decimal_comma_thousands",
	"en-US":       "dot_decimal_comma_thousands",
	"en-ZA":       "dot_decimal_comma_thousands",
	"en-ZM":       "dot_decimal_comma_thousands",
	"eo":          "dot_decimal_no_thousands",
	"es-ES":       "comma_decimal_dot_thousands",
	"fa":          "arabic_numerals_comma_decimal",
	"fa-AF":       "arabic_numerals_comma_decimal",
	"fa-IR":       "arabic_numerals_comma_decimal",
	"fr-FR":       "comma_decimal_space_thousands",
	"fr-CA":       "dot_decimal_comma_thousands",
	"hu-HU":       "comma_decimal_space_thousands",
	"id-ID":       "dot_decimal_no_thousands",
	"it-IT":       "comma_decimal_dot_thousands",
	"ku-IQ":       "arabic_numerals_dot_decimal",
	"nb-NO":       "comma_decimal_space_thousands",
	"nl-NL":       "comma_decimal_dot_thousands",
	"nn-NO":       "comma_decimal_space_thousands",
	"pl-PL":       "comma_decimal_space_thousands",
	"pl-Pl":       "comma_decimal_space_thousands",
	"pt-BR":       "comma_decimal_dot_thousands",
	"pt-PT":       "comma_decimal_dot_thousands",
	"ru-RU":       "comma_decimal_space_thousands",
	"sl-SI":       "comma_decimal_dot_thousands",
	"sr-RS":       "comma_decimal_dot_thousands",
	"sr-RS@latin": "comma_decimal_dot_thousands",
	"sv-SE":       "comma_decimal_space_thousands",
	"tr-TR":       "comma_decimal_dot_thousands",
	"uk-UA":       "comma_decimal_space_thousands",
	"vi-VN":       "dot_decimal_no_thousands",

	"zh-CN": "dot_decimal_comma_thousands",
	"zh-TW": "dot_decimal_comma_thousands",
	"ja-JP": "dot_decimal_no_thousands",
	"ko-KR": "dot_decimal_no_thousands",
	"th-TH": "dot_decimal_no_thousands",
}

const (
	defaultIsDecimalMinDigit = 1
	defaultIsDecimalMaxDigit = 10

	defaultIsDecimalFormat = "dot_decimal_comma_thousands"
)

// A validator that check if the string represents a decimal number, such as 0.1, .3, 1.1, 1.00003, 4.0, etc.
//
// IsDecimalOpts is a struct which defaults to {ForceDecimal: false, DecimalDigits: {Min: 1, Max: 10}, locale: 'en-US'}, if any are invalid or logically impossible (Min greater than Max).
//
// locale determines the decimal separator and is one of ('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'cs-CZ', 'da-DK', 'de-DE', 'el-GR', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fa', 'fa-AF', 'fa-IR', 'fr-FR', 'fr-CA', 'hu-HU', 'id-ID', 'it-IT', 'ku-IQ', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pl-Pl', 'pt-BR', 'pt-PT', 'ru-RU', 'sl-SI', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'tr-TR', 'uk-UA', 'vi-VN').
//
// ForceDecimal simply means a decimal must be present "123" will not pass but "123.45" will pass
func IsDecimal(str string, opts IsDecimalOpts) bool {
	if opts.DecimalDigits.Min <= 0 {
		opts.DecimalDigits.Min = defaultIsDecimalMinDigit
	}
	if opts.DecimalDigits.Max <= 0 {
		opts.DecimalDigits.Max = defaultIsDecimalMaxDigit
	}
	if opts.DecimalDigits.Min > opts.DecimalDigits.Max {
		opts.DecimalDigits.Min, opts.DecimalDigits.Max = opts.DecimalDigits.Max, opts.DecimalDigits.Min
	}

	format, ok := localeDecimalFormats[opts.Locale]

	if !ok {
		format = defaultIsDecimalFormat
	}

	validReFunc := decimalFormats[format]

	return validReFunc(opts).MatchString(str)
}
