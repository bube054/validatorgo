package validatorgo

import "testing"

func TestIsJSON(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid JSON
		{name: "Basic json object", param1: `{"name": "John", "age": 30, "city": "New York"}`, want: true},
		{name: "Nested JSON object", param1: `{"person": {"name": "Alice", "age": 25}, "status": true}`, want: true},
		{name: "JSON array", param1: `[{"id": 1, "value": "A"}, {"id": 2, "value": "B"}]`, want: true},
		{name: "Empty object", param1: `{}`, want: true},
		{name: "Empty array", param1: `[]`, want: true},
		{name: "Only string", param1: `"Hello World"`, want: true},
		{name: "Only number", param1: `30`, want: true},
		{name: "Only boolean True", param1: `true`, want: true},
		{name: "Only boolean False", param1: `false`, want: true},
		{name: "Only null", param1: `null`, want: true},
		// Invalid JSON
		{name: "Single quotes around strings", param1: `{'name': 'John', 'age': 30}`, want: false},
		{name: "Trailing comma", param1: `{"name": "John", "age": 30,}`, want: false},
		{name: "Unquoted keys", param1: `{name: "John", age: 30}`, want: false},
		{name: "Incorrect boolean values", param1: `{"status": True}`, want: false},
		{name: "Unterminated JSON object or array", param1: `{"name": "John", "age": 30`, want: false},
		{name: "Mixed data types without proper structure", param1: `{123: "abc", "key": value}`, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsJSON(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
