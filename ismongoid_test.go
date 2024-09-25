package validatorgo

import "testing"

func TestIsMongoID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid mongo id
		{name: "Valid mongo id", param1: "507f1f77bcf86cd799439011", want: true},
		{name: "Valid mongo id", param1: "5f2a6c69e1d7a4e0077b4e6b", want: true},
		{name: "Valid mongo id", param1: "60ad7c6d5f98bf2e6c1453c7", want: true},
		{name: "Valid mongo id", param1: "000000000000000000000000", want: true},

		// Invalid mongo id
		{name: "Invalid mongo id, too short", param1: "507f1f77bcf86cd79943901", want: false},
		{name: "Invalid mongo id, too long", param1: "507f1f77bcf86cd7994390110", want: false},
		{name: "Invalid mongo id, contains invalid characters", param1: "507f1f77bcf86cd79943901G", want: false},
		{name: "Invalid mongo id, invalid format only letters", param1: "ZZZZZZZZZZZZZZZZZZZZZZZZ", want: false},
		{name: "Invalid mongo id, too many digits", param1: "1234567890123456789012345", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMongoID(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
