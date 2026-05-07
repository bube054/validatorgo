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

// A validator that check if a string is a number.
//
// IsNumericOpts is a struct which defaults to { NoSymbols: false, Locale: ""}.
//
// If NoSymbols is true, the validator will reject numeric strings that feature a symbol (e.g. +, -, or .).
//
// Locale determines the numeric format and is one of ("ar", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IQ", "ar-JO", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-QA", "ar-QM", "ar-SA", "ar-SD", "ar-SY", "ar-TN", "ar-YE", "bg-BG", "cs-CZ", "da-DK", "de-DE", "en-AU", "en-GB", "en-HK", "en-IN", "en-NZ", "en-US", "en-ZA", "en-ZM", "eo", "es-ES", "fr-FR", "fr-CA", "hu-HU", "it-IT", "nb-NO", "nl-NL", "nn-NO", "pl-PL", "pt-BR", "pt-PT", "ru-RU", "sl-SI", "sr-RS", "sr-RS@latin", "sv-SE", "tr-TR", "uk-UA"). Locale will default to "en-US" if not present.
//
//	ok, _ := validatorgo.IsNumeric("12345", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsNumeric("12.34.56", nil)
//	fmt.Println(ok) // false
func IsNumeric(str string, opts *IsNumericOpts) (bool, error) {
	if opts == nil {
		opts = setIsNumericOptsToDefault()
	}

	var re *regexp.Regexp

	// has symbols and no Locale
	if !opts.NoSymbols && opts.Locale == "" {
		re = regexp.MustCompile(`^[+-]?\d+(\.\d+)?$`)
		if re.MatchString(str) {
			return true, nil
		}
		return false, newValidationError("IsNumeric", ErrInvalidFormat, "invalid numeric")
	}

	// no symbols and no Locale
	if opts.NoSymbols && opts.Locale == "" {
		re = regexp.MustCompile(`^\d+$`)
		if re.MatchString(str) {
			return true, nil
		}
		return false, newValidationError("IsNumeric", ErrInvalidFormat, "invalid numeric")
	}

	if opts.Locale == "" {
		opts.Locale = isNumericOptsDefaultLocale
	}

	// Locale is present and plus or minus symbols are optional(NoSymbol does not matter)
	codeNumForm := codeNumericFormats[opts.Locale]
	valFunc := numericFormatsRegex[codeNumForm]
	re = valFunc(*opts)

	if re.MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsNumeric", ErrInvalidFormat, "invalid numeric")
}

func setIsNumericOptsToDefault() *IsNumericOpts {
	return &IsNumericOpts{
		NoSymbols: isNumericOptsDefaultNoSymbols,
		Locale:    isNumericOptsDefaultLocale,
	}
}
