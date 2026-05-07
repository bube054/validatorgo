package validatorgo

// A validator that checks if the string contains any full-width chars.
//
//	ok, _ := validatorgo.IsFullWidth("テスト")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsFullWidth("abc123")
//	fmt.Println(ok) // false
func IsFullWidth(str string) (bool, error) {
	for _, char := range str {
		if (char >= '\uFF00' && char <= '\uFFEF') ||
			(char >= '\u4E00' && char <= '\u9FFF') ||
			(char >= '\u3040' && char <= '\u309F') ||
			(char >= '\u30A0' && char <= '\u30FF') {
			return true, nil
		}
	}
	return false, newValidationError("IsFullWidth", ErrInvalidFormat, "invalid fullwidth")
}
