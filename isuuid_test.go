package validatorgo

import "testing"

func TestIsUUID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Valid UUIDs for specific versions
		{name: "Valid UUID v1", param1: "550e8400-e29b-11d4-a716-446655440000", param2: "1", want: true},
		{name: "Valid UUID v2", param1: "6fa459ea-ee8a-2ca4-894e-db77e160355e", param2: "2", want: true},
		{name: "Valid UUID v3", param1: "f47ac10b-58cc-3372-a567-0e02b2c3d479", param2: "3", want: true},
		{name: "Valid UUID v4", param1: "f47ac10b-58cc-4372-a567-0e02b2c3d479", param2: "4", want: true},
		{name: "Valid UUID v5", param1: "f47ac10b-58cc-5372-a567-0e02b2c3d479", param2: "5", want: true},

		// Valid UUIDs without specifying a version
		{name: "Valid UUID any version (v1)", param1: "550e8400-e29b-11d4-a716-446655440000", param2: "", want: true},
		{name: "Valid UUID any version (v4)", param1: "f47ac10b-58cc-4372-a567-0e02b2c3d479", param2: "", want: true},

		// Invalid UUIDs for specific versions
		{name: "Invalid UUID for v1 (v4 format)", param1: "f47ac10b-58cc-4372-a567-0e02b2c3d479", param2: "1", want: false},
		{name: "Invalid UUID for v3 (v4 format)", param1: "f47ac10b-58cc-4372-a567-0e02b2c3d479", param2: "3", want: false},
		{name: "Invalid UUID for v5 (v3 format)", param1: "f47ac10b-58cc-3372-a567-0e02b2c3d479", param2: "5", want: false},

		// Invalid UUID due to incorrect structure
		{name: "Invalid UUID due to wrong structure (missing hyphen)", param1: "f47ac10b58cc4372a5670e02b2c3d479", param2: "", want: false},
		{name: "Invalid UUID due to incorrect length", param1: "f47ac10b-58cc-4372-a567-0e02b2c3d47", param2: "", want: false},

		// Invalid UUID due to invalid characters
		{name: "Invalid UUID with invalid characters", param1: "g47ac10b-58cc-4372-a567-0e02b2c3d479", param2: "", want: false},
		{name: "Invalid UUID with special characters", param1: "f47ac10b-58cc-4372-a567-0e02b2c3d479!", param2: "", want: false},

		// Edge cases with empty strings
		{name: "Empty string", param1: "", param2: "", want: false},
		{name: "Empty string with version", param1: "", param2: "4", want: false},

		// UUID for unsupported versions (outside of 1-5)
		{name: "Unsupported version (v6)", param1: "f47ac10b-58cc-6372-a567-0e02b2c3d479", param2: "6", want: false},
		{name: "Unsupported version (v8)", param1: "f47ac10b-58cc-8372-a567-0e02b2c3d479", param2: "8", want: false},

		// Valid UUID with version mismatch
		{name: "Valid UUID with version mismatch (v4 UUID but checking for v3)", param1: "f47ac10b-58cc-4372-a567-0e02b2c3d479", param2: "3", want: false},
		{name: "Valid UUID with version mismatch (v1 UUID but checking for v4)", param1: "550e8400-e29b-11d4-a716-446655440000", param2: "4", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsUUID(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
