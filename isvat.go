package validatorgo

import "regexp"

// countryCodeVATRegex is the set of country codes and their validating regex
var countryCodeVATRegex = map[string]*regexp.Regexp{
	"AL": regexp.MustCompile(`^[KJ]\d{8}[A-Z]$`),                // Albania
	"AT": regexp.MustCompile(`^U\d{8}$`),                        // Austria
	"AU": regexp.MustCompile(`^\d{11}$`),                        // Australia
	"BE": regexp.MustCompile(`^0\d{9}$`),                        // Belgium
	"BG": regexp.MustCompile(`^\d{9,10}$`),                      // Bulgaria
	"BO": regexp.MustCompile(`^\d{7,8}$`),                       // Bolivia
	"BR": regexp.MustCompile(`^\d{14}$`),                        // Brazil (CNPJ)
	"BY": regexp.MustCompile(`^\d{9}$`),                         // Belarus
	"CA": regexp.MustCompile(`^\d{9}$`),                         // Canada
	"CH": regexp.MustCompile(`^CHE\d{9}(MWST|TVA|IVA)?$`),       // Switzerland
	"CL": regexp.MustCompile(`^\d{8}-\d{1}$`),                   // Chile
	"CO": regexp.MustCompile(`^\d{9}$`),                         // Colombia
	"CR": regexp.MustCompile(`^\d{9}$`),                         // Costa Rica
	"CY": regexp.MustCompile(`^\d{8}[A-Z]$`),                    // Cyprus
	"CZ": regexp.MustCompile(`^\d{8,10}$`),                      // Czech Republic
	"DE": regexp.MustCompile(`^DE\d{9}$`),                       // Germany
	"DK": regexp.MustCompile(`^\d{8}$`),                         // Denmark
	"DO": regexp.MustCompile(`^\d{11}$`),                        // Dominican Republic
	"EC": regexp.MustCompile(`^\d{13}$`),                        // Ecuador
	"EE": regexp.MustCompile(`^\d{9}$`),                         // Estonia
	"ES": regexp.MustCompile(`^[A-Z]\d{7}[A-Z]$`),               // Spain
	"FI": regexp.MustCompile(`^FI\d{8}$`),                       // Finland
	"FR": regexp.MustCompile(`^[A-Z0-9]{2}\d{9}$`),              // France
	"GB": regexp.MustCompile(`^\d{9}$|^\d{12}$|^(GD|HA)\d{3}$`), // United Kingdom
	"GT": regexp.MustCompile(`^\d{8}$`),                         // Guatemala
	"HN": regexp.MustCompile(`^\d{14}$`),                        // Honduras
	"HR": regexp.MustCompile(`^\d{11}$`),                        // Croatia
	"HU": regexp.MustCompile(`^\d{8}$`),                         // Hungary
	"ID": regexp.MustCompile(`^\d{15}$`),                        // Indonesia
	"IE": regexp.MustCompile(`^\d{7}[A-Z]{1,2}$`),               // Ireland
	"IL": regexp.MustCompile(`^\d{9}$`),                         // Israel
	"IN": regexp.MustCompile(`^\d{11}$`),                        // India
	"IS": regexp.MustCompile(`^\d{5,6}$`),                       // Iceland
	"IT": regexp.MustCompile(`^IT\d{11}$`),                      // Italy
	"KZ": regexp.MustCompile(`^\d{12}$`),                        // Kazakhstan
	"LT": regexp.MustCompile(`^LT\d{9}$|^LT\d{12}$`),            // Lithuania
	"LU": regexp.MustCompile(`^\d{8}$`),                         // Luxembourg
	"LV": regexp.MustCompile(`^\d{11}$`),                        // Latvia
	"MK": regexp.MustCompile(`^\d{13}$`),                        // North Macedonia
	"MT": regexp.MustCompile(`^\d{8}$`),                         // Malta
	"MX": regexp.MustCompile(`^[A-Z]{4}\d{6}[A-Z0-9]{3}$`),      // Mexico
	"NG": regexp.MustCompile(`^\d{8,9}$`),                       // Nigeria
	"NI": regexp.MustCompile(`^\d{3}-\d{6}-\d{3}$`),             // Nicaragua
	"NL": regexp.MustCompile(`^NL\d{9}B\d{2}$`),                 // Netherlands
	"NO": regexp.MustCompile(`^\d{9}MVA$`),                      // Norway
	"NZ": regexp.MustCompile(`^\d{9}$`),                         // New Zealand
	"PA": regexp.MustCompile(`^\d{8}$`),                         // Panama
	"PE": regexp.MustCompile(`^\d{11}$`),                        // Peru
	"PH": regexp.MustCompile(`^\d{12}$`),                        // Philippines
	"PL": regexp.MustCompile(`^PL\d{10}$`),                      // Poland
	"PT": regexp.MustCompile(`^PT\d{9}$`),                       // Portugal
	"PY": regexp.MustCompile(`^\d{6,8}-\d{1}$`),                 // Paraguay
	"RO": regexp.MustCompile(`^RO\d{2,10}$`),                    // Romania
	"RS": regexp.MustCompile(`^\d{9}$`),                         // Serbia
	"RU": regexp.MustCompile(`^\d{10}$`),                        // Russia
	"SA": regexp.MustCompile(`^\d{15}$`),                        // Saudi Arabia
	"SE": regexp.MustCompile(`^\d{12}$`),                        // Sweden
	"SI": regexp.MustCompile(`^\d{8}$`),                         // Slovenia
	"SK": regexp.MustCompile(`^\d{10}$`),                        // Slovakia
	"SM": regexp.MustCompile(`^\d{5}$`),                         // San Marino
	"SV": regexp.MustCompile(`^\d{14}$`),                        // El Salvador
	"TR": regexp.MustCompile(`^\d{10}$`),                        // Turkey
	"UA": regexp.MustCompile(`^\d{12}$`),                        // Ukraine
	"UY": regexp.MustCompile(`^\d{12}$`),                        // Uruguay
	"UZ": regexp.MustCompile(`^\d{9}$`),                         // Uzbekistan
	"VE": regexp.MustCompile(`^[JGVE]-\d{8}-\d{1}$`),            // Venezuela
}

// A validator that checks if the string is a valid VAT number if validation is available for the given country code matching ISO 3166-1 alpha-2.
//
// countryCode is one of ("AL", "AR", "AT", "AU", "BE", "BG", "BO", "BR", "BY", "CA", "CH", "CL", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "EC", "EE", "EL", "ES", "FI", "FR", "GB", "GT", "HN", "HR", "HU", "ID", "IE", "IL", "IN", "IS", "IT", "KZ", "LT", "LU", "LV", "MK", "MT", "MX", "NG", "NI", "NL", "NO", "NZ", "PA", "PE", "PH", "PL", "PT", "PY", "RO", "RS", "RU", "SA", "SE", "SI", "SK", "SM", "SV", "TR", "UA", "UY", "UZ", "VE"). If no value or "any" is provided will match every one.
//
//	ok := validatorgo.IsVAT("DE123456789", "DE")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsVAT("DE12345678", "DE")
//	fmt.Println(ok) // false
func IsVAT(str, countryCode string) bool {
	re, ok := countryCodeVATRegex[countryCode]

	if !ok {
		if countryCode != "" && countryCode != "any" {
			return false
		}

		for _, reg := range countryCodeVATRegex {
			match := reg.MatchString(str)

			if match {
				return true
			}
		}

		return false
	}

	return re.MatchString(str)
}
