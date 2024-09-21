package validatorgo

import "regexp"

// countryCodePassportNumberRegex is the set of country codes and their validating regex
var countryCodePassportNumberRegex = map[string]*regexp.Regexp{
	"AM": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Armenia
	"AR": regexp.MustCompile(`^[A-Z]{3}\d{6}$`),      // Argentina
	"AT": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Austria
	"AU": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Australia
	"AZ": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Azerbaijan
	"BE": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Belgium
	"BG": regexp.MustCompile(`^\d{9}$`),              // Bulgaria
	"BY": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Belarus
	"BR": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Brazil
	"CA": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Canada
	"CH": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Switzerland
	"CN": regexp.MustCompile(`^G\d{8}$|^E\d{8}$`),    // China
	"CY": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Cyprus
	"CZ": regexp.MustCompile(`^\d{8}$`),              // Czech Republic
	"DE": regexp.MustCompile(`^[C-FH]\d{8}$`),        // Germany
	"DK": regexp.MustCompile(`^\d{9}$`),              // Denmark
	"DZ": regexp.MustCompile(`^\d{9}$`),              // Algeria
	"EE": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Estonia
	"ES": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Spain
	"FI": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Finland
	"FR": regexp.MustCompile(`^\d{2}[A-Z]{2}\d{5}$`), // France
	"GB": regexp.MustCompile(`^\d{9}$`),              // United Kingdom
	"GR": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Greece
	"HR": regexp.MustCompile(`^\d{9}$`),              // Croatia
	"HU": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Hungary
	"IE": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Ireland
	"IN": regexp.MustCompile(`^[A-Z]\d{7}$`),         // India
	"IR": regexp.MustCompile(`^[A-Z]\d{8}$`),         // Iran
	"ID": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Indonesia
	"IS": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Iceland
	"IT": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Italy
	"JM": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Jamaica
	"JP": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Japan
	"KR": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // South Korea
	"KZ": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Kazakhstan
	"LI": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Liechtenstein
	"LT": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Lithuania
	"LU": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Luxembourg
	"LV": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Latvia
	"LY": regexp.MustCompile(`^\d{8}$`),              // Libya
	"MT": regexp.MustCompile(`^\d{7}$`),              // Malta
	"MX": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Mexico
	"MY": regexp.MustCompile(`^[A-Z]\d{8}$`),         // Malaysia
	"MZ": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Mozambique
	"NL": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Netherlands
	"NZ": regexp.MustCompile(`^[A-Z]\d{7}$`),         // New Zealand
	"PH": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Philippines
	"PK": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Pakistan
	"PL": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Poland
	"PT": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Portugal
	"RO": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Romania
	"RU": regexp.MustCompile(`^\d{9}$`),              // Russia
	"SE": regexp.MustCompile(`^\d{8}$`),              // Sweden
	"SL": regexp.MustCompile(`^\d{8}$`),              // Sierra Leone
	"SK": regexp.MustCompile(`^\d{8}$`),              // Slovakia
	"TH": regexp.MustCompile(`^[A-Z]\d{7}$`),         // Thailand
	"TR": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // Turkey
	"UA": regexp.MustCompile(`^[A-Z]{2}\d{6}$`),      // Ukraine
	"US": regexp.MustCompile(`^\d{9}$`),              // United States
	"ZA": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),      // South Africa
}

// A validator that checks if the string is a valid passport number.
//
// countryCode is one of ("AM", "AR", "AT", "AU", "AZ", "BE", "BG", "BY", "BR", "CA", "CH", "CN", "CY", "CZ", "DE", "DK", "DZ", "EE", "ES", "FI", "FR", "GB", "GR", "HR", "HU", "IE", "IN", "IR", "ID", "IS", "IT", "JM", "JP", "KR", "KZ", "LI", "LT", "LU", "LV", "LY", "MT", "MX", "MY", "MZ", "NL", "NZ", "PH", "PK", "PL", "PT", "RO", "RU", "SE", "SL", "SK", "TH", "TR", "UA", "US", "ZA").
//
//	ok := validatorgo.IsPassportNumber("123456789", "US")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsPassportNumber("A12345678", "US")
//	fmt.Println(ok) // false
func IsPassportNumber(str, countryCode string) bool {
	re, exists := countryCodePassportNumberRegex[countryCode]

	if !exists {
		return false
	}

	return re.MatchString(str)
}
