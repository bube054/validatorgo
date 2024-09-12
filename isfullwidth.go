package validatorgo

import "regexp"

// A validator that checks if the string contains any full-width chars.
func IsFullWidth(str string) bool {
	fullWidthRegex := regexp.MustCompile(`[^\x{0020}-\x{007E}\x{FF61}-\x{FF9F}\x{FFA0}-\x{FFDC}\x{FFE8}-\x{FFEE}]`)
	return fullWidthRegex.MatchString(str)
}
