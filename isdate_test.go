package validatorgo

import "testing"

func TestIsDate(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsDateOpts
		want   bool
	}{
		// Valid cases
		{name: "Valid Standard Date", param1: "2007-01-02", param2: nil, want: true},
		{name: "Valid Slash Date", param1: "2007/01/02", param2: &IsDateOpts{Format: SlashDateLayout}, want: true},
		{name: "Valid DateTime", param1: "2007-01-02 15:04:05", param2: &IsDateOpts{Format: DateTimeLayout}, want: true},
		{name: "Valid ISO8601", param1: "2007-01-02T15:04:05", param2: &IsDateOpts{Format: ISO8601Layout}, want: true},
		{name: "Valid ISO8601 Zulu", param1: "2007-01-02T15:04:05Z", param2: &IsDateOpts{Format: ISO8601ZuluLayout}, want: true},
		{name: "Valid ISO8601 with Milliseconds", param1: "2007-01-02T15:04:05.000Z", param2: &IsDateOpts{Format: ISO8601WithMillisecondsLayout}, want: true},
		{name: "Valid Any Format", param1: "2007-01-02", param2: &IsDateOpts{Format: "any"}, want: true},

		// Invalid cases
		{name: "Invalid Date Format", param1: "01-02-2007", param2: nil, want: false},
		{name: "Invalid Slash Date", param1: "2007/01-02", param2: &IsDateOpts{Format: SlashDateLayout}, want: false},
		{name: "Invalid ISO8601", param1: "2007-01-02T15:04", param2: &IsDateOpts{Format: ISO8601Layout}, want: false},
		{name: "Invalid Strict Mode", param1: "2007/01/02", param2: &IsDateOpts{Format: StandardDateLayout, StrictMode: true}, want: false},
		{name: "Invalid Format Option", param1: "2007-01-02", param2: &IsDateOpts{Format: "invalid"}, want: false},
		{name: "Empty String", param1: "", param2: nil, want: false},
		{name: "Whitespace String", param1: "   ", param2: nil, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsDate(test.param1, test.param2)
			assertValidation(t, result, test.want, err)
		})
	}
}
