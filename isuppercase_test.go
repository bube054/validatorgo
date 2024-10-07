package validatorgo

import "testing"

func TestIsUpperCase(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Example
		{name: "Is uppercase", param1: "HELLO", want: true},
		// Invalid Example
		{name: "Empty string is not uppercase", param1: "", want: false},
		{name: "Is not uppercase", param1: "hello", want: false},
		{name: "Only few letter are uppercase", param1: "ExaMPle", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsUpperCase(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
