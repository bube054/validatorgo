package validatorgo

import "strings"

// A validator that checks if the string is lowercase.
//
//	ok, _ := validatorgo.IsLowerCase("hello")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsLowerCase("WORLD")
//	fmt.Println(ok) // false
func IsLowerCase(str string) (bool, error) {
	if str == strings.ToLower(str) && str != strings.ToUpper(str) {
		return true, nil
	}
	return false, newValidationError("IsLowerCase", ErrInvalidFormat, "invalid lowercase")
}
