package validatorgo

import "testing"

func TestIsHSL(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid HSL Colors
		{name: "Valid HSL with integer values", param1: "hsl(360, 100%, 50%)", want: true},
		{name: "Valid HSL with decimal values", param1: "hsl(180.5, 50%, 75%)", want: true},
		{name: "Valid HSLA with alpha value", param1: "hsla(120, 100%, 50%, 0.5)", want: true},
		{name: "Valid HSL with degree unit", param1: "hsl(120deg, 50%, 50%)", want: true},
		{name: "Valid HSL with radian unit", param1: "hsl(3.1415rad, 50%, 50%)", want: true},
		{name: "Valid HSL with turn unit", param1: "hsl(1turn, 50%, 50%)", want: true},
		{name: "Valid HSL with space", param1: "hsl(240deg 100% 50%)", want: true},
		{name: "Valid HSLA with percentage alpha", param1: "hsla(180, 100%, 50%, 50%)", want: true},
		{name: "Valid HSLA with percentage alpha and space", param1: "hsla(180 100% 50% / 0.5)", want: true},

		// // Invalid HSL Colors
		{name: "Invalid HSL missing values", param1: "hsl(360, 100%)", want: false},
		{name: "Invalid HSL with incorrect percent", param1: "hsl(360, 100, 50%)", want: false},
		{name: "Invalid HSL with out of range hue", param1: "hsl(361, 100%, 50%)", want: false},
		{name: "Invalid HSL with incorrect alpha value", param1: "hsla(120, 100%, 50%, 1.5)", want: false},
		{name: "Invalid HSL with invalid unit", param1: "hsl(120xyz, 50%, 50%)", want: false},
		{name: "Invalid HSL with incorrect syntax", param1: "hsl(120, 50%, 50%", want: false},
		{name: "Invalid HSLA with misplaced comma", param1: "hsla(120, 100%, 50% 0.5)", want: false},
		{name: "Invalid HSLA with too many values", param1: "hsla(120, 100%, 50%, 0.5, 10%)", want: false},
		{name: "Invalid HSL with wrong degrees value", param1: "hsl(360deg, 100%, 50%)", want: false},
		{name: "Invalid HSL with missing unit", param1: "hsl(120, 100%, 50)", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsHSL(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
