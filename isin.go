package validatorgo

// A validator that checks if the string is in a slice of allowed values.
func IsIn(str string, values []string) bool {
	for _, val := range values {
		if str == val {
			return true
		}
	}

	return false
}
