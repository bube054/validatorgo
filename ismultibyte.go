package validatorgo

import "regexp"

// A validator that checks if the string contains one or more multibyte chars.
//
//	ok := validatorgo.IsMultibyte("こんにちは")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsMultibyte("hello")
//	fmt.Println(ok) // false
func IsMultibyte(str string) bool {
	return regexp.MustCompile(`[^\x00-\x7F]`).MatchString(str)
}
