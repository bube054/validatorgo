package sanitizers

import (
	"testing"
)

func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{
			"test 1",
			"vasya+pupkin@gmail.com",
			"vasya@gmail.com",
		},
		{
			"test 2",
			"t.e-St+vasya@gmail.com",
			"te-st@gmail.com",
		},
		{
			"test 3",
			"t+.-est@yahoo.com",
			"t+.est@yahoo.com",
		},
		{
			"test 4",
			"t.e-St+@googlemail.com",
			"te-st@gmail.com",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := NormalizeEmail(test.param1)

			if result != test.want {
				t.Errorf("%s got `%s`, wanted `%s`", test.name, result, test.want)
			}
		})
	}
}
