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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsEmail(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
