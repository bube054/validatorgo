package validatorgo

import (
	"regexp"
)

// IsBase64Opts is used to configure IsBase64
type IsBase64Opts struct {
	UrlSafe bool // checks whether string is url safe.
}

// A validator that checks if the string is base64 encoded. 
//
// IsBase64Opts is an optional struct which defaults to { UrlSafe: false }.
// When UrlSafe is true it tests the given base64 encoded string is [url safe].
//
//	ok := validatorgo.IsBase64("SGVsbG8gd29ybGQ", validatorgo.IsBase64Opts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsBase64("SGVsbG8g@d29ybGQ=", validatorgo.IsBase64Opts{})
//	fmt.Println(ok) // false
//
// [url safe]: https://base64.guru/standards/base64url
func IsBase64(str string, opts IsBase64Opts) bool {
	if len(str) < 2{
		return false
	}

	if opts.UrlSafe {
		return regexp.MustCompile(`^(?:[A-Za-z0-9_-]{4})*(?:[A-Za-z0-9_-]{2}(?:==)?|[A-Za-z0-9_-]{3}=?)?$`).MatchString(str)
	} else {
		return regexp.MustCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`).MatchString(str)
	}
}
