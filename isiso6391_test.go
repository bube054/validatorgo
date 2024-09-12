package govalidator

import "testing"

func TestIsISO6391(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid IsISO6391
		{name: "Valid IsISO6391, english", param1: "en", want: true},
		{name: "Valid IsISO6391, french", param1: "fr", want: true},
		{name: "Valid IsISO6391, spanish", param1: "es", want: true},
		{name: "Valid IsISO6391, german", param1: "de", want: true},
		{name: "Valid IsISO6391, chinese", param1: "zh", want: true},
		// Invalid IsISO6391
		{name: "Not two letters", param1: "eng", want: false},
		{name: "Not a valid code", param1: "frz", want: false},
		{name: "Contains numbers", param1: "123", want: false},
		{name: "Only one letter", param1: "e", want: false},
		{name: "Uppercase, should be lowercase", param1: "FR", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISO6391(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
