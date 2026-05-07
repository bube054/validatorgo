package validatorgo

import (
	"regexp"
	"unicode/utf8"
)

// IsBase64Opts is used to configure IsBase64
type IsBase64Opts struct {
	UrlSafe *bool // checks whether string is url safe.
}

func (o *IsBase64Opts) mergeDefaults() {
	if o.UrlSafe == nil {
		o.UrlSafe = Bool(false)
	}
}

// A validator that checks if the string is base64 encoded.
//
// IsBase64Opts is an optional struct which defaults to { UrlSafe: false }.
// When UrlSafe is true it tests the given base64 encoded string is [url safe].
//
//	ok, _ := validatorgo.IsBase64("SGVsbG8gd29ybGQ", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsBase64("SGVsbG8g@d29ybGQ=", nil)
//	fmt.Println(ok) // false
//
// [url safe]: https://base64.guru/standards/base64url
func IsBase64(str string, opts *IsBase64Opts) (bool, error) {
	if opts == nil {
		opts = &IsBase64Opts{}
	}
	opts.mergeDefaults()

	if utf8.RuneCountInString(str) < 2 {
		return false, newValidationError("IsBase64", ErrTooShort, "string is too short to be valid base64")
	}

	if *opts.UrlSafe {
		if !regexp.MustCompile(`^(?:[A-Za-z0-9_-]{4})*(?:[A-Za-z0-9_-]{2}(?:==)?|[A-Za-z0-9_-]{3}=?)?$`).MatchString(str) {
			return false, newValidationError("IsBase64", ErrInvalidFormat, "string is not valid url-safe base64")
		}
		return true, nil
	} else {
		if !regexp.MustCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`).MatchString(str) {
			return false, newValidationError("IsBase64", ErrInvalidFormat, "string is not valid standard base64")
		}
		return true, nil
	}
}

