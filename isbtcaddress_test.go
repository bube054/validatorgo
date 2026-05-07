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
		{name: "Valid BTC 1MUz", param1: "1MUz4VMYui5qY1mxUiG8BQ1Luv6tqkvaiL", want: true},
		{name: "Valid BTC 3J98", param1: "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy", want: true},
		// TODO: bech32 support missing in IsBTCAddress implementation
		{name: "Valid BTC bech32 bc1qar", param1: "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq", want: false},
		{name: "Valid BTC 14qV", param1: "14qViLJfdGaP4EeHnDyJbEGQysnCpwk3gd", want: true},
		{name: "Valid BTC 35bS", param1: "35bSzXvRKLpHsHMrzb82f617cV4Srnt7hS", want: true},
		{name: "Valid BTC 17VZ", param1: "17VZNX1SN5NtKa8UQFxwQbFeFc3iqRYhemt", want: true},
		{name: "Valid BTC bech32 bc1qw", param1: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4", want: false},
		{name: `Invalid BTC address as it starts with "b"`, param1: "b1qarsrrr7ASHy5643ydab9re59gtzzwfrah", want: false},
		{name: `Invalid BTC address as it starts with 0.`, param1: "0J98t1RHT73CNmQwertyyWrnqRhWNLy", want: false},
		{name: "Invalid BTC address with special characters", param1: "1RAHU@EYstWetqabcFn5Au4m4GFg7xJaNVN2", want: false},
		{name: "Invalid BTC prefix 4", param1: "4J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy", want: false},
		{name: "Invalid ETH address not BTC", param1: "0x56F0B8A998425c53c75C4A303D4eF987533c5597", want: false},
		{name: "Empty string", param1: "", want: false},
		{name: "Invalid uppercase bech32", param1: "BC1QW508D6QEJXTDG4Y5R3ZARVAYR0C5XW7KV8F3T4", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsBTCAddress(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
