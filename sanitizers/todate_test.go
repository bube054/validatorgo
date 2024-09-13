package sanitizers

import (
	"testing"
)

func TestToDate(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{
			"test 1",
			"01/02 03:04:05PM '06 -0700",
			"2006-01-02 15:04:05 -0700 -0700",
		},
		{
			"test 2",
			"07/21 02:30:00PM '24 -0700",
			"2006-01-02 15:04:05 -0700 -0700",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			date := ToDate(test.param1)

			if date == nil && test.want != "" {
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

// Layout: "01/02 03:04:05PM '06 -0700"

// Example: "07/21 02:30:00PM '24 -0700"
// ANSIC: "Mon Jan _2 15:04:05 2006"

// Example: "Sun Jul 21 14:30:00 2024"
// UnixDate: "Mon Jan _2 15:04:05 MST 2006"

// Example: "Sun Jul 21 14:30:00 PST 2024"
// RubyDate: "Mon Jan 02 15:04:05 -0700 2006"

// Example: "Sun Jul 21 14:30:00 -0700 2024"
// RFC822: "02 Jan 06 15:04 MST"

// Example: "21 Jul 24 14:30 PST"
// RFC822Z: "02 Jan 06 15:04 -0700"

// Example: "21 Jul 24 14:30 -0700"
// RFC850: "Monday, 02-Jan-06 15:04:05 MST"

// Example: "Sunday, 21-Jul-24 14:30:00 PST"
// RFC1123: "Mon, 02 Jan 2006 15:04:05 MST"

// Example: "Sun, 21 Jul 2024 14:30:00 PST"
// RFC1123Z: "Mon, 02 Jan 2006 15:04:05 -0700"

// Example: "Sun, 21 Jul 2024 14:30:00 -0700"
// RFC3339: "2006-01-02T15:04:05Z07:00"

// Example: "2024-07-21T14:30:00Z07:00"
// RFC3339Nano: "2006-01-02T15:04:05.999999999Z07:00"

// Example: "2024-07-21T14:30:00.999999999Z07:00"
// Kitchen: "3:04PM"

// Example: "2:30PM"
// Stamp: "Jan _2 15:04:05"

// Example: "Jul 21 14:30:00"
// StampMilli: "Jan _2 15:04:05.000"

// Example: "Jul 21 14:30:00.000"
// StampMicro: "Jan _2 15:04:05.000000"

// Example: "Jul 21 14:30:00.000000"
// StampNano: "Jan _2 15:04:05.000000000"

// Example: "Jul 21 14:30:00.000000000"
// DateTime: "2006-01-02 15:04:05"

// Example: "2024-07-21 14:30:00"
// DateOnly: "2006-01-02"

// Example: "2024-07-21"
// TimeOnly: "15:04:05"

// Example: "14:30:00"
