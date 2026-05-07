package validatorgo

import (
	"regexp"
	"testing"
)

func TestMatches(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *regexp.Regexp
		want   bool
	}{
		{name: "Matches regex", param1: "foo", param2: regexp.MustCompile(`^foo$`), want: true},
		{name: "Does not match regex", param1: "foo", param2: regexp.MustCompile(`^foobar$`), want: false},
		{name: "nil regex", param1: "foo", param2: nil, want: false},
		// /abc/ pattern tests
		{name: "abc exact match", param1: "abc", param2: regexp.MustCompile(`abc`), want: true},
		{name: "abc in abcdef", param1: "abcdef", param2: regexp.MustCompile(`abc`), want: true},
		{name: "abc in 123abc", param1: "123abc", param2: regexp.MustCompile(`abc`), want: true},
		{name: "abc no match acb", param1: "acb", param2: regexp.MustCompile(`abc`), want: false},
		{name: "abc no match Abc", param1: "Abc", param2: regexp.MustCompile(`abc`), want: false},
		// /abc/i case insensitive tests
		{name: "abc case insensitive abc", param1: "abc", param2: regexp.MustCompile(`(?i)abc`), want: true},
		{name: "abc case insensitive abcdef", param1: "abcdef", param2: regexp.MustCompile(`(?i)abc`), want: true},
		{name: "abc case insensitive 123abc", param1: "123abc", param2: regexp.MustCompile(`(?i)abc`), want: true},
		{name: "abc case insensitive AbC", param1: "AbC", param2: regexp.MustCompile(`(?i)abc`), want: true},
		{name: "abc case insensitive no match acb", param1: "acb", param2: regexp.MustCompile(`(?i)abc`), want: false},
		// Edge cases
		{name: "Empty string no match", param1: "", param2: regexp.MustCompile(`abc`), want: false},
		{name: "Empty pattern matches", param1: "anything", param2: regexp.MustCompile(``), want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Matches(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
