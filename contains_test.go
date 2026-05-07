package validatorgo

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		param3 *ContainsOpt
		want   bool
	}{
		// Valid default config
		{name: "Basic match", param1: "hello world", param2: "world", param3: &ContainsOpt{}, want: true},
		{name: "Basic match digits", param1: "abc123", param2: "123", param3: &ContainsOpt{}, want: true},

		// Valid with ignoreCase true
		{name: "Case-insensitive match", param1: "Hello World", param2: "hello", param3: &ContainsOpt{IgnoreCase: true}, want: true},
		{name: "Case-insensitive match mixed", param1: "FOOBAR", param2: "bar", param3: &ContainsOpt{IgnoreCase: true}, want: true},

		// Valid with minimum occurrences
		{name: "Minimum occurrences met", param1: "hello hello world", param2: "hello", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: true},
		{name: "Minimum occurrences default", param1: "abc123", param2: "123", param3: &ContainsOpt{MinOccurrences: Int(1)}, want: true},

		// Invalid default config
		{name: "No match", param1: "hello world", param2: "earth", param3: &ContainsOpt{}, want: false},
		{name: "No match digits", param1: "abc123", param2: "xyz", param3: &ContainsOpt{}, want: false},

		// Invalid with ignoreCase false
		{name: "Case-sensitive no match", param1: "Hello World", param2: "WORLD", param3: &ContainsOpt{IgnoreCase: false}, want: false},
		{name: "Case-insensitive fail", param1: "FOOBAR", param2: "baz", param3: &ContainsOpt{IgnoreCase: true}, want: false},

		// Invalid with minimum occurrences
		{name: "Minimum occurrences not met", param1: "hello world", param2: "hello", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: false},
		{name: "Zero occurrences required", param1: "abc123", param2: "123", param3: &ContainsOpt{MinOccurrences: Int(0)}, want: true},

		// Test for nil param3
		{name: "Nil default, basic match", param1: "hello world", param2: "world", param3: nil, want: true},
		{name: "Nil default, no match", param1: "hello world", param2: "earth", param3: nil, want: false},

		// Valid default opts with seed "foo"
		{name: "Seed at start", param1: "foobar", param2: "foo", param3: nil, want: true},
		{name: "Seed at end", param1: "bazfoo", param2: "foo", param3: nil, want: true},
		// Invalid default opts with seed "foo"
		{name: "Partial seed", param1: "fobar", param2: "foo", param3: nil, want: false},

		// Valid with ignoreCase true
		{name: "Case-insensitive FOO match", param1: "FOObar", param2: "foo", param3: &ContainsOpt{IgnoreCase: true}, want: true},
		{name: "Case-insensitive Foo match", param1: "Foo", param2: "foo", param3: &ContainsOpt{IgnoreCase: true}, want: true},
		{name: "Case-insensitive BAZfoo", param1: "BAZfoo", param2: "foo", param3: &ContainsOpt{IgnoreCase: true}, want: true},
		// Invalid with ignoreCase true
		{name: "Case-insensitive baxoof", param1: "baxoof", param2: "foo", param3: &ContainsOpt{IgnoreCase: true}, want: false},

		// Valid with minOccurrences 2
		{name: "Three occurrences", param1: "foofoofoo", param2: "foo", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: true},
		{name: "Separated occurrences", param1: "12foo124foo", param2: "foo", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: true},
		{name: "Overlapping-like occurrences", param1: "fofooofoooofoooo", param2: "foo", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: true},
		{name: "Adjacent with separator", param1: "foo1foo", param2: "foo", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: true},
		// Invalid with minOccurrences 2
		{name: "Only one occurrence with min 2", param1: "foobar", param2: "foo", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: false},
		{name: "Case-sensitive miss with min 2", param1: "Fooofoo", param2: "foo", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: false},
		{name: "Partial match foofo", param1: "foofo", param2: "foo", param3: &ContainsOpt{MinOccurrences: Int(2)}, want: false},

		// Edge cases
		{name: "Empty string empty seed", param1: "", param2: "", param3: nil, want: true},
		{name: "Empty seed in string", param1: "hello", param2: "", param3: nil, want: true},
		{name: "Seed longer than string", param1: "hi", param2: "hello", param3: nil, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Contains(test.param1, test.param2, test.param3)

			assertValidation(t, result, test.want, err)
		})
	}
}
