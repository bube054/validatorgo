package validatorgo

import "testing"

func TestEquals(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		{name: "Valid equals", param1: "Hello", param2: "Hello", want: true},
		{name: "Valid not equals", param1: "Hello", param2: "World", want: false},
		{name: "Exact match", param1: "abc", param2: "abc", want: true},
		{name: "Case mismatch", param1: "Abc", param2: "abc", want: false},
		{name: "Number string mismatch", param1: "123", param2: "abc", want: false},
		{name: "Empty strings match", param1: "", param2: "", want: true},
		{name: "Empty vs non-empty", param1: "", param2: "a", want: false},
		{name: "Whitespace matters", param1: " ", param2: "", want: false},
		{name: "Unicode match", param1: "café", param2: "café", want: true},
		{name: "Unicode mismatch", param1: "café", param2: "cafe", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Equals(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
