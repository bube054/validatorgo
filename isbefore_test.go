package validatorgo

import (
	"testing"
)

func TestIsBefore(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsBeforeOpts
		want   bool
	}{
		// Valid cases
		{name: "Date before another date (Date only)", param1: "2023-01-01", param2: IsBeforeOpts{ComparisonDate: "2023-09-15"}, want: true},
		{name: "Same month, earlier day", param1: "2023-06-15", param2: IsBeforeOpts{ComparisonDate: "2023-12-31"}, want: true},
		{name: "Same day, earlier time", param1: "2024-01-03 15:04:05", param2: IsBeforeOpts{ComparisonDate: "2024-01-03 18:04:05"}, want: true},
		{name: "Earlier date in RFC1123 format", param1: "Sat, 01 Jan 2023 00:00:00 GMT", param2: IsBeforeOpts{ComparisonDate: "Mon, 01 Jan 2024 00:00:00 GMT"}, want: true},

		// Invalid cases
		{name: "Date is the same as ComparisonDate", param1: "2024-01-01", param2: IsBeforeOpts{ComparisonDate: "2024-01-01"}, want: false},
		{name: "Date after ComparisonDate", param1: "2024-01-01", param2: IsBeforeOpts{ComparisonDate: "2023-01-01"}, want: false},
		{name: "Same day, later time", param1: "2023-09-10T15:00:00Z", param2: IsBeforeOpts{ComparisonDate: "2023-09-10T14:00:00Z"}, want: false},
		{name: "Date in different format", param1: "01-01-2023", param2: IsBeforeOpts{ComparisonDate: "2023-01-01"}, want: false},

		// Edge cases
		{name: "Empty comparison date", param1: "2050-01-01", param2: IsBeforeOpts{ComparisonDate: ""}, want: false},
		{name: "Invalid date format", param1: "InvalidDate", param2: IsBeforeOpts{ComparisonDate: "2023-01-01"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBefore(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
