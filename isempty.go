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
//	ok, _ := validatorgo.IsEmpty("", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsEmpty("abc", nil)
//	fmt.Println(ok) // false
func IsEmpty(str string, opts *IsEmptyOpts) (bool, error) {
	if opts == nil {
		opts = setIsEmptyOptsToDefault()
	}

	if opts.IgnoreWhitespace {
		if !regexp.MustCompile(`^(\s+)?$`).MatchString(str) {
			return false, newValidationError("IsEmpty", ErrInvalidValue, "string is not empty")
		}
		return true, nil
	} else {
		if !regexp.MustCompile(`^$`).MatchString(str) {
			return false, newValidationError("IsEmpty", ErrInvalidValue, "string is not empty")
		}
		return true, nil
	}
}

func setIsEmptyOptsToDefault() *IsEmptyOpts {
	return &IsEmptyOpts{
		IgnoreWhitespace: isEmptyOptsDefaultIgnoreWhitespace,
	}
}
