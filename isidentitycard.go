package validatorgo

import (
	"regexp"
)

var identityCardLocaleRegex = map[string]*regexp.Regexp{
	"LK":    regexp.MustCompile(`^\d{9}[VXvx]$`),              // Sri Lanka
	"PL":    regexp.MustCompile(`^[A-Z]{3}\d{6}$`),            // Poland
	"ES":    regexp.MustCompile(`^\d{8}[A-Z]$`),               // Spain
	"FI":    regexp.MustCompile(`^\d{6}[+\-A]\d{3}[A-Z0-9]$`), // Finland
	"IN":    regexp.MustCompile(`^\d{12}$`),                   // India
	"IT":    regexp.MustCompile(`^[A-Z0-9]{9}$`),              // Italy
	"IR":    regexp.MustCompile(`^\d{10}$`),                   // Iran
	"MZ":    regexp.MustCompile(`^\d{13}$`),                   // Mozambique
	"NO":    regexp.MustCompile(`^\d{11}$`),                   // Norway
	"TH":    regexp.MustCompile(`^\d{13}$`),                   // Thailand
	"zh-TW": regexp.MustCompile(`^[A-Z][12]\d{8}$`),           // Taiwan
	"he-IL": regexp.MustCompile(`^\d{9}$`),                    // Israel
	"ar-LY": regexp.MustCompile(`^\d{12}$`),                   // Libya
	"ar-TN": regexp.MustCompile(`^\d{8}$`),                    // Tunisia
	"zh-CN": regexp.MustCompile(`^\d{17}[\dX]$`),              // China
	"zh-HK": regexp.MustCompile(`^[A-Z]\d{6}[\dA]$`),          // Hong Kong
	"PK":    regexp.MustCompile(`^\d{13}$`),                   // Pakistan
}

// A validator that checks if the string is a valid identity card code.
//
// locale is one of ['LK', 'PL', 'ES', 'FI', 'IN', 'IT', 'IR', 'MZ', 'NO', 'TH', 'zh-TW', 'he-IL', 'ar-LY', 'ar-TN', 'zh-CN', 'zh-HK', 'PK'] OR "any". If "any" is used, function will check if any of the locales match. Defaults to "any" if locale not present. No checksums calculated.
//
//	ok := validatorgo.IsIdentityCard("123456789V", "LK")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsIdentityCard("12345678X", "LK")
//	fmt.Println(ok) // false
func IsIdentityCard(str, locale string) bool {
	if locale == "" {
		locale = "any"
	}

	re, ok := identityCardLocaleRegex[locale]

	if ok {
		return re.MatchString(str)
	} else {
		if locale != "any" {
			return false
		}

		for _, reg := range identityCardLocaleRegex {
			matches := reg.MatchString(str)
			if matches {
				return true
			}
		}
		return false
	}
}
