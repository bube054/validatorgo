package validatorgo

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

// IsEmailOpts is used to configure IsEmail
type IsEmailOpts struct {
	AllowDisplayName         bool
	RequireDisplayName       bool
	AllowUTF8LocalPart       bool
	RequireTld               bool
	IgnoreMaxLength          bool
	AllowIpDomain            bool
	DomainSpecificValidation bool
	BlacklistedChars         string
	HostBlacklist            []string
	HostWhitelist            []string
}

// A validator that checks if the string is an email.
// IsEmailOpts is a struct which defaults to { AllowDisplayName: false, RequireDisplayName: false, AllowUTF8LocalPart: true, RequireTld: true, AllowIpDomain: false, allow_underscores: false, DomainSpecificValidation: false, BlacklistedChars: ”, HostBlacklist: [] }.
// If AllowDisplayName is set to true, the validator will also match Display Name <email-address>.
// If RequireDisplayName is set to true, the validator will reject strings without the format Display Name <email-address>.
// If AllowUTF8LocalPart is set to false, the validator will not allow any non-English UTF8 character in email address' local part.
// If RequireTld is set to false, email addresses without a TLD in their domain will also be matched.
// If IgnoreMaxLength is set to true, the validator will not check for the standard max length of an email.
// If AllowIpDomain is set to true, the validator will allow IP addresses in the host part.
// If DomainSpecificValidation is true, some additional validation will be enabled, e.g. disallowing certain syntactically valid email addresses that are rejected by Gmail.
// If BlacklistedChars receives a string, then the validator will reject emails that include any of the characters in the string, in the name part.
// If HostBlacklist is set to an array of strings and the part of the email after the @ symbol matches one of the strings defined in it, the validation fails.
// If HostWhitelist is set to an array of strings and the part of the email after the @ symbol matches none of the strings defined in it, the validation fails.
func IsEmail(str string, opts IsEmailOpts) bool {
	// Basic email regex, allowing for UTF-8 if specified
	var emailRegex *regexp.Regexp
	if opts.AllowUTF8LocalPart {
		emailRegex = regexp.MustCompile(`^[\p{L}0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}|(\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])$`)
	} else {
		emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}|(\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])$`)
	}

	// Handle display name if required
	if opts.RequireDisplayName || opts.AllowDisplayName {
		displayNameRegex := regexp.MustCompile(`^.+\s<(.+)>$`)
		matches := displayNameRegex.FindStringSubmatch(str)
		if len(matches) == 2 {
			str = matches[1]
		} else if opts.RequireDisplayName {
			return false
		}
	}

	// Basic email validation
	if !emailRegex.MatchString(str) {
		fmt.Println(emailRegex.String())
		fmt.Println("Basic email validation")
		return false
	}

	// Extract the domain part
	parts := strings.Split(str, "@")
	domain := parts[1]

	// Handle IP domain if allowed
	if opts.AllowIpDomain && strings.HasPrefix(domain, "[") && strings.HasSuffix(domain, "]") {
		ip := domain[1 : len(domain)-1]
		if net.ParseIP(ip) != nil {
			return true
		}
		return false
	}

	// Handle TLD requirement
	if opts.RequireTld && !strings.Contains(domain, ".") {
		return false
	}

	// Domain-specific validation
	if opts.DomainSpecificValidation && domain == "gmail.com" {
		local := parts[0]
		if strings.Contains(local, "+") {
			return false
		}
	}

	// Blacklisted characters
	if opts.BlacklistedChars != "" {
		for _, char := range opts.BlacklistedChars {
			if strings.ContainsRune(str, char) {
				return false
			}
		}
	}

	// Host blacklist/whitelist
	if len(opts.HostBlacklist) > 0 {
		for _, blacklisted := range opts.HostBlacklist {
			if strings.HasSuffix(domain, blacklisted) {
				return false
			}
		}
	}
	if len(opts.HostWhitelist) > 0 {
		allowed := false
		for _, whitelisted := range opts.HostWhitelist {
			if strings.HasSuffix(domain, whitelisted) {
				allowed = true
				break
			}
		}
		if !allowed {
			return false
		}
	}

	return true
}
