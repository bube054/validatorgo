package validatorgo

import "testing"

func TestIsISO8601(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsISO8601Opts
		want   bool
	}{
		// Valid with default options
		{name: "Valid date", param1: "2023-09-05", want: true},
		{name: "Valid date with slashes", param1: "2023/09/05", want: true},
		{name: "Valid date with dots for both date and time", param1: "2023.09.05T14.30.00", want: true},
		{name: "Valid date and time with spaces", param1: "2023 09 05 14 30 00", want: true},
		{name: "September 31st is valid with default Strict: false", param1: "2023-09-31T14:30:00", want: true},
		{name: "Valid UTC time", param1: "2023-09-05T14:30:00Z", want: true},
		{name: "Valid with time zone offset", param1: "2023-09-05T14:30:00+02:00", want: true},
		// Invalid with default options
		{name: "Invalid, month 13 does not exist", param1: "2023-13-05T14:30:00", want: false},
		{name: "Invalid, minutes can't be 60", param1: "2023-09-05T14:60:00", want: false},
		{name: "Invalid, hour 25 is not valid", param1: "2023-09-05T25:00:00", want: false},

		// Valid with Strict true
		{name: "Valid date", param1: "2023-09-05", param2: &IsISO8601Opts{Strict: true}, want: true},
		{name: "Valid leap year date", param1: "1888-02-29", param2: &IsISO8601Opts{Strict: true}, want: true},
		{name: "Valid date and UTC time", param1: "2023-09-05T14:30:00Z", param2: &IsISO8601Opts{Strict: true}, want: true},
		{name: "Valid date and time with time zone offset", param1: "2023-09-05T14:30:00+02:00", param2: &IsISO8601Opts{Strict: true}, want: true},
		// Invalid with Strict true
		{name: "Invalid, February 30th does not exist", param1: "2023-02-30", param2: &IsISO8601Opts{Strict: true}, want: false},
		{name: "Invalid, September has only 30 days", param1: "2023-09-31", param2: &IsISO8601Opts{Strict: true}, want: false},
		{name: "Invalid, month 13 does not exist", param1: "2023-13-05", param2: &IsISO8601Opts{Strict: true}, want: false},
		{name: "Invalid if it's not a leap year", param1: "2023-02-29", param2: &IsISO8601Opts{Strict: true}, want: false},

		// Valid with StrictSeparator true
		{name: "Valid with strict separators", param1: "2023-09-05T14:30:00", param2: &IsISO8601Opts{StrictSeparator: true}, want: true},
		{name: "Valid UTC time with strict separators", param1: "2023-09-05T14:30:00Z", param2: &IsISO8601Opts{StrictSeparator: true}, want: true},
		{name: "Valid date and time with strict time zone offset", param1: "2023-09-05T14:30:00+02:00", param2: &IsISO8601Opts{StrictSeparator: true}, want: true},
		// Invalid with StrictSeparator true
		{name: "Invalid, slashes / used in date", param1: "2023/09/05T14:30:00", param2: &IsISO8601Opts{StrictSeparator: true}, want: false},
		{name: "Invalid, dots . used in date", param1: "2023.09.05T14:30:00", param2: &IsISO8601Opts{StrictSeparator: true}, want: false},
		{name: "Invalid, dots . used in time", param1: "2023-09-05T14.30.00", param2: &IsISO8601Opts{StrictSeparator: true}, want: false},
		{name: "Invalid, spaces used in time", param1: "2023-09-05T14 30 00", param2: &IsISO8601Opts{StrictSeparator: true}, want: false},

		// Valid with Strict true and StrictSeparator true
		{name: "Valid date and time with strict separators", param1: "2023-09-05T14:30:00", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: true},
		{name: "Valid leap year date and UTC time with strict separators", param1: "2020-02-29T14:30:00Z", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: true},
		{name: "Valid date and time with strict time zone offset", param1: "2023-09-05T14:30:00+02:00", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: true},
		// Invalid with Strict true and StrictSeparator true
		{name: "Invalid, strict date validation disallows February 30th", param1: "2023-02-30", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: false},
		{name: "Invalid, September has only 30 days", param1: "2023-09-31", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: false},
		{name: "Invalid, slashes / used in date", param1: "2023/09/05T14:30:00", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: false},
		{name: "Invalid, dots used in date", param1: "2023.09.05T14:30:00", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: false},
		{name: "Invalid, spaces used in time", param1: "2023-09-05T14 30 00", param2: &IsISO8601Opts{StrictSeparator: true, Strict: true}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISO8601(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
