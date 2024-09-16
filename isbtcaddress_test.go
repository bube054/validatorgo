package validatorgo

import (
	"testing"
)

func TestIsBTCAddress(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
	{name: "Valid BTC", param1: "1RAHUEYstWetqabcFn5Au4m4GFg7xJaNVN2", want: true},
	{name: "Valid BTC", param1: "3J98t1RHT73CNmQwertyyWrnqRhWNLy", want: true},
	{name: "Valid BTC", param1: "bc1qarsrrr7ASHy5643ydab9re59gtzzwfrah", want: true},
	{name: `Invalid BTC address as it starts with "b"`, param1: "b1qarsrrr7ASHy5643ydab9re59gtzzwfrah", want: false},
	{name: `Invalid BTC address as it starts with 0.`, param1: "0J98t1RHT73CNmQwertyyWrnqRhWNLy", want: false},
	{name: "Invalid BTC address with special characters", param1: "1RAHU@EYstWetqabcFn5Au4m4GFg7xJaNVN2", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBTCAddress(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
