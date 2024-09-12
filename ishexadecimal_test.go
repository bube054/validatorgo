package validatorgo

import "testing"

func TestIsHexadecimal(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Hexadecimal
		{name: "Lowercase Hexadecimal", param1: "abcdef", want: true},
		{name: "Uppercase Hexadecimal", param1: "ABCDEF", want: true},
		{name: "Mixed Case Hexadecimal", param1: "aBcDeF", want: true},
		{name: "Hexadecimal with Digits", param1: "1234567890abcdef", want: true},

		// Invalid Hexadecimal
		{name: "Non-Hexadecimal Characters", param1: "abcdez", want: false},
		{name: "Empty String", param1: "", want: false},
		{name: "Contains Spaces", param1: "a b c d e f", want: false},
		{name: "Special Characters", param1: "abcd@123", want: false},
		{name: "Alphabetic Characters Beyond 'f'", param1: "abcdefg", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsHexadecimal(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
