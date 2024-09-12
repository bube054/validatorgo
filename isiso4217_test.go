package govalidator

import "testing"

func TestIsIso4217(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid IsIso4217
		{name: "Valid IsIso4217, UNITED ARAB EMIRATES (THE)", param1: "AED", want: true},
		{name: "Valid IsIso4217, TURKEY", param1: "TRY", want: true},
		{name: "Valid IsIso4217,YEMEN", param1: "YER", want: true},
		{name: "Valid IsIso4217, TUNISA", param1: "TND", want: true},
		{name: "Valid IsIso4217, SWEDEN", param1: "SEK", want: true},
		// Invalid IsIso4217
		{name: "Not three letters", param1: "AE", want: false},
		{name: "Not a valid code", param1: "CHD", want: false},
		{name: "Contains numbers", param1: "123", want: false},
		{name: "Only one letter", param1: "A", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIso4217(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
