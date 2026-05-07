package validatorgo

// regexEscapes is the set of special regex characters escaped
var regexEscapes = map[string]string{
	".":  `\.`,
	"*":  `\*`,
	"+":  `\+`,
	"?":  `\?`,
	"^":  `\^`,
	"$":  `\$`,
	"(":  `\(`,
	")":  `\)`,
	"[":  `\[`,
	"]":  `\]`,
	"{":  `\{`,
	"}":  `\}`,
	"|":  `\|`,
	"\\": `\\`,
	"-":  `\-`,
	"<":  `\<`,
	">":  `\>`,
	`"`:  `\"`,
	`'`:  `\'`,
}

var (
	// writingSystemAlphaRegex is the set of common writing systems and their validating alphabetical regex
	writingSystemAlphaRegex = map[string]string{
		"arabic":          "^[\u0600-\u06FF\u0750-\u077F]+$",              // Arabic
		"latin":           "^[A-Za-z]+$",                                  // Latin
		"latin_czech":     "^[A-Za-zÁČĎÉĚÍŇÓŘŠŤÚŮÝŽáčďéěíňóřšťúůýž]+$",    // Latin with Czech-specific characters
		"latin_danish":    "^[A-Za-zÆØÅæøå]+$",                            // Latin with Danish-specific characters
		"latin_german":    "^[A-Za-zÄÖÜßäöü]+$",                           // Latin with German-specific characters
		"latin_spanish":   "^[A-Za-zÁÉÍÑÓÚÜáéíñóúü]+$",                    // Latin with Spanish-specific characters
		"latin_french":    "^[A-Za-zÀÂÇÉÈÊËÎÏÔŒÙÛÜŸàâçéèêëîïôœùûüÿ]+$",    // Latin with French-specific characters
		"latin_esperanto": "^[A-Za-zĈĜĤĴŜŬĉĝĥĵŝŭ]+$",                      // Latin with Esperanto-specific characters
		"cyrillic":        "^[А-Яа-яЁё]+$",                                // Cyrillic
		"greek":           "^[Α-Ωα-ω]+$",                                  // Greek
		"armenian":        "^[\u0530-\u058F]+$",                           // Armenian
		"bengali":         "^[\u0980-\u09FF]+$",                           // Bengali
		"tibetan":         "^[\u0F00-\u0FFF]+$",                           // Tibetan
		"myanmar":         "^[\u1000-\u109F]+$",                           // Myanmar
		"chinese":         "^[\u4E00-\u9FFF]+$",                           // Chinese
		"hebrew":          "^[\u0590-\u05FF]+$",                           // Hebrew
		"devanagari":      "^[\u0900-\u097F]+$",                           // Devanagari
		"hangul":          "^[\uAC00-\uD7AF\u1100-\u11FF\u3130-\u318F]+$", // Hangul
		"japanese":        "^[\u4E00-\u9FFF\u3040-\u309F\u30A0-\u30FF]+$", // Kanji, Hiragana, Katakana
		"thai":            "^[\u0E00-\u0E7F]+$",                           // Thai
		"ukrainian":       "^[А-Яа-яЁёІЇЄҐіїєґ]+$",                        // Cyrillic with Ukrainian-specific characters
	}

	// localeWritingSystems is the set of locales and their writing systems
	localeWritingSystems = map[string]string{
		"ar":          "arabic",
		"ar-AE":       "arabic",
		"ar-BH":       "arabic",
		"ar-DZ":       "arabic",
		"ar-EG":       "arabic",
		"ar-IQ":       "arabic",
		"ar-JO":       "arabic",
		"ar-KW":       "arabic",
		"ar-LB":       "arabic",
		"ar-LY":       "arabic",
		"ar-MA":       "arabic",
		"ar-QA":       "arabic",
		"ar-QM":       "arabic",
		"ar-SA":       "arabic",
		"ar-SD":       "arabic",
		"ar-SY":       "arabic",
		"ar-TN":       "arabic",
		"ar-YE":       "arabic",
		"bg-BG":       "cyrillic",
		"bn":          "bengali",
		"cs-CZ":       "latin_czech",
		"da-DK":       "latin_danish",
		"de-DE":       "latin_german",
		"el-GR":       "greek",
		"en-AU":       "latin",
		"en-GB":       "latin",
		"en-HK":       "latin",
		"en-IN":       "latin",
		"en-NZ":       "latin",
		"en-US":       "latin",
		"en-ZA":       "latin",
		"en-ZM":       "latin",
		"eo":          "latin_esperanto",
		"es-ES":       "latin_spanish",
		"fa-IR":       "arabic",
		"fi-FI":       "latin_finnish",
		"fr-CA":       "latin_french",
		"fr-FR":       "latin_french",
		"he":          "hebrew",
		"hi-IN":       "devanagari",
		"hu-HU":       "latin_hungarian",
		"it-IT":       "latin_italian",
		"kk-KZ":       "cyrillic",
		"ko-KR":       "hangul",
		"ja-JP":       "japanese",
		"ku-IQ":       "arabic",
		"nb-NO":       "latin_norwegian",
		"nl-NL":       "latin_dutch",
		"nn-NO":       "latin_norwegian",
		"pl-PL":       "latin_polish",
		"pt-BR":       "latin_portuguese",
		"pt-PT":       "latin_portuguese",
		"ru-RU":       "cyrillic",
		"si-LK":       "sinhala",
		"sl-SI":       "latin_slovenian",
		"sk-SK":       "latin_slovak",
		"sr-RS":       "cyrillic",
		"sr-RS@latin": "latin_serbian",
		"sv-SE":       "latin_swedish",
		"th-TH":       "thai",
		"tr-TR":       "latin_turkish",
		"uk-UA":       "ukrainian",
	}
)
