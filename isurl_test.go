package validatorgo

import (
	"strings"
	"testing"
)

func TestIsURL(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsURLOpts
		want   bool
	}{
		// Test: Valid URL with default settings (http protocol, has TLD, valid)
		{
			name:   "Valid URL - http",
			param1: "http://example.com",
			param2: nil, // Defaults apply
			want:   true,
		},
		// Test: Valid URL with https protocol (allowed by default protocols)
		{
			name:   "Valid URL - https",
			param1: "https://example.com",
			param2: nil, // Defaults apply
			want:   true,
		},
		// Test: Valid URL with ftp protocol (allowed by default protocols)
		{
			name:   "Valid URL - ftp",
			param1: "ftp://example.com",
			param2: nil, // Defaults apply
			want:   true,
		},
		// Test: Invalid protocol (not in default protocols)
		{
			name:   "Invalid Protocol - mailto",
			param1: "mailto://example.com",
			param2: nil, // Defaults apply
			want:   false,
		},
		// Test: Missing TLD (valid if RequireTld is false)
		{
			name:   "Missing TLD",
			param1: "http://localhost",
			param2: &IsURLOpts{RequireTld: false, AllowProtocolRelativeUrls: true}, // Override to allow missing TLD
			want:   true,
		},
		// Test: Missing TLD (valid if RequireTld is false)
		{
			name:   "Missing TLD",
			param1: "http://localhost",
			param2: &IsURLOpts{RequireTld: true, AllowProtocolRelativeUrls: true}, // Override to allow missing TLD
			want:   false,
		},
		// Test: Require protocol (valid only if protocol is included)
		{
			name:   "Missing Protocol",
			param1: "example.com",
			param2: &IsURLOpts{RequireProtocol: true, AllowProtocolRelativeUrls: true}, // Require protocol
			want:   false,
		},
		// Test: Allow protocol-relative URL (//example.com)
		{
			name:   "Protocol-relative URL",
			param1: "//example.com",
			param2: &IsURLOpts{AllowProtocolRelativeUrls: true}, // Allow protocol-relative
			want:   true,
		},
		// Test: Invalid due to underscores in domain
		{
			name:   "Invalid with Underscore in Domain",
			param1: "http://example_domain.com",
			param2: nil, // Defaults apply (AllowUnderscores: false)
			want:   false,
		},
		// Test: Valid with underscores in domain (when allowed)
		{
			name:   "Valid with Underscore in Domain",
			param1: "http://example_domain.com",
			param2: &IsURLOpts{AllowUnderscores: true, AllowProtocolRelativeUrls: true}, // Allow underscores
			want:   true,
		},
		// Test: URL with trailing dot (invalid by default)
		{
			name:   "Invalid with Trailing Dot",
			param1: "http://example.com.",
			param2: nil, // Defaults apply (AllowTrailingDot: false)
			want:   false,
		},
		// Test: URL with trailing dot (valid when allowed)
		{
			name:   "Valid with Trailing Dot",
			param1: "http://example.com.",
			param2: &IsURLOpts{AllowTrailingDot: true, AllowProtocolRelativeUrls: true}, // Allow trailing dot
			want:   true,
		},
		// Test: Disallow query components (by default, query components are allowed)
		{
			name:   "Disallow Query Components",
			param1: "http://example.com/?key=value",
			param2: &IsURLOpts{AllowQueryComponents: false, AllowProtocolRelativeUrls: true}, // Disallow query components
			want:   false,
		},
		// Test: Disallow fragments (fragments allowed by default)
		{
			name:   "Disallow Fragments",
			param1: "http://example.com/#section",
			param2: &IsURLOpts{AllowFragments: false, AllowProtocolRelativeUrls: true}, // Disallow fragments
			want:   false,
		},
		// Test: URL with auth (valid when allowed by default)
		{
			name:   "Valid URL with Auth",
			param1: "http://user:pass@example.com",
			param2: nil, // Defaults apply (DisallowAuth: false)
			want:   true,
		},
		// Test: URL with auth (invalid when disallowed)
		{
			name:   "Invalid URL with Auth",
			param1: "http://user:pass@example.com",
			param2: &IsURLOpts{DisallowAuth: true, AllowProtocolRelativeUrls: true}, // Disallow authentication info
			want:   false,
		},
		// Test: Validate length (exceeds default limit)
		{
			name:   "Valid URL within default Length",
			param1: "http://" + strings.Repeat("example", 2048) + ".com",
			param2: &IsURLOpts{ValidateLength: true},
			want:   false,
		},
		// Test: Validate length (within allowed limit)
		{
			name:   "Valid URL within Length Limit",
			param1: "http://example.com/short-path",
			param2: &IsURLOpts{MaxAllowedLength: 50, ValidateLength: true, AllowProtocolRelativeUrls: true}, // Set max allowed length
			want:   true,
		},
		// Test: Validate length (exceeds allowed limit)
		{
			name:   "Invalid URL exceeds Length Limit",
			param1: "http://example.com/this-is-a-very-long-path-that-exceeds-the-length-limit-by-to-much-characters",
			param2: &IsURLOpts{MaxAllowedLength: 50, ValidateLength: true, AllowProtocolRelativeUrls: true}, // Set max allowed length
			want:   false,
		},
		// Test: Host is in whitelist (valid)
		{
			name:   "Valid - Host in Whitelist",
			param1: "http://example.com",
			param2: &IsURLOpts{HostWhitelist: []string{"example.com"}, AllowProtocolRelativeUrls: true}, // Only allow example.com
			want:   true,
		},
		// Test: Host not in whitelist (invalid)
		{
			name:   "Invalid - Host not in Whitelist",
			param1: "http://notallowed.com",
			param2: &IsURLOpts{HostWhitelist: []string{"example.com"}, AllowProtocolRelativeUrls: true}, // Only allow example.com
			want:   false,
		},
		// Test: Host is in blacklist (invalid)
		{
			name:   "Invalid - Host in Blacklist",
			param1: "http://badhost.com",
			param2: &IsURLOpts{HostBlacklist: []string{"badhost.com"}, AllowProtocolRelativeUrls: true}, // Disallow badhost.com
			want:   false,
		},
		// Test: Host not in blacklist (valid)
		{
			name:   "Valid - Host not in Blacklist",
			param1: "http://goodhost.com",
			param2: &IsURLOpts{HostBlacklist: []string{"badhost.com"}, AllowProtocolRelativeUrls: true}, // Disallow badhost.com
			want:   true,
		},
		// Test: Host is present (valid)
		{
			name:   "Valid - Host present",
			param1: "http://example.com",
			param2: &IsURLOpts{RequireHost: true, AllowProtocolRelativeUrls: true}, // Host is required
			want:   true,
		},
		// Test: Host is missing (invalid because RequireHost is true)
		{
			name:   "Invalid - Host missing",
			param1: "http:///path-only",
			param2: &IsURLOpts{RequireHost: true, AllowProtocolRelativeUrls: true}, // Host is required
			want:   false,
		},
		// Test: Host is missing but not required (valid)
		{
			name:   "Valid - Host missing but not required",
			param1: "http:///path-only",
			param2: &IsURLOpts{RequireHost: false, AllowProtocolRelativeUrls: true}, // Host is not required
			want:   true,
		},
		// Test: Port is present (valid)
		{
			name:   "Valid - Port present",
			param1: "http://example.com:8080",
			param2: &IsURLOpts{RequirePort: true, AllowProtocolRelativeUrls: true}, // Port is required
			want:   true,
		},
		// Test: Port is missing (invalid because RequirePort is true)
		{
			name:   "Invalid - Port missing",
			param1: "http://example.com",
			param2: &IsURLOpts{RequirePort: true, AllowProtocolRelativeUrls: true}, // Port is required
			want:   false,
		},
		// Test: Port is missing but not required (valid)
		{
			name:   "Valid - Port missing but not required",
			param1: "http://example.com",
			param2: &IsURLOpts{RequirePort: false, AllowProtocolRelativeUrls: true}, // Port is not required
			want:   true,
		},
		// Test: Not a url
		{
			name:   "Valid - Port missing but not required",
			param1: "xyz.abc.def",
			param2: nil, // Port is not required
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsURL(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
