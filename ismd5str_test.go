package validatorgo

import "testing"

func TestIsMD5(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid MD5 hashes
		{name: "Valid MD5 - all lowercase", param1: "9e107d9d372bb6826bd81d3542a419d6", want: true},
		{name: "Valid MD5 - mixed case", param1: "9E107D9D372BB6826BD81D3542A419D6", want: true},
		{name: "Valid MD5 - numeric", param1: "12345678901234567890123456789012", want: true},

		// Invalid MD5 hashes
		{name: "Invalid MD5 - too short", param1: "9e107d9d372bb6826bd81d3542a419d", want: false},
		{name: "Invalid MD5 - too long", param1: "9e107d9d372bb6826bd81d3542a419d61", want: false},
		{name: "Invalid MD5 - contains non-hex character", param1: "9e107d9d372bb6826bd81d3542a419dz", want: false},
		{name: "Invalid MD5 - empty string", param1: "", want: false},
		{name: "Invalid MD5 - non-hex characters", param1: "g3fcd3d76192e4007dfb496cca67e13b", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMD5(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
