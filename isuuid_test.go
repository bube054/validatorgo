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

		// validator.js: Version 3
		{name: "Valid UUID v3 (validator.js)", param1: "9deb20fe-a6e0-355c-81ea-288b009e4f6d", param2: "3", want: true},
		{name: "Invalid UUID v3 variant bits wrong (validator.js)", param1: "A987FBC9-4BED-3078-CF07-9141BA07C9F3", param2: "3", want: false},
		{name: "Invalid UUID v3 empty string (validator.js)", param1: "", param2: "3", want: false},
		{name: "Invalid UUID v3 numeric string (validator.js)", param1: "934859", param2: "3", want: false},
		{name: "Invalid UUID v3 is actually v4 (validator.js)", param1: "A987FBC9-4BED-4078-8F07-9141BA07C9F3", param2: "3", want: false},

		// validator.js: Version 4
		{name: "Valid UUID v4 #1 (validator.js)", param1: "713ae7e3-cb32-45f9-adcb-7c4fa86b90c1", param2: "4", want: true},
		{name: "Valid UUID v4 #2 (validator.js)", param1: "625e63f3-58f5-40b7-83a1-a72ad31acffb", param2: "4", want: true},
		{name: "Valid UUID v4 #3 (validator.js)", param1: "57b73598-8764-4ad0-a76a-679bb6640eb1", param2: "4", want: true},
		{name: "Valid UUID v4 #4 (validator.js)", param1: "9c858901-8a57-4791-81fe-4c455b099bc9", param2: "4", want: true},
		{name: "Invalid UUID v4 empty string (validator.js)", param1: "", param2: "4", want: false},
		{name: "Invalid UUID v4 numeric string (validator.js)", param1: "934859", param2: "4", want: false},
		{name: "Invalid UUID v4 is actually v5 (validator.js)", param1: "A987FBC9-4BED-5078-AF07-9141BA07C9F3", param2: "4", want: false},
		{name: "Invalid UUID v4 is actually v3 (validator.js)", param1: "A987FBC9-4BED-3078-CF07-9141BA07C9F3", param2: "4", want: false},

		// validator.js: Version 5
		{name: "Valid UUID v5 #1 (validator.js)", param1: "987FBC97-4BED-5078-AF07-9141BA07C9F3", param2: "5", want: true},
		{name: "Valid UUID v5 #2 (validator.js)", param1: "987FBC97-4BED-5078-BF07-9141BA07C9F3", param2: "5", want: true},
		{name: "Valid UUID v5 #3 (validator.js)", param1: "987FBC97-4BED-5078-8F07-9141BA07C9F3", param2: "5", want: true},
		{name: "Valid UUID v5 #4 (validator.js)", param1: "987FBC97-4BED-5078-9F07-9141BA07C9F3", param2: "5", want: true},
		{name: "Invalid UUID v5 empty string (validator.js)", param1: "", param2: "5", want: false},
		{name: "Invalid UUID v5 numeric string (validator.js)", param1: "934859", param2: "5", want: false},
		{name: "Invalid UUID v5 is actually v4 (validator.js)", param1: "9c858901-8a57-4791-81fe-4c455b099bc9", param2: "5", want: false},

		// validator.js: Version 1
		{name: "Valid UUID v1 (validator.js)", param1: "E034B584-7D89-11E9-9669-1AECF481A97B", param2: "1", want: true},
		{name: "Invalid UUID v1 is actually v4 (validator.js)", param1: "A987FBC9-4BED-4078-8F07-9141BA07C9F3", param2: "1", want: false},
		{name: "Invalid UUID v1 is actually v5 (validator.js)", param1: "A987FBC9-4BED-5078-AF07-9141BA07C9F3", param2: "1", want: false},

		// validator.js: Default/All (empty string = any version)
		{name: "Valid UUID any v3 (validator.js)", param1: "9deb20fe-a6e0-355c-81ea-288b009e4f6d", param2: "", want: true},
		{name: "Valid UUID any v4 (validator.js)", param1: "A987FBC9-4BED-4078-8F07-9141BA07C9F3", param2: "", want: true},
		{name: "Valid UUID any v5 (validator.js)", param1: "A987FBC9-4BED-5078-AF07-9141BA07C9F3", param2: "", want: true},
		{name: "Invalid UUID any empty string (validator.js)", param1: "", param2: "", want: false},
		{name: "Invalid UUID any extra prefix (validator.js)", param1: "xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", param2: "", want: false},
		{name: "Invalid UUID any no dashes (validator.js)", param1: "A987FBC94BED3078CF079141BA07C9F3", param2: "", want: false},
		{name: "Invalid UUID any numeric string (validator.js)", param1: "934859", param2: "", want: false},
		{name: "Invalid UUID any invalid char G (validator.js)", param1: "AAAAAAAA-1111-1111-AAAG-111111111111", param2: "", want: false},

		// validator.js: Nil UUID - TODO: nil UUID has version 0 and variant 0, not matched by any version regex
		{name: "Nil UUID any version (validator.js)", param1: "00000000-0000-0000-0000-000000000000", param2: "", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsUUID(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
