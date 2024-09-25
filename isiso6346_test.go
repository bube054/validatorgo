package validatorgo

import "testing"

func TestIsISO6346(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid ISO6346
		{name: "Valid ISO6346", param1: "CSQU3054383", want: true},
		{name: "Valid ISO6346", param1: "HLCU1234568", want: true},
		{name: "Valid ISO6346", param1: "MAEU9876542", want: true},
		{name: "Valid ISO6346", param1: "APLZ3216542", want: true},
		{name: "Valid ISO6346", param1: "CMAU7654327", want: true},
		// Invalid ISO6346
		{name: "Invalid owner code", param1: "CSQX3054383", want: false},
		{name: "Contains letters in the serial number", param1: "HLCU12345AB", want: false},
		{name: "Serial number too short", param1: "MAEU98765", want: false},
		{name: "Serial number too long", param1: "APLZ32165488", want: false},
		{name: "Missing one digit in serial number", param1: "CMAU765432", want: false},
		{name: "Wrong check digit", param1: "CMAU7654329", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISO6346(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
