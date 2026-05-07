package validatorgo

// A validator that checks if the string matches the comparison.
//
//	ok, _ := validatorgo.Equals("Hello", "Hello")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.Equals("Hello", "World")
//	fmt.Println(ok) // false
func Equals(str, comparison string) (bool, error) {
	if str == comparison {
		return true, nil
	}
	return false, newValidationError("Equals", ErrInvalidValue, "strings are not equal")
}
