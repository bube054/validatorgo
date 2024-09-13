package validatorgo

import "testing"

func TestIsWhitelisted(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		{name: "Empty string provided and empty list provided", param1: "", param2: "", want: true},
		{name: "Empty string provided but list provided", param1: "", param2: "thf", want: true},

		{name: "All characters are present in list", param1: "stop", param2: "post", want: true},
		{name: "No characters are present in list", param1: "bang", param2: "take", want: false},
		{name: "At least one character (c) is not present in list", param1: "cake", param2: "take", want: false},
		{name: "Characters present but empty list", param1: "cake", param2: "", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsWhitelisted(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
