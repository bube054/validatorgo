package validatorgo

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

const defaultMaxFQDNLength = 253

// IsFQDNoptsOpts is used to configure IsFQDNopts
type IsFQDNopts struct {
	RequireTld       bool
	AllowUnderscores bool
	AllowTrailingDot bool
	AllowNumericTld  bool
	IgnoreMaxLength  bool
}

// A validator that checks if the string is a fully qualified domain name (e.g. domain.com).
// If allow_wildcard is set to true, the validator will allow domain starting with *. (e.g. *.example.com or *.shop.example.com).
func IsFQDN(str string, opts IsFQDNopts) bool {
	ignMaxLength := true
	len := utf8.RuneCountInString(str)

	if !opts.IgnoreMaxLength && len > defaultMaxFQDNLength {
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
	// fmt.Println(reStr)
	re := regexp.MustCompile(reStr)
	isValid := re.MatchString(str)
	fmt.Printf("length is, %d while defleng is %d and isvalid is %t\n", len, defaultMaxFQDNLength, isValid)
	return isValid && ignMaxLength
}
