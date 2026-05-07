package validatorgo

import (
	"regexp"
)

// IsISO8601Opts is used to configure IsISO8601
type IsISO8601Opts struct {
	Strict          *bool // must be a date that has actually happened, is happening or will happen.
	StrictSeparator *bool // must be delimited by T
}

func (o *IsISO8601Opts) mergeDefaults() {
	if o.Strict == nil {
		o.Strict = Bool(false)
	}
	if o.StrictSeparator == nil {
		o.StrictSeparator = Bool(false)
	}
}

// A validator that checks if the string is a valid [ISO 8601] date.
//
// IsISO8601Opts is a struct which defaults to { Strict: false, StrictSeparator: false }.
//
// If Strict is true, date strings with invalid dates like 2009-02-29 will be invalid.
//
// If StrictSeparator is true, date strings with date and time separated by anything other than a T will be invalid.
//
//	ok, _ := validatorgo.IsISO8601("2023-09-05", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsISO8601("2023-13-05T14:30:00", nil)
//	fmt.Println(ok) // false
//
// [ISO 8601]: https://en.wikipedia.org/wiki/ISO_8601
func IsISO8601(str string, opts *IsISO8601Opts) (bool, error) {
	if opts == nil {
		opts = &IsISO8601Opts{}
	}
	opts.mergeDefaults()

	var re *regexp.Regexp

	if *opts.StrictSeparator {
		re = regexp.MustCompile(`^(\d{4})(-(0[1-9]|1[0-2])(-([12]\d|0[1-9]|3[01]))([T\s]((([01]\d|2[0-3])((:)[0-5]\d))([\:]\d+)?)?(:[0-5]\d([\.]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)$`)
	} else {
		re = regexp.MustCompile(`^(\d{4})([-\/\. ](0[1-9]|1[0-2])([-\/\. ]([12]\d|0[1-9]|3[01]))([T\s]((([01]\d|2[0-3])([: \.])[0-5]\d)(([: \.])\d+)?([: \.][0-5]\d([\.]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3])[: \.]?([0-5]\d)?)?)?)?)$`)
	}

	capGrps := re.FindStringSubmatch(str)

	if capGrps == nil {
		return false, newValidationError("IsISO8601", ErrInvalidFormat, "invalid iso8601")
	}

	if *opts.Strict {
		year, month, day := capGrps[1], capGrps[3], capGrps[5]

		if validYearMonthDay(year, month, day) {
			return true, nil
		}
		return false, newValidationError("IsISO8601", ErrInvalidFormat, "invalid iso8601")
	}

	return true, nil
}

