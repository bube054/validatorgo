package validatorgo

import (
	"regexp"
)

// localePostalCodeRegex is the set of locales and postal code regex
var localePostalCodeRegex = map[string]*regexp.Regexp{
	"AD": regexp.MustCompile(`^AD\d{3}$`),                                  // Andorra
	"AT": regexp.MustCompile(`^\d{4}$`),                                    // Austria
	"AU": regexp.MustCompile(`^\d{4}$`),                                    // Australia
	"AZ": regexp.MustCompile(`^\d{4}$`),                                    // Azerbaijan
	"BA": regexp.MustCompile(`^\d{5}$`),                                    // Bosnia and Herzegovina
	"BE": regexp.MustCompile(`^\d{4}$`),                                    // Belgium
	"BG": regexp.MustCompile(`^\d{4}$`),                                    // Bulgaria
	"BR": regexp.MustCompile(`^\d{5}-\d{3}$`),                              // Brazil
	"BY": regexp.MustCompile(`^\d{6}$`),                                    // Belarus
	"CA": regexp.MustCompile(`^[A-Za-z]\d[A-Za-z] \d[A-Za-z]\d$`),          // Canada
	"CH": regexp.MustCompile(`^\d{4}$`),                                    // Switzerland
	"CN": regexp.MustCompile(`^\d{6}$`),                                    // China
	"CO": regexp.MustCompile(`^\d{6}$`),                                    // Colombia
	"CZ": regexp.MustCompile(`^\d{3} \d{2}$`),                              // Czech Republic
	"DE": regexp.MustCompile(`^\d{5}$`),                                    // Germany
	"DK": regexp.MustCompile(`^\d{4}$`),                                    // Denmark
	"DO": regexp.MustCompile(`^\d{5}$`),                                    // Dominican Republic
	"DZ": regexp.MustCompile(`^\d{5}$`),                                    // Algeria
	"EE": regexp.MustCompile(`^\d{5}$`),                                    // Estonia
	"ES": regexp.MustCompile(`^\d{5}$`),                                    // Spain
	"FI": regexp.MustCompile(`^\d{5}$`),                                    // Finland
	"FR": regexp.MustCompile(`^\d{5}$`),                                    // France
	"GB": regexp.MustCompile(`^[A-Za-z]{1,2}\d[A-Za-z\d]? \d[A-Za-z]{2}$`), // United Kingdom
	"GR": regexp.MustCompile(`^\d{3} \d{2}$`),                              // Greece
	"HR": regexp.MustCompile(`^\d{5}$`),                                    // Croatia
	"HT": regexp.MustCompile(`^\d{4}$`),                                    // Haiti
	"HU": regexp.MustCompile(`^\d{4}$`),                                    // Hungary
	"ID": regexp.MustCompile(`^\d{5}$`),                                    // Indonesia
	"IE": regexp.MustCompile(`^[A-Za-z]\d{2}[A-Za-z0-9]{4}$`),              // Ireland
	"IL": regexp.MustCompile(`^\d{7}$`),                                    // Israel
	"IN": regexp.MustCompile(`^\d{6}$`),                                    // India
	"IR": regexp.MustCompile(`^\d{5}-\d{5}$`),                              // Iran
	"IS": regexp.MustCompile(`^\d{3}$`),                                    // Iceland
	"IT": regexp.MustCompile(`^\d{5}$`),                                    // Italy
	"JP": regexp.MustCompile(`^\d{3}-\d{4}$`),                              // Japan
	"KE": regexp.MustCompile(`^\d{5}$`),                                    // Kenya
	"KR": regexp.MustCompile(`^\d{5}$`),                                    // South Korea
	"LI": regexp.MustCompile(`^\d{4}$`),                                    // Liechtenstein
	"LK": regexp.MustCompile(`^\d{5}$`),                                    // Sri Lanka
	"LT": regexp.MustCompile(`^LT-\d{5}$`),                                 // Lithuania
	"LU": regexp.MustCompile(`^\d{4}$`),                                    // Luxembourg
	"LV": regexp.MustCompile(`^LV-\d{4}$`),                                 // Latvia
	"MG": regexp.MustCompile(`^\d{3}$`),                                    // Madagascar
	"MT": regexp.MustCompile(`^[A-Za-z]{3} \d{4}$`),                        // Malta
	"MX": regexp.MustCompile(`^\d{5}$`),                                    // Mexico
	"MY": regexp.MustCompile(`^\d{5}$`),                                    // Malaysia
	"NL": regexp.MustCompile(`^\d{4} [A-Za-z]{2}$`),                        // Netherlands
	"NO": regexp.MustCompile(`^\d{4}$`),                                    // Norway
	"NP": regexp.MustCompile(`^\d{5}$`),                                    // Nepal
	"NZ": regexp.MustCompile(`^\d{4}$`),                                    // New Zealand
	"PL": regexp.MustCompile(`^\d{2}-\d{3}$`),                              // Poland
	"PR": regexp.MustCompile(`^\d{5}(-\d{4})?$`),                           // Puerto Rico
	"PT": regexp.MustCompile(`^\d{4}-\d{3}$`),                              // Portugal
	"RO": regexp.MustCompile(`^\d{6}$`),                                    // Romania
	"RU": regexp.MustCompile(`^\d{6}$`),                                    // Russia
	"SA": regexp.MustCompile(`^\d{5}$`),                                    // Saudi Arabia
	"SE": regexp.MustCompile(`^\d{3} \d{2}$`),                              // Sweden
	"SG": regexp.MustCompile(`^\d{6}$`),                                    // Singapore
	"SI": regexp.MustCompile(`^\d{4}$`),                                    // Slovenia
	"SK": regexp.MustCompile(`^\d{3} \d{2}$`),                              // Slovakia
	"TH": regexp.MustCompile(`^\d{5}$`),                                    // Thailand
	"TN": regexp.MustCompile(`^\d{4}$`),                                    // Tunisia
	"TW": regexp.MustCompile(`^\d{3}(?:-\d{2})?$`),                         // Taiwan
	"UA": regexp.MustCompile(`^\d{5}$`),                                    // Ukraine
	"US": regexp.MustCompile(`^\d{5}(-\d{4})?$`),                           // United States
	"ZA": regexp.MustCompile(`^\d{4}$`),                                    // South Africa
	"ZM": regexp.MustCompile(`^\d{5}$`),                                    // Zambia
}

