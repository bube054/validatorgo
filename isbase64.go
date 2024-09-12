package validatorgo

import (
	"regexp"
)

// IsBase64Opts is used to configure IsBase64
type IsBase64Opts struct {
	UrlSafe bool
}

// A validator that checks if the string is base64 encoded. options is optional and defaults to { urlSafe: false }
// when urlSafe is true it tests the given base64 encoded string is url safe.
func IsBase64(str string, opts IsBase64Opts) bool {
	base64Regex := "^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"
	urlSafeRegex := "^(?:[A-Za-z0-9_-]{4})*(?:[A-Za-z0-9_-]{2}(?:==)?|[A-Za-z0-9_-]{3}=?)?$"

	if opts.UrlSafe {
		return regexp.MustCompile(urlSafeRegex).MatchString(str)
	} else {
		return regexp.MustCompile(base64Regex).MatchString(str)
	}
}
