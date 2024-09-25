package validatorgo

import "regexp"

// A validator that checks if the string is a hexadecimal number.
//
//	ok := validatorgo.IsHexadecimal("1234567890abcdef")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsHexadecimal("abcdefg")
//	fmt.Println(ok) // false
func IsHexadecimal(str string) bool {
	return regexp.MustCompile(`^[a-fA-F0-9]+$`).MatchString(str)
}
