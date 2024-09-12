package validatorgo

import (
	"testing"
)

func TestIsAfter(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		{
			name:   "2006-01-02 is not after 2007-01-02",
			param1: "2006-01-02",
			param2: "2007-01-02",
			want:   false,
		},
		{
			name:   "2009-06-02 is after 2007-01-02",
			param1: "2009-06-02",
			param2: "2007-01-02",
			want:   true,
		},
		{
			name:   "2020-04-03 is not after right now",
			param1: "2020-04-03",
			param2: "",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAfter(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
