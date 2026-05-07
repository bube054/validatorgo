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
			name:   "Invalid abc/3",
			param1: "abc",
			param2: 3,
			want:   false,
		},
		{
			name:   "Invalid divisible by 0",
			param1: "7",
			param2: 0,
			want:   false,
		},
		// Valid divisible by 2
		{name: "2 div by 2", param1: "2", param2: 2, want: true},
		{name: "4 div by 2", param1: "4", param2: 2, want: true},
		{name: "100 div by 2", param1: "100", param2: 2, want: true},
		{name: "1000 div by 2", param1: "1000", param2: 2, want: true},
		// Invalid divisible by 2
		{name: "1 not div by 2", param1: "1", param2: 2, want: false},
		{name: "2.5 not div by 2", param1: "2.5", param2: 2, want: false},
		{name: "101 not div by 2", param1: "101", param2: 2, want: false},
		{name: "foo not div by 2", param1: "foo", param2: 2, want: false},
		{name: "empty not div by 2", param1: "", param2: 2, want: false},
		// Extra edge cases
		{name: "Divisible by 5", param1: "25", param2: 5, want: true},
		{name: "Not divisible by 3", param1: "10", param2: 3, want: false},
		{name: "Zero divisible by any", param1: "0", param2: 7, want: true},
		{name: "Negative divisible", param1: "-10", param2: 5, want: true},
		{name: "Large number", param1: "1000000", param2: 1000, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsDivisibleBy(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
