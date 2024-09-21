package validatorgo

import "testing"

func TestIsRFC3339(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid RFC3339 timestamps
		{name: "Valid RFC3339 - UTC", param1: "2024-09-21T14:30:00Z", want: true},
		{name: "Valid RFC3339 - Positive Timezone Offset", param1: "2024-09-21T14:30:00+02:00", want: true},
		{name: "Valid RFC3339 - Negative Timezone Offset", param1: "2024-09-21T14:30:00-05:00", want: true},
		{name: "Valid RFC3339 - UTC Midnight", param1: "2024-01-01T00:00:00Z", want: true},
		{name: "Valid RFC3339 - Maximum Valid Time", param1: "2024-12-31T23:59:59+00:00", want: true},

		// Invalid RFC3339 timestamps
		{name: "Invalid RFC3339 - Missing T separator", param1: "2024-09-21 14:30:00Z", want: false},
		{name: "Invalid RFC3339 - Missing Timezone", param1: "2024-09-21T14:30:00", want: false},
		{name: "Invalid RFC3339 - Invalid Date", param1: "2024-02-30T14:30:00Z", want: false},                      // Invalid day
		{name: "Invalid RFC3339 - Invalid Time", param1: "2024-09-21T25:30:00Z", want: false},                      // Invalid hour
		{name: "Invalid RFC3339 - Invalid Seconds", param1: "2024-09-21T14:30:61Z", want: false},                   // Invalid seconds
		{name: "Invalid RFC3339 - Timezone Out of Range", param1: "2024-09-21T14:30:00+14:00", want: false},        // Timezone beyond valid range
		{name: "Invalid RFC3339 - Missing Timezone Offset Sign", param1: "2024-09-21T14:30:00 02:00", want: false}, // Missing '+' or '-' in offset
		{name: "Invalid RFC3339 - Extra Characters", param1: "2024-09-21T14:30:00Z123", want: false},               // Extra characters at the end

		// Edge cases
		{name: "Invalid RFC3339 - Empty String", param1: "", want: false},
		{name: "Invalid RFC3339 - Only Date", param1: "2024-09-21", want: false},
		{name: "Invalid RFC3339 - Only Time", param1: "14:30:00Z", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsRFC3339(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
