package validatorgo

import "testing"

func TestIsMacAddress(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsMacAddressOpts
		want   bool
	}{
		// Valid Mac address default options
		{name: "Colon-separated", param1: "00:1A:2B:3C:4D:5E", want: true},
		{name: "Hyphen-separated", param1: "00-1A-2B-3C-4D-5E", want: true},
		{name: "Without separators (optional)", param1: "001A2B3C4D5E", want: true},
		// Invalid Mac address default options
		{name: "Contains invalid characters ZZ", param1: "00:1A:2B:3C:4D:ZZ", want: false},
		{name: "Incorrect length", param1: "001A2B3C4D", want: false},
		{name: "Too many groups", param1: "00:1A:2B:3C:4D:5E:FF", want: false},
		// Valid Mac address NoSeparators true
		{name: "Correct no separator", param1: "001A2B3C4D5E", param2: &IsMacAddressOpts{NoSeparators: true}, want: true},
		// Invalid Mac address NoSeparators true
		{name: "Incorrect no separator", param1: "00:1A:2B:3C:4D:5E", param2: &IsMacAddressOpts{NoSeparators: true}, want: false},
		{name: "Incorrect wrong types", param1: "00:1A:2B:3C:4D:5E", param2: &IsMacAddressOpts{NoSeparators: true, Type: strPtr("56")}, want: false},
		// Valid Mac address EUI 64
		{name: "Colon-separated EUI 64", param1: "00:1A:2B:3C:4D:5E:6F:70", param2: &IsMacAddressOpts{Type: strPtr("64")}, want: true},
		{name: "Hyphen-separated EUI 64", param1: "00-1A-2B-3C-4D-5E-6F-70", param2: &IsMacAddressOpts{Type: strPtr("64")}, want: true},
		{name: "Without separators (optional) EUI 64", param1: "001A2B3C4D5E6F70", param2: &IsMacAddressOpts{Type: strPtr("64")}, want: true},
		// Invalid Mac address EUI 64
		{name: "Contains invalid characters ZZ EUI 64", param1: "00:1A:2B:3C:4D:5E:ZZ:70", want: false},
		{name: "Incorrect length EUI 64", param1: "00:1A:2B:3C:4D:5E:6F", want: false},
		// Valid Mac address EUI 64, NoSeparators true
		{name: "Correct no separator EUI 64", param1: "001A2B3C4D5E6F70", param2: &IsMacAddressOpts{NoSeparators: true, Type: strPtr("64")}, want: true},
		// Invalid Mac address EUI 64, NoSeparators true
		{name: "Incorrect no separator  EUI 64", param1: "00-1A-2B-3C-4D-5E-6F-70", param2: &IsMacAddressOpts{NoSeparators: true, Type: strPtr("64")}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMacAddress(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
