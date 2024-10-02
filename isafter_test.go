package validatorgo

import (
	"testing"
)

func TestIsAfter(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsAfterOpts
		want   bool
	}{
		// Valid comparisons
		{name: "Valid date is after", param1: "2023-09-15", param2: &IsAfterOpts{ComparisonDate: "2023-01-01"}, want: true},
		{name: "Valid date is not after", param1: "2023-01-01", param2: &IsAfterOpts{ComparisonDate: "2023-09-15"}, want: false},
		{name: "Same date", param1: "2023-09-15", param2: &IsAfterOpts{ComparisonDate: "2023-09-15"}, want: false},
		{name: "Future date comparison", param1: "2100-01-01", param2: &IsAfterOpts{ComparisonDate: "2050-01-01"}, want: true},

		// Valid with default (comparison against current date)
		{name: "Valid with default", param1: "2050-11-15", param2: &IsAfterOpts{}, want: true},

		// Invalid date formats
		{name: "Invalid date format", param1: "invalid", param2: &IsAfterOpts{ComparisonDate: "2023-09-15"}, want: false},
		{name: "Empty date string", param1: "", param2: &IsAfterOpts{ComparisonDate: "2023-09-15"}, want: false},
		{name: "Past date comparison", param1: "2000-01-01", param2: &IsAfterOpts{ComparisonDate: "2050-01-01"}, want: false},

		// Nil config cases
		{name: "Nil param2, valid date after", param1: "2050-11-15", param2: nil, want: true},
		{name: "Nil param2, valid date not after", param1: "2020-11-15", param2: nil, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAfter(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
