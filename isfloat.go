package validatorgo

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	isFloatOptsDefaultMin    *float64 = nil
	isFloatOptsDefaultMax    *float64 = nil
	isFloatOptsDefaultGt     *float64 = nil
	isFloatOptsDefaultLt     *float64 = nil
	isFloatOptsDefaultLocale string   = "en-US"
)

type IsFloatOpts struct {
	Min    *float64
	Max    *float64
	Gt     *float64
	Lt     *float64
	Locale string
}

// Define the map with the format as the key and a function as the value
var formatFloat = map[string]func(string) string{
	"dot_decimal_comma_thousands": func(s string) string {
		// Remove the thousands separator and replace the comma with a dot
		s = strings.Replace(s, ",", "", -1)
		return strings.Replace(s, ".", ".", 1)
	},
	"comma_decimal_dot_thousands": func(s string) string {
		// Remove the thousands separator and replace the comma with a dot
		s = strings.Replace(s, ".", "", -1)
		return strings.Replace(s, ",", ".", 1)
	},
	"arabic_numerals_dot_decimal": func(s string) string {
		// This is an example for converting Arabic numerals (assuming input is already in Arabic format)
		return s // Assuming Arabic numerals are correctly formatted
	},
	"comma_decimal_space_thousands": func(s string) string {
		// Remove the thousands separator (space) and replace the comma with a dot
		s = strings.Replace(s, " ", "", -1)
		return strings.Replace(s, ",", ".", 1)
	},
	"dot_decimal_no_thousands": func(s string) string {
		// No thousands separator, just replace the comma with a dot
		return strings.Replace(s, ",", ".", 1)
	},
	"arabic_numerals_comma_decimal": func(s string) string {
		// Assuming input is already in Arabic numerals with comma
		return strings.Replace(s, ",", ".", 1)
	},
	"indian_numbering_system": func(s string) string {
		// Remove the Indian grouping comma and replace it with the standard format
		s = strings.Replace(s, ",", "", -1)
		return strings.Replace(s, ".", ".", 1)
	},
}

// floatFormats is the set of decimal formats and their validating regular expressions
var floatFormats = map[string]*regexp.Regexp{
	"dot_decimal_comma_thousands":        regexp.MustCompile(`^[+-]?(\d{1,3}(,\d{3})*\.\d+|\d+\.\d+)$`),          // Matches 1,234.56 and 1234567.56
	"comma_decimal_dot_thousands":        regexp.MustCompile(`^[+-]?(\d{1,3}(\.\d{3})*,\d+|\d+,\d+)$`),           // Matches 1.234,56 and 1234567,56
	"comma_decimal_space_thousands":      regexp.MustCompile(`^[+-]?(\d{1,3}( \d{3})*,\d+|\d+,\d+)$`),            // Matches 1 234,56 and 1234567,56
	"dot_decimal_space_thousands":        regexp.MustCompile(`^[+-]?(\d{1,3}( \d{3})*\.\d+|\d+\.\d+)$`),          // Matches 1 234.56 and 1234567.56
	"dot_decimal_apostrophe_thousands":   regexp.MustCompile(`^[+-]?(\d{1,3}('\d{3})*\.\d+|\d+\.\d+)$`),          // Matches 1'234.56 and 1234567.56
	"comma_decimal_apostrophe_thousands": regexp.MustCompile(`^[+-]?(\d{1,3}('\d{3})*,\d+|\d+,\d+)$`),            // Matches 1'234,56 and 1234567,56
	"dot_decimal_no_thousands":           regexp.MustCompile(`^[+-]?\d*\.\d+$`),                                  // Matches 1234567.78
	"comma_decimal_no_thousands":         regexp.MustCompile(`^[+-]?\d*,\d+$`),                                   // Matches 1234567,78
	"indian_numbering_system":            regexp.MustCompile(`^[+-]?(\d{1,2}(,\d{2})*(,\d{3})*\.\d+|\d+\.\d+)$`), // Matches 1,23,456.78 and 1234567.78
	"arabic_numerals_dot_decimal":        regexp.MustCompile(`^[+-]?[٠١٢٣٤٥٦٧٨٩]*(\.\d+|\d+\.\d+)$`),             // Matches ١٢٣٤٥٦٧٫٥٦
	"arabic_numerals_comma_decimal":      regexp.MustCompile(`^[+-]?[٠١٢٣٤٥٦٧٨٩]*(\,\d+|\d+\,\d+)$`),             // Matches ١٢٣٤٥٦٧،٥٦
}

// floatDecimalFormats is the set of Locales and their decimal formats
var floatDecimalFormats = map[string]string{
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

// const defaultIsFloatFormat = "dot_decimal_comma_thousands"

// A validator that checks if the string is a float.
//
// IsFloatOpts is a struct which can contain the fields Min, Max, Gt, and/or Lt to validate the float is within boundaries it also has Locale as an option.
//
// Min and Max: are equivalent to 'greater or equal' and 'less or equal'.
//
// Gt and Lt: are their strict counterparts.
//
// Locale determines the decimal separator and is one of ('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'cs-CZ', 'da-DK', 'de-DE', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fr-CA', 'fr-FR', 'hu-HU', 'it-IT', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-BR', 'pt-PT', 'ru-RU', 'sl-SI', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'tr-TR', 'uk-UA').
//
//	ok := validatorgo.IsFloat("123.45",  &validatorgo.IsFloatOpts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsFloat("123", &validatorgo.IsFloatOpts{})
//	fmt.Println(ok) // false
func IsFloat(str string, opts *IsFloatOpts) bool {
	if opts == nil {
		opts = setIsFloatOptsToDefault()
	}

	if opts.Locale == "" {
		opts.Locale = isFloatOptsDefaultLocale
	}

	format, ok := floatDecimalFormats[opts.Locale]
	if !ok {
		return false
	}

	parsableFltFunc := formatFloat[format]
	parsableFlt := parsableFltFunc(str)

	flt, err := strconv.ParseFloat(parsableFlt, 64)

	if err != nil {
		return false
	}

	re := floatFormats[format]

	inRange := true

	if opts.Min != nil {
		inRange = flt >= *opts.Min && inRange
	}

	if opts.Max != nil {
		inRange = flt <= *opts.Max && inRange
	}

	if opts.Lt != nil {
		inRange = flt < *opts.Lt && inRange
	}

	if opts.Gt != nil {
		inRange = flt > *opts.Gt && inRange
	}

	return re.MatchString(str) && inRange
}

func setIsFloatOptsToDefault() *IsFloatOpts {
	return &IsFloatOpts{
		Min: isFloatOptsDefaultMin,
		Max: isFloatOptsDefaultMax,
		Gt:  isFloatOptsDefaultGt,
		Lt:  isFloatOptsDefaultLt,
	}
}
