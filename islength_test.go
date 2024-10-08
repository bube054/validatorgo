package validatorgo

import "testing"

func TestIsLength(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsLengthOpts
		want   bool
	}{
		// Valid examples
		{name: "Default config used", param1: "hello", param2: nil, want: true},
		{name: "Only minimum length specified", param1: "hello", param2: &IsLengthOpts{Min: 3}, want: true},
		{name: "Only minimum length specified", param1: "world", param2: &IsLengthOpts{Min: 5}, want: true},
		{name: "Both minimum and maximum length specified", param1: "example", param2: &IsLengthOpts{Min: 2, Max: uintPtr(10)}, want: true},
		{name: "Both minimum and maximum length specified", param1: "longword", param2: &IsLengthOpts{Min: 5, Max: uintPtr(10)}, want: true},
		{name: "No options, default to checking if string exists", param1: "", param2: &IsLengthOpts{Min: 0}, want: true},
		{name: "No options, default to checking if string exists", param1: "a", param2: &IsLengthOpts{Min: 1}, want: true},
		{name: "String with spaces", param1: "  spaced  ", param2: &IsLengthOpts{Min: 5, Max: uintPtr(15)}, want: true},
		// Invalid examples
		{name: "Below the minimum length", param1: "hi", param2: &IsLengthOpts{Min: 3}, want: false},
		{name: "Above the maximum length", param1: "toolongword", param2: &IsLengthOpts{Min: 2, Max: uintPtr(10)}, want: false},
		{name: "Empty string with a positive minimum", param1: "", param2: &IsLengthOpts{Min: 1}, want: false},
		{name: "String containing spaces (spaces count toward length)", param1: "  ", param2: &IsLengthOpts{Min: 3, Max: uintPtr(5)}, want: false},
		{name: "Missing max constraint but exceeding the default allowed length", param1: "overlimit", param2: &IsLengthOpts{Min: 1, Max: uintPtr(5)}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsLength(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
