package validatorgo

import (
	"regexp"
)

// A validator that checks if the string is base58 encoded.
func IsBase58(str string) bool {
	return regexp.MustCompile("^[A-HJ-NP-Za-km-z1-9]*$").MatchString(str)
}
