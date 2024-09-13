package validatorgo

import "testing"

func TestIsLocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Locales
		{name: "Is valid locale", param1: "ca_ES", want: true},
		{name: "Is valid locale", param1: "it_IT", want: true},
		{name: "Is valid locale", param1: "uk_UA", want: true},
		{name: "Is valid locale", param1: "zu_ZA", want: true},
		// Invalid Locales
		{name: "Is invalid locale", param1: "en_XY", want: false},
		{name: "Is invalid locale", param1: "fr_QQ", want: false},
		{name: "Is invalid locale", param1: "fs_ZZ", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsLocale(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
