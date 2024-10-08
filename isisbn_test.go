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
		{name: "Is invalid ISBN v10, wrong length", param1: "0-7168-0344-09", param2: "10", want: false},
		{name: "Is invalid ISBN v10, alphabets included", param1: "0-7rt8-0344-09", param2: "10", want: false},
		// v13
		{name: "Is valid ISBN v13", param1: "978-0-7167-0344-0", param2: "13", want: true},
		{name: "Is invalid ISBN v13", param1: "978-9-7167-0344-0", param2: "13", want: false},
		{name: "Is invalid ISBN v13, wrong length", param1: "978-0-7167-0344-07", param2: "13", want: false},
		{name: "Is invalid ISBN v13, alphabets included", param1: "abc-0-7167-0344-07", param2: "13", want: false},
		// not v10 or v13
		{name: "Version is not provided", param1: "978-0-7167-0344-0", param2: "", want: true},
		{name: "Version is also not provided", param1: "0-7167-0344-0", param2: "", want: true},
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
