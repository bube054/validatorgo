package sanitizer

import (
	"testing"
)

func TestBlacklist(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   string
	}{
		{name: "Basic case with one blacklisted character", param1: "Hello World!", param2: "!@#$%^&*()", want: "Hello World"},
		{name: "Multiple blacklisted characters", param1: "abc123!@#", param2: "!@#$%^&*()", want: "abc123"},
		{name: "No blacklisted characters present", param1: "CleanInput123", param2: "!@#$%^&*()", want: "CleanInput123"},
		{name: "All characters blacklisted", param1: "!@#$%^&*()", param2: "!@#$%^&*()", want: ""},
		{name: "Empty input string", param1: "", param2: "!@#$%^&*()", want: ""},
		{name: "Whitespace should be preserved", param1: " Hello  World! ", param2: "!@#$%^&*()", want: " Hello  World "},
		{name: "Long string with various characters", param1: "Testing with special chars #$%^&*() and numbers 123456", param2: "#$%^&*()", want: "Testing with special chars  and numbers 123456"},
		{name: "Input with unicode characters", param1: "Hello 😊 World!", param2: "!@#$%^&*()", want: "Hello 😊 World"},
		{name: "Blacklisted string is a substring", param1: "Hello!", param2: "Hello", want: "!"},
		{name: "Multiple instances of blacklisted characters", param1: "Goodbye!!!", param2: "!", want: "Goodbye"},
		{name: "Blacklisted characters mixed with allowed ones", param1: "Hello @ World!", param2: "@!", want: "Hello  World"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Blacklist(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}
