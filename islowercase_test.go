package validatorgo

import "testing"

func TestIsLowerCase(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Example
		{name: "Is lowercase", param1: "hello", want: true},
		{name: "Empty string is lowercase", param1: "", want: true},
		// Invalid Example
		{name: "Is not lowercase", param1: "WORLD", want: false},
		{name: "Only few letter are lowercase", param1: "ExaMPle", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsLowerCase(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
