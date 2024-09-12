package validatorgo

import "regexp"

// IsEmptyOpts is used to configure IsEmpty
type IsEmptyOpts struct {
	IgnoreWhitespace bool
}

// A validator check if the string has a length of zero.
// IsEmptyOpts is a struct which defaults to { IgnoreWhitespace: false }.
func IsEmpty(str string, opts IsEmptyOpts) bool {
	if opts.IgnoreWhitespace {
		return regexp.MustCompile(`^(\s+)?$`).MatchString(str)
	} else {
		return regexp.MustCompile(`^$`).MatchString(str)
	}
}
