package validatorgo

import "unicode/utf8"

// A validator that checks if the string is a valid hex-encoded representation of a MongoDB ObjectId.
func IsMongoID(str string) bool {
	return IsHexadecimal(str) && utf8.RuneCountInString(str) == 24
}
