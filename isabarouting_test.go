package validatorgo

import (
	"testing"
)

func TestIsAbaRouting(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "Input is less than 9 digits",
			param1: "123456",
			want:   false,
		},
		{
			name:   "Input is greater than 9 digits",
			param1: "1234567890",
			want:   false,
		},
		{
			name:   "Input is not numeric",
			param1: "abcdef6789",
			want:   false,
		},
		{
			name:   "Invalid aba routing but valid input length",
			param1: "123456789",
			want:   false,
		},
		{
			name:   "Valid aba routing and valid input length",
			param1: "021000021",
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAbaRouting(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
