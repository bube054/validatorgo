package validatorgo

import "regexp"

// A validator that checks if the string is a hexadecimal color.
//
//	ok := validatorgo.IsHexColor("#abc")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsHexColor("#xyz")
//	fmt.Println(ok) // false
func IsHexColor(str string) bool {
	return regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`).MatchString(str)
}
