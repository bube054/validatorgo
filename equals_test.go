package validatorgo

import "testing"

func TestEquals(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		{
			name:   "1",
			param1: "Hello",
			param2: "Hello",
			want:   true,
		},
		{
			name:   "2",
			param1: "Hello",
			param2: "Hello!",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Equals(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
