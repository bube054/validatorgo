package sanitizer

import (
	"testing"
)

func TestToDate(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{name: "Valid ANSIC format", param1: "Mon Jan  2 15:04:05 2006", want: "2006-01-02 15:04:05 +0000 UTC"},
		{name: "Valid UnixDate format", param1: "Mon Jan  2 15:04:05 MST 2006", want: "2006-01-02 15:04:05 +0000 MST"},
		{name: "Valid RubyDate format", param1: "Mon Jan 02 15:04:05 -0700 2006", want: "2006-01-02 15:04:05 -0700 -0700"},
		{name: "Valid RFC822 format", param1: "02 Jan 06 15:04 MST", want: "2006-01-02 15:04:00 +0000 MST"},
		{name: "Valid RFC822Z format", param1: "02 Jan 06 15:04 -0700", want: "2006-01-02 15:04:00 -0700 -0700"},
		{name: "Valid RFC850 format", param1: "Monday, 02-Jan-06 15:04:05 MST", want: "2006-01-02 15:04:05 +0000 MST"},
		{name: "Valid RFC1123 format", param1: "Mon, 02 Jan 2006 15:04:05 MST", want: "2006-01-02 15:04:05 +0000 MST"},
		{name: "Valid RFC1123Z format", param1: "Mon, 02 Jan 2006 15:04:05 -0700", want: "2006-01-02 15:04:05 -0700 -0700"},
		{name: "Valid Kitchen format", param1: "3:04PM", want: "0000-01-01 15:04:00 +0000 UTC"},
		{name: "Valid Stamp format", param1: "Jan 2 15:04:05", want: "0000-01-02 15:04:05 +0000 UTC"},
		{name: "Valid StampMilli format", param1: "Jan 2 15:04:05.000", want: "0000-01-02 15:04:05 +0000 UTC"},
		{name: "Valid StampMicro format", param1: "Jan 2 15:04:05.000000", want: "0000-01-02 15:04:05 +0000 UTC"},
		{name: "Valid StampNano format", param1: "Jan 2 15:04:05.000000000", want: "0000-01-02 15:04:05 +0000 UTC"},
		{name: "Valid DateTime format", param1: "2006-01-02 15:04:05", want: "2006-01-02 15:04:05 +0000 UTC"},
		{name: "Valid DateOnly format", param1: "2006-01-02", want: "2006-01-02 00:00:00 +0000 UTC"},
		{name: "Valid TimeOnly format", param1: "15:04:05", want: "0000-01-01 15:04:05 +0000 UTC"},
		{name: "Invalid format", param1: "invalid-date-string", want: "<nil>"},
		{name: "Empty string", param1: "", want: "<nil>"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			date := ToDate(test.param1)

			if date == nil && test.want != "" {
				if test.want == "<nil>" {
					return
				}

				t.Errorf("got `%v`, wanted `%s`", nil, test.want)
				return
			}

			dateAsStr := date.String()

			if dateAsStr != test.want {
				t.Errorf("got `%s`, wanted `%s`", dateAsStr, test.want)
			}
		})
	}
}