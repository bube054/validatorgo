package validatorgo

import "testing"

func TestIsIPrange(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Valid IPv4 Ranges
		{name: "Covers IPs from 192.168.0.0 to 192.168.0.255", param1: "192.168.0.0/24", param2: "4", want: true},
		{name: "Covers IPs from 10.0.0.0 to 10.255.255.255", param1: "10.0.0.0/8", param2: "4", want: true},
		{name: "Covers IPs from 172.16.0.0 to 172.31.255.255", param1: "172.16.0.0/12", param2: "4", want: true},
		{name: "Dash-separated Range 192.168.1.10 - 192.168.1.20", param1: "192.168.1.10 - 192.168.1.20", param2: "4", want: true},
		{name: "Dash-separated Range 10.0.0.5 - 10.0.0.50", param1: "10.0.0.5 - 10.0.0.50", param2: "4", want: true},
		// Invalid IPv4 Ranges
		{name: "Invalid prefix length; must be between 0 and 32", param1: "192.168.0.0/33", param2: "4", want: false},
		{name: "Ending IP is lower than the starting IP", param1: "192.168.1.100 - 192.168.1.10", param2: "4", want: false},
		{name: "IP address out of range valid octets are 0-255", param1: "192.168.1.1 - 192.168.1.500", param2: "4", want: false},
		{name: "Incomplete IP address", param1: "10.0.0.1 - 10.0.0", param2: "4", want: false},
		// Valid IPv6 Ranges
		{name: "Covers IPs from 2001:0db8:0000:0000:0000:0000:0000:0000 to 2001:0db8:ffff:ffff:ffff:ffff:ffff:ffff", param1: "2001:0db8::/32", param2: "6", want: true},
		{name: "Covers link-local IPs from fe80:0000:0000:0000:0000:0000:0000:0000 to febf:ffff:ffff:ffff:ffff:ffff:ffff:ffff)", param1: "2001:0db8::/32", param2: "6", want: true},
		{name: "Dash-separated Range 2001:0db8:85a3::8a2e:0370:7334 - 2001:0db8:85a3::8a2e:0370:7339", param1: "2001:0db8::/32", param2: "6", want: true},
		{name: "Dash-separated Range 2001:0db8::1 - 2001:0db8::ff", param1: "2001:0db8::/32", param2: "6", want: true},
		// Invalid IPv6 Ranges
		{name: "Invalid prefix length must be between 0 and 128", param1: "2001:0db8::/129", param2: "6", want: false},
		{name: "Ending IP is lower than the starting IP", param1: "2001:0db8::ffff - 2001:0db8::1", param2: "6", want: false},
		{name: "Invalid character 'g'; only hexadecimal digits are valid", param1: "2001:0db8::g234", param2: "6", want: false},
		{name: "Invalid IP format; multiple :: are not allowed", param1: "2001:0db8::85a3::8a2e:0370:7334 - 2001:0db8::85a3::8a2e:0370:7339", param2: "6", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIPRange(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
