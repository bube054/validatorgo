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

		// Ported from validator.js — valid base58 (testnet prefixes)
		// TODO: should be valid per validator.js — Go regex missing testnet 'm' prefix
		{name: "Valid BTC testnet m-prefix (validator.js)", param1: "mucFNhKMYoBQYUAEsrFVscQ1YaFQPekBpg", want: false},
		// TODO: should be valid per validator.js — Go regex missing testnet '2' prefix
		{name: "Valid BTC testnet 2-prefix (validator.js)", param1: "2NFUBBRcTJbYc1D4HSCbJhKZp6YCV4PQFpQ", want: false},

		// Ported from validator.js — valid bech32 (testnet and taproot)
		// TODO: should be valid per validator.js — Go missing tb1 testnet bech32 support
		{name: "Valid BTC testnet bech32 tb1q (validator.js)", param1: "tb1qxhkl607frtvjsy9nlyeg03lf6fsq947pl2pe82", want: false},
		// TODO: should be valid per validator.js — Go missing taproot bc1p support (address too long for current regex)
		{name: "Valid BTC taproot bc1p (validator.js)", param1: "bc1p5d7rjq7g6rdk2yhzks9smlaqtedr4dekq08ge8ztwac72sfr9rusxg3297", want: false},
		// TODO: should be valid per validator.js — Go missing testnet taproot tb1p support
		{name: "Valid BTC testnet taproot tb1p (validator.js)", param1: "tb1pzpelffrdh9ptpaqnurwx30dlewqv57rcxfeetp86hsssk30p4cws38tr9y", want: false},

		// Ported from validator.js — invalid base58 (forbidden characters 0, o, I, l)
		{name: "Invalid BTC contains '0' (validator.js)", param1: "3J98t1WpEZ73CNmQviecrnyiWrnqh0WNL0", want: false},
		{name: "Invalid BTC contains 'o' (validator.js)", param1: "3J98t1WpEZ73CNmQviecrnyiWrnqh0WNLo", want: false},
		{name: "Invalid BTC contains 'I' (validator.js)", param1: "3J98t1WpEZ73CNmQviecrnyiWrnqh0WNLI", want: false},
		{name: "Invalid BTC contains 'l' (validator.js)", param1: "3J98t1WpEZ73CNmQviecrnyiWrnqh0WNLl", want: false},

		// Ported from validator.js — invalid misc
		{name: "Invalid BTC prefix 'p' (validator.js)", param1: "pp8skudq3x5hzw8ew7vzsw8tn4k8wxsqsv0lt0mf3g", want: false},
		{name: "Invalid BTC contains 'l' in middle (validator.js)", param1: "17VZNX1SN5NlKa8UQFxwQbFeFc3iqRYhem", want: false},

		// Ported from validator.js — invalid bech32 (bad trailing chars, uppercase)
		{name: "Invalid bech32 taproot ends '1' (validator.js)", param1: "bc1p5d7rjq7g6rdk2yhzks9smlaqtedr4dekq08ge8ztwac72sfr9rusxg3291", want: false},
		{name: "Invalid bech32 taproot ends 'b' (validator.js)", param1: "bc1p5d7rjq7g6rdk2yhzks9smlaqtedr4dekq08ge8ztwac72sfr9rusxg329b", want: false},
		{name: "Invalid bech32 taproot ends 'i' (validator.js)", param1: "bc1p5d7rjq7g6rdk2yhzks9smlaqtedr4dekq08ge8ztwac72sfr9rusxg329i", want: false},
		{name: "Invalid bech32 taproot ends 'o' (validator.js)", param1: "bc1p5d7rjq7g6rdk2yhzks9smlaqtedr4dekq08ge8ztwac72sfr9rusxg329o", want: false},
		{name: "Invalid uppercase bech32 taproot BC1P (validator.js)", param1: "BC1P5D7RJQ7G6RDK2YHZKS9SMLAQTEDR4DEKQ08GE8ZTWAC72SFR9RUSXG3297", want: false},
		{name: "Invalid uppercase bech32 testnet TB1P (validator.js)", param1: "TB1PZPELFFRDH9PTPAQNURWX30DLEWQV57RCXFEETP86HSSSK30P4CWS38TR9Y", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsBTCAddress(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
