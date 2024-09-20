package validatorgo

import "regexp"

var localeLicensePlateRegex = map[string]*regexp.Regexp{
	"cs-CZ": regexp.MustCompile(`/^(([ABCDEFHKIJKLMNPRSTUVXYZ]|[0-9])-?){5,8}$/g`),
}

// (NOT IMPLEMENTED, DO NOT USE) A validator that checks if the string matches the format of a country's license plate.
//
// locale is one of 'cs-CZ', 'de-DE', 'de-LI', 'en-IN', 'en-SG', 'en-PK', 'es-AR', 'hu-HU', 'pt-BR', 'pt-PT', 'sq-AL', 'sv-SE', 'any'.
func IsLicensePlate(str string, locale string) bool {
	return true
}
