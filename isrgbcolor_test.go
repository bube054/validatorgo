package validatorgo

import "testing"

func TestIsRgbColor(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsRgbOpts
		want   bool
	}{
		// Valid RGB colors without spaces and without percent values
		{name: "Valid RGB - No Spaces, No Percent", param1: "rgb(255,0,0)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: true},
		{name: "Valid RGB - No Spaces, No Percent", param1: "rgb(0,128,255)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: true},

		// Valid RGB with spaces allowed
		{name: "Valid RGB - With Spaces Allowed", param1: "rgb(255, 0, 0)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: true}, want: true},
		{name: "Valid RGB - Mixed Spaces Allowed", param1: "rgb(255, 0, 128)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: true}, want: true},

		// Valid RGB with percentage values
		{name: "Valid RGB - Percent Values Allowed", param1: "rgb(100%,0%,0%)", param2: IsRgbOpts{IncludePercentValues: true, AllowSpaces: false}, want: true},
		{name: "Valid RGB - Percent and Spaces Allowed", param1: "rgb(100%, 0%, 50%)", param2: IsRgbOpts{IncludePercentValues: true, AllowSpaces: true}, want: true},

		// Invalid cases: Disallowing spaces or percent values
		{name: "Invalid RGB - Percent Not Allowed", param1: "rgb(100%, 0%, 0%)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Spaces Not Allowed", param1: "rgb(255, 0, 0)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},

		// Invalid RGB: Out of range values
		{name: "Invalid RGB - Values Out of Range", param1: "rgb(256, 0, 0)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Negative Value", param1: "rgb(-1, 255, 255)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},

		// Invalid RGB: Incorrect format
		{name: "Invalid RGB - Missing Comma", param1: "rgb(255 0 0)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Extra Characters", param1: "rgb(255,0,0,50)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Invalid Percent Format", param1: "rgb(100%, 50%, 101%)", param2: IsRgbOpts{IncludePercentValues: true, AllowSpaces: false}, want: false},

		// Edge cases
		{name: "Invalid RGB - Empty String", param1: "", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Only Letters", param1: "rgb(red,green,blue)", param2: IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsRgbColor(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
