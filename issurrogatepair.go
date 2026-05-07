package validatorgo

// A validator that checks if the string contains any surrogate pairs chars.
//
//	ok, _ := validatorgo.IsSurrogatePair("")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsSurrogatePair("")
//	fmt.Println(ok) // false
func IsSurrogatePair(str string) (bool, error) {
	for _, r := range str {
		if r > 0xFFFF {
			return true, nil
		}
	}
	return false, newValidationError("IsSurrogatePair", ErrInvalidFormat, "invalid surrogatepair")
}
