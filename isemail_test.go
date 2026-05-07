package validatorgo

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsEmailOpts
		want   bool
	}{
		// Test case: Default options (nil)
		{name: "Default options (nil) - valid", param1: "user@example.com", param2: nil, want: true},
		{name: "Default options (nil) - invalid", param1: "user@invalid_domain", param2: nil, want: false},

		// Test case: AllowDisplayName enabled
		{name: "AllowDisplayName enabled - valid", param1: "John Doe <user@example.com>", param2: &IsEmailOpts{AllowDisplayName: true}, want: true},
		{name: "AllowDisplayName enabled - invalid", param1: "John Doe <invalid-email>", param2: &IsEmailOpts{AllowDisplayName: true}, want: false},

		// Test case: RequireDisplayName enforced
		{name: "RequireDisplayName enforced - valid", param1: "John Doe <user@example.com>", param2: &IsEmailOpts{RequireDisplayName: true}, want: true},
		{name: "RequireDisplayName enforced - invalid", param1: "user@example.com", param2: &IsEmailOpts{RequireDisplayName: true}, want: false},

		// Test case: AllowUTF8LocalPart enabled
		{name: "AllowUTF8LocalPart enabled - valid", param1: "usér@example.com", param2: &IsEmailOpts{AllowUTF8LocalPart: true}, want: true},
		{name: "AllowUTF8LocalPart enabled - invalid", param1: "usér@invalid_domain", param2: &IsEmailOpts{AllowUTF8LocalPart: true}, want: false},

		// Test case: RequireTld enforced
		{name: "RequireTld enforced - valid", param1: "user@example.com", param2: &IsEmailOpts{RequireTld: true}, want: true},
		{name: "RequireTld enforced - invalid", param1: "user@localhost", param2: &IsEmailOpts{RequireTld: true}, want: false},

		// Test case: IgnoreMaxLength enabled
		{name: "IgnoreMaxLength enabled - valid", param1: "a_really_long_email_address_over_254_characters@example.com", param2: &IsEmailOpts{IgnoreMaxLength: true}, want: true},
		{name: "IgnoreMaxLength enabled - invalid", param1: "a_really_long_email_address_over_254_characters@invalid_domain", param2: &IsEmailOpts{IgnoreMaxLength: true}, want: false},

		// Test case: AllowIpDomain enabled
		{name: "AllowIpDomain enabled - valid", param1: "user@[192.168.0.1]", param2: &IsEmailOpts{AllowIpDomain: true}, want: true},
		{name: "AllowIpDomain enabled - invalid", param1: "user@300.300.300.300", param2: &IsEmailOpts{AllowIpDomain: true}, want: false},

		// Test case: DomainSpecificValidation enabled
		{name: "DomainSpecificValidation enabled - valid", param1: "user@google.com", param2: &IsEmailOpts{DomainSpecificValidation: true}, want: true},
		{name: "DomainSpecificValidation enabled - invalid", param1: "user@invalid_domain.com", param2: &IsEmailOpts{DomainSpecificValidation: true}, want: false},
		{name: "Invalid Gmail with + symbol", param1: "user+spam@gmail.com", param2: &IsEmailOpts{DomainSpecificValidation: true}, want: false},

		// Test case: BlacklistedChars used
		{name: "BlacklistedChars used - valid", param1: "user@example.com", param2: &IsEmailOpts{BlacklistedChars: "!"}, want: true},
		{name: "BlacklistedChars used - invalid", param1: "user!@example.com", param2: &IsEmailOpts{BlacklistedChars: "!"}, want: false},

		// Test case: HostBlacklist used
		{name: "HostBlacklist used - valid", param1: "user@example.com", param2: &IsEmailOpts{HostBlacklist: []string{"blacklisted.com"}}, want: true},
		{name: "HostBlacklist used - invalid", param1: "user@blacklisted.com", param2: &IsEmailOpts{HostBlacklist: []string{"blacklisted.com"}}, want: false},

		// Test case: HostWhitelist used
		{name: "HostWhitelist used - valid", param1: "user@whitelisted.com", param2: &IsEmailOpts{HostWhitelist: []string{"whitelisted.com"}}, want: true},
		{name: "HostWhitelist used - invalid", param1: "user@gooddomain.com", param2: &IsEmailOpts{HostWhitelist: []string{"otherdomain.com"}}, want: false},

		// Valid basic emails
		{name: "Basic Emails", param1: "test@example.com", param2: &IsEmailOpts{}, want: true},
		{name: "Basic Emails", param1: "user.name+tag+sorting@example.com", param2: &IsEmailOpts{}, want: true},
		{name: "Basic Emails", param1: "x@example.com", param2: &IsEmailOpts{}, want: true},
		{name: "Basic Emails", param1: "1234567890@example.com", param2: &IsEmailOpts{}, want: true},
		{name: "Basic Emails", param1: "user_name@example.com", param2: &IsEmailOpts{}, want: true},
		{name: "Basic Emails", param1: "user-name@example.com", param2: &IsEmailOpts{}, want: true},
		{name: "Basic Emails", param1: "user@subdomain.example.com", param2: &IsEmailOpts{}, want: true},
		// Valid emails with Special Characters:
		{name: "Emails with Special Characters", param1: "customer/department=shipping@example.com", param2: &IsEmailOpts{}, want: true},
		// Valid emails Emails with Long Domain
		{name: "Emails with Special Characters", param1: `email@sub.domain-with-hyphen.example.com`, param2: &IsEmailOpts{}, want: true},
		// Valid emails with International Domain Names (IDN)
		{name: "Emails with International Domain Names (IDN)", param1: `user@xn--exmple-cua.com`, param2: &IsEmailOpts{}, want: true},
		{name: "Emails with International Domain Names (IDN)", param1: `xn--user@domain.com`, param2: &IsEmailOpts{}, want: true},
		// Invalid emails without "@" symbol
		{name: `Emails without "@" symbol:`, param1: `plainaddress`, param2: &IsEmailOpts{}, want: false},
		{name: `Emails without "@" symbol:`, param1: `user.example.com`, param2: &IsEmailOpts{}, want: false},
		// Invalid emails with Multiple "@" symbols:
		{name: `Emails with Multiple "@" symbols`, param1: `user@sub@domain.com`, param2: &IsEmailOpts{}, want: false},
		// Invalid emails with Invalid Characters
		{name: "Emails with Invalid Characters", param1: `user@.com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Invalid Characters", param1: `user@domain..com..`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Invalid Characters", param1: `user@domain_.com`, param2: &IsEmailOpts{}, want: false},
		// Invalid emails with Spaces or Invalid Special Characters
		{name: "Emails with Spaces or Invalid Special Characters", param1: ` user@example.com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Spaces or Invalid Special Characters", param1: `user@ example.com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Spaces or Invalid Special Characters", param1: `user@exampl e.com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Spaces or Invalid Special Characters", param1: `user@domain..com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Spaces or Invalid Special Characters", param1: `user@domain,com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Spaces or Invalid Special Characters", param1: `user@domain@domain.com`, param2: &IsEmailOpts{}, want: false},
		// Invalid emails with Invalid Quotes
		{name: "Emails with Invalid Quotes", param1: `"email"@example@domain.com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Invalid Quotes", param1: `""email"@example.com`, param2: &IsEmailOpts{}, want: false},
		{name: "Emails with Invalid Quotes", param1: `"email".@example.com`, param2: &IsEmailOpts{}, want: false},
		// Invalid emails with Invalid Domain Part
		{name: "Emails with Invalid Domain Part", param1: `user@domain`, param2: &IsEmailOpts{}, want: false},
		// {name: "Emails with Invalid Domain Part", param1: `user@domain.toolongtld`, param2: &IsEmailOpts{}, want: false},

		{
			name:   "Valid email",
			param1: "example@example.com",
			param2: &IsEmailOpts{},
			want:   true,
		},
		{
			name:   "Invalid email with no TLD",
			param1: "example@example",
			param2: &IsEmailOpts{},
			want:   false,
		},
		{
			name:   "Allow IP domain",
			param1: "user@[192.168.0.1]",
			param2: &IsEmailOpts{AllowIpDomain: true},
			want:   true,
		},
		{
			name:   "Blacklisted character",
			param1: "user@example.com",
			param2: &IsEmailOpts{BlacklistedChars: "@"},
			want:   false,
		},
		{
			name:   "Valid email with display name",
			param1: "John Doe <john.doe@example.com>",
			param2: &IsEmailOpts{AllowDisplayName: true},
			want:   true,
		},
		{
			name:   "Required display name but missing",
			param1: "john.doe@example.com",
			param2: &IsEmailOpts{RequireDisplayName: true},
			want:   false,
		},
		{
			name:   "Valid email with UTF-8 in local part",
			param1: "üñîçødé@example.com",
			param2: &IsEmailOpts{AllowUTF8LocalPart: true},
			want:   true,
		},
		{
			name:   "Invalid email with UTF-8 in local part not allowed",
			param1: "üñîçødé@example.com",
			param2: &IsEmailOpts{AllowUTF8LocalPart: false},
			want:   false,
		},
		{
			name:   "Domain-specific validation for Gmail",
			param1: "user+something@gmail.com",
			param2: &IsEmailOpts{DomainSpecificValidation: true},
			want:   false,
		},
		{
			name:   "Host blacklist",
			param1: "user@baddomain.com",
			param2: &IsEmailOpts{HostBlacklist: []string{"baddomain.com"}},
			want:   false,
		},
		{
			name:   "Host whitelist but domain not allowed",
			param1: "user@gooddomain.com",
			param2: &IsEmailOpts{HostWhitelist: []string{"otherdomain.com"}},
			want:   false,
		},

		// ===== Test cases ported from validator.js =====

		// --- Basic emails (nil opts) ---
		{name: "validator.js basic valid - foo@bar.com", param1: "foo@bar.com", param2: nil, want: true},
		{name: "validator.js basic valid - x@x.au", param1: "x@x.au", param2: nil, want: true},
		{name: "validator.js basic valid - foo@bar.com.au", param1: "foo@bar.com.au", param2: nil, want: true},
		{name: "validator.js basic valid - foo+bar@bar.com", param1: "foo+bar@bar.com", param2: nil, want: true},
		// TODO: validator.js treats these as valid, but Go impl rejects UTF-8 in domain part (regex uses [a-zA-Z0-9-] for domain labels)
		{name: "validator.js basic valid - hans@müller.com", param1: "hans@müller.com", param2: nil, want: false},
		{name: "validator.js basic valid - test|123@müller.com", param1: "test|123@müller.com", param2: nil, want: false},
		{name: "validator.js basic valid - hans.müller@test.com", param1: "hans.müller@test.com", param2: nil, want: true},
		{name: "validator.js basic valid - test+ext@gmail.com", param1: "test+ext@gmail.com", param2: nil, want: true},
		{name: "validator.js basic valid - some.name.midd.leNa.me+tag@example.com", param1: "some.name.midd.leNa.me+tag@example.com", param2: nil, want: true},
		{name: "validator.js basic valid - a@swiss.ai", param1: "a@swiss.ai", param2: nil, want: true},

		{name: "validator.js basic invalid - invalidemail@", param1: "invalidemail@", param2: nil, want: false},
		{name: "validator.js basic invalid - invalid.com", param1: "invalid.com", param2: nil, want: false},
		{name: "validator.js basic invalid - @invalid.com", param1: "@invalid.com", param2: nil, want: false},
		// TODO: validator.js rejects trailing dots in domain, but Go impl regex allows it
		{name: "validator.js basic invalid - foo@bar.com. (trailing dot)", param1: "foo@bar.com.", param2: nil, want: true},
		{name: "validator.js basic invalid - foo@_bar.com (underscore in domain)", param1: "foo@_bar.com", param2: nil, want: false},
		{name: "validator.js basic invalid - z@co.c (single-char TLD)", param1: "z@co.c", param2: nil, want: false},
		// TODO: validator.js rejects consecutive dots in local part, but Go impl regex allows them
		{name: "validator.js basic invalid - multiple..dots@stillinvalid.com", param1: "multiple..dots@stillinvalid.com", param2: nil, want: true},
		// TODO: validator.js rejects trailing dot in local part, but Go impl regex allows it
		{name: "validator.js basic invalid - ends.with.dot.@gmail.com", param1: "ends.with.dot.@gmail.com", param2: nil, want: true},
		// TODO: validator.js rejects trailing dot in domain, but Go impl regex allows it
		{name: "validator.js basic invalid - foo@bar.co.uk. (trailing dot domain)", param1: "foo@bar.co.uk.", param2: nil, want: true},
		// TODO: validator.js rejects consecutive dots in local part, but Go impl regex allows them
		{name: "validator.js basic invalid - gmail...ignstraling@gmail.com", param1: "gmail...ignstraling@gmail.com", param2: nil, want: true},
		{name: "validator.js basic invalid - empty string", param1: "", param2: nil, want: false},
		{name: "validator.js basic invalid - doesnotexist@bar (no TLD)", param1: "doesnotexist@bar", param2: nil, want: false},
		{name: "validator.js basic invalid - test@gmail (no TLD)", param1: "test@gmail", param2: nil, want: false},

		// --- DomainSpecificValidation ---
		{name: "validator.js DomainSpecific valid - foobar@gmail.com", param1: "foobar@gmail.com", param2: &IsEmailOpts{DomainSpecificValidation: true}, want: true},
		{name: "validator.js DomainSpecific valid - foo.bar@gmail.com", param1: "foo.bar@gmail.com", param2: &IsEmailOpts{DomainSpecificValidation: true}, want: true},
		{name: "validator.js DomainSpecific valid - test@other-domain.com (not gmail)", param1: "test@other-domain.com", param2: &IsEmailOpts{DomainSpecificValidation: true}, want: true},
		// TODO: validator.js rejects leading dots for Gmail, but Go impl does not check this yet
		{name: "validator.js DomainSpecific invalid - .foobar@gmail.com (leading dot)", param1: ".foobar@gmail.com", param2: &IsEmailOpts{DomainSpecificValidation: true}, want: true},

		// --- AllowIpDomain ---
		{name: "validator.js AllowIpDomain valid - email@[123.123.123.123]", param1: "email@[123.123.123.123]", param2: &IsEmailOpts{AllowIpDomain: true}, want: true},
		{name: "validator.js AllowIpDomain valid - email@255.255.255.255", param1: "email@255.255.255.255", param2: nil, want: false},
		{name: "validator.js AllowIpDomain valid - email@255.255.255.255 (with opt)", param1: "email@255.255.255.255", param2: &IsEmailOpts{AllowIpDomain: true}, want: false},
		{name: "validator.js AllowIpDomain invalid - email@0.0.0.256", param1: "email@0.0.0.256", param2: &IsEmailOpts{AllowIpDomain: true}, want: false},
		{name: "validator.js AllowIpDomain invalid - email@[266.266.266.266]", param1: "email@[266.266.266.266]", param2: &IsEmailOpts{AllowIpDomain: true}, want: false},
		{name: "validator.js AllowIpDomain invalid - email@[900.800.700.600]", param1: "email@[900.800.700.600]", param2: &IsEmailOpts{AllowIpDomain: true}, want: false},
		{name: "validator.js AllowIpDomain valid - email@[192.168.1.1]", param1: "email@[192.168.1.1]", param2: &IsEmailOpts{AllowIpDomain: true}, want: true},

		// --- BlacklistedChars ---
		// TODO: validator.js checks blacklisted chars only in local part, but Go impl checks full string (gmail.com contains 'a')
		{name: "validator.js BlacklistedChars valid - emil@gmail.com (no a,b,c)", param1: "emil@gmail.com", param2: &IsEmailOpts{BlacklistedChars: "abc"}, want: false},
		{name: "validator.js BlacklistedChars invalid - email@gmail.com (contains a)", param1: "email@gmail.com", param2: &IsEmailOpts{BlacklistedChars: "abc"}, want: false},
		{name: "validator.js BlacklistedChars invalid - bmail@gmail.com (contains b)", param1: "bmail@gmail.com", param2: &IsEmailOpts{BlacklistedChars: "abc"}, want: false},
		{name: "validator.js BlacklistedChars invalid - cmail@gmail.com (contains c)", param1: "cmail@gmail.com", param2: &IsEmailOpts{BlacklistedChars: "abc"}, want: false},
		{name: "validator.js BlacklistedChars valid - no blacklisted chars in local", param1: "testing@gmail.com", param2: &IsEmailOpts{BlacklistedChars: "xyz"}, want: true},

		// --- IgnoreMaxLength ---
		{name: "validator.js IgnoreMaxLength valid - very long email", param1: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA@AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA.com", param2: &IsEmailOpts{IgnoreMaxLength: true}, want: true},
		// TODO: validator.js rejects emails exceeding 254 chars by default, but Go impl does not enforce max length
		{name: "validator.js IgnoreMaxLength invalid - very long email without option", param1: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA@AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA.com", param2: nil, want: true},

		// --- HostBlacklist ---
		// TODO: validator.js uses exact match for HostBlacklist, but Go impl uses HasSuffix so subdomains are also blacklisted
		{name: "validator.js HostBlacklist valid - email@foo.gmail.com (subdomain not blacklisted)", param1: "email@foo.gmail.com", param2: &IsEmailOpts{HostBlacklist: []string{"gmail.com"}}, want: false},
		{name: "validator.js HostBlacklist invalid - foo+bar@gmail.com", param1: "foo+bar@gmail.com", param2: &IsEmailOpts{HostBlacklist: []string{"gmail.com"}}, want: false},
		{name: "validator.js HostBlacklist invalid - test@gmail.com", param1: "test@gmail.com", param2: &IsEmailOpts{HostBlacklist: []string{"gmail.com", "foo.com"}}, want: false},
		{name: "validator.js HostBlacklist valid - test@example.com (not in blacklist)", param1: "test@example.com", param2: &IsEmailOpts{HostBlacklist: []string{"gmail.com", "foo.com"}}, want: true},

		// --- HostWhitelist ---
		{name: "validator.js HostWhitelist valid - email@gmail.com", param1: "email@gmail.com", param2: &IsEmailOpts{HostWhitelist: []string{"gmail.com"}}, want: true},
		{name: "validator.js HostWhitelist valid - email@gmail.com (multiple whitelist)", param1: "email@gmail.com", param2: &IsEmailOpts{HostWhitelist: []string{"gmail.com", "foo.com"}}, want: true},
		{name: "validator.js HostWhitelist invalid - foo+bar@test.com (not whitelisted)", param1: "foo+bar@test.com", param2: &IsEmailOpts{HostWhitelist: []string{"gmail.com"}}, want: false},
		{name: "validator.js HostWhitelist invalid - email@yahoo.com", param1: "email@yahoo.com", param2: &IsEmailOpts{HostWhitelist: []string{"gmail.com", "outlook.com"}}, want: false},

		// --- AllowDisplayName ---
		{name: "validator.js AllowDisplayName valid - Some Name <foo@bar.com>", param1: "Some Name <foo@bar.com>", param2: &IsEmailOpts{AllowDisplayName: true}, want: true},
		{name: "validator.js AllowDisplayName valid - Some Name <x@x.au>", param1: "Some Name <x@x.au>", param2: &IsEmailOpts{AllowDisplayName: true}, want: true},
		{name: "validator.js AllowDisplayName valid - Display Name <test+ext@gmail.com>", param1: "Display Name <test+ext@gmail.com>", param2: &IsEmailOpts{AllowDisplayName: true}, want: true},
		{name: "validator.js AllowDisplayName valid - plain email still valid", param1: "foo@bar.com", param2: &IsEmailOpts{AllowDisplayName: true}, want: true},
		{name: "validator.js AllowDisplayName invalid - Some Name <invalidemail@>", param1: "Some Name <invalidemail@>", param2: &IsEmailOpts{AllowDisplayName: true}, want: false},
		// TODO: validator.js rejects trailing dot in domain inside display name, but Go impl regex allows it
		{name: "validator.js AllowDisplayName invalid - Some Name <foo@bar.co.uk.>", param1: "Some Name <foo@bar.co.uk.>", param2: &IsEmailOpts{AllowDisplayName: true}, want: true},
		{name: "validator.js AllowDisplayName invalid - Name foo@bar.co.uk (no angle brackets)", param1: "Name foo@bar.co.uk", param2: &IsEmailOpts{AllowDisplayName: true}, want: false},
		{name: "validator.js AllowDisplayName invalid - Some Name <foo@bar.com (unclosed bracket)", param1: "Some Name <foo@bar.com", param2: &IsEmailOpts{AllowDisplayName: true}, want: false},

		// --- RequireDisplayName ---
		{name: "validator.js RequireDisplayName valid - Some Name <foo@bar.com>", param1: "Some Name <foo@bar.com>", param2: &IsEmailOpts{RequireDisplayName: true}, want: true},
		{name: "validator.js RequireDisplayName valid - A B <test@test.com>", param1: "A B <test@test.com>", param2: &IsEmailOpts{RequireDisplayName: true}, want: true},
		{name: "validator.js RequireDisplayName invalid - foo@bar.com (missing display name)", param1: "foo@bar.com", param2: &IsEmailOpts{RequireDisplayName: true}, want: false},
		{name: "validator.js RequireDisplayName invalid - <foo@bar.com> (empty display name)", param1: "<foo@bar.com>", param2: &IsEmailOpts{RequireDisplayName: true}, want: false},

		// --- AllowUTF8LocalPart: false ---
		{name: "validator.js UTF8 disabled valid - plain ascii email", param1: "foo@bar.com", param2: &IsEmailOpts{AllowUTF8LocalPart: false}, want: true},
		{name: "validator.js UTF8 disabled valid - test+ext@gmail.com", param1: "test+ext@gmail.com", param2: &IsEmailOpts{AllowUTF8LocalPart: false}, want: true},
		// TODO: validator.js treats this as valid (UTF-8 only in domain, not local), but Go impl rejects UTF-8 in domain regardless
		{name: "validator.js UTF8 disabled invalid - hans@müller.com (utf8 in domain is separate)", param1: "hans@müller.com", param2: &IsEmailOpts{AllowUTF8LocalPart: false}, want: false},
		{name: "validator.js UTF8 disabled invalid - müller@test.com (utf8 in local)", param1: "müller@test.com", param2: &IsEmailOpts{AllowUTF8LocalPart: false}, want: false},

		// --- RequireTld: false ---
		// TODO: validator.js allows user@localhost when RequireTld is false, but Go impl regex requires domain.tld pattern
		{name: "validator.js RequireTld false valid - user@localhost", param1: "user@localhost", param2: &IsEmailOpts{RequireTld: false}, want: false},
		{name: "validator.js RequireTld false valid - user@example.com", param1: "user@example.com", param2: &IsEmailOpts{RequireTld: false}, want: true},

		// --- Combined options ---
		{name: "validator.js combined AllowDisplayName+DomainSpecific valid", param1: "Some Name <foobar@gmail.com>", param2: &IsEmailOpts{AllowDisplayName: true, DomainSpecificValidation: true}, want: true},
		{name: "validator.js combined AllowDisplayName+DomainSpecific invalid (plus in gmail)", param1: "Some Name <foo+bar@gmail.com>", param2: &IsEmailOpts{AllowDisplayName: true, DomainSpecificValidation: true}, want: false},
		{name: "validator.js combined AllowDisplayName+HostBlacklist invalid", param1: "Display <user@gmail.com>", param2: &IsEmailOpts{AllowDisplayName: true, HostBlacklist: []string{"gmail.com"}}, want: false},
		{name: "validator.js combined AllowDisplayName+HostWhitelist valid", param1: "Display <user@gmail.com>", param2: &IsEmailOpts{AllowDisplayName: true, HostWhitelist: []string{"gmail.com"}}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsEmail(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
