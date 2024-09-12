package govalidator

import "testing"

func TestIsISRC(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 bool
		want   bool
	}{
		// Valid ISRC without hyphens
		{name: "Valid ISRC without hyphens", param1: "AASKG1912345", param2: false, want: true},
		{name: "Valid ISRC without hyphens", param1: "USABC2021234", param2: false, want: true},
		{name: "Valid ISRC without hyphens", param1: "GBXYZ2112345", param2: false, want: true},
		{name: "Valid ISRC without hyphens", param1: "FRDEF1912345", param2: false, want: true},
		// Invalid ISRC without hyphens
		{name: "Valid ISRC but contains hyphens", param1: "AA-SKG-19-12345", param2: false, want: false},
		{name: "Too few digits in the unique identifier section", param1: "USABC201234", param2: false, want: false},
		{name: "Starts with numbers instead of country code", param1: "1234ABC202001234", param2: false, want: false},
		{name: "Too many characters in the unique identifier", param1: "USABC20987654321", param2: false, want: false},
		// Valid ISRC with hyphens
		{name: "Valid ISRC with hyphens", param1: "AA-SKG-19-12345", param2: true, want: true},
		{name: "Valid ISRC without hyphens", param1: "USABC2021234", param2: true, want: true},
		{name: "Valid ISRC with hyphens", param1: "GB-XYZ-21-12345", param2: true, want: true},
		{name: "Valid ISRC without hyphens", param1: "FRDEF1912345", param2: true, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISRC(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}