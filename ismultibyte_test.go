package validatorgo

import "testing"

func TestIsMultibyte(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid multi byte
		{name: "Valid multi byte Japanese characters", param1: "こんにちは", want: true},
		{name: "Valid multi byte chinese characters", param1: "你好", want: true},
		{name: "Valid multi byte Cyrillic characters", param1: "Привет", want: true},
		{name: "Valid multi byte Unicode characters above U+FFF", param1: "𠜎𠜱", want: true},
		{name: "Valid multi byte emoji", param1: "😊🍕", want: true},

		// Invalid multi byte
		{name: "Invalid multi byte ASCII characters only", param1: "hello", want: false},
		{name: "Invalid multi byte numbers only", param1: "12345", want: false},
		{name: "Invalid multi byte single-byte symbol", param1: "!", want: false},
		{name: "Invalid multi byte ASCII letters and numbers", param1: "test123", want: false},
		{name: "Invalid multi byte single-byte symbols", param1: "@#%&", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMultibyte(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
