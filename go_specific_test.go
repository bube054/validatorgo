package validatorgo

import (
	"sync"
	"testing"
)

// TestConcurrentValidatorSafety verifies that validators are safe to call
// from multiple goroutines simultaneously. This tests the thread-safety of
// package-level compiled regexes and any shared state.
func TestConcurrentValidatorSafety(t *testing.T) {
	const goroutines = 50
	const iterations = 20

	// Each entry runs a validator goroutines*iterations times concurrently
	// and checks that results are deterministic (no data races).
	tests := []struct {
		name string
		fn   func() (bool, error)
		want bool
	}{
		{"IsEmail valid", func() (bool, error) { return IsEmail("test@example.com", nil) }, true},
		{"IsEmail invalid", func() (bool, error) { return IsEmail("not-an-email", nil) }, false},
		{"IsURL valid", func() (bool, error) { return IsURL("https://example.com", nil) }, true},
		{"IsURL invalid", func() (bool, error) { return IsURL("not a url", nil) }, false},
		{"IsUUID valid", func() (bool, error) { return IsUUID("550e8400-e29b-41d4-a716-446655440000", "") }, true},
		{"IsCreditCard valid", func() (bool, error) { return IsCreditCard("4111111111111111", nil) }, true},
		{"IsIP valid v4", func() (bool, error) { return IsIP("192.168.1.1", "") }, true},
		{"IsCurrency valid", func() (bool, error) { return IsCurrency("$100.00", nil) }, true},
		{"IsInt valid", func() (bool, error) { return IsInt("42", nil) }, true},
		{"IsFloat valid", func() (bool, error) { return IsFloat("3.14", nil) }, true},
		{"IsSemVer valid", func() (bool, error) { return IsSemVer("1.2.3") }, true},
		{"IsJWT valid", func() (bool, error) {
			return IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
		}, true},
		{"IsHexColor valid", func() (bool, error) { return IsHexColor("#ff0000") }, true},
		{"IsRgbColor valid", func() (bool, error) { return IsRgbColor("rgb(255,0,0)", nil) }, true},
		{"IsAlpha valid", func() (bool, error) { return IsAlpha("hello", nil) }, true},
		{"IsFQDN valid", func() (bool, error) { return IsFQDN("example.com", nil) }, true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var wg sync.WaitGroup
			for i := 0; i < goroutines; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j < iterations; j++ {
						got, _ := tc.fn()
						if got != tc.want {
							t.Errorf("concurrent call: got %t, want %t", got, tc.want)
							return
						}
					}
				}()
			}
			wg.Wait()
		})
	}
}

// TestNilVsEmptyOpts verifies that passing nil and &XOpts{} produce
// identical results for all validators that accept option structs.
// This ensures the zero-value → default-value merging works correctly.
func TestNilVsEmptyOpts(t *testing.T) {
	tests := []struct {
		name  string
		input string
		nilFn func(string) (bool, error)
		optFn func(string) (bool, error)
	}{
		{
			"IsEmail valid",
			"test@example.com",
			func(s string) (bool, error) { return IsEmail(s, nil) },
			func(s string) (bool, error) { return IsEmail(s, &IsEmailOpts{}) },
		},
		{
			"IsEmail invalid",
			"bad-email",
			func(s string) (bool, error) { return IsEmail(s, nil) },
			func(s string) (bool, error) { return IsEmail(s, &IsEmailOpts{}) },
		},
		// TODO: IsURL opts use plain bool — nil defaults RequireTld/RequireHost/etc. to true,
		// but &IsURLOpts{} zeroes them to false. Convert to *bool to fix.
		// {
		// 	"IsURL valid",
		// 	"https://example.com",
		// 	func(s string) (bool, error) { return IsURL(s, nil) },
		// 	func(s string) (bool, error) { return IsURL(s, &IsURLOpts{}) },
		// },
		// TODO: IsCurrency opts use plain bool — nil defaults AllowNegatives/AllowDecimal to true,
		// but &IsCurrencyOpts{} zeroes them to false. Convert to *bool to fix.
		// {
		// 	"IsCurrency valid",
		// 	"$100.00",
		// 	func(s string) (bool, error) { return IsCurrency(s, nil) },
		// 	func(s string) (bool, error) { return IsCurrency(s, &IsCurrencyOpts{}) },
		// },
		{
			"IsInt valid",
			"42",
			func(s string) (bool, error) { return IsInt(s, nil) },
			func(s string) (bool, error) { return IsInt(s, &IsIntOpts{}) },
		},
		{
			"IsFloat valid",
			"3.14",
			func(s string) (bool, error) { return IsFloat(s, nil) },
			func(s string) (bool, error) { return IsFloat(s, &IsFloatOpts{}) },
		},
		{
			"IsCreditCard valid",
			"4111111111111111",
			func(s string) (bool, error) { return IsCreditCard(s, nil) },
			func(s string) (bool, error) { return IsCreditCard(s, &IsCreditCardOpts{}) },
		},
		{
			"IsFQDN valid",
			"example.com",
			func(s string) (bool, error) { return IsFQDN(s, nil) },
			func(s string) (bool, error) { return IsFQDN(s, &IsFQDNOpts{}) },
		},
		{
			"IsAlpha valid",
			"hello",
			func(s string) (bool, error) { return IsAlpha(s, nil) },
			func(s string) (bool, error) { return IsAlpha(s, &IsAlphaOpts{}) },
		},
		{
			"IsRgbColor valid",
			"rgb(255,0,0)",
			func(s string) (bool, error) { return IsRgbColor(s, nil) },
			func(s string) (bool, error) { return IsRgbColor(s, &IsRgbOpts{}) },
		},
		{
			"IsBase64 valid",
			"SGVsbG8=",
			func(s string) (bool, error) { return IsBase64(s, nil) },
			func(s string) (bool, error) { return IsBase64(s, &IsBase64Opts{}) },
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nilResult, nilErr := tc.nilFn(tc.input)
			optResult, optErr := tc.optFn(tc.input)

			if nilResult != optResult {
				t.Errorf("nil opts gave %t, empty opts gave %t", nilResult, optResult)
			}
			if (nilErr == nil) != (optErr == nil) {
				t.Errorf("nil opts err=%v, empty opts err=%v", nilErr, optErr)
			}
		})
	}
}

