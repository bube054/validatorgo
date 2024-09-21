package validatorgo

// import "strings"

// A validator that checks if the string is lowercase.
//
//	ok := validatorgo.IsLowerCase("hello")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsLowerCase("WORLD")
//	fmt.Println(ok) // false
func IsLowerCase(str string) bool {
	// return str == strings.ToLower(str)
	return false // need to think about special characters
}
