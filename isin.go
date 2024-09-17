package validatorgo

// A validator that checks if the string is in a slice of allowed values.
//
//	ok := validatorgo.IsIn("apple", []string{"apple", "banana", "grape"})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsIn("orange", []string{"apple", "banana", "grape"})
//	fmt.Println(ok) // false
func IsIn(str string, values []string) bool {
	for _, val := range values {
		if str == val {
			return true
		}
	}

	return false
}
