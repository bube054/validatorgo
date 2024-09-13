package validatorgo

import "regexp"

// A validator that checks if the string is a valid octal number.
func IsOctal(str string) bool {
	return regexp.MustCompile(`^[0-7]+$`).MatchString(str)
}
