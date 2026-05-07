package validatorgo

// A validator that checks if the string is in a slice of allowed values.
//
//	ok, _ := validatorgo.IsIn("apple", []string{"apple", "banana", "grape"})
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsIn("orange", []string{"apple", "banana", "grape"})
//	fmt.Println(ok) // false
func IsIn(str string, values []string) (bool, error) {
	for _, val := range values {
		if str == val {
			return true, nil
		}
	}

	return false, newValidationError("IsIn", ErrNotFound, "value not found in allowed list")
}
