package sanitizer

import (
	"testing"
)

func TestToBoolean(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 bool
		want   bool
	}{
		{name: "Non-strict mode: empty string returns false", param1: "", param2: false, want: false},
		{name: "Non-strict mode: '0' returns false", param1: "0", param2: false, want: false},
		{name: "Non-strict mode: 'false' returns false", param1: "false", param2: false, want: false},
		{name: "Non-strict mode: 'true' returns true", param1: "true", param2: false, want: true},
		{name: "Non-strict mode: '1' returns true", param1: "1", param2: false, want: true},
		{name: "Non-strict mode: random string returns true", param1: "random", param2: false, want: true},
		{name: "Strict mode: '1' returns true", param1: "1", param2: true, want: true},
		{name: "Strict mode: 'true' returns true", param1: "true", param2: true, want: true},
		{name: "Strict mode: '0' returns false", param1: "0", param2: true, want: false},
		{name: "Strict mode: 'false' returns false", param1: "false", param2: true, want: false},
		{name: "Strict mode: random string returns false", param1: "random", param2: true, want: false},
		{name: "Strict mode: empty string returns false", param1: "", param2: true, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToBoolean(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
