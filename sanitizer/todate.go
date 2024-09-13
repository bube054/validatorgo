package sanitizer

import "time"

// A sanitizer that converts the input string to a pointer to time.Time, if the input is not of a time layout returns a nil pointer.
// e.g Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly
func ToDate(str string) *time.Time {
	layouts := []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, str); err == nil {
			return &t
		}
	}

	return nil
}
