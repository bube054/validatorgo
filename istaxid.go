package validatorgo

import "regexp"

// taxIdLocaleRegex is the set of locales and their tax id regex
var taxIdLocaleRegex = map[string]*regexp.Regexp{
	"bg-BG": regexp.MustCompile(`^\d{10}$`),                                        // Bulgaria: 10 digits
	"cs-CZ": regexp.MustCompile(`^\d{6}\/\d{3,4}$`),                                // Czech Republic: 6 digits / 3-4 digits
	"de-AT": regexp.MustCompile(`^\d{9}$`),                                         // Austria: 9 digits
	"de-DE": regexp.MustCompile(`^\d{11}$`),                                        // Germany: 11 digits
	"dk-DK": regexp.MustCompile(`^\d{8}$`),                                         // Denmark: 8 digits
	"el-CY": regexp.MustCompile(`^[0-59]\d{7}[A-Z]$`),                              // Cyprus: 8 digits and 1 letter
	"el-GR": regexp.MustCompile(`^\d{9}$`),                                         // Greece: 9 digits
	"en-CA": regexp.MustCompile(`^\d{9}$`),                                         // Canada: 9 digits (Social Insurance Number)
	"en-GB": regexp.MustCompile(`^\d{10}$`),                                        // UK: 10 digits (UTR)
	"en-IE": regexp.MustCompile(`^\d{7}[A-Z]{1,2}$`),                               // Ireland: 7 digits + 1-2 letters
	"en-US": regexp.MustCompile(`^\d{3}-?\d{2}-?\d{4}$`),                           // USA: 3-2-4 format (SSN)
	"es-AR": regexp.MustCompile(`^\d{11}$`),                                        // Argentina: 11 digits (CUIT)
	"es-ES": regexp.MustCompile(`^[0-9A-Z]\d{7}[0-9A-Z]$`),                         // Spain: 8 digits + 1 letter (NIF)
	"et-EE": regexp.MustCompile(`^\d{11}$`),                                        // Estonia: 11 digits
	"fi-FI": regexp.MustCompile(`^\d{6}[-+A]\d{3}[0-9A-FHJ-NPR-Y]$`),               // Finland: 6 digits, a separator, 3 digits + 1 character (personal identity code)
	"fr-BE": regexp.MustCompile(`^\d{11}$`),                                        // Belgium: 11 digits
	"fr-CA": regexp.MustCompile(`^\d{9}$`),                                         // Canada (French): 9 digits (Social Insurance Number)
	"fr-FR": regexp.MustCompile(`^\d{13}$`),                                        // France: 13 digits (INSEE number)
	"fr-LU": regexp.MustCompile(`^\d{13}$`),                                        // Luxembourg: 13 digits
	"hr-HR": regexp.MustCompile(`^\d{11}$`),                                        // Croatia: 11 digits (OIB)
	"hu-HU": regexp.MustCompile(`^\d{10}$`),                                        // Hungary: 10 digits
	"it-IT": regexp.MustCompile(`^[A-Z]{6}\d{2}[A-EHLMPR-T]\d{2}[A-Z]\d{3}[A-Z]$`), // Italy: Codice Fiscale (complex alphanumeric format)
	"lb-LU": regexp.MustCompile(`^\d{13}$`),                                        // Luxembourg: 13 digits
	"lt-LT": regexp.MustCompile(`^\d{11}$`),                                        // Lithuania: 11 digits
	"lv-LV": regexp.MustCompile(`^\d{11}$`),                                        // Latvia: 11 digits
	"mt-MT": regexp.MustCompile(`^\d{3,7}[A-Z]$`),                                  // Malta: 3-7 digits + 1 letter
	"nl-BE": regexp.MustCompile(`^\d{11}$`),                                        // Belgium (Dutch): 11 digits
	"nl-NL": regexp.MustCompile(`^\d{9}$`),                                         // Netherlands: 9 digits (BSN)
	"pl-PL": regexp.MustCompile(`^\d{10}$`),                                        // Poland: 10 digits (PESEL)
	"pt-BR": regexp.MustCompile(`^\d{11}$`),                                        // Brazil: 11 digits (CPF)
	"pt-PT": regexp.MustCompile(`^\d{9}$`),                                         // Portugal: 9 digits (NIF)
	"ro-RO": regexp.MustCompile(`^\d{13}$`),                                        // Romania: 13 digits
	"sk-SK": regexp.MustCompile(`^\d{10}$`),                                        // Slovakia: 10 digits (Rodné číslo without slash)
	"sl-SI": regexp.MustCompile(`^\d{8}$`),                                         // Slovenia: 8 digits
	"sv-SE": regexp.MustCompile(`^\d{6}-?\d{4}$`),                                  // Sweden: 10 digits (Personal Identity Number)
	"uk-UA": regexp.MustCompile(`^\d{10}$`),                                        // Ukraine: 10 digits (INN)
}

// A validator that checks if the string is a valid Tax Identification Number. Defaults locale is "en-US" and "any" will match any of them.
//
// Supported locales: ("bg-BG", "cs-CZ", "de-AT", "de-DE", "dk-DK", "el-CY", "el-GR", "en-CA", "en-GB", "en-IE", "en-US", "es-AR", "es-ES", "et-EE", "fi-FI", "fr-BE", "fr-CA", "fr-FR", "fr-LU", "hr-HR", "hu-HU", "it-IT", "lb-LU", "lt-LT", "lv-LV", "mt-MT", "nl-BE", "nl-NL", "pl-PL", "pt-BR", "pt-PT", "ro-RO", "sk-SK", "sl-SI", "sv-SE", "uk-UA").
func IsTaxID(str, locale string) bool {
	if locale == "" {
		locale = "en-US"
	}

	re, ok := taxIdLocaleRegex[locale]

	if !ok {
		if locale != "any" {
			return false
		}

		for _, reg := range taxIdLocaleRegex {
			match := reg.MatchString(str)

			if match {
				return true
			}
		}

		return false
	}

	return re.MatchString(str)
}
