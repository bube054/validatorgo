package validatorgo

// A validator that	checks if the string is a MD5 hash.
//
// Please note that you can also use the isHash(str, "md5") function.
//
// Keep in mind that MD5 has some collision weaknesses compared to other algorithms (e.g., SHA).
//
//	ok, _ := validatorgo.IsMD5("9e107d9d372bb6826bd81d3542a419d6")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsMD5("9e107d9d372bb6826bd81d3542a419d")
//	fmt.Println(ok) // false
func IsMD5(str string) (bool, error) {
	isHash, _ := IsHash(str, "md5")
	if !isHash {
		return false, newValidationError("IsMD5", ErrInvalidFormat, "string is not a valid MD5 hash")
	}
	return true, nil
}
