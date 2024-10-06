package validatorgo

import "testing"

func TestIsStrongPassword(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsStrongPasswordOpts
		want   bool
		score  float64
	}{
		// Valid passwords
		{name: "Valid password with nil config", param1: "Password123!", param2: nil, want: true, score: 130.5},
		{name: "Valid password with defaults", param1: "Password123!", param2: &IsStrongPasswordOpts{}, want: true, score: 130.5},
		{name: "Valid complex password", param1: "7Lcr;}!60q,a!@M", param2: &IsStrongPasswordOpts{}, want: true, score: 163.50},
		{name: "Short password", param1: "Pas12!", param2: &IsStrongPasswordOpts{}, want: false, score: 66.0},
		{name: "No numbers", param1: "Password!@", param2: &IsStrongPasswordOpts{}, want: false, score: 108.5},
		{name: "Only numbers", param1: "123456789", param2: &IsStrongPasswordOpts{}, want: false, score: 99.0},
		{name: "Custom MinLength, no symbols", param1: "StrongP4ssw0rd", param2: &IsStrongPasswordOpts{MinLength: intPtr(12), MinSymbols: intPtr(0)}, want: true, score: 151.0},
		// {name: "Custom PointsPerUnique and PointsForContainingLower", param1: "LowerCaseOnly123", param2: &IsStrongPasswordOpts{PointsPerUnique: floatPtr(2), PointsForContainingLower: floatPtr(20)}, want: true, score: 288.50},
		{name: "Weak password, only lowercase", param1: "weakpassword", param2: &IsStrongPasswordOpts{}, want: false, score: 127.50},
		{name: "Valid password, high score custom points", param1: "CompL3x$P@ss", param2: &IsStrongPasswordOpts{PointsPerUnique: floatPtr(2), PointsForContainingUpper: floatPtr(15)}, want: true, score: 155.5},

		// Invalid passwords
		{name: "Too short", param1: "P@ss1", param2: &IsStrongPasswordOpts{}, want: false, score: 53.50},
		{name: "No uppercase letters", param1: "password123!", param2: &IsStrongPasswordOpts{}, want: false, score: 130.50},
		{name: "No lowercase letters", param1: "PASSWORD123!", param2: &IsStrongPasswordOpts{}, want: false, score: 130.50},
		{name: "No symbols", param1: "Password123", param2: &IsStrongPasswordOpts{}, want: false, score: 119.50},
		{name: "No numbers", param1: "Password!@", param2: &IsStrongPasswordOpts{}, want: false, score: 108.5},
		{name: "Only symbols", param1: "!@#$%^&*", param2: &IsStrongPasswordOpts{}, want: false, score: 88.0},
		{name: "Empty password", param1: "", param2: &IsStrongPasswordOpts{}, want: false, score: 0.0},
		{name: "Only repeated characters", param1: "aaaaaaa", param2: &IsStrongPasswordOpts{}, want: false, score: 70.5},
		{name: "Only digits", param1: "12345678", param2: &IsStrongPasswordOpts{}, want: false, score: 88.0},
		{name: "Invalid with custom opts, missing symbols", param1: "ComplexPass123", param2: &IsStrongPasswordOpts{MinSymbols: intPtr(1)}, want: false, score: 152.50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, score := IsStrongPassword(test.param1, test.param2)

			if result != test.want || score != test.score {
				t.Errorf("got `%t` & `%.2f` but, wanted `%t` & `%.2f`", result, score, test.want, test.score)
			}
		})
	}
}
