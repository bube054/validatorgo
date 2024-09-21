package validatorgo

import "regexp"

// IsRgbOpts is used to configure IsRgbColor
type IsRgbOpts struct {
	IncludePercentValues bool // must use percent values 90% not 0-255
	AllowSpaces bool // whether to include spaces
}

// A validator that checks if the string is a rgb or rgba color.
//
// IsRgbOpts is a struct with the following properties:
//
// IncludePercentValues defaults to false. If you don't want to allow to set rgb or rgba values with percents, like rgb(5%,5%,5%), or rgba(90%,90%,90%,.3), then set it to false.
//
// AllowSpaces defaults to false, which prohibits whitespace. If set to false, whitespace between color values is allowed, such as rgb(255, 255, 255) or even rgba(255,       128,        0,      0.7).
//
//	ok := validatorgo.IsRgbColor("rgb(255,0,0)", validatorgo.IsRgbColor{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsRgbColor("rgb( 255 , 0 , 0 )", validatorgo.IsRgbColor{})
//	fmt.Println(ok) // false
func IsRgbColor(str string, opts IsRgbOpts) bool {
	if opts.IncludePercentValues && opts.AllowSpaces {
		return regexp.MustCompile(`^rgba?\((\d{0,100}(\.[0-9]*)?%|\d{0,255}),\s*(\d{0,100}(\.[0-9]*)?%|\d{0,255}),\s*(\d{0,100}(\.[0-9]*)?%|\d{0,255})(,\s*(1|0?\.[0-9])?)?\)$`).MatchString(str)
	} else if !opts.IncludePercentValues && opts.AllowSpaces {
		return regexp.MustCompile(`^rgba?\(\d{0,255},\s*\d{0,255},\s*\d{0,255}(,\s*(1|0?\.[1-9])*)?\)$`).MatchString(str)
	} else if opts.IncludePercentValues && !opts.AllowSpaces {
		return regexp.MustCompile(`^rgba?\((\d{0,100}(\.[0-9]*)?%|\d{0,255}),(\d{0,100}(\.[0-9]*)?%|\d{0,255}),(\d{0,100}(\.[0-9]*)?%|\d{0,255})(,(1|0?\.[0-9])?)?\)$`).MatchString(str)
	}else {
		return regexp.MustCompile(`^rgba?\(\d{0,255},\d{0,255},\d{0,255}(,(1|0?\.[1-9])*)?\)$`).MatchString(str)
	}
}