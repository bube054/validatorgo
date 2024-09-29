package sanitizer

import (
	"testing"
)

func TestUnescape(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{name: "Unescape HTML entity", param1: "&amp;", want: "&"},
		{name: "Unescape multiple HTML entities", param1: "&lt;&gt;&quot;&apos;", want: "<>\"'"},
		{name: "Unescape numeric entity", param1: "&#39;", want: "'"},
		{name: "Unescape mixed entities and text", param1: "Hello &amp; World!", want: "Hello & World!"},
		{name: "Unescape without entities", param1: "Hello World!", want: "Hello World!"},
		{name: "Unescape string with multiple spaces", param1: "Hello &nbsp;&nbsp;World", want: "Hello   World"},
		{name: "Unescape with no entities", param1: "Plain text", want: "Plain text"},
		{name: "Unescape hexadecimal entity", param1: "&#x27;", want: "'"},
		{name: "Unescape string with special character entities", param1: "&copy; 2024", want: "© 2024"},
		{name: "Unescape empty string", param1: "", want: ""},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Unescape(test.param1)

			if result != test.want {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}
