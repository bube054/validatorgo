package validatorgo

import "regexp"

// A validator that checks if the string contains any half-width chars.
//
//	ok := validatorgo.IsHalfWidth("Hello, ｶﾀｶﾅ")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsHalfWidth("Hello, World!")
//	fmt.Println(ok) // false
func IsHalfWidth(str string) bool {
	halfWidthRegex := regexp.MustCompile(`[\x{FF61}-\x{FFDC}\x{FFE8}-\x{FFEE}]`)
	return halfWidthRegex.MatchString(str)
}
