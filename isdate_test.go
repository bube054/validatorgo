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
		{name: "Valid ISO8601 date with default config", param1: "2007-01-02", param2: nil, want: true},
		{name: "Valid ISO8601 date no format and not strict", param1: "2006-01-02", param2: &IsDateOpts{}, want: true},
		{name: "Valid USFormat date no format and not strict", param1: "01/02/2006", param2: &IsDateOpts{}, want: true},
		{name: "Valid date multiple formats and not strict", param1: "2006-01-02", param2: &IsDateOpts{Format: ISO8601}, want: true},
		{name: "Valid date leap year strict", param1: "29/02/2020", param2: &IsDateOpts{Format: EUFormat, StrictMode: true}, want: true},

		// Invalid cases
		{name: "Invalid ISO8601 date no format and not strict", param1: "20006-01-02", param2: &IsDateOpts{}, want: false},
		{name: "Invalid USFormat date no format and not strict", param1: "01/023/2006", param2: &IsDateOpts{}, want: false},

		// With format and strict mode
		{name: "Valid ISO8601 date with USFormat format and not strict", param1: "2006-01-02", param2: &IsDateOpts{Format: USFormat}, want: true},
		{name: "Invalid ISO8601 date with USFormat format and not strict", param1: "20406-041-042", param2: &IsDateOpts{Format: USFormat}, want: false},
		{name: "Valid ISO8601 date with USFormat format and is strict", param1: "2006-01-02", param2: &IsDateOpts{Format: USFormat, StrictMode: true}, want: false},
		{name: "Valid EUFormat date with EUFormat format and is strict", param1: "02/01/2006", param2: &IsDateOpts{Format: EUFormat, StrictMode: true}, want: true},
		{name: "Invalid date multiple formats and strict", param1: "2006-01-02T15:04:05", param2: &IsDateOpts{Format: USFormat, StrictMode: true}, want: false},
		{name: "Invalid date with wrong format", param1: "2006-01-02T15:04:05", param2: &IsDateOpts{Format: "GSFormat", StrictMode: true}, want: false},

		// {
		// 	name:   "Valid ISO8601 date no format and not strict",
		// 	param1: "2006-01-02",
		// 	param2: IsDateOpts{},
		// 	want:   true,
		// },
		// {
		// 	name:   "Valid USFormat date no format and not strict",
		// 	param1: "01/02/2006",
		// 	param2: IsDateOpts{},
		// 	want:   true,
		// },
		// {
		// 	name:   "Invalid ISO8601 date no format and not strict",
		// 	param1: "20006-01-02",
		// 	param2: IsDateOpts{},
		// 	want:   false,
		// },
		// {
		// 	name:   "Invalid USFormat date no format and not strict",
		// 	param1: "01/023/2006",
		// 	param2: IsDateOpts{},
		// 	want:   false,
		// },

		// {
		// 	name:   "Valid ISO8601 date with USFormat format and not strict",
		// 	param1: "2006-01-02",
		// 	param2: IsDateOpts{Format: USFormat},
		// 	want:   true,
		// },
		// {
		// 	name:   "Invalid ISO8601 date with USFormat format and not strict",
		// 	param1: "20406-041-042",
		// 	param2: IsDateOpts{Format: USFormat},
		// 	want:   false,
		// },
		// {
		// 	name:   "Valid ISO8601 date with USFormat format and is strict",
		// 	param1: "2006-01-02",
		// 	param2: IsDateOpts{Format: USFormat, StrictMode: true},
		// 	want:   false,
		// },
		// {
		// 	name:   "Valid EUFormat date with EUFormat format and is strict",
		// 	param1: "02/01/2006",
		// 	param2: IsDateOpts{Format: USFormat, StrictMode: true},
		// 	want:   true,
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsDate(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
