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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Matches(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
