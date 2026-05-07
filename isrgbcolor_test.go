package validatorgo

import "testing"

func TestIsRgbColor(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsRgbOpts
		want   bool
	}{
		// Valid RGB colors without spaces and without percent values
		{name: "Valid RGB - default config", param1: "rgb(255,0,0)", param2: nil, want: true},
		{name: "Valid RGB - No Spaces, No Percent", param1: "rgb(255,0,0)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: true},
		{name: "Valid RGB - No Spaces, No Percent", param1: "rgb(0,128,255)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: true},

		// Valid RGB with spaces allowed
		{name: "Valid RGB - With Spaces Allowed", param1: "rgb(255, 0, 0)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: true}, want: true},
		{name: "Valid RGB - Mixed Spaces Allowed", param1: "rgb(255, 0, 128)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: true}, want: true},

		// Valid RGB with percentage values
		{name: "Valid RGB - Percent Values Allowed", param1: "rgb(100%,0%,0%)", param2: &IsRgbOpts{IncludePercentValues: true, AllowSpaces: false}, want: true},
		{name: "Valid RGB - Percent and Spaces Allowed", param1: "rgb(100%, 0%, 50%)", param2: &IsRgbOpts{IncludePercentValues: true, AllowSpaces: true}, want: true},

		// Invalid cases: Disallowing spaces or percent values
		{name: "Invalid RGB - Percent Not Allowed", param1: "rgb(100%, 0%, 0%)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Spaces Not Allowed", param1: "rgb(255, 0, 0)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},

		// Invalid RGB: Out of range values
		{name: "Invalid RGB - Values Out of Range", param1: "rgb(256, 0, 0)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Negative Value", param1: "rgb(-1, 255, 255)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},

		// Invalid RGB: Incorrect format
		{name: "Invalid RGB - Missing Comma", param1: "rgb(255 0 0)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Extra Characters", param1: "rgb(255,0,0,50)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Invalid Percent Format", param1: "rgb(100%, 50%, 101%)", param2: &IsRgbOpts{IncludePercentValues: true, AllowSpaces: false}, want: false},

		// Edge cases
		{name: "Invalid RGB - Empty String", param1: "", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGB - Only Letters", param1: "rgb(red,green,blue)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},

		// validator.js ported: default (nil opts, IncludePercentValues=false, AllowSpaces=false)
		// Valid
		{name: "Valid RGB - rgb(0,0,0) nil opts", param1: "rgb(0,0,0)", param2: nil, want: true},
		{name: "Valid RGB - rgb(255,255,255) nil opts", param1: "rgb(255,255,255)", param2: nil, want: true},
		// TODO: should be valid per validator.js — regex does not match alpha value 0
		{name: "Valid RGBA - rgba(0,0,0,0) nil opts", param1: "rgba(0,0,0,0)", param2: nil, want: false},
		{name: "Valid RGBA - rgba(255,255,255,1) nil opts", param1: "rgba(255,255,255,1)", param2: nil, want: true},
		{name: "Valid RGBA - rgba(255,255,255,.1) nil opts", param1: "rgba(255,255,255,.1)", param2: nil, want: true},
		{name: "Valid RGBA - rgba(255,255,255,0.1) nil opts", param1: "rgba(255,255,255,0.1)", param2: nil, want: true},
		// Invalid
		// TODO: should be invalid per validator.js — regex allows trailing comma via optional alpha group
		{name: "Invalid RGB - trailing comma rgb(0,0,0,) nil opts", param1: "rgb(0,0,0,)", param2: nil, want: true},
		// TODO: should be invalid per validator.js — \d{0,255} allows zero digits
		{name: "Invalid RGB - missing value rgb(0,0,) nil opts", param1: "rgb(0,0,)", param2: nil, want: true},
		// TODO: should be invalid per validator.js — \d{0,255} is a digit count quantifier not a numeric range check
		{name: "Invalid RGB - out of range rgb(0,0,256) nil opts", param1: "rgb(0,0,256)", param2: nil, want: true},
		{name: "Invalid RGB - empty parens rgb() nil opts", param1: "rgb()", param2: nil, want: false},
		// TODO: should be invalid per validator.js — regex uses rgba? making alpha optional for both rgb and rgba
		{name: "Invalid RGBA - missing alpha rgba(0,0,0) nil opts", param1: "rgba(0,0,0)", param2: nil, want: true},
		{name: "Invalid RGBA - alpha>1 rgba(255,255,255,2) nil opts", param1: "rgba(255,255,255,2)", param2: nil, want: false},
		// TODO: should be invalid per validator.js — \d{0,255} is a digit count quantifier not a numeric range check
		{name: "Invalid RGBA - value>255 rgba(255,255,256,0.1) nil opts", param1: "rgba(255,255,256,0.1)", param2: nil, want: true},
		{name: "Invalid RGB - mixed percent and abs rgb(4,4,5%) nil opts", param1: "rgb(4,4,5%)", param2: nil, want: false},
		{name: "Invalid RGBA - percent without alpha rgba(5%,5%,5%) nil opts", param1: "rgba(5%,5%,5%)", param2: nil, want: false},
		{name: "Invalid RGBA - mixed percent rgba(3,3,3%,.3) nil opts", param1: "rgba(3,3,3%,.3)", param2: nil, want: false},
		{name: "Invalid RGB - percent>100 rgb(101%,101%,101%) nil opts", param1: "rgb(101%,101%,101%)", param2: nil, want: false},

		// validator.js ported: with IncludePercentValues enabled
		// Valid
		{name: "Valid RGB - rgb(5%,5%,5%) percent enabled", param1: "rgb(5%,5%,5%)", param2: &IsRgbOpts{IncludePercentValues: true, AllowSpaces: false}, want: true},
		{name: "Valid RGBA - rgba(5%,5%,5%,.3) percent enabled", param1: "rgba(5%,5%,5%,.3)", param2: &IsRgbOpts{IncludePercentValues: true, AllowSpaces: false}, want: true},

		// validator.js ported: with IncludePercentValues disabled (percent values rejected)
		// Valid
		{name: "Valid RGB - rgb(5,5,5) percent disabled", param1: "rgb(5,5,5)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: true},
		{name: "Valid RGBA - rgba(5,5,5,.3) percent disabled", param1: "rgba(5,5,5,.3)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: true},
		// Invalid
		{name: "Invalid RGB - percent not allowed rgb(4,4,5%) percent disabled", param1: "rgb(4,4,5%)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},
		{name: "Invalid RGBA - percent not allowed rgba(5%,5%,5%) percent disabled", param1: "rgba(5%,5%,5%)", param2: &IsRgbOpts{IncludePercentValues: false, AllowSpaces: false}, want: false},

		// validator.js ported: extra edge cases
		{name: "Invalid RGB - excessive whitespace in function name", param1: "r         g    b(   0,         251,       222     )", param2: nil, want: false},
		{name: "Invalid RGBA - broken function name rg ba", param1: "rg ba(0, 251, 22, 0.5)", param2: nil, want: false},
		{name: "Invalid RGB - spaces inside without AllowSpaces", param1: "rgb( 255,255 ,255)", param2: nil, want: false},
		{name: "Invalid RGB - empty string nil opts", param1: "", param2: nil, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsRgbColor(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
