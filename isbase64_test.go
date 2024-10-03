package validatorgo

import (
	"testing"
)

func TestIsBase64(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsBase64Opts
		want   bool
	}{
		// Valid Base64 strings
		{name: "Valid Base64 standard", param1: "SGVsbG8gd29ybGQ=", param2: &IsBase64Opts{UrlSafe: false}, want: true},
		{name: "Valid Base64 URL-safe", param1: "SGVsbG8gd29ybGQ", param2: &IsBase64Opts{UrlSafe: true}, want: true},
		{name: "Valid Base64 URL-safe with padding", param1: "SGVsbG8gd29ybGQ=", param2: &IsBase64Opts{UrlSafe: true}, want: true},
		{name: "Valid Base64 standard with padding", param1: "U29tZSB0ZXN0IHN0cmluZw==", param2: &IsBase64Opts{UrlSafe: false}, want: true},
		{name: "Valid Base64 standard with nil config", param1: "SGVsbG8gd29ybGQ=", param2: nil, want: true},

		// Invalid Base64 strings
		{name: "Invalid Base64 string with special characters", param1: "SGVsbG8g@d29ybGQ=", param2: &IsBase64Opts{UrlSafe: false}, want: false},
		{name: "Invalid Base64 string with space", param1: "SGVsbG8gd29y bGQ=", param2: &IsBase64Opts{UrlSafe: false}, want: false},
		{name: "Invalid Base64 URL-safe with unsafe characters", param1: "SGVsbG8gd29ybGQ$", param2: &IsBase64Opts{UrlSafe: true}, want: false},
		{name: "Empty string", param1: "", param2: &IsBase64Opts{UrlSafe: false}, want: false},
		{name: "Invalid Base64 with special characters and nil config", param1: "SGVsbG8g@d29ybGQ=", param2: nil, want: false},
		{name: "Empty string with nil config", param1: "", param2: nil, want: false},

		// Edge cases
		{name: "Valid URL-safe Base64 without padding", param1: "U29tZVRlc3RTdHJpbmc", param2: &IsBase64Opts{UrlSafe: true}, want: true},
		{name: "Valid Base64 standard with numbers", param1: "MTIzNDU2Nzg5MA==", param2: &IsBase64Opts{UrlSafe: false}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBase64(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
