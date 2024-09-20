package validatorgo

import (
	"regexp"
	"strconv"
)

// IsISO8601Opts is used to configure IsISO8601
type IsISO8601Opts struct {
	Strict          bool // must be a date that has actually happened, is happening or will happen.
	StrictSeparator bool // must be delimited by T
}

// A validator that checks if the string is a valid [ISO 8601] date.
//
// IsISO8601Opts is a struct which defaults to { Strict: false, StrictSeparator: false }.
//
// If Strict is true, date strings with invalid dates like 2009-02-29 will be invalid.
//
// If StrictSeparator is true, date strings with date and time separated by anything other than a T will be invalid.
//
//	ok := validatorgo.IsISO8601("2023-09-05", validatorgo.IsISO8601Opts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsISO8601("2023-13-05T14:30:00", validatorgo.IsISO8601Opts{})
//	fmt.Println(ok) // false
//
// [ISO 8601]: https://en.wikipedia.org/wiki/ISO_8601
func IsISO8601(str string, opts IsISO8601Opts) bool {
	var re *regexp.Regexp

	if opts.StrictSeparator {
		re = regexp.MustCompile(`^(\d{4})(-(0[1-9]|1[0-2])(-([12]\d|0[1-9]|3[01]))([T\s]((([01]\d|2[0-3])((:)[0-5]\d))([\:]\d+)?)?(:[0-5]\d([\.]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3]):?([0-5]\d)?)?)?)$`)
	} else {
		re = regexp.MustCompile(`^(\d{4})([-\/\. ](0[1-9]|1[0-2])([-\/\. ]([12]\d|0[1-9]|3[01]))([T\s]((([01]\d|2[0-3])([: \.])[0-5]\d)(([: \.])\d+)?([: \.][0-5]\d([\.]\d+)?)?([zZ]|([\+-])([01]\d|2[0-3])[: \.]?([0-5]\d)?)?)?)?)$`)
	}

	capGrps := re.FindStringSubmatch(str)

	if capGrps == nil {
		return false
	}

	if opts.Strict {
		year, err := strconv.Atoi(capGrps[1])

		if err != nil {
			return false
		}

		month, err := strconv.Atoi(capGrps[3])

		if err != nil {
			return false
		}

		day, err := strconv.Atoi(capGrps[5])

		if err != nil {
			return false
		}

		monthLength := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

		// Adjust for leap years
		if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
			monthLength[1] = 29
		}

		if !(month > 0 && month < 13) {
			return false
		}

		if day < 0 || day > monthLength[month-1] {
			return false
		}

	}

	return true
}
