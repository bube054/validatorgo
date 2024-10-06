package validatorgo

import (
	"regexp"
)

var (
	isNumericOptsDefaultNoSymbols bool   = false
	isNumericOptsDefaultLocale    string = "en-US"
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
	// 8: func(opts IsNumericOpts) *regexp.Regexp {
	// 	if opts.NoSymbols {
	// 		return regexp.MustCompile(`^(\d{1,3})(٬\d{3})*(٫\d+)?$`) // Without symbols
	// 	}
	// 	return regexp.MustCompile(`^[+-]?(\d{1,3})(٬\d{3})*(٫\d+)?$`) // With symbols
	// }, // matches ١٬٢٣٤٬٥٦٧٫٨٩ (Arabic: comma-like for thousands, period-like for decimals)
	8: func(opts IsNumericOpts) *regexp.Regexp {
		if opts.NoSymbols {
			return regexp.MustCompile(`^[\x{0660}-\x{0669}]{1,3}(٬[\x{0660}-\x{0669}]{3})*(٫[\x{0660}-\x{0669}]+)?$`) // Without symbols
		}
		return regexp.MustCompile(`^[+-]?[\x{0660}-\x{0669}]{1,3}(٬[\x{0660}-\x{0669}]{3})*(٫[\x{0660}-\x{0669}]+)?$`) // With symbols
	},

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

// A validator that check if a string is a number.
//
// IsNumericOpts is a struct which defaults to { NoSymbols: false, Locale: ""}.
//
// If NoSymbols is true, the validator will reject numeric strings that feature a symbol (e.g. +, -, or .).
//
// Locale determines the numeric format and is one of ("ar", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IQ", "ar-JO", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-QA", "ar-QM", "ar-SA", "ar-SD", "ar-SY", "ar-TN", "ar-YE", "bg-BG", "cs-CZ", "da-DK", "de-DE", "en-AU", "en-GB", "en-HK", "en-IN", "en-NZ", "en-US", "en-ZA", "en-ZM", "eo", "es-ES", "fr-FR", "fr-CA", "hu-HU", "it-IT", "nb-NO", "nl-NL", "nn-NO", "pl-PL", "pt-BR", "pt-PT", "ru-RU", "sl-SI", "sr-RS", "sr-RS@latin", "sv-SE", "tr-TR", "uk-UA"). Locale will default to "en-US" if not present.
//
//	ok = validatorgo.IsNumeric("12345", &validatorgo.IsNumericOpts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsNumeric("12.34.56", &validatorgo.IsNumericOpts{})
//	fmt.Println(ok) // false
func IsNumeric(str string, opts *IsNumericOpts) bool {
	if opts == nil {
		opts = setIsNumericOptsToDefault()
	}

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
		opts.Locale = isNumericOptsDefaultLocale
	}

	// Locale is present and plus or minus symbols are optional(NoSymbol does not matter)
	codeNumForm := codeNumericFormats[opts.Locale]
	valFunc := numericFormatsRegex[codeNumForm]
	re = valFunc(*opts)

	return re.MatchString(str)
}

func setIsNumericOptsToDefault() *IsNumericOpts {
	return &IsNumericOpts{
		NoSymbols: isNumericOptsDefaultNoSymbols,
		Locale:    isNumericOptsDefaultLocale,
	}
}
