package validatorgo

import "testing"

func TestIsISO31661Alpha2(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid IsISO31661Alpha2
		{name: "Valid IsISO31661Alpha2, english", param1: "EN", want: true},
		{name: "Valid IsISO31661Alpha2, french", param1: "FR", want: true},
		{name: "Valid IsISO31661Alpha2, spanish", param1: "ES", want: true},
		{name: "Valid IsISO31661Alpha2, german", param1: "DE", want: true},
		{name: "Valid IsISO31661Alpha2, chinese", param1: "ZH", want: true},
		// Invalid IsISO31661Alpha2
		{name: "Not two letters", param1: "eng", want: false},
		{name: "Not a valid code", param1: "frz", want: false},
		{name: "Contains numbers", param1: "123", want: false},
		{name: "Only one letter", param1: "e", want: false},
		{name: "Lowercase, should be uppercase", param1: "fr", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISO31661Alpha2(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
