package validatorgo

import "testing"

func TestIsUpperCase(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Example
		{name: "Is uppercase", param1: "HELLO", want: true},
		// Invalid Example
		{name: "Empty string is not uppercase", param1: "", want: false},
		{name: "Is not uppercase", param1: "hello", want: false},
		{name: "Only few letter are uppercase", param1: "ExaMPle", want: false},

		// JS valid
		{name: "JS valid ABC", param1: "ABC", want: true},
		{name: "JS valid ABC123", param1: "ABC123", want: true},
		{name: "JS valid all caps sentence", param1: "ALL CAPS IS FUN.", want: true},
		{name: "JS invalid spaces and dot", param1: "   .", want: false},
		// JS invalid
		{name: "JS invalid fooBar", param1: "fooBar", want: false},
		{name: "JS invalid 123abc", param1: "123abc", want: false},

		// Extra edges
		{name: "Only digits", param1: "123", want: false},
		{name: "Only special chars", param1: "!@#", want: false},
		{name: "Single lowercase", param1: "ABc", want: false},
		{name: "Mixed with space", param1: "HELLO WORLD", want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsUpperCase(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
