package validatorgo

import "testing"

func TestIsDate(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsDateOpts
		want   bool
	}{
		{
			name:   "Valid ISO8601 date no format and not strict",
			param1: "2006-01-02",
			param2: IsDateOpts{},
			want:   true,
		},
		{
			name:   "Valid USFormat date no format and not strict",
			param1: "01/02/2006",
			param2: IsDateOpts{},
			want:   true,
		},
		{
			name:   "Invalid ISO8601 date no format and not strict",
			param1: "20006-01-02",
			param2: IsDateOpts{},
			want:   false,
		},
		{
			name:   "Invalid USFormat date no format and not strict",
			param1: "01/023/2006",
			param2: IsDateOpts{},
			want:   false,
		},

		{
			name:   "Valid ISO8601 date with USFormat format and not strict",
			param1: "2006-01-02",
			param2: IsDateOpts{Format: USFormat},
			want:   true,
		},
		{
			name:   "Invalid ISO8601 date with USFormat format and not strict",
			param1: "20406-041-042",
			param2: IsDateOpts{Format: USFormat},
			want:   false,
		},
		{
			name:   "Valid ISO8601 date with USFormat format and is strict",
			param1: "2006-01-02",
			param2: IsDateOpts{Format: USFormat, StrictMode: true},
			want:   false,
		},
		{
			name:   "Valid EUFormat date with EUFormat format and is strict",
			param1: "02/01/2006",
			param2: IsDateOpts{Format: USFormat, StrictMode: true},
			want:   true,
		},
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
