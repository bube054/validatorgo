package validatorgo

import "testing"

func TestIsLicensePlate(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// Valid examples
		// Invalid examples
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsLicensePlate(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
