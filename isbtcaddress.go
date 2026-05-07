package validatorgo

import "regexp"

// A validator that checks if the string is a valid BTC address.
//
//	ok, _ := validatorgo.IsBTCAddress("1RAHUEYstWetqabcFn5Au4m4GFg7xJaNVN2")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsBTCAddress("0J98t1RHT73CNmQwertyyWrnqRhWNLy")
//	fmt.Println(ok) // false
func IsBTCAddress(str string) (bool, error) {
	bech32Re := regexp.MustCompile(`^(bc1|tb1|bc1p|tb1p)[ac-hj-np-z02-9]{39,58}$`)
	base58Re := regexp.MustCompile(`^(1|2|3|m|n)[A-HJ-NP-Za-km-z1-9]{25,39}$`)
	legacyBech32Re := regexp.MustCompile(`^(bc1)[a-km-zA-HJ-NP-Z1-9]{25,34}$`)

	if bech32Re.MatchString(str) || base58Re.MatchString(str) || legacyBech32Re.MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsBTCAddress", ErrInvalidFormat, "invalid btcaddress")
}
