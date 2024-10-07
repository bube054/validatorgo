package validatorgo

import "testing"

func TestIsFreightContainerID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid container ID (assuming correct check digit calculation)
		{name: "Valid container ID", param1: "CSQU3054383", want: true},
		{name: "Valid container ID", param1: "ABCU1234560", want: true},
		{name: "Valid container ID", param1: "MSKU6011672", want: true},

		// Invalid container ID (wrong format - less than 11 characters)
		{name: "Too short", param1: "ABC12345", want: false},
		// Invalid container ID (wrong format - more than 11 characters)
		{name: "Too long", param1: "ABCU123456789", want: false},
		// Invalid container ID (wrong format - invalid characters in prefix)
		{name: "Invalid characters in prefix", param1: "AB11234567", want: false},
		// Invalid container ID (invalid UJZ code)
		{name: "Invalid UJZ code", param1: "ABCD1234567", want: false}, // 'D' is not a valid UJZ code
		// Invalid container ID (non-numeric check digit)
		{name: "Non-numeric check digit", param1: "ABCU123456X", want: false},
		// Invalid container ID (invalid check digit, different valid UJZ code)
		{name: "Invalid check digit with U", param1: "ABCU1234561", want: false},
		// Edge case: all zeros in the numeric part
		{name: "All zeros in the numeric part", param1: "ABCU0000000", want: false}, // Check digit is likely incorrect
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsFreightContainerID(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
