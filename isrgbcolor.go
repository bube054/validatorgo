package validatorgo

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

// Pieces of the rgb() and rgba() patterns, matching validator.js.
const (
	rgbNumber  = `(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])` // 0 to 255
	rgbPercent = `(100|[1-9]?[0-9])%`                          // 0% to 100%
	rgbAlpha   = `(0?\.[0-9]{1,2}|1(\.0)?|0(\.0)?)`            // 0 to 1, e.g. .5, 0.25, 1
)

// A validator that checks if the string is a rgb or rgba color.
//
// rgb() takes three values and rgba() takes three values plus an alpha from 0 to 1.
// Each value is a number from 0 to 255, or with IncludePercentValues a percentage
// from 0% to 100%. The three values must all be numbers or all be percentages.
//
// IsRgbOpts is a struct with the following properties:
//
// IncludePercentValues defaults to false. If set to true, percentage values such as rgb(5%,5%,5%) or rgba(90%,90%,90%,.3) are also allowed.
//
// AllowSpaces defaults to false, which prohibits whitespace. If set to true, whitespace around the values is allowed, such as rgb(255, 255, 255) or even rgba(255,       128,        0,      0.7).
//
//	ok, _ := validatorgo.IsRgbColor("rgb(255,0,0)", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsRgbColor("rgb( 255 , 0 , 0 )", nil)
//	fmt.Println(ok) // false
//	ok, _ = validatorgo.IsRgbColor("rgb(999,0,0)", nil)
//	fmt.Println(ok) // false
func IsRgbColor(str string, opts *IsRgbOpts) (bool, error) {
	if opts == nil {
		opts = &IsRgbOpts{}
	}
	optsCopy := *opts // fill in defaults on a copy, never on the caller's struct
	opts = &optsCopy
	opts.mergeDefaults()

	open, sep, closing := `\(`, `,`, `\)`
	if *opts.AllowSpaces {
		open, sep, closing = `\(\s*`, `\s*,\s*`, `\s*\)`
	}

	values := []string{rgbNumber}
	if *opts.IncludePercentValues {
		values = append(values, rgbPercent)
	}

	for _, v := range values {
		three := v + sep + v + sep + v
		pattern := `^(rgb` + open + three + closing + `|rgba` + open + three + sep + rgbAlpha + closing + `)$`
		if compileOnce(pattern).MatchString(str) {
			return true, nil
		}
	}

	return false, newValidationError("IsRgbColor", ErrInvalidFormat, "invalid rgbcolor")
}
