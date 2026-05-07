package validatorgo

import "testing"

func TestIsLowerCase(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Example
		{name: "Is lowercase", param1: "hello", want: true},
		// Invalid Example
		{name: "Empty string is not lowercase", param1: "", want: false},
		{name: "Is not lowercase", param1: "WORLD", want: false},
		{name: "Only few letter are lowercase", param1: "ExaMPle", want: false},

		// JS valid
		{name: "JS valid abc", param1: "abc", want: true},
		{name: "JS valid abc123", param1: "abc123", want: true},
		{name: "JS valid sentence", param1: "this is lowercase.", want: true},
		{name: "JS valid unicode", param1: "tr竪s 端ber", want: true},
		// JS invalid
		{name: "JS invalid fooBar", param1: "fooBar", want: false},
		{name: "JS invalid 123A", param1: "123A", want: false},

		// Extra edges
		{name: "Only digits", param1: "123", want: false},
		{name: "Only spaces", param1: "   ", want: false},
		{name: "With special chars lowercase", param1: "hello!", want: true},
		{name: "Single uppercase", param1: "aBc", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsLowerCase(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
