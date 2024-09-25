package validatorgo

import (
	"testing"
)

func TestIsBic(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid BIC codes
		{name: "Valid BIC 8 characters", param1: "DEUTDEFF", want: true},
		{name: "Valid BIC 11 characters", param1: "DEUTDEFF500", want: true},
		{name: "Valid BIC with digits", param1: "NEDSZAJJXXX", want: true},
		{name: "Valid barclays swift or bic code", param1: "BARCGB22XXX", want: true},
		{name: "Valid HSBC swift or bic code", param1: "HBUKGB4106W", want: true},
		{name: "Valid Goldman Sachs swift or bic code", param1: "GOSGGB21", want: true},
		{name: "Valid Union Bank swift or bic code", param1: "UBNINGLA", want: true},

		// Invalid BIC codes
		{name: "Invalid BIC with too few characters", param1: "DEUTDE", want: false},
		{name: "Invalid BIC with too many characters", param1: "DEUTDEFF500XYZ", want: false},
		{name: "Invalid BIC with special characters", param1: "DEUT!EFF500", want: false},
		{name: "Invalid BIC with spaces", param1: "DEUT DEFF 500", want: false},
		{name: "Invalid swift or bic code length", param1: "ABCDUS12XXXX", want: false},
		{name: "Invalid swift or bic code bank code", param1: "ABCDEFGH", want: false},

		// Edge cases
		{name: "Empty string", param1: "", want: false},
		{name: "Valid BIC with lowercase letters", param1: "deutdeff500", want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBic(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}

}
