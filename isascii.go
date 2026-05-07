package validatorgo

// A validator that checks if the string contains ASCII chars only.
//
//	ok, _ := validatorgo.IsAscii("Hello")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsAscii("こんにちは")
//	fmt.Println(ok) // false
func IsAscii(str string) (bool, error) {
	for _, char := range str {
		if !(char >= 0 && char <= 127) {
			return false, newValidationError("IsAscii", ErrInvalidFormat, "invalid ascii")
		}
	}

	return true, nil
}
