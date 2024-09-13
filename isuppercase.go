package validatorgo

import "strings"

// A validator that checks if the string is uppercase.
func IsUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}
