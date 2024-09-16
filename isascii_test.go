package validatorgo

import "testing"

func TestIsAscii(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{name: "Basic ASCII check", param1: "hello", want: true},
		{name: "Uppercase ASCII check", param1: "HELLO", want: true},
		{name: "Alphanumeric ASCII", param1: "Hello123", want: true},
		{name: "ASCII special characters", param1: "!@#$%^&*()", want: true},
		{name: "ASCII space", param1: "Hello World", want: true},
		{name: "Numeric ASCII", param1: "1234567890", want: true},
		{name: "Empty string", param1: "", want: true},
		{name: "Newline character", param1: "Hello\nWorld", want: true},
		{name: "Tab character", param1: "Hello\tWorld", want: true},
		{name: "Extended ASCII (non-ASCII)", param1: "Olá", want: false},
		{name: "Non-ASCII Unicode character", param1: "こんにちは", want: false},
		{name: "Emoji character", param1: "Hello🙂", want: false},
		{name: "Non-ASCII accented character", param1: "Café", want: false},
		{name: "Non-ASCII control character", param1: "Hello\u200BWorld", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAscii(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
