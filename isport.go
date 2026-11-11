package validatorgo

// A validator that checks if the string is a valid port number: a whole number
// from 0 to 65535 written without leading zeros, matching validator.js.
//
//	ok, _ := validatorgo.IsPort("8080")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsPort("65536")
//	fmt.Println(ok) // false
//	ok, _ = validatorgo.IsPort("080")
//	fmt.Println(ok) // false
func IsPort(str string) (bool, error) {
	if compileOnce(`^([0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`).MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsPort", ErrInvalidFormat, "invalid port")
}
