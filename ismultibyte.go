package validatorgo

import "regexp"

// A validator that checks if the string contains one or more multibyte chars.
//
//	ok, _ := validatorgo.IsMultibyte("こんにちは")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsMultibyte("hello")
//	fmt.Println(ok) // false
func IsMultibyte(str string) (bool, error) {
	if regexp.MustCompile(`[^\x00-\x7F]`).MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsMultibyte", ErrInvalidFormat, "invalid multibyte")
}
