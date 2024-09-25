package validatorgo

import "testing"

func TestIsOctal(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid octal numbers
		{name: "Valid Octal - Simple", param1: "07", want: true},
		{name: "Valid Octal - Complex", param1: "0754321", want: true},
		{name: "Valid Octal - Zero", param1: "0", want: true},
		{name: "Valid Octal - Max", param1: "07777777777", want: true},

		// Invalid octal numbers
		{name: "Invalid Octal - Contains 8", param1: "078", want: false},
		{name: "Invalid Octal - Contains 9", param1: "079", want: false},
		{name: "Invalid Octal - Non-numeric character", param1: "075a321", want: false},
		{name: "Invalid Octal - Negative sign", param1: "-0754", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsOctal(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
