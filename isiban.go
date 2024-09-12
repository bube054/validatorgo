package govalidator

import "regexp"

// iban regex's for country codes
var ibanRegex = map[string]*regexp.Regexp{
	"AD": regexp.MustCompile(`^AD\d{2}\d{4}\d{4}\d{12}$`),               // Andorra
	"AE": regexp.MustCompile(`^AE\d{2}\d{3}\d{16}$`),                    // United Arab Emirates
	"AL": regexp.MustCompile(`^AL\d{2}\d{8}\d{16}$`),                    // Albania
	"AT": regexp.MustCompile(`^AT\d{2}\d{5}\d{11}$`),                    // Austria
	"AZ": regexp.MustCompile(`^AZ\d{2}[A-Z]{4}\d{20}$`),                 // Azerbaijan
	"BA": regexp.MustCompile(`^BA\d{2}\d{3}\d{3}\d{8}\d{2}$`),           // Bosnia and Herzegovina
	"BE": regexp.MustCompile(`^BE\d{2}\d{3}\d{7}\d{2}$`),                // Belgium
	"BG": regexp.MustCompile(`^BG\d{2}[A-Z]{4}\d{4}\d{10}$`),            // Bulgaria
	"BH": regexp.MustCompile(`^BH\d{2}[A-Z]{4}\d{14}$`),                 // Bahrain
	"BR": regexp.MustCompile(`^BR\d{2}\d{8}\d{5}\d{10}[A-Z]{1}\d{1}$`),  // Brazil
	"BY": regexp.MustCompile(`^BY\d{2}[A-Z]{4}\d{4}\d{16}$`),            // Belarus
	"CH": regexp.MustCompile(`^CH\d{2}\d{5}\d{12}$`),                    // Switzerland
	"CR": regexp.MustCompile(`^CR\d{2}\d{4}\d{14}$`),                    // Costa Rica
	"CY": regexp.MustCompile(`^CY\d{2}\d{3}\d{5}\d{16}$`),               // Cyprus
	"CZ": regexp.MustCompile(`^CZ\d{2}\d{4}\d{6}\d{10}$`),               // Czech Republic
	"DE": regexp.MustCompile(`^DE\d{2}\d{8}\d{10}$`),                    // Germany
	"DK": regexp.MustCompile(`^DK\d{2}\d{4}\d{10}$`),                    // Denmark
	"DO": regexp.MustCompile(`^DO\d{2}[A-Z]{4}\d{20}$`),                 // Dominican Republic
	"EE": regexp.MustCompile(`^EE\d{2}\d{2}\d{2}\d{11}$`),               // Estonia
	"ES": regexp.MustCompile(`^ES\d{2}\d{4}\d{4}\d{2}\d{10}$`),          // Spain
	"FI": regexp.MustCompile(`^FI\d{2}\d{6}\d{7}\d{1}$`),                // Finland
	"FO": regexp.MustCompile(`^FO\d{2}\d{4}\d{9}\d{1}$`),                // Faroe Islands
	"FR": regexp.MustCompile(`^FR\d{2}\d{5}\d{5}[A-Z0-9]{11}\d{2}$`),    // France
	"GB": regexp.MustCompile(`^GB\d{2}[A-Z]{4}\d{6}\d{8}$`),             // United Kingdom
	"GE": regexp.MustCompile(`^GE\d{2}[A-Z]{2}\d{16}$`),                 // Georgia
	"GI": regexp.MustCompile(`^GI\d{2}[A-Z]{4}\d{15}$`),                 // Gibraltar
	"GL": regexp.MustCompile(`^GL\d{2}\d{4}\d{10}$`),                    // Greenland
	"GR": regexp.MustCompile(`^GR\d{2}\d{3}\d{4}[A-Z0-9]{16}$`),         // Greece
	"GT": regexp.MustCompile(`^GT\d{2}[A-Z0-9]{4}\d{20}$`),              // Guatemala
	"HR": regexp.MustCompile(`^HR\d{2}\d{7}\d{10}$`),                    // Croatia
	"HU": regexp.MustCompile(`^HU\d{2}\d{3}\d{4}\d{1}\d{15}\d{1}$`),     // Hungary
	"IE": regexp.MustCompile(`^IE\d{2}[A-Z]{4}\d{6}\d{8}$`),             // Ireland
	"IL": regexp.MustCompile(`^IL\d{2}\d{3}\d{3}\d{13}$`),               // Israel
	"IS": regexp.MustCompile(`^IS\d{2}\d{4}\d{2}\d{6}\d{10}$`),          // Iceland
	"IT": regexp.MustCompile(`^IT\d{2}[A-Z]{1}\d{5}\d{5}[A-Z0-9]{12}$`), // Italy
	"JO": regexp.MustCompile(`^JO\d{2}[A-Z]{4}\d{4}\d{18}$`),            // Jordan
	"KW": regexp.MustCompile(`^KW\d{2}[A-Z]{4}\d{22}$`),                 // Kuwait
	"KZ": regexp.MustCompile(`^KZ\d{2}\d{3}\d{13}$`),                    // Kazakhstan
	"LB": regexp.MustCompile(`^LB\d{2}\d{4}\d{20}$`),                    // Lebanon
	"LC": regexp.MustCompile(`^LC\d{2}[A-Z]{4}[A-Z0-9]{24}$`),           // Saint Lucia
	"LI": regexp.MustCompile(`^LI\d{2}\d{5}\d{12}$`),                    // Liechtenstein
	"LT": regexp.MustCompile(`^LT\d{2}\d{5}\d{11}$`),                    // Lithuania
	"LU": regexp.MustCompile(`^LU\d{2}\d{3}\d{13}$`),                    // Luxembourg
	"LV": regexp.MustCompile(`^LV\d{2}[A-Z]{4}\d{13}$`),                 // Latvia
	"MC": regexp.MustCompile(`^MC\d{2}\d{5}\d{5}[A-Z0-9]{11}\d{2}$`),    // Monaco
	"MD": regexp.MustCompile(`^MD\d{2}[A-Z0-9]{2}\d{18}$`),              // Moldova
	"ME": regexp.MustCompile(`^ME\d{2}\d{3}\d{13}\d{2}$`),               // Montenegro
	"MK": regexp.MustCompile(`^MK\d{2}\d{3}[A-Z0-9]{10}\d{2}$`),         // North Macedonia
	"MR": regexp.MustCompile(`^MR\d{2}\d{5}\d{5}\d{11}\d{2}$`),          // Mauritania
	"MT": regexp.MustCompile(`^MT\d{2}[A-Z]{4}\d{5}\d{18}$`),            // Malta
	"MU": regexp.MustCompile(`^MU\d{2}[A-Z]{4}\d{19}[A-Z]{3}$`),         // Mauritius
	"NL": regexp.MustCompile(`^NL\d{2}[A-Z]{4}\d{10}$`),                 // Netherlands
	"NO": regexp.MustCompile(`^NO\d{2}\d{4}\d{6}\d{1}$`),                // Norway
	"PK": regexp.MustCompile(`^PK\d{2}[A-Z]{4}\d{16}$`),                 // Pakistan
	"PL": regexp.MustCompile(`^PL\d{2}\d{8}\d{16}$`),                    // Poland
	"PS": regexp.MustCompile(`^PS\d{2}[A-Z]{4}\d{21}$`),                 // Palestine
	"PT": regexp.MustCompile(`^PT\d{2}\d{4}\d{4}\d{11}\d{2}$`),          // Portugal
	"QA": regexp.MustCompile(`^QA\d{2}[A-Z]{4}\d{21}$`),                 // Qatar
	"RO": regexp.MustCompile(`^RO\d{2}[A-Z]{4}\d{16}$`),                 // Romania
	"RS": regexp.MustCompile(`^RS\d{2}\d{3}\d{13}\d{2}$`),               // Serbia
	"SA": regexp.MustCompile(`^SA\d{2}\d{2}\d{18}$`),                    // Saudi Arabia
	"SC": regexp.MustCompile(`^SC\d{2}[A-Z]{4}\d{20}[A-Z]{3}$`),         // Seychelles
	"SE": regexp.MustCompile(`^SE\d{2}\d{3}\d{16}\d{1}$`),               // Sweden
	"SI": regexp.MustCompile(`^SI\d{2}\d{5}\d{8}\d{2}$`),                // Slovenia
	"SK": regexp.MustCompile(`^SK\d{2}\d{4}\d{6}\d{10}$`),               // Slovakia
	"SM": regexp.MustCompile(`^SM\d{2}[A-Z]{1}\d{5}\d{5}[A-Z0-9]{12}$`), // San Marino
	"ST": regexp.MustCompile(`^ST\d{2}\d{4}\d{4}\d{11}\d{2}$`),          // Sao Tome and Principe
	"SV": regexp.MustCompile(`^SV\d{2}[A-Z]{4}\d{20}$`),                 // El Salvador
	"TL": regexp.MustCompile(`^TL\d{2}\d{3}\d{14}$`),                    // East Timor
	"TN": regexp.MustCompile(`^TN\d{2}\d{2}\d{3}\d{13}\d{2}$`),          // Tunisia
	"TR": regexp.MustCompile(`^TR\d{2}\d{5}\d{1}\d{16}$`),               // Turkey
	"UA": regexp.MustCompile(`^UA\d{2}\d{6}\d{19}$`),                    // Ukraine
	"VG": regexp.MustCompile(`^VG\d{2}[A-Z]{4}\d{16}$`),                 // British Virgin Islands
	"XK": regexp.MustCompile(`^XK\d{2}\d{4}\d{10}\d{2}$`),               // Kosovo
	"IQ": regexp.MustCompile(`^IQ\d{2}[A-Z]{4}\d{15}$`),                 // Iraq
	"MZ": regexp.MustCompile(`^MZ\d{2}[A-Z0-9]{21}$`),                   // Mozambique
	"VA": regexp.MustCompile(`^VA\d{2}[A-Z0-9]{18}$`),                   // Vatican City
}

// A validator that checks if the string is an IBAN (International Bank Account Number).
// these are the allowed country codes ('AD','AE','AL','AT','AZ','BA','BE','BG','BH','BR','BY','CH','CR','CY','CZ','DE','DK','DO','EE','EG','ES','FI','FO','FR','GB','GE','GI','GL','GR','GT','HR','HU','IE','IL','IQ','IR','IS','IT','JO','KW','KZ','LB','LC','LI','LT','LU','LV','MC','MD','ME','MK','MR','MT','MU','MZ','NL','NO','PK','PL','PS','PT','QA','RO','RS','SA','SC','SE','SI','SK','SM','SV','TL','TN','TR','UA','VA','VG','XK').
// Abeg no checksum 😭😭😭😭
func IsIBAN(str string, countryCode string) bool {
	re, exist := ibanRegex[countryCode]

	if !exist {
		return false
	}

	return re.MatchString(str)
}
