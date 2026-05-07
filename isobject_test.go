package validatorgo

import (
	"testing"
)

func TestIsObject(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsObjectOpts
		want   bool
	}{
		// Valid json default config with strictness
		{name: "Valid JSON object", param1: `{"name": "John", "age": 30}`, param2: nil, want: true},
		{name: "Valid empty JSON", param1: `{}`, param2: nil, want: true},
		// Invalid json default config with strictness
		{name: "Invalid JSON object, no arrays", param1: `["item1", "item2"]`, param2: nil, want: false},
		{name: "Invalid JSON object, no null", param1: `null`, param2: nil, want: false},
		{name: "Invalid JSON object", param1: `{"name": "John", "age": 30`, param2: nil, want: false},

		// Valid json without strictness
		{name: "Valid JSON object", param1: `{"name": "John", "age": 30}`, param2: &IsObjectOpts{Strict: Bool(false)}, want: true},
		{name: "Valid JSON array", param1: `["item1", "item2"]`, param2: &IsObjectOpts{Strict: Bool(false)}, want: true},
		{name: "Valid JSON null", param1: `null`, param2: &IsObjectOpts{Strict: Bool(false)}, want: true},
		// Invalid json without strictness
		{name: "Valid JSON, is string", param1: `"Just a string"`, param2: &IsObjectOpts{Strict: Bool(false)}, want: false},
		{name: "Valid JSON format", param1: `{invalid json}`, param2: &IsObjectOpts{Strict: Bool(false)}, want: false},
		{name: "Valid JSON format misspell null", param1: `nil`, param2: &IsObjectOpts{Strict: Bool(false)}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsObject(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
