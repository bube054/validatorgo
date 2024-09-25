package validatorgo

import "regexp"

// A validator that checks if the string is a valid octal number.
//
//	ok := validatorgo.IsOctal("07")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsOctal("078")
//	fmt.Println(ok) // false
func IsOctal(str string) bool {
	return regexp.MustCompile(`^[0-7]+$`).MatchString(str)
}
