package validatorgo

import "testing"

func TestIsFullWidth(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{name: "Only full-width Japanese characters", param1: "テスト", want: true}, // All full-width Katakana
		{name: "Only full-width CJK characters", param1: "漢字", want: true},       // Full-width Chinese characters
		{name: "Full-width and half-width mix", param1: "テストabc", want: true},    // Full-width Katakana with half-width English
		{name: "Full-width with numbers", param1: "１２３", want: true},             // Full-width numbers
		{name: "Half-width characters only", param1: "abc123", want: false},      // All half-width
		{name: "Mixed full-width and half-width", param1: "ＡＢＣｄｅｆ", want: true},  // Full-width Latin with half-width Latin
		{name: "Empty string", param1: "", want: false},                          // Empty string
		{name: "Only half-width punctuation", param1: ".,!?", want: false},       // All half-width punctuation
		{name: "Full-width punctuation", param1: "！＃＄％", want: true},             // Full-width punctuation
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsFullWidth(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
