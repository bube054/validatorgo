package validatorgo

import "testing"

func TestIsIP(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		{name: "Valid version but valid IPv4", param1: "192.168.0.1", param2: "4", want: true},
		{name: "Valid version but valid IPv6, no", param1: "2001:0db8:85a3:0000:0000:8a2e:0370:7334", param2: "6", want: true},
		{name: "Valid version 4, no version provided", param1: "192.168.0.1", param2: "", want: true},
		{name: "Valid version 6, no version provided", param1: "2001:0db8:85a3:0000:0000:8a2e:0370:7334", param2: "", want: true},
		{name: "Invalid version 4, no version provided", param1: "a2.168.0.1", param2: "", want: false},
		{name: "Invalid version 6, no version provided", param1: "z001:0dbz:85a3:0000:0000:8a2e:0370:7334", param2: "", want: false},
		{name: "valid version 4, incorrect version provided", param1: "a2.168.0.1", param2: "20", want: false},
		{name: "valid version 6, incorrect version provided", param1: "z001:0dbz:85a3:0000:0000:8a2e:0370:7334", param2: "xyz", want: false},

		{name: "Valid IPv4 Google's public DNS", param1: "8.8.8.8", param2: "4", want: true},
		{name: "Invalid IPv4 Octets must be between 0 and 255", param1: "256.256.256.256", param2: "4", want: false},
		{name: "Valid IPv4 Leading zeros are allowed", param1: "192.168.001.001", param2: "4", want: true},
		{name: "Invalid IPv4 CIDR notation is not a valid pure IP address", param1: "192.168.0.1/24", param2: "4", want: false},

		{name: "Valid IPv6", param1: "2001:0db8:85a3:0000:0000:8a2e:0370:7334", param2: "6", want: true},
		{name: "Valid IPv6 Shortened form using `::`", param1: "2001:db8::2:1", param2: "6", want: true},
		{name: "Invalid IPv6 should have 8 groups", param1: "2001:0db8:85a3:0000:0000:8a2e:0370:7334:abcd", param2: "6", want: false},
		{name: "Invalid IPv6 hexadecimal character 'g'", param1: "2001:db8:85a3::g", param2: "6", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIP(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
