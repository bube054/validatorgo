package validatorgo

import "regexp"

// A validator that checks if the string is of type slug.
//
//	ok := validatorgo.IsSlug("rgb(255,0,0)")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsSlug("rgb( 255 , 0 , 0 )")
//	fmt.Println(ok) // false
func IsSlug(str string) bool {
	return regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`).MatchString(str)
}