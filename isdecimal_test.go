package validatorgo

import "testing"

func TestIsDecimal(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsDecimalOpts
		want   bool
	}{
		{
			name:   "Is valid with no decimal and default option set",
			param1: "123",
			param2: IsDecimalOpts{},
			want:   true,
		},
		{
			name:   "Is valid 2 decimal digit and default option set",
			param1: "123.00",
			param2: IsDecimalOpts{},
			want:   true,
		},
		{
			name:   "Is invalid max decimal digits exceeded and default option set",
			param1: "123.00000000000",
			param2: IsDecimalOpts{},
			want:   false,
		},
		{
			name:   "Is valid with no digits before decimal point and default option set",
			param1: ".123",
			param2: IsDecimalOpts{},
			want:   true,
		},
		{
			name:   "Test british valid range",
			param1: ".123",
			param2: IsDecimalOpts{},
			want:   true,
		},
		{
			name:   "Is invalid with no decimal and force decimal is on",
			param1: "123",
			param2: IsDecimalOpts{ForceDecimal: true},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsDecimal(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
