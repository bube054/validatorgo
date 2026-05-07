package validatorgo

import (
	"regexp"
)

// IsMailToURIOpts is used to configure IsMailtoURI
type IsMailToURIOpts struct {
	IsEmailOpts
}

// A validator that checks if the string is a [Mailto URI] format.
//
// IsMailToURIOpts is a struct that directly embeds IsEmailOpts.
//
// IsMailToURIOpts validates emails inside the URI (check IsEmailOpts for details).
//
//	ok, _ := validatorgo.IsMailtoURI("mailto:someone@example.com", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsMailtoURI("someone@example.com", nil)
//	fmt.Println(ok) // false
//
// [Mailto URI]: https://en.wikipedia.org/wiki/Mailto
func IsMailtoURI(str string, opts *IsMailToURIOpts) (bool, error) {
	re := regexp.MustCompile(`^(mailto:)([^\?]+)(\?.*)?$`)

	capGrp := re.FindStringSubmatch(str)

	if capGrp == nil {
		return false, newValidationError("IsMailtoURI", ErrInvalidFormat, "invalid mailtouri")
	}

	email := capGrp[2]

	var emailOpts *IsEmailOpts
	if opts != nil {
		emailOpts = &opts.IsEmailOpts
	}

	ok, _ := IsEmail(email, emailOpts)
	if ok {
		return true, nil
	}
	return false, newValidationError("IsMailtoURI", ErrInvalidFormat, "invalid mailtouri")
}
