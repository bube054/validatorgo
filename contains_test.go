package validatorgo

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		param3 ContainsOpt
		want   bool
	}{
		{
			name:   "1",
			param1: "Hello, world! Hello, universe!",
			param2: "hello",
			param3: ContainsOpt{IgnoreCase: true},
			want:   true,
		},
		{
			name:   "2",
			param1: "Hello, world! Hello, universe!",
			param2: "Hello",
			param3: ContainsOpt{MinOccurrences: 2},
			want:   true,
		},
		{
			name:   "3",
			param1: "Hello, world! Hello, universe!",
			param2: "Hello",
			param3: ContainsOpt{MinOccurrences: 3},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Contains(test.param1, test.param2, test.param3)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
