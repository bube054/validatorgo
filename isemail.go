package validatorgo

import (
	"net"
	"regexp"
	"strings"
)

// IsEmailOpts is used to configure IsEmail
type IsEmailOpts struct {
	AllowDisplayName         *bool
	RequireDisplayName       *bool
	AllowUTF8LocalPart       *bool
	RequireTld               *bool
	IgnoreMaxLength          *bool
	AllowIpDomain            *bool
	DomainSpecificValidation *bool
	BlacklistedChars         *string
	HostBlacklist            []string
	HostWhitelist            []string
}

func (o *IsEmailOpts) mergeDefaults() {
	if o.AllowDisplayName == nil {
		o.AllowDisplayName = Bool(false)
	}
	if o.RequireDisplayName == nil {
		o.RequireDisplayName = Bool(false)
	}
	if o.AllowUTF8LocalPart == nil {
		o.AllowUTF8LocalPart = Bool(true)
	}
	if o.RequireTld == nil {
		o.RequireTld = Bool(true)
	}
	if o.IgnoreMaxLength == nil {
		o.IgnoreMaxLength = Bool(false)
	}
	if o.AllowIpDomain == nil {
		o.AllowIpDomain = Bool(false)
	}
	if o.DomainSpecificValidation == nil {
		o.DomainSpecificValidation = Bool(false)
	}
	if o.BlacklistedChars == nil {
		o.BlacklistedChars = String("")
	}
}

// A validator that checks if the string is an email.
//
// IsEmailOpts is a struct which defaults to { AllowDisplayName: false, RequireDisplayName: false, AllowUTF8LocalPart: true, RequireTld: true, AllowIpDomain: false, isEmailDefaultIgnoreMaxLength: false, DomainSpecificValidation: false, BlacklistedChars: "", HostBlacklist: [] }.
//
// AllowDisplayName: if set to true, the validator will also match Display Name <email-address>.
// RequireDisplayName: if set to true, the validator will reject strings without the format Display Name <email-address>.
// AllowUTF8LocalPart: if set to false, the validator will not allow any non-English UTF8 character in email address' local part.
// RequireTld: if set to false, email addresses without a TLD in their domain will also be matched.
// IgnoreMaxLength: if set to true, the validator will not check for the standard max length of an email.
// AllowIpDomain: if set to true, the validator will allow IP addresses in the host part.
// DomainSpecificValidation: is true, some additional validation will be enabled, e.g. disallowing certain syntactically valid email addresses that are rejected by Gmail.
// BlacklistedChars: receives a string, then the validator will reject emails that include any of the characters in the string, in the name part.
// HostBlacklist: if set to an slice of strings and the part of the email after the @ symbol matches one of the strings defined in it, the validation fails.
// HostWhitelist: if set to an slice of strings and the part of the email after the @ symbol matches none of the strings defined in it, the validation fails.
//
//	ok, _ := validatorgo.IsEmail("test@example.com", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsEmail("user.example.com", nil)
//	fmt.Println(ok) // false
func IsEmail(str string, opts *IsEmailOpts) (bool, error) {
	if opts == nil {
		opts = &IsEmailOpts{}
	}
	optsCopy := *opts // fill in defaults on a copy, never on the caller's struct
	opts = &optsCopy
	opts.mergeDefaults()

	// Basic email regex, allowing for UTF-8 if specified. The two domain forms, a
	// hostname or a bracketed IP, are grouped so ^ and $ apply to the whole address:
	// without the group, "a@example.com<script>" matched and "junk [1.2.3.4]" passed
	// with no @ at all.
	var emailRegex *regexp.Regexp
	if *opts.AllowUTF8LocalPart {
		emailRegex = compileOnce(`^[\p{L}0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@(([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}|\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])$`)
	} else {
		emailRegex = compileOnce(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@(([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}|\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])$`)
	}

	// Handle display name if required
	if *opts.RequireDisplayName || *opts.AllowDisplayName {
		displayNameRegex := compileOnce(`^.+\s<(.+)>$`)
		matches := displayNameRegex.FindStringSubmatch(str)
		if len(matches) == 2 {
			str = matches[1]
		} else if *opts.RequireDisplayName {
			return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
		}
	}

	// Basic email validation
	if !emailRegex.MatchString(str) {
		// fmt.Println(emailRegex.String())
		// fmt.Println("Basic email validation")
		return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
	}

	// Split into the part before and after the @. The regex above guarantees there is
	// exactly one.
	at := strings.LastIndexByte(str, '@')
	local, domain := str[:at], str[at+1:]

	// Same limits as validator.js: at most 64 bytes before the @ and 254 after it
	if !*opts.IgnoreMaxLength && (len(local) > 64 || len(domain) > 254) {
		return false, newValidationError("IsEmail", ErrTooLong, "email is too long")
	}

	// Dots before the @ may not come first, last or twice in a row
	if strings.HasPrefix(local, ".") || strings.HasSuffix(local, ".") || strings.Contains(local, "..") {
		return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
	}

	// An IP domain such as [192.168.0.1] is only accepted with AllowIpDomain
	if strings.HasPrefix(domain, "[") && strings.HasSuffix(domain, "]") {
		if !*opts.AllowIpDomain {
			return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
		}
		ip := domain[1 : len(domain)-1]
		if net.ParseIP(ip) != nil {
			return true, nil
		}
		return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
	}

	// Handle TLD requirement
	if *opts.RequireTld && !strings.Contains(domain, ".") {
		return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
	}

	// Domain-specific validation
	if *opts.DomainSpecificValidation && domain == "gmail.com" {
		if strings.Contains(local, "+") {
			return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
		}
	}

	// Blacklisted characters
	if *opts.BlacklistedChars != "" {
		for _, char := range *opts.BlacklistedChars {
			if strings.ContainsRune(str, char) {
				return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
			}
		}
	}

	// Host blacklist/whitelist
	if len(opts.HostBlacklist) > 0 {
		for _, blacklisted := range opts.HostBlacklist {
			if strings.HasSuffix(domain, blacklisted) {
				return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
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
			return false, newValidationError("IsEmail", ErrInvalidFormat, "invalid email")
		}
	}

	return true, nil
}
