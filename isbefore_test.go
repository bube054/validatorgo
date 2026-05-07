package validatorgo

import (
	"testing"
)

func TestIsBefore(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsBeforeOpts
		want   bool
	}{
		// Valid cases with StandardDateLayout
		{name: "StandardDateLayout - param1 before", param1: "2023-01-01", param2: &IsBeforeOpts{ComparisonDate: "2023-09-15"}, want: true},
		{name: "StandardDateLayout - param1 not before", param1: "2023-09-15", param2: &IsBeforeOpts{ComparisonDate: "2023-01-01"}, want: false},

		// Valid cases with SlashDateLayout
		{name: "SlashDateLayout - param1 before", param1: "2023/01/01", param2: &IsBeforeOpts{ComparisonDate: "2023/09/15"}, want: true},
		{name: "SlashDateLayout - param1 not before", param1: "2023/09/15", param2: &IsBeforeOpts{ComparisonDate: "2023/01/01"}, want: false},

		// ISO8601 Layouts
		{name: "ISO8601Layout - param1 before", param1: "2023-09-15T11:59:59", param2: &IsBeforeOpts{ComparisonDate: "2023-09-15T12:00:00"}, want: true},
		{name: "ISO8601ZuluLayout - param1 not before", param1: "2023-09-15T12:00:01Z", param2: &IsBeforeOpts{ComparisonDate: "2023-09-15T12:00:00Z"}, want: false},
		{name: "ISO8601WithMillisecondsLayout - param1 before", param1: "2023-09-15T12:00:00.000Z", param2: &IsBeforeOpts{ComparisonDate: "2023-09-15T12:00:00.001Z"}, want: true},

		// Invalid formats
		{name: "Invalid param1 format", param1: "15-09-2023", param2: &IsBeforeOpts{ComparisonDate: "2023-01-01"}, want: false},
		{name: "Invalid ComparisonDate format", param1: "2023-09-15", param2: &IsBeforeOpts{ComparisonDate: "01/01/2023"}, want: false},

		// Default to current time
		{name: "Empty ComparisonDate - param1 before current time", param1: "2000-01-01", param2: &IsBeforeOpts{ComparisonDate: ""}, want: true},
		{name: "Empty ComparisonDate - param1 after current time", param1: "3000-01-01", param2: &IsBeforeOpts{ComparisonDate: ""}, want: false},

		// Edge cases
		{name: "Leap year comparison", param1: "2019-12-31", param2: &IsBeforeOpts{ComparisonDate: "2020-02-29"}, want: true},
		{name: "Non-leap year comparison", param1: "2021-03-01", param2: &IsBeforeOpts{ComparisonDate: "2021-02-28"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := IsBefore(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
