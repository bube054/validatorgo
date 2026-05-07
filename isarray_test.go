package validatorgo

import (
	"testing"
)

func TestIsArray(t *testing.T) {
	var one uint = 1
	var two uint = 2
	var three uint = 3

	tests := []struct {
		name   string
		param1 string
		param2 *IsArrayOpts
		want   bool
	}{
		// Valid json array
		{name: "Valid array", param1: `["item1", "item2"]`, param2: nil, want: true},
		{name: "Valid empty array", param1: `[]`, param2: nil, want: true},
		{name: "Valid array with Min", param1: `["item1", "item2"]`, param2: &IsArrayOpts{Min: &two}, want: true},
		{name: "Valid array with Max", param1: `["item1", "item2"]`, param2: &IsArrayOpts{Max: &three}, want: true},
		{name: "Valid array with MIn & Max", param1: `["item1", "item2"]`, param2: &IsArrayOpts{Min: &two, Max: &three}, want: true},
		// Invalid json array
		{name: "Invalid array is object", param1: `{"name": "John", "age": 30}`, param2: nil, want: false},
		{name: "Invalid array bad format", param1: `[`, param2: nil, want: false},
		{name: "Invalid array precedes Min", param1: `["item1", "item2"]`, param2: &IsArrayOpts{Min: &three}, want: false},
		{name: "Invalid array exceeds Max", param1: `["item1", "item2"]`, param2: &IsArrayOpts{Max: &one}, want: false},
		{name: "Invalid array precedes Min and exceeds Max", param1: `["item1", "item2"]`, param2: &IsArrayOpts{Min: &three, Max: &one}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := IsArray(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
