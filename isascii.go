package validatorgo

import "regexp"

// A validator that checks if the string contains ASCII chars only.
//	ok := govalidator.IsAscii("Hello")
//	fmt.Println(ok) // true
func IsAscii(str string) bool {
	return regexp.MustCompile("^[ -~]+$").MatchString(str)
}