// A validator that checks if the string is a postal code.
//
// locale is one of ("AD", "AT", "AU", "AZ", "BA", "BE", "BG", "BR", "BY", "CA", "CH", "CN", "CO", "CZ", "DE", "DK", "DO", "DZ", "EE", "ES", "FI", "FR", "GB", "GR", "HR", "HT", "HU", "ID", "IE", "IL", "IN", "IR", "IS", "IT", "JP", "KE", "KR", "LI", "LK", "LT", "LU", "LV", "MG", "MT", "MX", "MY", "NL", "NO", "NP", "NZ", "PL", "PR", "PT", "RO", "RU", "SA", "SE", "SG", "SI", "SK", "TH", "TN", "TW", "UA", "US", "ZA", "ZM"). If "any" or no locale is used, function will check if any of the locales match.
//
//	ok := validatorgo.IsPostalCode("90210", "US")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsPostalCode("902101", "DE")
//	fmt.Println(ok) // false
func IsPostalCode(str, locale string) bool {
	re, exists := localePostalCodeRegex[locale]

	if !exists {
		if locale != "" && locale != "any" {
			return false
		}

		for _, reg := range localePostalCodeRegex {
			match := reg.MatchString(str)

			if match {
				return true
			}
		}

		return false
	}

	return re.MatchString(str)
}
