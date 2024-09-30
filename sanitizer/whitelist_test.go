package sanitizer

import (
	"testing"
)

func TestWhitelist(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   string
	}{
		{name: "Whitelist only letters", param1: "Hello123 World!", param2: "a-zA-Z", want: "HelloWorld"},
		{name: "Whitelist only digits", param1: "Hello123 World!", param2: "0-9", want: "123"},
		{name: "Whitelist letters and spaces", param1: "Hello123 World!", param2: "a-zA-Z ", want: "Hello World"},
		{name: "Whitelist alphanumeric characters", param1: "Hello123 World!", param2: "a-zA-Z0-9", want: "Hello123World"},
		{name: "Whitelist specific characters (symbols)", param1: "Hello123@World!", param2: "@!", want: "@!"},
		{name: "Whitelist spaces only", param1: "Hello123 World!", param2: " ", want: " "},
		{name: "Whitelist letters with accents", param1: "Olá123 Mundo!", param2: "a-zA-ZáéíóúÁÉÍÓÚ", want: "OláMundo"},
		{name: "Whitelist empty string", param1: "", param2: "a-zA-Z", want: ""},
		{name: "Whitelist with no matching characters", param1: "123456789", param2: "a-zA-Z", want: ""},
		{name: "Whitelist all characters", param1: "Hello World!", param2: `.`, want: "Hello World!"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Whitelist(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}
