package validatorgo

import (
	"time"

	"github.com/bube054/govalidator/sanitizers"
)

// A validator that checks if the string is a date that is before the specified date.
// comparisonDate defaults to the current time.
// string layouts for str and comparisonDate can be different layout.
// these are the only valid layouts from the time package e.g Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly
//
//	ok := govalidator.IsBefore("2009-06-02", "2007-01-02")
//	fmt.Println(ok) // false
//	ok = govalidator.IsBefore("2020-04-03", "")
//	fmt.Println(ok) // true
func IsBefore(str string, comparisonDate string) bool {
	date1 := sanitizers.ToDate(str)

	var date2 *time.Time
	if comparisonDate == "" {
		now := time.Now()
		date2 = &now
	} else {
		date2 = sanitizers.ToDate(comparisonDate)
	}

	if date1 == nil || date2 == nil {
		return false
	}

	return date1.Before(*date2)
}
