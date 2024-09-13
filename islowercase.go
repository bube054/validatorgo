package validatorgo

import "strings"

// A validator that checks if the string is lowercase.
func IsLowerCase(str string) bool {
	return str == strings.ToLower(str)
}
