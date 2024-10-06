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
//	isAlpha := validatorgo.IsBase32("JBSWY3DPEBLW64TMMQ", &validatorgo.IsBase32Opts{})
//	fmt.Println(isAlpha) // true
//	isAlpha := validatorgo.IsBase32("jbswy3dpeblw64tmmq======", &validatorgo.IsBase32Opts{})
//	fmt.Println(isAlpha) // false
//
// [crockford's]: http://www.crockford.com/base32.html
func IsBase32(str string, opts *IsBase32Opts) bool {
	if opts == nil {
		opts = setIsBase32OptsToDefault()
	}

	strWithoutEq := strings.TrimRight(str, "=")
	strWithoutHyp := stripHyphens(strWithoutEq)

	if len(strWithoutHyp) < 2 {
		return false
	}

	if opts.Crockford {
		return regexp.MustCompile(`^[A-HJ-KM-NP-TV-Z0-9]+$`).MatchString(strings.ToUpper(strWithoutHyp))
	} else {
		return regexp.MustCompile(`^[A-Z2-7]+$`).MatchString(strWithoutHyp)
	}
}

func setIsBase32OptsToDefault() (opts *IsBase32Opts) {
	return &IsBase32Opts{
		Crockford: isBase32OptsDefaultCrockford,
	}
}
