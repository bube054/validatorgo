package validatorgo

import "testing"

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Hex Colors
		{name: "Valid 3-digit Hex Color", param1: "#abc", want: true},
		{name: "Valid 6-digit Hex Color", param1: "#abcdef", want: true},
		{name: "Valid 6-digit Hex Color Uppercase", param1: "#ABCDEF", want: true},
		{name: "Valid 3-digit Hex Color Mixed Case", param1: "#AbC", want: true},

		// Invalid Hex Colors
		{name: "No Hash Symbol", param1: "abcdef", want: false},
		{name: "Invalid Length", param1: "#abcd", want: false},
		{name: "Invalid Characters", param1: "#xyz", want: false},
		{name: "Empty String", param1: "", want: false},
		{name: "Contains Spaces", param1: "#a b c", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsHexColor(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
