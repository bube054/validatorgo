package validatorgo

import "regexp"

// A validator that checks if the string is a valid octal number.
//
//	ok, _ := validatorgo.IsOctal("07")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsOctal("078")
//	fmt.Println(ok) // false
func IsOctal(str string) (bool, error) {
	if regexp.MustCompile(`^[0-7]+$`).MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsOctal", ErrInvalidFormat, "invalid octal")
}
