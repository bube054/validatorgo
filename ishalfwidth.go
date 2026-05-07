package validatorgo

// A validator that checks if the string contains any half-width chars.
//
//	ok, _ := validatorgo.IsHalfWidth("abc123")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsHalfWidth("漢字テスト")
//	fmt.Println(ok) // false
func IsHalfWidth(str string) (bool, error) {
	for _, char := range str {
		if char >= '\u0020' && char <= '\u007E' { // ASCII (half-width characters)
			return true, nil
		}
	}
	return false, newValidationError("IsHalfWidth", ErrInvalidFormat, "invalid halfwidth")
}
