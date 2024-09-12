package govalidator

import "testing"

func TestIsISO31661Alpha3(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid IsISO31661Alpha3
		{name: "Valid IsISO31661Alpha3, Aruba", param1: "ABW", want: true},
		{name: "Valid IsISO31661Alpha3, Benin", param1: "BEN", want: true},
		{name: "Valid IsISO31661Alpha3, Colombia", param1: "COL", want: true},
		{name: "Valid IsISO31661Alpha3, Guam", param1: "GUM", want: true},
		{name: "Valid IsISO31661Alpha3, Zambia", param1: "ZMB", want: true},
		// Invalid IsISO31661Alpha3
		{name: "Not three letters", param1: "ES", want: false},
		{name: "Not a valid code", param1: "RGS", want: false},
		{name: "Contains numbers", param1: "123", want: false},
		{name: "Only one letter", param1: "e", want: false},
		{name: "Lowercase, should be uppercase", param1: "sdn", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISO31661Alpha3(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
