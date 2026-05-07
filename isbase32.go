package validatorgo

import (
	"regexp"
	"strings"
)

var (
	isBase32OptsDefaultCrockford bool = false
)

// IsBase32Opts is used to configure IsBase32
type IsBase32Opts struct {
	Crockford bool // whether to use crockfords base32 alternative encoding scheme

	// ZBase bool // whether to use crockfords base32 alternative encoding scheme
}

// A validator that checks if the string is base32 encoded.
//
// IsBase32Opts defaults to { Crockford: false }.
// When Crockford is true it tests the given base32 encoded string using [crockford's] base32 alternative.
//
//	ok, _ := validatorgo.IsBase32("JBSWY3DPEBLW64TMMQ", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsBase32("jbswy3dpeblw64tmmq======", nil)
//	fmt.Println(ok) // false
//
// [crockford's]: http://www.crockford.com/base32.html
func IsBase32(str string, opts *IsBase32Opts) (bool, error) {
	if opts == nil {
		opts = setIsBase32OptsToDefault()
	}

	strWithoutEq := strings.TrimRight(str, "=")
	strWithoutHyp := stripHyphens(strWithoutEq)

	if len(strWithoutHyp) < 2 {
		return false, newValidationError("IsBase32", ErrTooShort, "string is too short to be valid base32")
	}

	if opts.Crockford {
		if !regexp.MustCompile(`^[A-HJ-KM-NP-TV-Z0-9]+$`).MatchString(strings.ToUpper(strWithoutHyp)) {
			return false, newValidationError("IsBase32", ErrInvalidFormat, "string is not valid crockford base32")
		}
		return true, nil
	} else {
		if !regexp.MustCompile(`^[A-Z2-7]+$`).MatchString(strWithoutHyp) {
			return false, newValidationError("IsBase32", ErrInvalidFormat, "string is not valid standard base32")
		}
		return true, nil
	}
}

func setIsBase32OptsToDefault() (opts *IsBase32Opts) {
	return &IsBase32Opts{
		Crockford: isBase32OptsDefaultCrockford,
	}
}
