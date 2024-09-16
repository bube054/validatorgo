package validatorgo

// A validator that checks if the string contains ASCII chars only.
//
//	ok := validatorgo.IsAscii("Hello")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsAscii("こんにちは")
//	fmt.Println(ok) // false
func IsAscii(str string) bool {
	for _, char := range str {
		if !(char >= 0 && char <= 127) {
			return false
		}
	}

	return true
}
