package validatorgo

import "testing"

func TestIsISO31661Numeric(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid IsISO31661Numeric
		{name: "Valid IsISO31661Numeric, Argentina", param1: "032", want: true},
		{name: "Valid IsISO31661Numeric, Belgium", param1: "056", want: true},
		{name: "Valid IsISO31661Numeric, Cuba", param1: "192", want: true},
		{name: "Valid IsISO31661Numeric, India", param1: "356", want: true},
		{name: "Valid IsISO31661Numeric, Oman", param1: "512", want: true},
		// Invalid IsISO31661Numeric
		{name: "Not three numbers", param1: "56", want: false},
		{name: "Not a valid code", param1: "525", want: false},
		{name: "Contains letters", param1: "abc", want: false},
		{name: "Only one number", param1: "9", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISO31661Numeric(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