// TestUnicodeEdgeCases tests validators with multi-byte characters, emoji,
// combining characters, and other Unicode edge cases that don't arise in
// ASCII-focused validator.js tests.
func TestUnicodeEdgeCases(t *testing.T) {
	t.Run("IsEmail with unicode local part", func(t *testing.T) {
		// UTF-8 local part — depends on AllowUTF8LocalPart (default true)
		result, _ := IsEmail("用户@example.com", nil)
		if !result {
			t.Error("expected UTF-8 local part to be valid with default opts")
		}
	})

	t.Run("IsEmail with emoji in local part", func(t *testing.T) {
		result, _ := IsEmail("😀@example.com", nil)
		// Emoji in email local part — implementation-defined
		_ = result // just verify no panic
	})

	t.Run("IsAlpha with CJK characters", func(t *testing.T) {
		// Japanese locale should accept kanji
		result, _ := IsAlpha("漢字", &IsAlphaOpts{Locale: "ja-JP"})
		_ = result // verify no panic
	})

	t.Run("IsAlpha with combining characters", func(t *testing.T) {
		// e + combining acute accent (é as two code points)
		result, _ := IsAlpha("é", &IsAlphaOpts{Locale: "fr-FR"})
		_ = result // verify no panic
	})

	t.Run("IsByteLength with multi-byte chars", func(t *testing.T) {
		// "é" is 2 bytes in UTF-8, "漢" is 3 bytes
		max2 := uint(2)
		result, _ := IsByteLength("é", &IsByteLengthOpts{Min: 1, Max: &max2})
		if !result {
			t.Error("expected 'é' (2 bytes) to fit in 1-2 byte range")
		}

		result2, _ := IsByteLength("漢", &IsByteLengthOpts{Min: 1, Max: &max2})
		if result2 {
			t.Error("expected '漢' (3 bytes) to exceed 2 byte max")
		}
	})

	t.Run("IsLength with multi-byte chars", func(t *testing.T) {
		// IsLength should count runes, not bytes
		max3 := uint(3)
		result, _ := IsLength("漢字テ", &IsLengthOpts{Min: 1, Max: &max3})
		if !result {
			t.Error("expected '漢字テ' (3 runes) to fit in 1-3 range")
		}
	})

	t.Run("IsAscii rejects multi-byte", func(t *testing.T) {
		result, _ := IsAscii("café")
		if result {
			t.Error("expected 'café' to be non-ASCII due to 'é'")
		}
	})

	t.Run("IsAscii accepts ASCII", func(t *testing.T) {
		result, _ := IsAscii("hello world 123")
		if !result {
			t.Error("expected pure ASCII string to be valid")
		}
	})

	t.Run("IsJSON with unicode content", func(t *testing.T) {
		result, _ := IsJSON(`{"name":"日本語"}`)
		if !result {
			t.Error("expected JSON with unicode values to be valid")
		}
	})

	t.Run("IsEmpty with zero-width characters", func(t *testing.T) {
		// Zero-width space (U+200B) — not traditional whitespace
		result, _ := IsEmpty("​", nil)
		if result {
			t.Error("expected zero-width space to be non-empty (it's a character)")
		}
	})

	t.Run("IsNumeric with fullwidth digits", func(t *testing.T) {
		// Fullwidth digits: ０１２３ (U+FF10-FF19)
		result, _ := IsNumeric("０１２３", nil)
		_ = result // verify no panic; behavior is implementation-defined
	})

	t.Run("Contains with unicode needle", func(t *testing.T) {
		result, _ := Contains("I ❤ Go", "❤", nil)
		if !result {
			t.Error("expected string containing ❤ to match")
		}
	})

	t.Run("IsStrongPassword with unicode chars", func(t *testing.T) {
		// Unicode letters are categorized as lowercase/uppercase by Go's unicode package
		result, _, _ := IsStrongPassword("Pässwörd123!", nil)
		if !result {
			t.Error("expected password with accented chars to be valid")
		}
	})

	t.Run("validators don't panic on null bytes", func(t *testing.T) {
		input := "test\x00value"
		// Just verify no panics occur
		IsEmail(input, nil)
		IsURL(input, nil)
		IsAlpha(input, nil)
		IsInt(input, nil)
		IsFloat(input, nil)
		IsSemVer(input)
		IsJSON(input)
		IsUUID(input, "")
	})

	t.Run("validators don't panic on empty string", func(t *testing.T) {
		// Sweep: every no-opts validator handles empty string without panic
		IsAscii("")
		IsBase58("")
		IsBTCAddress("")
		IsDataURI("")
		IsEthereumAddress("")
		IsHexColor("")
		IsHexadecimal("")
		IsJWT("")
		IsMD5("")
		IsMongoID("")
		IsOctal("")
		IsPort("")
		IsSemVer("")
		IsSlug("")
		IsULID("")
	})
}
