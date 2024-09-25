package validatorgo

import "testing"

func TestIsTaxID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Valid Tax id's
		{want: true},
		// Invalid Tax id's
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsTaxID(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
