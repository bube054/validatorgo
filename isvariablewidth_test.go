package validatorgo

import "testing"

func TestIsVariableWidth(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid cases with both full-width and half-width characters
		{name: "Mixed full-width and half-width characters", param1: "ａｂｃ123", want: true},       // Full-width letters with half-width numbers
		{name: "Mixed full-width Japanese and half-width English", param1: "テストabc", want: true}, // Full-width Japanese with half-width English
		{name: "Full-width Chinese and half-width punctuation", param1: "你好！!", want: true},      // Full-width Chinese with half-width punctuation
		{name: "Mixed full-width numbers and half-width letters", param1: "１２３abc", want: true},  // Full-width numbers with half-width letters

		// Cases without mixed width
		{name: "All full-width characters", param1: "ａｂｃ１２３", want: false}, // Only full-width characters
		{name: "All half-width characters", param1: "abc123", want: false}, // Only half-width characters
		{name: "Full-width punctuation only", param1: "！＠＃", want: false},  // Only full-width punctuation
		{name: "Half-width punctuation only", param1: "!@#", want: false},  // Only half-width punctuation

		// Edge cases
		{name: "Empty string", param1: "", want: false},                 // Empty string should return false
		{name: "Single full-width character", param1: "ａ", want: false}, // Single full-width character
		{name: "Single half-width character", param1: "a", want: false}, // Single half-width character

		// Mixed cases with special characters
		{name: "Full-width and half-width special characters", param1: "＠a！", want: true}, // Full-width punctuation with half-width letter
		{name: "Mixed width including spaces", param1: "１２３ abc", want: true},             // Full-width numbers with half-width letters and space
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsVariableWidth(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
