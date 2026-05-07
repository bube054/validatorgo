package validatorgo

import (
	"regexp"
	"strings"
)

var (
	isAlphaOptsDefaultIgnore string = ""
	isAlphaOptsDefaultLocale string = "en-US"
)

// IsAlphaOpts is used to configure IsAlpha
type IsAlphaOpts struct {
	Ignore string  // string to be ignored
	Locale *string // a locale
}

// escapeRegexChars returns escaped regex special characters
func escapeRegexChars(str string) string {
	var escIgnChars strings.Builder

	for _, char := range str {
		charStr := string(char)
		escChar, ok := regexEscapes[charStr]

		if ok {
			escIgnChars.WriteString(escChar)
		} else {
			escIgnChars.WriteString(charStr)
		}
	}

	return escIgnChars.String()
}

// A validator that checks if the string contains only letters (a-zA-Z).
//
// IsAlphaOpts is an optional struct that can be supplied with the following key(s):
//
// Ignore: is the string to be ignored e.g. " -" will ignore spaces and -'s.
//
// Locale: one of ("ar", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IQ", "ar-JO", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-QA", "ar-QM", "ar-SA", "ar-SD", "ar-SY", "ar-TN", "ar-YE", "bg-BG", "bn", "cs-CZ", "da-DK", "de-DE", "el-GR", "en-AU", "en-GB", "en-HK", "en-IN", "en-NZ", "en-US", "en-ZA", "en-ZM", "eo", "es-ES", "fa-IR", "fi-FI", "fr-CA", "fr-FR", "he", "hi-IN", "hu-HU", "it-IT", "kk-KZ", "ko-KR", "ja-JP", "ku-IQ", "nb-NO", "nl-NL", "nn-NO", "pl-PL", "pt-BR", "pt-PT", "ru-RU", "si-LK", "sl-SI", "sk-SK", "sr-RS", "sr-RS@latin", "sv-SE", "th-TH", "tr-TR", "uk-UA") and defaults to en-US if none is provided.
//
//	ok, _ := validatorgo.IsAlpha("hello", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsAlpha("hello123", nil)
//	fmt.Println(ok) // false
func IsAlpha(str string, opts *IsAlphaOpts) (bool, error) {
	if opts == nil {
		opts = &IsAlphaOpts{}
	}

	opts.mergeDefaults()

	var (
		re                *regexp.Regexp
		lenClsCharFromEnd = 3
	)

	if opts.Ignore == "" && *opts.Locale == "" {
		re = regexp.MustCompile(`^[a-zA-z]+$`)
	}

	if opts.Ignore == "" && *opts.Locale != "" {
		wrtSys, ok := localeWritingSystems[*opts.Locale]
		if !ok {
			return false, newValidationError("IsAlpha", ErrInvalidFormat, "invalid alpha")
		}
		re = regexp.MustCompile(writingSystemAlphaRegex[wrtSys])
	}

	if opts.Ignore != "" && *opts.Locale == "" {
		charsToIgn := escapeRegexChars(opts.Ignore)
		rec := regexp.MustCompile(`^[a-zA-z` + charsToIgn + `]+$`)
		re = rec
	}

	if opts.Ignore != "" && *opts.Locale != "" {
		charsToIgn := escapeRegexChars(opts.Ignore)
		wrtSys := localeWritingSystems[*opts.Locale]
		wrtSysRe := writingSystemAlphaRegex[wrtSys]
		divLen := len(wrtSysRe) - lenClsCharFromEnd
		fstPrtRe, secPrtRe := wrtSysRe[:divLen], wrtSysRe[divLen:]
		rec := regexp.MustCompile(fstPrtRe + charsToIgn + secPrtRe)
		re = rec
	}

	if re.MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsAlpha", ErrInvalidFormat, "invalid alpha")
}

func (o *IsAlphaOpts) mergeDefaults() {
	if o.Locale == nil {
		o.Locale = String("en-US")
	}
}
