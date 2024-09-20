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
//	ok := validatorgo.IsMailtoURI("mailto:someone@example.com")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsMailtoURI("someone@example.com")
//	fmt.Println(ok) // false
//
// [Mailto URI]: https://en.wikipedia.org/wiki/Mailto
func IsMailtoURI(str string, opts IsMailToURIOpts) bool {
	re := regexp.MustCompile(`^(mailto:)([^\?]+)(\?.*)?$`)

	capGrp := re.FindStringSubmatch(str)

	if capGrp == nil {
		return false
	}

	email := capGrp[2]

	return IsEmail(email, IsEmailOpts{
		AllowDisplayName:         opts.AllowDisplayName,
		RequireDisplayName:       opts.RequireDisplayName,
		AllowUTF8LocalPart:       opts.AllowUTF8LocalPart,
		RequireTld:               opts.RequireTld,
		IgnoreMaxLength:          opts.IgnoreMaxLength,
		AllowIpDomain:            opts.AllowIpDomain,
		DomainSpecificValidation: opts.DomainSpecificValidation,
		BlacklistedChars:         opts.BlacklistedChars,
		HostBlacklist:            opts.HostBlacklist,
		HostWhitelist:            opts.HostWhitelist,
	})
}
