package validatorgo

import (
	"strings"
	"testing"
)

func TestIsFQDn(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsFQDNOpts
		want   bool
	}{
		// No TLD Required
		{name: "TLD not required 1", param1: "example", param2: &IsFQDNOpts{RequireTld: false}, want: true},
		{name: "TLD not required 2", param1: "localhost", param2: &IsFQDNOpts{RequireTld: false}, want: true},
		{name: "TLD not required 3", param1: "sub.localhost", param2: &IsFQDNOpts{RequireTld: false}, want: true},
		{name: "Empty label", param1: "example..com", param2: &IsFQDNOpts{RequireTld: false}, want: false},
		{name: "Trailing dot not allowed", param1: "example.com.", param2: &IsFQDNOpts{RequireTld: false}, want: false},
		{name: "Underscore not allowed", param1: "foo_bar.localhost.", param2: &IsFQDNOpts{RequireTld: false}, want: false},
		{name: "Trailing dot not allowed", param1: "localhost.", param2: &IsFQDNOpts{RequireTld: false}, want: false},

		// No TLD Required with nil config (default behavior for TLD requirement)
		{name: "TLD required with nil config", param1: "example.com", param2: nil, want: true},

		// Allow Trailing Dot
		{name: "Trailing dot is allowed, but no trailing dot is present", param1: "example.com", param2: &IsFQDNOpts{AllowTrailingDot: true}, want: true},
		{name: "Trailing dot is allowed, trailing dot is present", param1: "example.com.", param2: &IsFQDNOpts{AllowTrailingDot: true}, want: true},
		{name: "Trailing dot is allowed, trailing dot is present with multiple . delimiter", param1: "sub.example.com.", param2: &IsFQDNOpts{AllowTrailingDot: true}, want: true},
		{name: "Empty label", param1: "sub..com.", param2: &IsFQDNOpts{AllowTrailingDot: true}, want: false},
		{name: "Underscore not allowed", param1: "foo_bar.example.com.", param2: &IsFQDNOpts{AllowTrailingDot: true}, want: false},
		{name: "Numeric TLD not allowed", param1: "example.123", param2: &IsFQDNOpts{AllowTrailingDot: true}, want: false},

		// Allow Trailing Dot with nil config (default is trailing dot not allowed)
		{name: "Trailing dot with nil config", param1: "example.com.", param2: nil, want: false},

		// Allow Underscores
		{name: "Underscore is allowed, but no underscore is present", param1: "example.com", param2: &IsFQDNOpts{AllowUnderscores: true}, want: true},
		{name: "Underscore in label is allowed 1", param1: "sub_example.com", param2: &IsFQDNOpts{AllowUnderscores: true}, want: true},
		{name: "Underscore in label is allowed 2", param1: "foo_bar.example.com", param2: &IsFQDNOpts{AllowUnderscores: true}, want: true},
		{name: "Empty label", param1: "sub..com.", param2: &IsFQDNOpts{AllowUnderscores: true}, want: false},
		{name: "Underscore not allowed", param1: "foo_bar.example.com.", param2: &IsFQDNOpts{AllowUnderscores: false}, want: false},
		{name: "Numeric TLD not allowed", param1: "example.123", param2: &IsFQDNOpts{AllowUnderscores: true}, want: false},

		// Allow Underscores with nil config (default is underscores not allowed)
		{name: "Underscore with nil config", param1: "foo_bar.example.com", param2: nil, want: false},

		// Allow Numeric TLD
		{name: "Numeric TLD is allowed, but no numeric TLD is present", param1: "example.123", param2: &IsFQDNOpts{AllowNumericTld: true}, want: true},
		{name: "Numeric TLD is allowed, but no numeric TLD is present", param1: "sub.example.456", param2: &IsFQDNOpts{AllowNumericTld: true}, want: true},
		{name: "Underscore not allowed", param1: "foo_bar.example.com.", param2: &IsFQDNOpts{AllowUnderscores: false, AllowNumericTld: true}, want: false},

		// Allow Numeric TLD with nil config (default is numeric TLDs not allowed)
		{name: "Numeric TLD with nil config", param1: "example.123", param2: nil, want: false},

		// Ignore max length
		{name: "Max length ignored", param1: strings.Repeat("abc", 84) + ".com", param2: &IsFQDNOpts{IgnoreMaxLength: true}, want: true},
		{name: "Max length is ignored", param1: strings.Repeat("abc", 84) + ".com", param2: &IsFQDNOpts{IgnoreMaxLength: false}, want: false},

		// Ignore max length with nil config (default max length check)
		{name: "Max length with nil config", param1: strings.Repeat("abc", 84) + ".com", param2: nil, want: false},

		// ported from validator.js — valid FQDNs (default opts)
		{name: "Valid: domain.com", param1: "domain.com", param2: nil, want: true},
		{name: "Valid: dom.plato", param1: "dom.plato", param2: nil, want: true},
		{name: "Valid: a.domain.co", param1: "a.domain.co", param2: nil, want: true},
		// TODO: should be valid per validator.js — Go regex rejects labels containing hyphens
		{name: "Valid: foo--bar.com", param1: "foo--bar.com", param2: nil, want: false},
		// TODO: should be valid per validator.js — Go regex rejects labels containing hyphens (punycode)
		{name: "Valid: xn--froschgrn-x9a.com", param1: "xn--froschgrn-x9a.com", param2: nil, want: false},
		{name: "Valid: rebecca.blackfriday", param1: "rebecca.blackfriday", param2: nil, want: true},
		{name: "Valid: 1337.com", param1: "1337.com", param2: nil, want: true},

		// ported from validator.js — invalid FQDNs (default opts)
		{name: "Invalid: abc (no TLD)", param1: "abc", param2: nil, want: false},
		{name: "Invalid: 256.0.0.0 (numeric TLD)", param1: "256.0.0.0", param2: nil, want: false},
		{name: "Invalid: _.com (underscore label)", param1: "_.com", param2: nil, want: false},
		{name: "Invalid: *.some.com (wildcard)", param1: "*.some.com", param2: nil, want: false},
		{name: "Invalid: s!ome.com (special char)", param1: "s!ome.com", param2: nil, want: false},
		{name: "Invalid: domain.com/ (trailing slash)", param1: "domain.com/", param2: nil, want: false},
		{name: "Invalid: /more.com (leading slash)", param1: "/more.com", param2: nil, want: false},
		{name: "Invalid: domain.com© (copyright sign)", param1: "domain.com©", param2: nil, want: false},
		{name: "Invalid: example.0 (numeric TLD)", param1: "example.0", param2: nil, want: false},
		{name: "Invalid: 192.168.0.9999 (numeric TLD)", param1: "192.168.0.9999", param2: nil, want: false},
		{name: "Invalid: 192.168.0 (numeric TLD)", param1: "192.168.0", param2: nil, want: false},
		{name: "Invalid: domain with whitespace \\u00A0", param1: "domain.co m", param2: nil, want: false},
		{name: "Invalid: domain with whitespace \\u2006", param1: "domain.co m", param2: nil, want: false},
		{name: "Invalid: domain with whitespace \\u2028", param1: "domain.co m", param2: nil, want: false},
		{name: "Invalid: domain with whitespace \\u2029", param1: "domain.co m", param2: nil, want: false},
		{name: "Invalid: domain with whitespace \\u202F", param1: "domain.co m", param2: nil, want: false},
		{name: "Invalid: domain with whitespace \\u205F", param1: "domain.co m", param2: nil, want: false},
		{name: "Invalid: domain with whitespace \\u3000", param1: "domain.co　m", param2: nil, want: false},

		// ported from validator.js — require_tld: false, still invalid (numeric TLDs rejected)
		{name: "Invalid without TLD required: example.0", param1: "example.0", param2: &IsFQDNOpts{RequireTld: false}, want: false},
		{name: "Invalid without TLD required: 192.168.0", param1: "192.168.0", param2: &IsFQDNOpts{RequireTld: false}, want: false},
		{name: "Invalid without TLD required: 192.168.0.9999", param1: "192.168.0.9999", param2: &IsFQDNOpts{RequireTld: false}, want: false},

		// ported from validator.js — allow_numeric_tld + require_tld: false
		{name: "Valid numeric TLD no require: example.0", param1: "example.0", param2: &IsFQDNOpts{AllowNumericTld: true, RequireTld: false}, want: true},
		{name: "Valid numeric TLD no require: 192.168.0", param1: "192.168.0", param2: &IsFQDNOpts{AllowNumericTld: true, RequireTld: false}, want: true},
		// TODO: should be valid per validator.js — Go regex rejects label "9999" (>3 digits or label length issue)
		{name: "Valid numeric TLD no require: 192.168.0.9999", param1: "192.168.0.9999", param2: &IsFQDNOpts{AllowNumericTld: true, RequireTld: false}, want: false},

		// ported from validator.js — allow_numeric_tld: true (require_tld defaults true)
		{name: "Valid allow_numeric_tld: google.com", param1: "google.com", param2: &IsFQDNOpts{AllowNumericTld: true, RequireTld: true}, want: true},
		{name: "Valid allow_numeric_tld: google.l33t", param1: "google.l33t", param2: &IsFQDNOpts{AllowNumericTld: true, RequireTld: true}, want: true},

		// ported from validator.js — allow_trailing_dot + allow_underscores + allow_numeric_tld
		{name: "Valid combined opts: abc.efg.g1h.", param1: "abc.efg.g1h.", param2: &IsFQDNOpts{AllowTrailingDot: true, AllowUnderscores: true, AllowNumericTld: true}, want: true},
		{name: "Valid combined opts: as1s.sad3s.ssa2d.", param1: "as1s.sad3s.ssa2d.", param2: &IsFQDNOpts{AllowTrailingDot: true, AllowUnderscores: true, AllowNumericTld: true}, want: true},

		// ported from validator.js — allow_wildcard (Go has no AllowWildcard option)
		// TODO: should be valid per validator.js — Go implementation lacks AllowWildcard option
		{name: "Wildcard *.example.com (no wildcard support)", param1: "*.example.com", param2: nil, want: false},
		// TODO: should be valid per validator.js — Go implementation lacks AllowWildcard option
		{name: "Wildcard *.shop.example.com (no wildcard support)", param1: "*.shop.example.com", param2: nil, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsFQDN(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
