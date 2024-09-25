package validatorgo

import "testing"

func TestIsHalfWidth(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{name: "Only half-width English characters", param1: "abc123", want: true},  // All half-width
		{name: "Half-width punctuation", param1: ".,!?", want: true},                // Half-width punctuation
		{name: "Mixed full-width and half-width", param1: "テストabc", want: true},     // Full-width Katakana with half-width English
		{name: "Only full-width Japanese characters", param1: "テスト", want: false},   // All full-width
		{name: "Only full-width numbers", param1: "１２３", want: false},               // All full-width numbers
		{name: "Empty string", param1: "", want: false},                             // Empty string
		{name: "Mixed characters with no half-width", param1: "漢字テスト", want: false}, // All full-width CJK characters
		{name: "Only half-width special characters", param1: "@#$%^&*", want: true}, // Half-width special characters
		{name: "Only full-width punctuation", param1: "！＃＄％", want: false},          // All full-width punctuation
		{name: "Half-width symbols", param1: "©®", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsHalfWidth(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
