package validatorgo

import "testing"

func TestIsULID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid ULID
		{name: "Valid ULID - All Uppercase", param1: "01ARZ3NDEKTSV4RRFFQ69G5FAV", want: true},                   // Proper ULID format
		{name: "Valid ULID - Starts with Timestamp", param1: "01B3EAFYJXG24N2H5YAFVYTHZV", want: true},           // Correct ULID structure
		{name: "Valid ULID - Mixed with Timestamp and Random", param1: "01BX5ZZKBKACTAV9WEVGEMMVRY", want: true}, // Valid timestamp and random parts

		// Invalid ULID
		{name: "Invalid ULID - Too Short", param1: "01ARZ3NDEKTSV4RRFFQ69G5FA", want: false},                    // Only 25 characters (1 missing)
		{name: "Invalid ULID - Contains Invalid Characters", param1: "01ARZ3NDEKTSV4RRFFQ69G5FA@", want: false}, // Invalid character '@'
		{name: "Invalid ULID - Lowercase Letters", param1: "01arz3ndektsv4rrffq69g5fav", want: false},           // Lowercase letters are not allowed
		{name: "Invalid ULID - Non-base32 Characters", param1: "01ARZ3NDEKTSV4RRFFQ69G5FAL", want: false},       // Contains 'L', which is not allowed in Crockford's Base32
		{name: "Invalid ULID - Empty String", param1: "", want: false},                                          // Empty string
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsULID(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
