package validatorgo

import (
	"testing"
)

func TestIsCurrency(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsCurrencyOpts
		want   bool
	}{
		{
			name:   "Test symbol",
			param1: "£10.50",
			param2: IsCurrencyOpts{Symbol: "£"},
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsCurrency(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
