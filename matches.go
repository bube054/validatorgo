package validatorgo

import "regexp"

// A validator that checks if the string matches the regex.
//
//	ok := validatorgo.Matches("foo", regexp.MustCompile(`^foo$`))
//	fmt.Println(ok) // true
//	ok := validatorgo.Matches("foo", regexp.MustCompile(`^foobar$`))
//	fmt.Println(ok) // false
func Matches(str string, re *regexp.Regexp) bool {
	if re == nil {
		return false
	}

	return re.MatchString(str)
}
