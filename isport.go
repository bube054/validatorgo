package validatorgo

import "regexp"

// A validator that checks if the string is a valid port number.
//
//	ok := validatorgo.IsPort("07")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsPort("078")
//	fmt.Println(ok) // false
func IsPort(str string) bool {
	return regexp.MustCompile(`^0*([0-9]|[1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`).MatchString(str)
}
