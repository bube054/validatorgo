package sanitizers

import "testing"

func TestToFloat(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   float64
	}{
		{
			"test 1",
			"1.00",
			1.00,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToFloat(test.param1)

			if result != test.want {
				t.Errorf("got `%.2f`, wanted `%.2f`", result, test.want)
			}
		})
	}
}
