package govalidator

import "testing"

func TestIsISIN(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid ISIN
		{name: "Valid ISIN", param1: "US0378331005", want: true},
		{name: "Valid ISIN", param1: "AU0000XVGZA3", want: true},
		{name: "Valid ISIN", param1: "AU0000VXGZA3", want: true},
		// Invalid ISIN
		{name: "The transposition typo is caught by the checksum constraint.", param1: "US0373831005", want: false},
		{name: "The substitution typo is caught by the format constraint.", param1: "U50378331005", want: false},
		{name: "The duplication typo is caught by the format constraint.", param1: "US03378331005", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISIN(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
