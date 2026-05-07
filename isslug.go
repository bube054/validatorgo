package validatorgo

import "regexp"

// A validator that checks if the string is of type slug.
//
//	ok, _ := validatorgo.IsSlug("rgb(255,0,0)")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsSlug("rgb( 255 , 0 , 0 )")
//	fmt.Println(ok) // false
func IsSlug(str string) (bool, error) {
	if regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`).MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsSlug", ErrInvalidFormat, "invalid slug")
}
