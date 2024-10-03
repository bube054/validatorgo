package validatorgo

import (
	"time"

	"github.com/bube054/validatorgo/sanitizer"
)

const (
	isBeforeOptsDefaultComparisonDate string = ""
)

// IsBeforeOpts is used to configure IsBefore
type IsBeforeOpts struct {
	ComparisonDate string // date to be compared to. Valid layouts are from the time package e.g Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly
}

// A validator that checks if the string is a date that is before the specified date.
//
// IsBeforeOpts is a struct that defaults to { ComparisonDate: "" }.
//
// IsBeforeOpts:
//
// ComparisonDate: defaults to the current time.
// string layouts for str and ComparisonDate can be different layout.
// these are the only valid layouts from the time package e.g Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly.
//
//	ok := validatorgo.IsBefore("2023-01-01", &IsBeforeOpts{ComparisonDate: "2023-09-15"})
//	fmt.Println(ok) // true
//	ok = validatorgo.IsBefore("2024-01-01", &IsBeforeOpts{ComparisonDate: "2023-01-01"})
//	fmt.Println(ok) // false
func IsBefore(str string, opts *IsBeforeOpts) bool {
	if opts == nil {
		opts = setIsBeforeOptsToDefault()
	}

	date1 := sanitizer.ToDate(str)

	var date2 *time.Time
	if opts.ComparisonDate == "" {
		now := time.Now()
		date2 = &now
	} else {
		date2 = sanitizer.ToDate(opts.ComparisonDate)
	}

	if date1 == nil || date2 == nil {
		return false
	}

	return date1.Before(*date2)
}

func setIsBeforeOptsToDefault() (opts *IsBeforeOpts) {
	return &IsBeforeOpts{
		ComparisonDate: isBeforeOptsDefaultComparisonDate,
	}
}
