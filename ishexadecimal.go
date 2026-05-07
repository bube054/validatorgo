package validatorgo

import "regexp"

// A validator that checks if the string is a hexadecimal number.
//
//	ok, _ := validatorgo.IsHexadecimal("1234567890abcdef")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsHexadecimal("abcdefg")
//	fmt.Println(ok) // false
func IsHexadecimal(str string) (bool, error) {
	if regexp.MustCompile(`^[a-fA-F0-9]+$`).MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsHexadecimal", ErrInvalidFormat, "invalid hexadecimal")
}
