package validatorgo

import "strings"

// A validator that checks if the string is uppercase.
//
//	ok, _ := validatorgo.IsUpperCase("HELLO")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsUpperCase("world")
//	fmt.Println(ok) // false
func IsUpperCase(str string) (bool, error) {
	if str == strings.ToUpper(str) && str != strings.ToLower(str) {
		return true, nil
	}
	return false, newValidationError("IsUpperCase", ErrInvalidFormat, "invalid uppercase")
}
