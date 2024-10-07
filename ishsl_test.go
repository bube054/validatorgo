package validatorgo

import "testing"

func TestIsHSL(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid HSL Colors
		{name: "Valid hsl color all values in range", param1: "hsl(170, 50%, 50%)", want: true},
		{name: "Valid hsl color hue not preceding range", param1: "hsl(0, 50%, 50%)", want: true},
		{name: "Valid hsl color saturation not preceding range", param1: "hsl(0, 0%, 50%)", want: true},
		{name: "Valid hsl color lightness not preceding range", param1: "hsl(10, 50%, 0%)", want: true},
		{name: "Valid hsl color hue not exceeding range", param1: "hsl(360, 50%, 50%)", want: true},
		{name: "Valid hsl color saturation not exceeding range", param1: "hsl(0, 100%, 50%)", want: true},
		{name: "Valid hsl color lightness not exceeding range", param1: "hsl(10, 50%, 100%)", want: true},
		{name: "Valid hsl color no spaces", param1: "hsl(0,50%,70%)", want: true},

		// Invalid HSL Colors
		{name: "Invalid hsl color all values in range", param1: "hsl(170%, 50, 50)", want: false},
		{name: "Invalid hsl color hue is not a digit", param1: "hsl(abc, 50%, 50%)", want: false},
		{name: "Invalid hsl color hue preceding range", param1: "hsl(-1, 50%, 50%)", want: false},
		{name: "Invalid hsl color saturation preceding range", param1: "hsl(0, -0%, 50%)", want: false},
		{name: "Invalid hsl color lightness preceding range", param1: "hsl(10, 50%, -0%)", want: false},
		{name: "Invalid hsl color hue exceeding range", param1: "hsl(361, 50%, 50%)", want: false},
		{name: "Invalid hsl color saturation exceeding range", param1: "hsl(0, 101%, 50%)", want: false},
		{name: "Invalid hsl color lightness exceeding range", param1: "hsl(10, 50%, 101%)", want: false},

		// Valid HSLA Colors
		{name: "Valid hsla color alpha at low", param1: "hsla(170, 50%, 50%, 0)", want: true},
		{name: "Valid hsla color alpha at high", param1: "hsla(170, 50%, 50%, 1)", want: true},
		{name: "Valid hsla color alpha at mid", param1: "hsla(170, 50%, 50%, 0.5)", want: true},
		{name: "Valid hsla color no spaces", param1: "hsla(170,50%,50%,0.5)", want: true},

		// Invalid HSLA Colors
		{name: "Invalid hsla alpha present but hsla was not specified", param1: "hsl(170, 50%, 50%, 0)", want: false},
		{name: "Invalid hsla alpha value exceeded", param1: "hsla(170, 50%, 50%, 1.1)", want: false},
		{name: "Invalid hsla alpha value exceeded", param1: "hsla(170, 50%, 50%, -1)", want: false},
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
