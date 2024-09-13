package validatorgo

// A validator that checks if the string matches the comparison.
//
//	ok := govalidator.Equals("Hello", "Hello")
//	fmt.Println(ok) // true
//	ok := govalidator.Equals("Hello", "World")
//	fmt.Println(ok) // false
func Equals(str, comparison string) bool {
	return str == comparison
}
