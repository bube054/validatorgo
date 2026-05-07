package validatorgo

import (
	"regexp"
)

// A validator that checks if the string is base58 encoded.
func IsBase58(str string) (bool, error) {
	if regexp.MustCompile("^[A-HJ-NP-Za-km-z1-9]+$").MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsBase58", ErrInvalidFormat, "invalid base58")
}
