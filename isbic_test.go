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
		{
			name:   "Valid barclays swift or bic code",
			param1: "BARCGB22XXX",
			want:   true,
		},
		{
			name:   "Valid HSBC swift or bic code",
			param1: "HBUKGB4106W",
			want:   true,
		},
		{
			name:   "Valid Goldman Sachs swift or bic code",
			param1: "GOSGGB21",
			want:   true,
		},
		{
			name:   "Invalid Union Bank swift or bic code",
			param1: "UBNINGLA",
			want:   true,
		},
		{
			name:   "Invalid swift or bic code length",
			param1: "ABCDUS12XXXX",
			want:   false,
		},
		{
			name:   "Invalid swift or bic code bank code",
			param1: "ABCDEFGH",
			want:   false,
		},
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
