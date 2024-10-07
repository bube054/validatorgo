package validatorgo

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

var (
	isFQDNOptsDefaultRequireTld       bool = true
	isFQDNOptsDefaultAllowUnderscores bool = false
	isFQDNOptsDefaultAllowTrailingDot bool = false
	isFQDNOptsDefaultAllowNumericTld  bool = false
	isFQDNOptsDefaultIgnoreMaxLength  bool = false

	isFQDNMaxLength int = 253
)

// IsFQDNOptsOpts is used to configure IsFQDNOpts
type IsFQDNOpts struct {
	RequireTld       bool
	AllowUnderscores bool
	AllowTrailingDot bool
	AllowNumericTld  bool
	IgnoreMaxLength  bool
}

// A validator that checks if the string is a fully qualified domain name (e.g. domain.com).
//
// IsFQDNOpts is a struct which defaults to { RequireTld: true, AllowUnderscores: false, AllowTrailingDot: false, AllowNumericTld: false, allow_wildcard: false, IgnoreMaxLength: false }.
//
//	ok := validatorgo.IsFQDN("localhost",  &validatorgo.IsFQDNOpts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsFQDN("example..com", &validatorgo.IsFQDNOpts{})
//	fmt.Println(ok) // false
func IsFQDN(str string, opts *IsFQDNOpts) bool {
	if opts == nil {
		opts = setIsFQDNOptsToDefault()
	}

	ignMaxLength := true
	ln := utf8.RuneCountInString(str)

	if !opts.IgnoreMaxLength && ln > isFQDNMaxLength {
		ignMaxLength = false
	}

	allowUnderScoreRe := `a-zA-Z0-9`
	if opts.AllowUnderscores {
		allowUnderScoreRe = `\w`
	}
	requireTldRe := "?"
	if opts.RequireTld {
		requireTldRe = ""
	}
	allowTrailingDotRe := ``
	if opts.AllowTrailingDot {
		allowTrailingDotRe = `\.?`
	}
	allowNumTldRe := `[a-zA-Z_]`
	if opts.AllowNumericTld {
		allowNumTldRe = `\w`
	}

	reStr := fmt.Sprintf(`^([%s])+(\.[%s]+)?\.%s%s+%s$`, allowUnderScoreRe, allowUnderScoreRe, requireTldRe, allowNumTldRe, allowTrailingDotRe)
	re := regexp.MustCompile(reStr)
	isValid := re.MatchString(str)
	return isValid && ignMaxLength
}

func setIsFQDNOptsToDefault() *IsFQDNOpts {
	return &IsFQDNOpts{
		RequireTld:       isFQDNOptsDefaultRequireTld,
		AllowUnderscores: isFQDNOptsDefaultAllowUnderscores,
		AllowTrailingDot: isFQDNOptsDefaultAllowTrailingDot,
		AllowNumericTld:  isFQDNOptsDefaultAllowNumericTld,
		IgnoreMaxLength:  isFQDNOptsDefaultIgnoreMaxLength,
	}
}
