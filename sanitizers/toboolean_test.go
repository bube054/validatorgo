package sanitizers

import (
	"testing"
)

func TestToBoolean(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 bool
		want   bool
	}{
		{
			"test 1",
			"true",
			false,
			true,
		},
		{
			"test 2",
			"true",
			true,
			true,
		},
		{
			"test 3",
			"false",
			false,
			false,
		},
		{
			"test 4",
			"false",
			true,
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToBoolean(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
