package validatorgo

import (
	"fmt"
	"regexp"
	"strconv"
)

// decimalFormats is the set of decimal formats and their validating functions
var decimalFormats = map[string]func(IsDecimalOpts) (*regexp.Regexp, error){
	"dot_decimal_comma_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 1,234.56
		var q = "?"
		if *ido.ForceDecimal {
			q = ""
		}
		var maxStr string
		if ido.DecimalDigits.Max == nil {
			maxStr = ""
		} else {
			maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
		}
		return regexp.Compile(fmt.Sprintf(`^[+-]?(\d{1,3}(,\d{3})*)(\.\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	},
	"comma_decimal_dot_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 1.234,56
		var q = "?"
		if *ido.ForceDecimal {
			q = ""
		}
		var maxStr string
		if ido.DecimalDigits.Max == nil {
			maxStr = ""
		} else {
			maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
		}
		return regexp.Compile(fmt.Sprintf(`^[+-]?(\d{1,3}(\.\d{3})*)(\,\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	},
	"comma_decimal_space_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 1 234,56
		var q = "?"
		if *ido.ForceDecimal {
			q = ""
		}
		var maxStr string
		if ido.DecimalDigits.Max == nil {
			maxStr = ""
		} else {
			maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
		}
		return regexp.Compile(fmt.Sprintf(`^[+-]?(\d{1,3}( \d{3})*)(\,\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	},
	// "dot_decimal_space_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 1 234.56
	// 	var q = "?"
	// 	if *ido.ForceDecimal {
	// 		q = ""
	// 	}
	// 	var maxStr string
	// 	if ido.DecimalDigits.Max == nil {
	// 		maxStr = ""
	// 	} else {
	// 		maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
	// 	}
	// 	return regexp.Compile(fmt.Sprintf(`^[+-]?(\d{1,3}( \d{3})*)(\.\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	// },
	// "dot_decimal_apostrophe_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 1'234.56
	// 	var q = "?"
	// 	if *ido.ForceDecimal {
	// 		q = ""
	// 	}
	// 	var maxStr string
	// 	if ido.DecimalDigits.Max == nil {
	// 		maxStr = ""
	// 	} else {
	// 		maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
	// 	}
	// 	return regexp.Compile(fmt.Sprintf(`^[+-]?(\d{1,3}('\d{3})*)(\.\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	// },
	// "comma_decimal_apostrophe_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 1'234,56
	// 	var q = "?"
	// 	if *ido.ForceDecimal {
	// 		q = ""
	// 	}
	// 	var maxStr string
	// 	if ido.DecimalDigits.Max == nil {
	// 		maxStr = ""
	// 	} else {
	// 		maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
	// 	}
	// 	return regexp.Compile(fmt.Sprintf(`^[+-]?(\d{1,3}('\d{3})*)(\,\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	// },
	"dot_decimal_no_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 123456.78
		var q = "?"
		if *ido.ForceDecimal {
			q = ""
		}
		var maxStr string
		if ido.DecimalDigits.Max == nil {
			maxStr = ""
		} else {
			maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
		}
		return regexp.Compile(fmt.Sprintf(`^[+-]?(\d+)?(\.\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	},
	// "comma_decimal_no_thousands": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 123456,78
	// 	var q = "?"
	// 	if *ido.ForceDecimal {
	// 		q = ""
	// 	}
	// 	var maxStr string
	// 	if ido.DecimalDigits.Max == nil {
	// 		maxStr = ""
	// 	} else {
	// 		maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
	// 	}
	// 	return regexp.Compile(fmt.Sprintf(`^[+-]?(\d+)?(\,\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	// },
	"indian_numbering_system": func(ido IsDecimalOpts) (*regexp.Regexp, error) { // 1,23,456.78
		var q = "?"
		if *ido.ForceDecimal {
			q = ""
		}
		var maxStr string
		if ido.DecimalDigits.Max == nil {
			maxStr = ""
		} else {
			maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
		}
		return regexp.Compile(fmt.Sprintf(`^[+-]?(\d{1,2}(,\d{2})*(,\d{3}))?(\.\d{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	},
	"arabic_numerals_dot_decimal": func(ido IsDecimalOpts) (*regexp.Regexp, error) {
		var q = "?"
		if *ido.ForceDecimal {
			q = ""
		}
		var maxStr string
		if ido.DecimalDigits.Max == nil {
			maxStr = ""
		} else {
			maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
		}
		return regexp.Compile(fmt.Sprintf(`^[+-]?[٠١٢٣٤٥٦٧٨٩]+(٫[٠١٢٣٤٥٦٧٨٩]{%d,%s})%s$`, *ido.DecimalDigits.Min, maxStr, q))
	},
	"arabic_numerals_comma_decimal": func(ido IsDecimalOpts) (*regexp.Regexp, error) {
		var q = "?" // Optional decimal part by default
		if *ido.ForceDecimal {
			q = "" // Force the decimal part to appear
		}
		var maxStr string
		if ido.DecimalDigits.Max == nil {
			maxStr = "" // No upper limit for decimal digits
		} else {
			maxStr = strconv.Itoa(int(*ido.DecimalDigits.Max))
		}

		// General regex for decimal numbers with comma as a thousand separator and locale-based numerals
		return regexp.Compile(fmt.Sprintf(
			`^[+-]?[۰۱۲۳۴۵۶۷۸۹١٢٣٤٥٦٧٨٩0-9]+(\,?[۰۱۲۳۴۵۶۷۸۹١٢٣٤٥٦٧٨٩0-9]{3})*(\.|٫)[۰۱۲۳۴۵۶۷۸۹١٢٣٤٥٦٧٨٩0-9]{%d,%s}%s$`,
			*ido.DecimalDigits.Min, maxStr, q))
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
