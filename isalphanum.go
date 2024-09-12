package validatorgo

import (
	"regexp"
)

// IsAlphanumericOpts is used to configure IsAlpha
type IsAlphanumericOpts struct {
	Ignore string
	Locale string
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
// locale is one of ('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'bn', 'cs-CZ', 'da-DK', 'de-DE', 'el-GR', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fa-IR', 'fi-FI', 'fr-CA', 'fr-FR', 'he', 'hi-IN', 'hu-HU', 'it-IT', 'kk-KZ', 'ko-KR', 'ja-JP', 'ku-IQ', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-BR', 'pt-PT', 'ru-RU', 'si-LK', 'sl-SI', 'sk-SK', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'th-TH', 'tr-TR', 'uk-UA') and defaults to en-US. Locale list is validator.isAlphaLocales. options is an optional object that can be supplied with the following key(s): ignore is the string to be ignored e.g. " -" will ignore spaces and -'s.
//
//	isAlpha := govalidator.IsAlphanumeric("helloWORLD", govalidator.IsAlphanumericOpts{})
//	fmt.Println(isAlpha) // true
//	isAlpha := govalidator.IsAlphanumeric("7889gygu", govalidator.IsAlphanumericOpts{Locale: "bg-BG"})
//	fmt.Println(isAlpha) // true
//	isAlpha := govalidator.IsAlphanumeric("1Hello23", govalidator.IsAlphanumericOpts{Ignore: "^"})
//	fmt.Println(isAlpha) // true
//	isAlpha := govalidator.IsAlphanumeric("Привет89a", govalidator.IsAlphanumericOpts{Ignore: "89a", Locale: "bg-BG"})
//	fmt.Println(isAlpha) // true
func IsAlphanumeric(str string, opts IsAlphanumericOpts) bool {
	var (
		re                *regexp.Regexp
		defLocWrtSys      = "latin"
		lenClsCharFromEnd = 3
	)

	if opts.Ignore == "" && opts.Locale == "" {
		re = regexp.MustCompile(`^[a-zA-z0-9]+$`)
	}

	if opts.Ignore == "" && opts.Locale != "" {
		wrtSys, ok := localeWritingSystems[opts.Locale]
		if !ok {
			wrtSys = defLocWrtSys
		}
		re = regexp.MustCompile(writingSystemAlphaNumRegex[wrtSys])
	}

	if opts.Ignore != "" && opts.Locale == "" {
		charsToIgn := escapeRegexChars(opts.Ignore)
		rec, err := regexp.Compile(`^[a-zA-z0-9` + charsToIgn + `]+$`)
		if err != nil {
			return false
		}
		re = rec
	}

	if opts.Ignore != "" && opts.Locale != "" {
		charsToIgn := escapeRegexChars(opts.Ignore)
		wrtSys, ok := localeWritingSystems[opts.Locale]
		if !ok {
			wrtSys = defLocWrtSys
		}
		wrtSysRe := writingSystemAlphaNumRegex[wrtSys]
		divLen := len(wrtSysRe) - lenClsCharFromEnd
		fstPrtRe, secPrtRe := wrtSysRe[:divLen], wrtSysRe[divLen:]
		rec, err := regexp.Compile(fstPrtRe + charsToIgn + secPrtRe)
		if err != nil {
			return false
		}
		re = rec
	}

	return re.MatchString(str)
}
