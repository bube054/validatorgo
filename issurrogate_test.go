package validatorgo

import "testing"

func TestIsSurrogatePair(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid surrogates
		{name: "Valid - Single Surrogate Pair", param1: "𝌆", want: true},                   // Single surrogate pair character (U+1D306)
		{name: "Valid - Multiple Surrogate Pairs", param1: "𝌆𝌇", want: true},               // Multiple surrogate pair characters
		{name: "Valid - Surrogate Pair with Regular Chars", param1: "abc𝌆def", want: true}, // Surrogate pair with regular characters
		// Invalid surrogates
		{name: "Invalid - Regular Characters", param1: "abcdef", want: false},                 // No surrogate pairs, only regular characters
		{name: "Invalid - Empty String", param1: "", want: false},                             // Empty string
		{name: "Invalid - High Surrogate Alone", param1: `\uD834`, want: false},               // Unpaired high surrogate
		{name: "Invalid - Low Surrogate Alone", param1: `\uDF06`, want: false},                // Unpaired low surrogate
		{name: "Invalid - Mixed Regular and Surrogates", param1: `abc\uD834xyz`, want: false}, // High surrogate without matching low surrogate
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsSurrogatePair(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
