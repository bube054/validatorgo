package validatorgo

import "regexp"

var localeLicensePlateRegex = map[string]*regexp.Regexp{
	"cs-CZ": regexp.MustCompile(`^[A-Z]{1,2}[0-9]{1,5}[A-Z]{0,2}$`),       // Czech Republic (e.g., AA12345)
	"de-DE": regexp.MustCompile(`^[A-Z]{1,3}-[A-Z]{1,2}-[0-9]{1,4}$`),     // Germany (e.g., B-MA-1234)
	"de-LI": regexp.MustCompile(`^[0-9]{1,5}$`),                           // Liechtenstein (e.g., 12345)
	"en-IN": regexp.MustCompile(`^[A-Z]{2}[0-9]{1,2}[A-Z]{1,2}[0-9]{4}$`), // India (e.g., DL12AB1234)
	"en-SG": regexp.MustCompile(`^[A-Z]{1,3}[0-9]{1,4}[A-Z]{1}$`),         // Singapore (e.g., SGP1234A)
	"en-PK": regexp.MustCompile(`^[A-Z]{2,3}-[0-9]{1,4}$`),                // Pakistan (e.g., ABC-1234)
	"es-AR": regexp.MustCompile(`^[A-Z]{2}[0-9]{3}[A-Z]{2}$`),             // Argentina (e.g., AB123CD)
	"hu-HU": regexp.MustCompile(`^[A-Z]{3}-[0-9]{3}$`),                    // Hungary (e.g., ABC-123)
	"pt-BR": regexp.MustCompile(`^[A-Z]{3}-[0-9]{4}$`),                    // Brazil (e.g., ABC-1234)
	"pt-PT": regexp.MustCompile(`^[0-9]{2}-[0-9]{2}-[A-Z]{2}$`),           // Portugal (e.g., 12-34-AB)
	"sq-AL": regexp.MustCompile(`^[A-Z]{2}-[0-9]{3}-[A-Z]{2}$`),           // Albania (e.g., AB-123-CD)
	"sv-SE": regexp.MustCompile(`^[A-Z]{3}[0-9]{3}$`),                     // Sweden (e.g., ABC123)
}

// A validator that checks if the string matches the format of a country's license plate.
//
// locale is one of ("cs-CZ", "de-DE", "de-LI", "en-IN", "en-SG", "en-PK", "es-AR", "hu-HU", "pt-BR", "pt-PT", "sq-AL", "sv-SE", "any")
//
//	ok := validatorgo.IsLength("hello", &validatorgo.IsLengthOpts{Min: 3})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsLength("hi", &validatorgo.IsLengthOpts{Min: 3})
//	fmt.Println(ok) // false
func IsLicensePlate(str string, locale string) bool {
	re, ok := localeLicensePlateRegex[locale]

	if !ok {
		if locale != "any" && locale != "" {
			return false
		}

		for _, reg := range localeLicensePlateRegex {
			match := reg.MatchString(str)

			if match {
				return true
			}

		}
		return false
	}

	return re.MatchString(str)
}
