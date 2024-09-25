package validatorgo

// A validator that checks if the string contains any half-width chars.
//
//	ok := validatorgo.IsHalfWidth("abc123")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsHalfWidth("漢字テスト")
//	fmt.Println(ok) // false
func IsHalfWidth(str string) bool {
	for _, char := range str {
		if char >= '\u0020' && char <= '\u007E' { // ASCII (half-width characters)
			return true
		}
	}
	return false
}
