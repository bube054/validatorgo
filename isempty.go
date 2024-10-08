package validatorgo

import "regexp"

var (
	isEmptyOptsDefaultIgnoreWhitespace bool = false
)

// IsEmptyOpts is used to configure IsEmpty
type IsEmptyOpts struct {
	IgnoreWhitespace bool
}

// A validator check if the string has a length of zero.
//
// IsEmptyOpts is a struct which defaults to { IgnoreWhitespace: false }.
//
//	ok := validatorgo.IsEmpty("", &validatorgo.IsEmpty{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsEmpty("abc", &validatorgo.IsEmpty{})
//	fmt.Println(ok) // false
func IsEmpty(str string, opts *IsEmptyOpts) bool {
	if opts == nil {
		opts = setIsEmptyOptsToDefault()
	}

	if opts.IgnoreWhitespace {
		return regexp.MustCompile(`^(\s+)?$`).MatchString(str)
	} else {
		return regexp.MustCompile(`^$`).MatchString(str)
	}
}

func setIsEmptyOptsToDefault() *IsEmptyOpts {
	return &IsEmptyOpts{
		IgnoreWhitespace: isEmptyOptsDefaultIgnoreWhitespace,
	}
}
