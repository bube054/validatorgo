package validatorgo

import "regexp"

// A validator that checks if the string is a hexadecimal color.
//
//	ok, _ := validatorgo.IsHexColor("#abc")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsHexColor("#xyz")
//	fmt.Println(ok) // false
func IsHexColor(str string) (bool, error) {
	if regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`).MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsHexColor", ErrInvalidFormat, "invalid hexcolor")
}
