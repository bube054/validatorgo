package validatorgo

import "testing"

func TestIsIn(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 []string
		want   bool
	}{
		{name: "Value is present", param1: "apple", param2: []string{"apple", "banana", "grape"}, want: true},
		{name: "Value is not present", param1: "orange", param2: []string{"apple", "banana", "grape"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIn(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
