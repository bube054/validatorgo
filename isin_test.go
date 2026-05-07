package validatorgo

import "testing"

func TestIsIn(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 []string
		want   bool
	}{
		{name: "Value is present", param1: "apple", param2: []string{"apple", "banana", "grape"}, want: true},
		{name: "Value is not present", param1: "orange", param2: []string{"apple", "banana", "grape"}, want: false},

		// Valid - found in slice
		{name: "Found in slice", param1: "foo", param2: []string{"foo", "bar"}, want: true},
		{name: "Found bar in slice", param1: "bar", param2: []string{"foo", "bar"}, want: true},
		{name: "Found in number strings", param1: "1", param2: []string{"1", "2", "3"}, want: true},
		{name: "Found 2", param1: "2", param2: []string{"1", "2", "3"}, want: true},
		{name: "Found 3", param1: "3", param2: []string{"1", "2", "3"}, want: true},

		// Invalid - not found in slice
		{name: "Not in slice", param1: "foobar", param2: []string{"foo", "bar"}, want: false},
		{name: "Reversed not found", param1: "barfoo", param2: []string{"foo", "bar"}, want: false},
		{name: "Empty not in slice", param1: "", param2: []string{"foo", "bar"}, want: false},
		{name: "4 not in 1,2,3", param1: "4", param2: []string{"1", "2", "3"}, want: false},
		{name: "Empty string not in numbers", param1: "", param2: []string{"1", "2", "3"}, want: false},

		// Edge cases
		{name: "Nil values slice", param1: "foo", param2: nil, want: false},
		{name: "Empty values slice", param1: "foo", param2: []string{}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsIn(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
