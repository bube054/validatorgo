package validatorgo

// A validator that checks if the string contains a mixture of full and half-width chars.
//
//	ok := validatorgo.IsVariableWidth("ａｂｃ123")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsVariableWidth("ａｂｃ１２３")
//	fmt.Println(ok) // false
func IsVariableWidth(str string) bool {
	hasFullWidth := false
	hasHalfWidth := false

	for _, char := range str {
		if (char >= '\uFF00' && char <= '\uFFEF') ||
			(char >= '\u4E00' && char <= '\u9FFF') ||
			(char >= '\u3040' && char <= '\u309F') ||
			(char >= '\u30A0' && char <= '\u30FF') {
			hasFullWidth = true
		}

		if char >= '\u0020' && char <= '\u007E' {
			hasHalfWidth = true
		}

		if hasFullWidth && hasHalfWidth {
			return true
		}
	}

	return false
}
