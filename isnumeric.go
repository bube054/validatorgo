package validatorgo

import (
	"regexp"
)

// IsNumericOpts is used to configure IsNumeric
type IsNumericOpts struct {
	NoSymbols bool
	Locale    string
}

// numericFormatsRegex is the set of number validating regex for each cldr code
var numericFormatsRegex = map[int]func(opts IsNumericOpts) *regexp.Regexp{
	0: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})(,\d{3})*(\.\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})(,\d{3})*(\.\d+)?$`) // With symbols
	}, // matches 1,234,567.89 (comma for thousands, dot for decimals)
	1: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d+)(\.\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d+)(\.\d+)?$`) // With symbols
	}, // matches 1234567.89 (no thousands separator, dot for decimals)
	2: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})( \d{1,3})*(,\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})( \d{1,3})*(,\d+)?$`) // With symbols
	}, // matches 1 234 567,89 (space for thousands, comma for decimals)
	3: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})(\.\d{3})*(,\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})(\.\d{3})*(,\d+)?$`) // With symbols
	}, // matches 1.234.567,89 (dot for thousands, comma for decimals)
	4: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})(,\d{3})*(·\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})(,\d{3})*(·\d+)?$`) // With symbols
	}, // matches 1,234,567·89 (comma for thousands, interpunct for decimals)
	5: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,2},)*(\d{1,3})(\.\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,2},)*(\d{1,3})(\.\d+)?$`) // With symbols
	}, // matches 12,34,567.89 (Indian system: commas for thousands, dot for decimals)
	6: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,2} )*(\d{1,3})(\.\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,2} )*(\d{1,3})(\.\d+)?$`) // With symbols
	}, // matches 12 34 567.89 (Indian system: space for thousands, dot for decimals)
	7: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})('\d{3})*(\.\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})('\d{3})*(\.\d+)?$`) // With symbols
	}, // matches 1'234'567.89 (apostrophe for thousands, dot for decimals)
	8: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})(٬\d{3})*(٫\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})(٬\d{3})*(٫\d+)?$`) // With symbols
	}, // matches ١٬٢٣٤٬٥٦٧٫٨٩ (Arabic: comma-like for thousands, period-like for decimals)
	9: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})( \d{3})*(\.\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})( \d{3})*(\.\d+)?$`) // With symbols
	}, // matches 1 234 567.89 (space for thousands, dot for decimals)
	10: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})(\.\d{3})*(·\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})(\.\d{3})*(·\d+)?$`) // With symbols
	}, // matches 1234567,89 (no thousands separator, comma for decimals)
	11: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d+)('\d{3})*(٫\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d+)('\d{3})*(٫\d+)?$`) // With symbols
	}, // matches 12,34,567.89 (variant of Indian system with space and comma separators)
	12: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})( \d{3})*(,\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})( \d{3})*(,\d+)?$`) // With symbols
	}, // matches Arabic format without spaces for thousands separators
	13: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d+)(٬\d{3})*(\.\d+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?(\d+)(٬\d{3})*(\.\d+)?$`) // With symbols
	}, // matches numbers without thousand separators, e.g., in East Asia (e.g., 1234.56)
	14: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^(\d{1,3})(,\d{3})*(\.\d{2})?$`) // Without symbols, 2 decimals
		}
		return regexp.MustCompile(`^[+-]?(\d{1,3})(,\d{3})*(\.\d{2})?$`) // With symbols, 2 decimals
	}, // matches 1 2345.67 (thousands separator every four digits, dot for decimals, variant)
}

// codeNumericFormats is the set of locales pointing to numericFormatsRegex
var codeNumericFormats = map[string]int{
	"ar":          8, // Arabic numeral system (1٬234٬567٫89)
	"ar-AE":       8,
	"ar-BH":       8,
	"ar-DZ":       8,
	"ar-EG":       8,
	"ar-IQ":       8,
	"ar-JO":       8,
	"ar-KW":       8,
	"ar-LB":       8,
	"ar-LY":       8,
	"ar-MA":       8,
	"ar-QA":       8,
	"ar-QM":       8,
	"ar-SA":       8,
	"ar-SD":       8,
	"ar-SY":       8,
	"ar-TN":       8,
	"ar-YE":       8,
	"bg-BG":       3, // Bulgarian (1.234.567,89)
	"cs-CZ":       3, // Czech (1.234.567,89)
	"da-DK":       3, // Danish (1.234.567,89)
	"de-DE":       3, // German (1.234.567,89)
	"en-AU":       0, // Australian English (1,234,567.89)
	"en-GB":       0, // British English (1,234,567.89)
	"en-HK":       0, // Hong Kong English (1,234,567.89)
	"en-IN":       5, // Indian English (12,34,567.89)
	"en-NZ":       0, // New Zealand English (1,234,567.89)
	"en-US":       0, // American English (1,234,567.89)
	"en-ZA":       0, // South African English (1,234,567.89)
	"en-ZM":       0, // Zambian English (1,234,567.89)
	"eo":          1, // Esperanto (1234567.89)
	"es-ES":       3, // Spanish (1.234.567,89)
	"fr-FR":       3, // French (1.234.567,89)
	"fr-CA":       0, // Canadian French (1,234,567.89)
	"hu-HU":       3, // Hungarian (1.234.567,89)
	"it-IT":       3, // Italian (1.234.567,89)
	"nb-NO":       3, // Norwegian Bokmål (1.234.567,89)
	"nl-NL":       3, // Dutch (1.234.567,89)
	"nn-NO":       3, // Norwegian Nynorsk (1.234.567,89)
	"pl-PL":       3, // Polish (1.234.567,89)
	"pt-BR":       3, // Brazilian Portuguese (1.234.567,89)
	"pt-PT":       3, // Portuguese (1.234.567,89)
	"ru-RU":       3, // Russian (1.234.567,89)
	"sl-SI":       3, // Slovenian (1.234.567,89)
	"sr-RS":       3, // Serbian (1.234.567,89)
	"sr-RS@latin": 3, // Serbian Latin (1.234.567,89)
	"sv-SE":       3, // Swedish (1.234.567,89)
	"tr-TR":       3, // Turkish (1.234.567,89)
	"uk-UA":       3, // Ukrainian (1.234.567,89)
}

// var codeNumericFormats = map[string]int{
// 	"ar":          3, // Arabic (general)
// 	"ar-AE":       3, // Arabic (United Arab Emirates)
// 	"ar-BH":       3, // Arabic (Bahrain)
// 	"ar-DZ":       3, // Arabic (Algeria)
// 	"ar-EG":       3, // Arabic (Egypt)
// 	"ar-IQ":       3, // Arabic (Iraq)
// 	"ar-JO":       3, // Arabic (Jordan)
// 	"ar-KW":       3, // Arabic (Kuwait)
// 	"ar-LB":       3, // Arabic (Lebanon)
// 	"ar-LY":       3, // Arabic (Libya)
// 	"ar-MA":       3, // Arabic (Morocco)
// 	"ar-QA":       3, // Arabic (Qatar)
// 	"ar-QM":       3, // Arabic (Oman)
// 	"ar-SA":       3, // Arabic (Saudi Arabia)
// 	"ar-SD":       3, // Arabic (Sudan)
// 	"ar-SY":       3, // Arabic (Syria)
// 	"ar-TN":       3, // Arabic (Tunisia)
// 	"ar-YE":       3, // Arabic (Yemen)
// 	"bg-BG":       2, // Bulgarian (Bulgaria)
// 	"cs-CZ":       2, // Czech (Czech Republic)
// 	"da-DK":       3, // Danish (Denmark)
// 	"de-DE":       3, // German (Germany)
// 	"en-AU":       0, // English (Australia)
// 	"en-GB":       0, // English (United Kingdom)
// 	"en-HK":       0, // English (Hong Kong)
// 	"en-IN":       5, // English (India)
// 	"en-NZ":       0, // English (New Zealand)
// 	"en-US":       0, // English (United States)
// 	"en-ZA":       2, // English (South Africa)
// 	"en-ZM":       0, // English (Zambia)
// 	"eo":          0, // Esperanto
// 	"es-ES":       3, // Spanish (Spain)
// 	"fr-FR":       2, // French (France)
// 	"fr-CA":       0, // French (Canada)
// 	"hu-HU":       2, // Hungarian (Hungary)
// 	"it-IT":       3, // Italian (Italy)
// 	"nb-NO":       3, // Norwegian Bokmål (Norway)
// 	"nl-NL":       3, // Dutch (Netherlands)
// 	"nn-NO":       3, // Norwegian Nynorsk (Norway)
// 	"pl-PL":       2, // Polish (Poland)
// 	"pt-BR":       2, // Portuguese (Brazil)
// 	"pt-PT":       2, // Portuguese (Portugal)
// 	"ru-RU":       2, // Russian (Russia)
// 	"sl-SI":       3, // Slovenian (Slovenia)
// 	"sr-RS":       3, // Serbian (Cyrillic, Serbia)
// 	"sr-RS@latin": 3, // Serbian (Latin, Serbia)
// 	"sv-SE":       3, // Swedish (Sweden)
// 	"tr-TR":       3, // Turkish (Turkey)
// 	"uk-UA":       2, // Ukrainian (Ukraine)
// }

// A validator that check if a string is a number.
// options is an struct which defaults to { NoSymbols: false, Locale: ""}.
// If NoSymbols is true, the validator will reject numeric strings that feature a symbol (e.g. +, -, or .).
// Locale determines the numeric format and is one of ('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'cs-CZ', 'da-DK', 'de-DE', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fr-FR', 'fr-CA', 'hu-HU', 'it-IT', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-BR', 'pt-PT', 'ru-RU', 'sl-SI', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'tr-TR', 'uk-UA'). If Locale is present but not in this list, Locale will default to en-US.
//
//	isNumber := govalidator.IsNumeric("+123.45", govalidator.IsNumericOpts{NoSymbols: true, Locale: ""})
//	fmt.Println(isNumber) // false
//	isNumber = govalidator.IsNumeric("-123.45", govalidator.IsNumericOpts{NoSymbols: false, Locale: ""})
//	fmt.Println(isNumber) // true
//	isNumber = govalidator.IsNumeric("+123.45", govalidator.IsNumericOpts{Locale: "en"})
//	fmt.Println(isNumber) // true
func IsNumeric(str string, opts IsNumericOpts) bool {
	var re *regexp.Regexp

	// has symbols and no Locale
	if !opts.NoSymbols && opts.Locale == "" {
		re = regexp.MustCompile(`^[+-]?\d+(\.\d+)?$`)
		return re.MatchString(str)
	}

	// no symbols and no Locale
	if opts.NoSymbols && opts.Locale == "" {
		re = regexp.MustCompile(`^\d+$`)
		return re.MatchString(str)
	}

	if opts.Locale == "" {
		opts.Locale = "en-US"
	}

	// Locale is present and plus or minus symbols are optional(NoSymbol does not matter)
	codeNumForm := codeNumericFormats[opts.Locale]
	valFunc := numericFormatsRegex[codeNumForm]
	re = valFunc(opts)

	return re.MatchString(str)
}

// // codeNumericFormats is the set of cldr chart codes
// var codeNumericFormats = map[string]int{
// 	"af":      2, // Afrikaans
// 	"sq":      2, // Albanian
// 	"am":      0, // Amharic
// 	"ar":      0, // Arabic
// 	"hy":      2, // Armenian
// 	"as":      5, // Assamese
// 	"az":      3, // Azerbaijani
// 	"bn":      5, // Bengali (India)
// 	"eu":      3, // Basque
// 	"be":      2, // Belarusian
// 	"bs":      3, // Bosnian
// 	"bg":      2, // Bulgarian
// 	"de_CH":   7, // Swiss German
// 	"fr_CH":   7, // Swiss French
// 	"it_CH":   7, // Swiss Italian
// 	"my":      0, // Burmese
// 	"yue":     5, // Cantonese
// 	"ca":      3, // Catalan
// 	"chr":     5, // Cherokee
// 	"zh":      0, // Chinese (Simplified)
// 	"hr":      3, // Croatian
// 	"cs":      2, // Czech
// 	"da":      3, // Danish
// 	"nl":      3, // Dutch
// 	"en":      0, // English
// 	"et":      2, // Estonian
// 	"pt_PT":   2, // Portuguese (Portugal)
// 	"fil":     0, // Filipino
// 	"fi":      2, // Finnish
// 	"fr":      2, // French
// 	"gl":      3, // Galician
// 	"ka":      2, // Georgian
// 	"de":      3, // German
// 	"el":      3, // Greek
// 	"gu":      5, // Gujarati
// 	"ha":      0, // Hausa
// 	"hi":      0, // Hindi
// 	"hi_Latn": 5, // Hindi (Latin script)
// 	"hu":      2, // Hungarian
// 	"is":      3, // Icelandic
// 	"ig":      0, // Igbo
// 	"id":      3, // Indonesian
// 	"ga":      0, // Irish
// 	"it":      3, // Italian
// 	"ja":      0, // Japanese
// 	"jv":      0, // Javanese
// 	"kn":      5, // Kannada
// 	"kk":      3, // Kazakh
// 	"km":      0, // Khmer
// 	"ko":      0, // Korean
// 	"lo":      0, // Lao
// 	"lv":      2, // Latvian
// 	"lt":      2, // Lithuanian
// 	"mk":      3, // Macedonian
// 	"ms":      3, // Malay
// 	"ml":      5, // Malayalam
// 	"mr":      5, // Marathi
// 	"mn":      0, // Mongolian
// 	"ne":      5, // Nepali
// 	"no":      3, // Norwegian
// 	"fa":      3, // Persian
// 	"pl":      2, // Polish
// 	"pt_BR":   2, // Portuguese (Brazil)
// 	"pa":      5, // Punjabi
// 	"ro":      2, // Romanian
// 	"ru":      2, // Russian
// 	"sr":      3, // Serbian
// 	"si":      5, // Sinhala
// 	"sk":      2, // Slovak
// 	"sl":      3, // Slovenian
// 	"es":      3, // Spanish
// 	"sw":      0, // Swahili
// 	"sv":      3, // Swedish
// 	"ta":      5, // Tamil
// 	"te":      5, // Telugu
// 	"th":      0, // Thai
// 	"tr":      3, // Turkish
// 	"uk":      2, // Ukrainian
// 	"ur":      0, // Urdu
// 	"uz":      3, // Uzbek
// 	"vi":      0, // Vietnamese
// 	"cy":      0, // Welsh
// 	"yo":      0, // Yoruba
// 	"zu":      0, // Zulu
// }
