package validatorgo

import "testing"

func TestIsStrongPassword(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsStrongPasswordOpts
		want   bool
	}{
		// Valid passwords
		{name: "Valid password with nil config", param1: "Password123!", param2: nil, want: true},
		{name: "Valid password with defaults", param1: "Password123!", param2: &IsStrongPasswordOpts{}, want: true},
		{name: "Valid complex password", param1: "7Lcr;}!60q,a!@M", param2: &IsStrongPasswordOpts{}, want: true},
		{name: "Custom MinLength, no symbols", param1: "StrongP4ssw0rd", param2: &IsStrongPasswordOpts{MinLength: intPtr(12), MinSymbols: intPtr(0)}, want: true},
		{name: "Valid password, high score custom points", param1: "CompL3x$P@ss", param2: &IsStrongPasswordOpts{PointsPerUnique: floatPtr(2), PointsForContainingUpper: floatPtr(15)}, want: true},

		// Invalid passwords
		{name: "Short password", param1: "Pas12!", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Too short", param1: "P@ss1", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "No uppercase letters", param1: "password123!", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "No lowercase letters", param1: "PASSWORD123!", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "No symbols", param1: "Password123", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "No numbers", param1: "Password!@", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Only numbers", param1: "123456789", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Only symbols", param1: "!@#$%^&*", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Only lowercase", param1: "weakpassword", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Only repeated characters", param1: "aaaaaaa", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Only digits", param1: "12345678", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Empty password", param1: "", param2: &IsStrongPasswordOpts{}, want: false},
		{name: "Invalid with custom opts, missing symbols", param1: "ComplexPass123", param2: &IsStrongPasswordOpts{MinSymbols: intPtr(1)}, want: false},

		// validator.js ported: valid passwords
		{name: "JS valid: special chars and mixed case", param1: `%2%k{7BsL"M%Kd6e`, param2: nil, want: true},
		{name: "JS valid: long password with underscore", param1: "EXAMPLE of very long_password123!", param2: nil, want: true},
		{name: "JS valid: mixed symbols", param1: "mxH_+2vs&54_+H3P", param2: nil, want: true},
		{name: "JS valid: symbols and digits", param1: "+&DxJ=X7-4L8jRCD", param2: nil, want: true},
		{name: "JS valid: percent and ampersand", param1: "etV*p%Nr6w&H%FeF", param2: nil, want: true},
		{name: "JS valid: pound sign", param1: "£3.ndSau_7", param2: nil, want: true},
		{name: "JS valid: backslash as symbol", param1: `VaLIDWith\Symb0l`, param2: nil, want: true},

		// validator.js ported: invalid passwords
		{name: "JS invalid: all lowercase", param1: "password", param2: nil, want: false},
		{name: "JS invalid: short lowercase+digit", param1: "hunter2", param2: nil, want: false},
		{name: "JS invalid: lowercase with spaces", param1: "hello world", param2: nil, want: false},
		{name: "JS invalid: lowercase+digits no symbol", param1: "passw0rd", param2: nil, want: false},
		{name: "JS invalid: lowercase+symbol no upper/digit", param1: "password!", param2: nil, want: false},
		{name: "JS invalid: uppercase+symbol no lower/digit", param1: "PASSWORD!", param2: nil, want: false},

		// Custom MinLowercase
		{name: "MinLowercase 0, all uppercase+digits+symbols", param1: "STRONG1!", param2: &IsStrongPasswordOpts{MinLowercase: intPtr(0)}, want: true},
		{name: "MinLowercase 3, not enough lowercase", param1: "ABcd12!@", param2: &IsStrongPasswordOpts{MinLowercase: intPtr(3)}, want: false},

		// Custom MinUppercase
		{name: "MinUppercase 0, all lowercase+digits+symbols", param1: "strong1!", param2: &IsStrongPasswordOpts{MinUppercase: intPtr(0)}, want: true},
		{name: "MinUppercase 3, not enough uppercase", param1: "ABcdef1!", param2: &IsStrongPasswordOpts{MinUppercase: intPtr(3)}, want: false},

		// Custom MinNumbers
		{name: "MinNumbers 0, no digits needed", param1: "Password!", param2: &IsStrongPasswordOpts{MinNumbers: intPtr(0)}, want: true},
		{name: "MinNumbers 3, not enough digits", param1: "Password12!", param2: &IsStrongPasswordOpts{MinNumbers: intPtr(3)}, want: false},

		// Custom MinSymbols
		{name: "MinSymbols 0, no symbols needed", param1: "Password123", param2: &IsStrongPasswordOpts{MinSymbols: intPtr(0)}, want: true},
		{name: "MinSymbols 3, not enough symbols", param1: "Password1!@", param2: &IsStrongPasswordOpts{MinSymbols: intPtr(3)}, want: false},

		// Minimum length edge cases
		{name: "MinLength 4, short valid", param1: "Ab1!", param2: &IsStrongPasswordOpts{MinLength: intPtr(4)}, want: true},
		{name: "MinLength 20, too short", param1: "Password123!", param2: &IsStrongPasswordOpts{MinLength: intPtr(20)}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _, err := IsStrongPassword(test.param1, test.param2)
			assertValidation(t, result, test.want, err)
		})
	}
}

