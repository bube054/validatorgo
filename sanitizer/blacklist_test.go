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
		{name: "Blacklist digits", param1: "Hello123 World!", param2: "0-9", want: "Hello World!"},
		{name: "Blacklist letters", param1: "Hello123 World!", param2: "a-zA-Z", want: "123 !"},
		{name: "Blacklist spaces", param1: "Hello123 World!", param2: " ", want: "Hello123World!"},
		{name: "Blacklist alphanumeric characters", param1: "Hello123 World!", param2: "a-zA-Z0-9", want: " !"},
		{name: "Blacklist specific characters (symbols)", param1: "Hello123@World!", param2: "@!", want: "Hello123World"},
		{name: "Blacklist letters with accents", param1: "Olá123 Mundo!", param2: "áéíóúÁÉÍÓÚ", want: "Ol123 Mundo!"},
		{name: "Blacklist empty string", param1: "Hello World!", param2: "", want: "Hello World!"},
		{name: "Blacklist with no matching characters", param1: "Hello World!", param2: "123", want: "Hello World!"},
		{name: "Blacklist all characters", param1: "Hello World!", param2: ".", want: ""},
		{name: "Blacklist multiple types (letters, digits, spaces)", param1: "Hello123 World!", param2: "a-zA-Z0-9 ", want: "!"},
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
