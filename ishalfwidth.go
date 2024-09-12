package validatorgo

import "regexp"

// A validator that checks if the string contains any half-width chars.
func IsHalfWidth(str string) bool {
	halfWidthRegex := regexp.MustCompile(`[\x{FF61}-\x{FFDC}\x{FFE8}-\x{FFEE}]`)
	return halfWidthRegex.MatchString(str)
}
