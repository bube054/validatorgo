package validatorgo

import "testing"

func TestIsHalfWidth(t *testing.T) {
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
			name:   "Half-width Katakana",
			param1: "ｶﾀｶﾅ",
			want:   true,
		},
		{
			name:   "Full-width characters",
			param1: "こんにちは",
			want:   false,
		},
		{
			name:   "Mixed ASCII and half-width characters",
			param1: "Hello, ｶﾀｶﾅ",
			want:   true,
		},
		{
			name:   "Half-width punctuation",
			param1: "｡｢｣､",
			want:   true,
		},
		{
			name:   "Full-width alphanumeric",
			param1: "ＡＢＣ１２３",
			want:   false,
		},
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