func TestIsStrongPassword_ScoreOrdering(t *testing.T) {
	// Verify intuitive scoring properties without hardcoding exact values.
	// These tests survive scoring formula tweaks as long as relative ordering holds.

	score := func(pw string, opts *IsStrongPasswordOpts) float64 {
		_, s, _ := IsStrongPassword(pw, opts)
		return s
	}

	t.Run("empty password scores zero", func(t *testing.T) {
		s := score("", nil)
		if s != 0 {
			t.Errorf("empty password: got score %.2f, want 0", s)
		}
	})

	t.Run("non-empty password scores above zero", func(t *testing.T) {
		s := score("a", nil)
		if s <= 0 {
			t.Errorf("single char password: got score %.2f, want > 0", s)
		}
	})

	t.Run("adding characters from new classes increases score", func(t *testing.T) {
		// Note: Go uses count-based scoring (points per char), not presence-based.
		// Same-length passwords with different class distributions score identically.
		// Adding MORE characters always increases the score.
		base := score("abcdefgh", nil)         // 8 lower
		withUpper := score("abcdefghAB", nil)   // 8 lower + 2 upper
		withNum := score("abcdefghAB12", nil)   // 8 lower + 2 upper + 2 num
		withSym := score("abcdefghAB12!@", nil) // 8 lower + 2 upper + 2 num + 2 sym

		if base >= withUpper {
			t.Errorf("base (%.2f) should < +upper (%.2f)", base, withUpper)
		}
		if withUpper >= withNum {
			t.Errorf("+upper (%.2f) should < +num (%.2f)", withUpper, withNum)
		}
		if withNum >= withSym {
			t.Errorf("+num (%.2f) should < +sym (%.2f)", withNum, withSym)
		}
	})

	t.Run("unique chars score higher than repeated", func(t *testing.T) {
		repeated := score("aaaaaaa", nil)
		unique := score("abcdefg", nil)
		if repeated >= unique {
			t.Errorf("repeated (%.2f) should < unique (%.2f)", repeated, unique)
		}
	})

	t.Run("longer password scores higher than shorter with same pattern", func(t *testing.T) {
		short := score("aB1!", nil)
		long := score("aaBB11!!", nil)
		if short >= long {
			t.Errorf("short (%.2f) should < long (%.2f)", short, long)
		}
	})

	t.Run("complex password scores higher than simple", func(t *testing.T) {
		simple := score("password", nil)
		medium := score("Password1", nil)
		complex := score("Password123!", nil)
		veryComplex := score("7Lcr;}!60q,a!@M", nil)

		if simple >= medium {
			t.Errorf("simple (%.2f) should < medium (%.2f)", simple, medium)
		}
		if medium >= complex {
			t.Errorf("medium (%.2f) should < complex (%.2f)", medium, complex)
		}
		if complex >= veryComplex {
			t.Errorf("complex (%.2f) should < very complex (%.2f)", complex, veryComplex)
		}
	})

	t.Run("custom points affect scoring", func(t *testing.T) {
		defaultScore := score("Password123!", nil)
		highUpper := score("Password123!", &IsStrongPasswordOpts{
			PointsForContainingUpper: floatPtr(50),
		})
		if defaultScore >= highUpper {
			t.Errorf("default (%.2f) should < high upper points (%.2f)", defaultScore, highUpper)
		}
	})

	t.Run("custom PointsPerUnique affects scoring", func(t *testing.T) {
		defaultScore := score("Password123!", nil)
		highUnique := score("Password123!", &IsStrongPasswordOpts{
			PointsPerUnique: floatPtr(5),
		})
		if defaultScore >= highUnique {
			t.Errorf("default (%.2f) should < high unique points (%.2f)", defaultScore, highUnique)
		}
	})
}
