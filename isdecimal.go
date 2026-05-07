package validatorgo

var (
	isDecimalOptsDefaultForceDecimal bool   = false
	isDecimalOptsDefaultLocale       string = "en-US"

	isDecimalOptsDefaultMin uint  = 0
	isDecimalOptsDefaultMax *uint = nil
)

// DecimalDigits is used to configure IsDecimalOpts
type DecimalDigits struct {
	Min *uint // minimum allowed decimal range
	Max *uint // maximum allowed decimal range
}

// IsDecimalOpts is used to configure IsDecimal
type IsDecimalOpts struct {
	DecimalDigits

	ForceDecimal *bool   // decimal/radix point must be present
	Locale       *string // locale used
}


// A validator that check if the string represents a decimal number, such as 0.1, .3, 1.1, 1.00003, 4.0, etc.
//
// IsDecimalOpts is a struct which defaults to {ForceDecimal: false, DecimalDigits: {Min: 0, Max: nil}, locale: 'en-US'}.
//
// locale: determines the decimal separator and is one of ("ar", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IQ", "ar-JO", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-QA", "ar-QM", "ar-SA", "ar-SD", "ar-SY", "ar-TN", "ar-YE", "bg-BG", "cs-CZ", "da-DK", "de-DE", "el-GR", "en-AU", "en-GB", "en-HK", "en-IN", "en-NZ", "en-US", "en-ZA", "en-ZM", "eo", "es-ES", "fa", "fa-AF", "fa-IR", "fr-FR", "fr-CA", "hu-HU", "id-ID", "it-IT", "ku-IQ", "nb-NO", "nl-NL", "nn-NO", "pl-PL", "pl-Pl", "pt-BR", "pt-PT", "ru-RU", "sl-SI", "sr-RS", "sr-RS@latin", "sv-SE", "tr-TR", "uk-UA", "vi-VN"). Locale defaults to "en-US"
//
// ForceDecimal: simply means a decimal must be present "123" will not pass but "123.45" will pass
//
// DecimalDigits: is the allowed decimal range e.g {Min: 3, Max: nil} 123.456
//
//	ok, _ := validatorgo.IsDecimal("123", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsDecimal("abc", nil)
//	fmt.Println(ok) // false
func IsDecimal(str string, opts *IsDecimalOpts) (bool, error) {
	if opts == nil {
		opts = &IsDecimalOpts{}
	}

	opts.mergeDefaults()

	format, ok := localeDecimalFormats[*opts.Locale]

	if !ok {
		return false, newValidationError("IsDecimal", ErrInvalidFormat, "invalid decimal")
	}

	validReFunc := decimalFormats[format]
	re, err := validReFunc(*opts)

	if err != nil {
		return false, newValidationError("IsDecimal", ErrInvalidFormat, "invalid decimal")
	}

	if re.MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsDecimal", ErrInvalidFormat, "invalid decimal")
}

func (o *IsDecimalOpts) mergeDefaults() {
	if o.ForceDecimal == nil {
		o.ForceDecimal = Bool(false)
	}
	if o.DecimalDigits.Min == nil {
		o.DecimalDigits.Min = Uint(0)
	}
	if o.Locale == nil {
		o.Locale = String("en-US")
	}
}
