package validatorgo

import (
	"strings"
	"testing"
)

func TestIsFQDn(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsFQDNopts
		want   bool
	}{
		// No TLD Required
		{name: "TLD not required 1", param1: "example", param2: IsFQDNopts{RequireTld: false}, want: true},
		{name: "TLD not required 2", param1: "localhost", param2: IsFQDNopts{RequireTld: false}, want: true},
		{name: "TLD not required 3", param1: "sub.localhost", param2: IsFQDNopts{RequireTld: false}, want: true},
		{name: "Empty label", param1: "example..com", param2: IsFQDNopts{RequireTld: false}, want: false},
		{name: "Trailing dot not allowed", param1: "example.com.", param2: IsFQDNopts{RequireTld: false}, want: false},
		{name: "Underscore not allowed", param1: "foo_bar.localhost.", param2: IsFQDNopts{RequireTld: false}, want: false},
		{name: "Trailing dot not allowed", param1: "localhost.", param2: IsFQDNopts{RequireTld: false}, want: false},

		// Allow Trailing Dot
		{name: "Trailing dot is allowed, but no trailing dot is present", param1: "example.com", param2: IsFQDNopts{AllowTrailingDot: true}, want: true},
		{name: "Trailing dot is allowed, trailing dot is present", param1: "example.com.", param2: IsFQDNopts{AllowTrailingDot: true}, want: true},
		{name: "Trailing dot is allowed, trailing dot is present with multiple . delimiter", param1: "sub.example.com.", param2: IsFQDNopts{AllowTrailingDot: true}, want: true},
		{name: "Empty label", param1: "sub..com.", param2: IsFQDNopts{AllowTrailingDot: true}, want: false},
		{name: "Underscore not allowed", param1: "foo_bar.example.com.", param2: IsFQDNopts{AllowTrailingDot: true}, want: false},
		{name: "Numeric TLD not allowed", param1: "example.123", param2: IsFQDNopts{AllowTrailingDot: true}, want: false},

		// Allow Underscores
		{name: "Underscore is allowed, but no underscore is present", param1: "example.com", param2: IsFQDNopts{AllowUnderscores: true}, want: true},
		{name: "Underscore in label is allowed 1", param1: "sub_example.com", param2: IsFQDNopts{AllowUnderscores: true}, want: true},
		{name: "Underscore in label is allowed 2", param1: "foo_bar.example.com", param2: IsFQDNopts{AllowUnderscores: true}, want: true},
		{name: "Empty label", param1: "sub..com.", param2: IsFQDNopts{AllowUnderscores: true}, want: false},
		{name: "Underscore not allowed", param1: "foo_bar.example.com.", param2: IsFQDNopts{AllowUnderscores: false}, want: false},
		{name: "Numeric TLD not allowed", param1: "example.123", param2: IsFQDNopts{AllowUnderscores: true}, want: false},

		// Allow Numeric TLD
		{name: "Numeric TLD is allowed, but no numeric TLD is present", param1: "example.123", param2: IsFQDNopts{AllowNumericTld: true}, want: true},
		{name: "Numeric TLD is allowed, but no numeric TLD is present", param1: "sub.example.456", param2: IsFQDNopts{AllowNumericTld: true}, want: true},
		{name: "Underscore not allowed", param1: "foo_bar.example.com.", param2: IsFQDNopts{AllowUnderscores: false, AllowNumericTld: true}, want: false},

		// Ignore max length
		{name: "Max length ignored", param1: strings.Repeat("abc", 84) + ".com", param2: IsFQDNopts{IgnoreMaxLength: true}, want: true},
		{name: "Max length is ignored", param1: strings.Repeat("abc", 84) + ".com", param2: IsFQDNopts{IgnoreMaxLength: false}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsFQDN(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
