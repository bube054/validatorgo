package validatorgo

import "regexp"

// A validator that checks if the string is a valid BTC address.
//
//	ok := validatorgo.IsBTCAddress("1RAHUEYstWetqabcFn5Au4m4GFg7xJaNVN2")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsBTCAddress("0J98t1RHT73CNmQwertyyWrnqRhWNLy")
//	fmt.Println(ok) // false
func IsBTCAddress(str string) bool {
	return regexp.MustCompile("^(bc1|[13])[a-km-zA-HJ-NP-Z1-9]{25,34}$").MatchString(str)
}
