package validatorgo

import "testing"

func TestIsDivisibleBy(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 int
		want   bool
	}{
		{
			name:   "Valid 10/2",
			param1: "10",
			param2: 2,
			want:   true,
		},
		{
			name:   "Invalid 10/3",
			param1: "10",
			param2: 3,
			want:   false,
		},
		{
			name:   "Invalid divisible by 0",
			param1: "7",
			param2: 0,
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsDivisibleBy(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
