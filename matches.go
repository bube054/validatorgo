package validatorgo

import "regexp"

// A validator that checks if the string matches the regex.
//
//	ok, _ := validatorgo.Matches("foo", regexp.MustCompile(`^foo$`))
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.Matches("foo", regexp.MustCompile(`^foobar$`))
//	fmt.Println(ok) // false
func Matches(str string, re *regexp.Regexp) (bool, error) {
	if re == nil {
		return false, newValidationError("Matches", ErrInvalidFormat, "string does not match pattern")
	}

	if re.MatchString(str) {
		return true, nil
	}
	return false, newValidationError("Matches", ErrInvalidFormat, "string does not match pattern")
}
