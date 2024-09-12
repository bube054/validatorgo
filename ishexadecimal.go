package validatorgo

import "regexp"

// A validator that checks if the string is a hexadecimal number.
func IsHexadecimal(str string) bool {
	return regexp.MustCompile(`^[a-fA-F0-9]+$`).MatchString(str)
}
