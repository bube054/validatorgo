package validatorgo

import "testing"

func TestIsFullWidth(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "Empty string",
			param1: "",
			want:   false,
		},
		{
			name:   "ASCII characters only",
			param1: "Hello, World!",
			want:   false,
		},
		{
			name:   "Full-width characters only",
			param1: "こんにちは",
			want:   true,
		},
		{
			name:   "Mixed ASCII and full-width characters",
			param1: "Hello, こんにちは",
			want:   true,
		},
		{
			name:   "Full-width punctuation",
			param1: "！＄％＆",
			want:   true,
		},
		{
			name:   "Half-width Katakana",
			param1: "ｶﾀｶﾅ",
			want:   false,
		},
		{
			name:   "Full-width alphanumeric",
			param1: "ＡＢＣ１２３",
			want:   true,
		},
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
