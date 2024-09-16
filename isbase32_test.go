package validatorgo

import (
	"testing"
)

func TestIsBase32(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsBase32Opts
		want   bool
	}{
		// Standard Base32 (RFC 4648) tests
		{name: "Basic Base32 valid", param1: "JBSWY3DPEBLW64TMMQ======", param2: IsBase32Opts{}, want: true},
		{name: "Valid Base32 no padding", param1: "JBSWY3DPEBLW64TMMQ", param2: IsBase32Opts{}, want: true},
		{name: "Valid Base32 but short", param1: "JBSWY3DP", param2: IsBase32Opts{}, want: true},
		{name: "Valid Base32", param1: "JBSWY3DPEBLW64TMMQ=====", param2: IsBase32Opts{}, want: true},
		{name: "Valid Base32 with $ characters", param1: "JBSWY3DP$BLW64TMMQ======", param2: IsBase32Opts{}, want: false},
		{name: "Invalid Base32 lowercase", param1: "jbswy3dpeblw64tmmq======", param2: IsBase32Opts{}, want: false},

		// Crockford Base32 tests
		{name: "Crockford Base32 valid", param1: "91JPRV3F41VPYWKCCG", param2: IsBase32Opts{Crockford: true}, want: true},
		{name: "Crockford Base32 valid with hyphens", param1: "91-JP-RV-3F-41-VP-YW-KC-CG", param2: IsBase32Opts{Crockford: true}, want: true},
		{name: "Crockford Base32 lowercase", param1: "91jprv3f41vpywkccg", param2: IsBase32Opts{Crockford: true}, want: true},
		{name: "Crockford Base32 with special characters", param1: "91JPRV3F!41VPYWKCCG", param2: IsBase32Opts{Crockford: true}, want: false},
		{name: "Crockford Base32 invalid due to incorrect padding", param1: "A1B2C3D4E5F6G7H8I9J", param2: IsBase32Opts{Crockford: true}, want: false},

		// Edge cases
		{name: "Empty string", param1: "", param2: IsBase32Opts{}, want: false},
		{name: "Single Base32 character", param1: "J", param2: IsBase32Opts{}, want: false},
		{name: "Valid Base32 with spaces", param1: "JBSWY 3DPE BLW64 TMMQ======", param2: IsBase32Opts{}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBase32(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
