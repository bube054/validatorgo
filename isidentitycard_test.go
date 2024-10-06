package validatorgo

import "testing"

func TestIsIdentityCard(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{
		{name: "Valid Sri Lanka", param1: "123456789V", param2: "LK", want: true},
		{name: "Invalid Sri Lanka", param1: "12345678X", param2: "LK", want: false},

		{name: "Valid Poland", param1: "ABC123456", param2: "PL", want: true},
		{name: "Invalid Poland", param1: "AB123456", param2: "PL", want: false},

		{name: "Valid Spain", param1: "12345678Z", param2: "ES", want: true},
		{name: "Invalid Spain", param1: "1234567Z", param2: "ES", want: false},

		{name: "Valid Finland", param1: "131052-308T", param2: "FI", want: true},
		{name: "Invalid Finland", param1: "131052308T", param2: "FI", want: false},

		{name: "Valid India", param1: "123456789012", param2: "IN", want: true},
		{name: "Invalid India", param1: "12345678901", param2: "IN", want: false},

		{name: "Valid Italy", param1: "A12345678", param2: "IT", want: true},
		{name: "Invalid Italy", param1: "A1234567", param2: "IT", want: false},

		{name: "Valid Iran", param1: "0123456789", param2: "IR", want: true},
		{name: "Invalid Iran", param1: "12345678", param2: "IR", want: false},

		{name: "Valid Mozambique", param1: "1234567890123", param2: "MZ", want: true},
		{name: "Invalid Mozambique", param1: "123456789012", param2: "MZ", want: false},

		{name: "Valid Norway", param1: "12345678901", param2: "NO", want: true},
		{name: "Invalid Norway", param1: "1234567890", param2: "NO", want: false},

		{name: "Valid Thailand", param1: "1234567890123", param2: "TH", want: true},
		{name: "Invalid Thailand", param1: "123456789012", param2: "TH", want: false},

		{name: "Valid Taiwan", param1: "A123456789", param2: "zh-TW", want: true},
		{name: "Invalid Taiwan", param1: "B023456789", param2: "zh-TW", want: false},

		{name: "Valid Israel", param1: "123456789", param2: "he-IL", want: true},
		{name: "Invalid Israel", param1: "12345678", param2: "he-IL", want: false},

		{name: "Valid Libya", param1: "123456789012", param2: "ar-LY", want: true},
		{name: "Invalid Libya", param1: "12345678901", param2: "ar-LY", want: false},

		{name: "Valid Tunisia", param1: "12345678", param2: "ar-TN", want: true},
		{name: "Invalid Tunisia", param1: "1234567", param2: "ar-TN", want: false},

		{name: "Valid China", param1: "12345678901234567X", param2: "zh-CN", want: true},
		{name: "Invalid China", param1: "12345678901234567A", param2: "zh-CN", want: false},

		{name: "Valid Hong Kong", param1: "A1234567", param2: "zh-HK", want: true},
		{name: "Invalid Hong Kong", param1: "A123456", param2: "zh-HK", want: false},

		{name: "Valid Pakistan", param1: "1234567890123", param2: "PK", want: true},
		{name: "Invalid Pakistan", param1: "123456789012", param2: "PK", want: false},

		// Tests for locales not in the list
		{name: "Valid with Unrecognized Locale", param1: "1234567890123", param2: "XX", want: false},
		{name: "Valid with any Locale", param1: "1234567890123", param2: "any", want: true},
		{name: "Valid with no locale specified", param1: "1234567890123", param2: "", want: true},
		{name: "Invalid with Unrecognized Locale", param1: "12345678901239009", param2: "YY", want: false},

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIdentityCard(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
