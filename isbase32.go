package validatorgo

import (
	"regexp"
	"strings"
)

// IsBase32Opts is used to configure IsBase32
type IsBase32Opts struct {
	Crockford bool
}

// A validator that checks if the string is base32 encoded.
// IsBase32Opts defaults to { crockford: false }.
// When crockford is true it tests the given base32 encoded string using Crockford's base32 alternative(http://www.crockford.com/base32.html).
func IsBase32(str string, opts IsBase32Opts) bool {
	str = strings.TrimRight(str, "=")

	base32Regex := "^[A-Z2-7]+$"
	crockfordRegex := "^[A-HJ-KM-NP-TV-Z0-9]+$"

	if opts.Crockford {
		return regexp.MustCompile(crockfordRegex).MatchString(strings.ToUpper(str))
	} else {
		return regexp.MustCompile(base32Regex).MatchString(str)
	}
}
