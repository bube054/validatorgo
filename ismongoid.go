package validatorgo

import "unicode/utf8"

// A validator that checks if the string is a valid hex-encoded representation of a MongoDB ObjectId.
//
//	ok := validatorgo.IsMongoID("507f1f77bcf86cd799439011")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsMongoID("507f1f77bcf86cd79943901")
//	fmt.Println(ok) // false
func IsMongoID(str string) bool {
	return IsHexadecimal(str) && utf8.RuneCountInString(str) == 24
}
