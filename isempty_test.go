package validatorgo

import "testing"

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsEmptyOpts
		want   bool
	}{
		{name: "Is empty", param1: "", param2: &IsEmptyOpts{IgnoreWhitespace: false}, want: true},
		{name: "Is not empty", param1: "viyyv", param2: &IsEmptyOpts{IgnoreWhitespace: false}, want: false},
		{name: "Is empty and contains tabs and spaces with ignore whitespace is false", param1: "   	", param2: &IsEmptyOpts{IgnoreWhitespace: false}, want: false},
		{name: "Is empty and contains tabs and spaces with ignore whitespace is true", param1: "   	", param2: &IsEmptyOpts{IgnoreWhitespace: true}, want: true},
		{name: "Is empty with nil config", param1: "", param2: nil, want: true},
		{name: "Is not empty with nil config", param1: "text", param2: nil, want: false},
		{name: "Is empty with tabs and spaces and nil config", param1: "   	", param2: nil, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsEmpty(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
