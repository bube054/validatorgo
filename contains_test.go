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
		// Valid default config
		{name: "Basic match", param1: "hello world", param2: "world", param3: ContainsOpt{}, want: true},
		{name: "Basic match digits", param1: "abc123", param2: "123", param3: ContainsOpt{}, want: true},

		// Valid with ignoreCase true
		{name: "Case-insensitive match", param1: "Hello World", param2: "hello", param3: ContainsOpt{IgnoreCase: true}, want: true},
		{name: "Case-insensitive match mixed", param1: "FOOBAR", param2: "bar", param3: ContainsOpt{IgnoreCase: true}, want: true},

		// Valid with minimum occurrences
		{name: "Minimum occurrences met", param1: "hello hello world", param2: "hello", param3: ContainsOpt{MinOccurrences: 2}, want: true},
		{name: "Minimum occurrences default", param1: "abc123", param2: "123", param3: ContainsOpt{MinOccurrences: 1}, want: true},

		// Invalid default config
		{name: "No match", param1: "hello world", param2: "earth", param3: ContainsOpt{}, want: false},
		{name: "No match digits", param1: "abc123", param2: "xyz", param3: ContainsOpt{}, want: false},

		// Invalid with ignoreCase false
		{name: "Case-sensitive no match", param1: "Hello World", param2: "WORLD", param3: ContainsOpt{IgnoreCase: false}, want: false},
		{name: "Case-insensitive fail", param1: "FOOBAR", param2: "baz", param3: ContainsOpt{IgnoreCase: true}, want: false},

		// Invalid with minimum occurrences
		{name: "Minimum occurrences not met", param1: "hello world", param2: "hello", param3: ContainsOpt{MinOccurrences: 2}, want: false},
		{name: "Zero occurrences required", param1: "abc123", param2: "123", param3: ContainsOpt{MinOccurrences: 0}, want: true},
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
