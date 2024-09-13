package validatorgo

import "regexp"

// A validator that checks if the string matches the pattern.
func Matches(str string, re *regexp.Regexp) bool {
	return re.MatchString(str)
}
