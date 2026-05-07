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

		// JS default (no ignore_whitespace)
		{name: "JS default valid empty", param1: "", param2: nil, want: true},
		{name: "JS default invalid space", param1: " ", param2: nil, want: false},
		{name: "JS default invalid foo", param1: "foo", param2: nil, want: false},
		{name: "JS default invalid 3", param1: "3", param2: nil, want: false},
		// JS with ignore_whitespace=true
		{name: "JS ignore_ws valid empty", param1: "", param2: &IsEmptyOpts{IgnoreWhitespace: true}, want: true},
		{name: "JS ignore_ws valid space", param1: " ", param2: &IsEmptyOpts{IgnoreWhitespace: true}, want: true},
		{name: "JS ignore_ws invalid foo", param1: "foo", param2: &IsEmptyOpts{IgnoreWhitespace: true}, want: false},
		{name: "JS ignore_ws invalid 3", param1: "3", param2: &IsEmptyOpts{IgnoreWhitespace: true}, want: false},

		// Extra edges
		{name: "Tab character not ignored", param1: "\t", param2: nil, want: false},
		{name: "Tab ignored with whitespace", param1: "\t", param2: &IsEmptyOpts{IgnoreWhitespace: true}, want: true},
		{name: "Newline not ignored", param1: "\n", param2: nil, want: false},
		{name: "Multiple spaces ignored", param1: "   ", param2: &IsEmptyOpts{IgnoreWhitespace: true}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsEmpty(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
