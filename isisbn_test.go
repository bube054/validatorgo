package validatorgo

import "testing"

func TestIsValidISBN(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		// v10
		{name: "Is valid ISBN v10", param1: "0-7167-0344-0", param2: "10", want: true},
		{name: "Is invalid ISBN v10", param1: "0-7168-0344-0", param2: "10", want: false},
		// v13
		{name: "Is valid ISBN v13", param1: "978-0-7167-0344-0", param2: "13", want: true},
		{name: "Is invalid ISBN v13", param1: "978-9-7167-0344-0", param2: "13", want: false},
		// not v10 or v13
		{name: "Version is 6", param1: "978-0-7167-0344-0", param2: "6", want: true},
		{name: "Version is any", param1: "0-7167-0344-0", param2: "any", want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISBN(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
