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

		// JS valid
		// "507f1f77bcf86cd799439011" already covered above
		// JS invalid
		{name: "JS invalid too short", param1: "507f1f77bcf86cd7994390", want: false},
		{name: "JS invalid char z", param1: "507f1f77bcf86cd79943901z", want: false},
		{name: "JS invalid empty", param1: "", want: false},
		{name: "JS invalid trailing space", param1: "507f1f77bcf86cd799439011 ", want: false},

		// Extra edges
		{name: "Uppercase hex valid", param1: "507F1F77BCF86CD799439011", want: true},
		{name: "Too long", param1: "507f1f77bcf86cd7994390111", want: false},
		{name: "Non-hex chars", param1: "507f1f77bcf86cd79943901g", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsMongoID(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
