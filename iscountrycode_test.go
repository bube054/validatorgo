package validatorgo

import "testing"

func TestIsCountryCode(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid County-Codes
		{name: "Is valid county code", param1: "AL", want: true},
		{name: "Is valid county code", param1: "BF", want: true},
		{name: "Is valid county code", param1: "PM", want: true},
		{name: "Is valid county code", param1: "TR", want: true},
		// Invalid County-Codes
		{name: "Is invalid county code", param1: "XX", want: false},
		{name: "Is invalid county code", param1: "YY", want: false},
		{name: "Is invalid county code", param1: "YY", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsCountryCode(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
