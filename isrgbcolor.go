package validatorgo

import "regexp"

// IsRgbOpts is used to configure IsRgbColor
type IsRgbOpts struct {
	IncludePercentValues *bool // must use percent values 90% not 0-255
	AllowSpaces          *bool // whether to include spaces
}

func (o *IsRgbOpts) mergeDefaults() {
	if o.IncludePercentValues == nil {
		o.IncludePercentValues = Bool(false)
	}
	if o.AllowSpaces == nil {
		o.AllowSpaces = Bool(false)
	}
}

// A validator that checks if the string is a rgb or rgba color.
//
// IsRgbOpts is a struct with the following properties:
//
// IncludePercentValues defaults to false. If you don't want to allow to set rgb or rgba values with percents, like rgb(5%,5%,5%), or rgba(90%,90%,90%,.3), then set it to false.
//
// AllowSpaces defaults to false, which prohibits whitespace. If set to false, whitespace between color values is allowed, such as rgb(255, 255, 255) or even rgba(255,       128,        0,      0.7).
//
//	ok, _ := validatorgo.IsRgbColor("rgb(255,0,0)", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsRgbColor("rgb( 255 , 0 , 0 )", nil)
//	fmt.Println(ok) // false
func IsRgbColor(str string, opts *IsRgbOpts) (bool, error) {
	if opts == nil {
		opts = &IsRgbOpts{}
	}
	opts.mergeDefaults()

	if *opts.IncludePercentValues && *opts.AllowSpaces {
		if regexp.MustCompile(`^rgba?\((\d{0,100}(\.[0-9]*)?%|\d{0,255}),\s*(\d{0,100}(\.[0-9]*)?%|\d{0,255}),\s*(\d{0,100}(\.[0-9]*)?%|\d{0,255})(,\s*(1|0?\.[0-9])?)?\)$`).MatchString(str) {
			return true, nil
		}
		return false, newValidationError("IsRgbColor", ErrInvalidFormat, "invalid rgbcolor")
	} else if !*opts.IncludePercentValues && *opts.AllowSpaces {
		if regexp.MustCompile(`^rgba?\(\d{0,255},\s*\d{0,255},\s*\d{0,255}(,\s*(1|0?\.[1-9])*)?\)$`).MatchString(str) {
			return true, nil
		}
		return false, newValidationError("IsRgbColor", ErrInvalidFormat, "invalid rgbcolor")
	} else if *opts.IncludePercentValues && !*opts.AllowSpaces {
		if regexp.MustCompile(`^rgba?\((\d{0,100}(\.[0-9]*)?%|\d{0,255}),(\d{0,100}(\.[0-9]*)?%|\d{0,255}),(\d{0,100}(\.[0-9]*)?%|\d{0,255})(,(1|0?\.[0-9])?)?\)$`).MatchString(str) {
			return true, nil
		}
		return false, newValidationError("IsRgbColor", ErrInvalidFormat, "invalid rgbcolor")
	} else {
		if regexp.MustCompile(`^rgba?\(\d{0,255},\d{0,255},\d{0,255}(,(1|0?\.[1-9])*)?\)$`).MatchString(str) {
			return true, nil
		}
		return false, newValidationError("IsRgbColor", ErrInvalidFormat, "invalid rgbcolor")
	}
}

