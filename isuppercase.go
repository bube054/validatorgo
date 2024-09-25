package validatorgo

import "strings"

// A validator that checks if the string is uppercase.
//
//	ok := validatorgo.IsUpperCase("HELLO")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsUpperCase("world")
//	fmt.Println(ok) // false
func IsUpperCase(str string) bool {
	return str == strings.ToUpper(str)
}
