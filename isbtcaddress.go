package validatorgo

import "regexp"

// A validator that checks if the string is a valid BTC address.
func IsBTCAddress(str string) bool {
	return regexp.MustCompile("^(bc1|[13])[a-km-zA-HJ-NP-Z1-9]{25,34}$").MatchString(str)
}
