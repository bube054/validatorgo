package validatorgo

import "regexp"

// IsEmptyOpts is used to configure IsEmpty
type IsEmptyOpts struct {
	IgnoreWhitespace *bool
}

func (o *IsEmptyOpts) mergeDefaults() {
	if o.IgnoreWhitespace == nil {
		o.IgnoreWhitespace = Bool(false)
	}
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
		opts = &IsEmptyOpts{}
	}
	opts.mergeDefaults()

	if *opts.IgnoreWhitespace {
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

