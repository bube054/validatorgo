package validatorgo

import (
	"time"

	"github.com/bube054/validatorgo/sanitizer"
)

var (
	isAfterOptsDefaultComparisonDate string = ""
)

// IsAfterOpts is used to configure IsAfter
type IsAfterOpts struct {
	ComparisonDate string // date to be compared to. Valid layouts are from the time package e.g Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly, StandardDateLayout, SlashDateLayout, DateTimeLayout, ISO8601Layout, ISO8601ZuluLayout, ISO8601WithMillisecondsLayout.
}

// A validator that checks if the string is a date that is after the specified date.
//
// IsAfterOpts is a struct that defaults to { ComparisonDate: "" }.
//
// IsAfterOpts:
//
// ComparisonDate: defaults to the current time.
//
// string layouts for str and ComparisonDate can be different layout.
//
// these are the only valid layouts from the time package
// e.g Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly, StandardDateLayout, SlashDateLayout, DateTimeLayout, ISO8601Layout, ISO8601ZuluLayout, ISO8601WithMillisecondsLayout.
//
//	ok, _ := validatorgo.IsAfter("2023-09-15", &validatorgo.IsAfterOpts{ComparisonDate: "2023-01-01"})
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsAfter("2023-01-01", &validatorgo.IsAfterOpts{ComparisonDate: "2023-09-15"})
//	fmt.Println(ok) // false
func IsAfter(str string, opts *IsAfterOpts) (bool, error) {
	if opts == nil {
		opts = setIsAfterOptsToDefault()
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
		return false, newValidationError("IsAfter", ErrInvalidFormat, "invalid after")
	}
	if date1.After(*date2) {
		return true, nil
	}
	return false, newValidationError("IsAfter", ErrInvalidFormat, "invalid after")
}

func setIsAfterOptsToDefault() *IsAfterOpts {

	return &IsAfterOpts{
		ComparisonDate: isAfterOptsDefaultComparisonDate,
	}
}
