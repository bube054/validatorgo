package validatorgo

import (
	"fmt"
	"regexp"
)

var (
	isAlphanumericOptsDefaultIgnore string = ""
	isAlphanumericOptsDefaultLocale string = "en-US"
)

// IsAlphanumericOpts is used to configure IsAlphanumeric
type IsAlphanumericOpts struct {
	Ignore string // string to be ignored
	Locale string // a locale
}

// writingSystemAlphaNumRegex is the set of common writing systems and their validating alphanumeric regex
var writingSystemAlphaNumRegex = map[string]string{
	"arabic":          "^[A-Za-z0-9\u0600-\u06FF\u0750-\u077F]+$",              // Alphanumeric + Arabic
	"latin":           "^[A-Za-z0-9]+$",                                        // Alphanumeric (Latin)
	"latin_czech":     "^[A-Za-z0-9ÁČĎÉĚÍŇÓŘŠŤÚŮÝŽáčďéěíňóřšťúůýž]+$",          // Alphanumeric + Latin with Czech-specific characters
	"latin_danish":    "^[A-Za-z0-9ÆØÅæøå]+$",                                  // Alphanumeric + Latin with Danish-specific characters
	"latin_german":    "^[A-Za-z0-9ÄÖÜßäöü]+$",                                 // Alphanumeric + Latin with German-specific characters
	"latin_spanish":   "^[A-Za-z0-9ÁÉÍÑÓÚÜáéíñóúü]+$",                          // Alphanumeric + Latin with Spanish-specific characters
	"latin_french":    "^[A-Za-z0-9ÀÂÇÉÈÊËÎÏÔŒÙÛÜŸàâçéèêëîïôœùûüÿ]+$",          // Alphanumeric + Latin with French-specific characters
	"latin_esperanto": "^[A-Za-z0-9ĈĜĤĴŜŬĉĝĥĵŝŭ]+$",                            // Alphanumeric + Latin with Esperanto-specific characters
	"cyrillic":        "^[A-Za-z0-9А-Яа-яЁё]+$",                                // Alphanumeric + Cyrillic
	"greek":           "^[A-Za-z0-9Α-Ωα-ω]+$",                                  // Alphanumeric + Greek
	"armenian":        "^[A-Za-z0-9\u0530-\u058F]+$",                           // Alphanumeric + Armenian
	"bengali":         "^[A-Za-z0-9\u0980-\u09FF]+$",                           // Alphanumeric + Bengali
	"tibetan":         "^[A-Za-z0-9\u0F00-\u0FFF]+$",                           // Alphanumeric + Tibetan
	"myanmar":         "^[A-Za-z0-9\u1000-\u109F]+$",                           // Alphanumeric + Myanmar
	"chinese":         "^[A-Za-z0-9\u4E00-\u9FFF]+$",                           // Alphanumeric + Chinese
	"hebrew":          "^[A-Za-z0-9\u0590-\u05FF]+$",                           // Alphanumeric + Hebrew
	"devanagari":      "^[A-Za-z0-9\u0900-\u097F]+$",                           // Alphanumeric + Devanagari
	"hangul":          "^[A-Za-z0-9\uAC00-\uD7AF\u1100-\u11FF\u3130-\u318F]+$", // Alphanumeric + Hangul
	"japanese":        "^[A-Za-z0-9\u4E00-\u9FFF\u3040-\u309F\u30A0-\u30FF]+$", // Alphanumeric + Kanji, Hiragana, Katakana
	"thai":            "^[A-Za-z0-9\u0E00-\u0E7F]+$",                           // Alphanumeric + Thai
	"ukrainian":       "^[A-Za-z0-9А-Яа-яЁёІЇЄҐіїєґ]+$",                        // Alphanumeric + Cyrillic with Ukrainian-specific characters
}

// A validator that checks if the string contains only letters and numbers (a-zA-Z0-9).
//
// IsAlphaOpts is an optional struct that can be supplied with the following key(s):
//
// Ignore: is the string to be ignored e.g. " -" will ignore spaces and -'s.
//
// locale is one of ('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bn', 'bg-BG', 'cs-CZ', 'da-DK', 'de-DE', 'el-GR', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fa-IR', 'fi-FI', 'fr-CA', 'fr-FR', 'he', 'hi-IN', 'hu-HU', 'it-IT', 'kk-KZ', 'ko-KR', 'ja-JP','ku-IQ', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-BR', 'pt-PT', 'ru-RU', 'si-LK', 'sl-SI', 'sk-SK', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'th-TH', 'tr-TR', 'uk-UA') and defaults to en-US.
//
//	isAlpha := validatorgo.IsAlphanumeric("hello123", &validatorgo.IsAlphanumericOpts{})
//	fmt.Println(isAlpha) // true
//	isAlpha := validatorgo.IsAlphanumeric("hello!", &validatorgo.IsAlphanumericOpts{})
//	fmt.Println(isAlpha) // false
func IsAlphanumeric(str string, opts *IsAlphanumericOpts) bool {
	if opts == nil {
		opts = setIsAlphanumericOptsToDefault()
	}

	var (
		re                *regexp.Regexp
		lenClsCharFromEnd = 3
	)

	if opts.Ignore == "" && opts.Locale == "" {
		re = regexp.MustCompile(`^[a-zA-z0-9]+$`)
	}

	if opts.Ignore == "" && opts.Locale != "" {
		wrtSys, ok := localeWritingSystems[opts.Locale]
		if !ok {
			return false
		}
		re = regexp.MustCompile(writingSystemAlphaNumRegex[wrtSys])
	}

	if opts.Ignore != "" && opts.Locale == "" {
		charsToIgn := escapeRegexChars(opts.Ignore)
		rec := regexp.MustCompile(`^[a-zA-z0-9` + charsToIgn + `]+$`)
		re = rec
	}

	if opts.Ignore != "" && opts.Locale != "" {
		charsToIgn := escapeRegexChars(opts.Ignore)
		wrtSys := localeWritingSystems[opts.Locale]
		wrtSysRe := writingSystemAlphaNumRegex[wrtSys]
		divLen := len(wrtSysRe) - lenClsCharFromEnd
		fstPrtRe, secPrtRe := wrtSysRe[:divLen], wrtSysRe[divLen:]
		rec := regexp.MustCompile(fstPrtRe + charsToIgn + secPrtRe)
		fmt.Println(rec.String())
		re = rec
	}

	return re.MatchString(str)
}

func setIsAlphanumericOptsToDefault() *IsAlphanumericOpts {
	return &IsAlphanumericOpts{
		Ignore: isAlphanumericOptsDefaultIgnore,
		Locale: isAlphanumericOptsDefaultLocale,
	}
}
