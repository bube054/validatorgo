package validatorgo

import (
	"testing"
)

func TestIsBase58(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Base58 test cases
		{name: "Basic Base58 valid", param1: "JDEe4eBcQWTe1GV34By3dz7", want: true},
		{name: "Base58 valid with mixed case", param1: "JDEE4EBCQWTE1GV34BY3DZ7", want: true},
		{name: "Base58 valid with varying length", param1: "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz", want: true},
		{name: "Base58 valid with short length", param1: "1A3", want: true},

		// Invalid Base58 test cases
		{name: "Invalid Base58 with non-base58 characters", param1: "JDEe4eBcQWTe1GV34By3dz7@", want: false},
		{name: "Invalid Base58 with spaces", param1: "JDEe 4eBcQWTe1GV34By3dz7", want: false},
		{name: "Invalid Base58 with special characters", param1: "JDEe4eBcQWTe1GV34By3dz7!", want: false},
		{name: "Invalid Base58 incorrect characters", param1: "JDEe4eBcQWTe1GV34BY3DZ#", want: false},
		{name: "Empty string", param1: "", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBase58(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
