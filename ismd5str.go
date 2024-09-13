package validatorgo

// A validator that	checks if the string is a MD5 hash.
// Please note that you can also use the isHash(str, 'md5') function.
// Keep in mind that MD5 has some collision weaknesses compared to other algorithms (e.g., SHA).
func IsMD5(str string) bool {
	return IsHash(str, "md5")
}
