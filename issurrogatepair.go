package validatorgo

// A validator that checks if the string contains any surrogate pairs chars.
//
//	ok := validatorgo.IsSurrogatePair("")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsSurrogatePair("")
//	fmt.Println(ok) // false
func IsSurrogatePair(str string) bool {
	for _, r := range str {
		if r > 0xFFFF {
			return true
		}
	}
	return false
}
