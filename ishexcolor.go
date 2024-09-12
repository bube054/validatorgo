package validatorgo

import "regexp"

// A validator that checks if the string is a hexadecimal color.
func IsHexColor(str string) bool {
	return regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`).MatchString(str)
}
