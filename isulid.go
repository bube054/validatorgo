package validatorgo

import "regexp"

// A validator that checks if the string is a [ULID].
//
//	ok := validatorgo.IsULID("01ARZ3NDEKTSV4RRFFQ69G5FAV")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsULID("01ARZ3NDEKTSV4RRFFQ69G5FA")
//	fmt.Println(ok) // false
//
// [ULID]: https://github.com/ulid/spec
func IsULID(str string) bool {
	return regexp.MustCompile(`[0-7][0-9A-HJKMNP-TV-Z]{25}`).MatchString(str)
}
