package validatorgo

import "testing"

func TestIsStrongPassword(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsStrongPasswordOpts
		want   bool
		score  int
	}{
		// Valid passwords
		{name: "Test password", param1: "Password123!", param2: IsStrongPasswordOpts{}, want: true, score: 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, score := IsStrongPassword(test.param1, test.param2)

			if result != test.want || score != test.score {
				t.Errorf("got `%t` & `%d` but, wanted `%t` & `%d`", result, score, test.want, test.score)
			}
		})
	}
}
