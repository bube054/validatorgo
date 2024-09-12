package validatorgo

import (
	"testing"
)

func TestIsBase58(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "Valid base 58 string",
			param1: "3mJr7AoUCHxNqd",
			want:   true,
		},
		{
			name:   "Invalid base 58 string without padding",
			param1: "3mJr7Ao0UCHxNqd",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBase58(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
