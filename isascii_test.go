package validatorgo

import "testing"

func TestIsAscii(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "Special characters ascii",
			param1: "a - = ? / ~",
			want:   true,
		},
		{
			name:   "Normal letters ascii",
			param1: "normal letters ascii",
			want:   true,
		},
		{
			name:   "Normal letters but containing one single non ascii",
			param1: "lor難em",
			want:   false,
		},
		{
			name:   "Non ascii",
			param1: "說上難車中防水回大石在該是並",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAscii(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
