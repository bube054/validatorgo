package validatorgo

import (
	"regexp"
)

type IsDateOpts struct {
	Format     string
	StrictMode bool
}

const (
	ISO8601       = "2006-01-02"              // YYYY-MM-DD
	USFormat      = "01/02/2006"              // MM/DD/YYYY
	EUFormat      = "02/01/2006"              // DD/MM/YYYY
	JapanFormat   = "2006年01月02日"             // YYYY年MM月DD日
	LongForm      = "January 02, 2006"        // Month DD, YYYY
	ShortForm     = "02-Jan-2006"             // DD-MMM-YYYY
	NoDelim       = "20060102"                // YYYYMMDD
	WeekDay       = "Monday, 02 January 2006" // Day, DD Month YYYY
	YearMonth     = "2006-01"                 // YYYY-MM
	UnixTimestamp = "2006-01-02 15:04:05"     // Full date and time
)

var dateFormatRegex = map[string]*regexp.Regexp{
	ISO8601:       regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`),                   // YYYY-MM-DD
	USFormat:      regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`),                   // MM/DD/YYYY
	EUFormat:      regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`),                   // DD/MM/YYYY
	JapanFormat:   regexp.MustCompile(`^\d{4}年\d{2}月\d{2}日$`),                  // YYYY年MM月DD日
	LongForm:      regexp.MustCompile(`^[A-Za-z]+ \d{2}, \d{4}$`),              // Month DD, YYYY
	ShortForm:     regexp.MustCompile(`^\d{2}-[A-Za-z]{3}-\d{4}$`),             // DD-MMM-YYYY
	NoDelim:       regexp.MustCompile(`^\d{8}$`),                               // YYYYMMDD
	WeekDay:       regexp.MustCompile(`^[A-Za-z]+, \d{2} [A-Za-z]+ \d{4}$`),    // Day, DD Month YYYY
	YearMonth:     regexp.MustCompile(`^\d{4}-\d{2}$`),                         // YYYY-MM
	UnixTimestamp: regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`), // Full date and time
}

func dateMatchesAnyFormat(str string) bool {
	for _, format := range dateFormatRegex {
		matches := format.MatchString(str)
		if matches {
			return true
		}
	}
	return false
}

// A Validator that checks if the string is a valid date. e.g. 2002-07-15.
// ISO8601 is a struct which can contain the keys format, strictMode.
// format is a string and defaults to ginvalidator.ISO8601 if not provided or found.
// strictMode is a boolean and defaults to false. If strictMode is set to true, the validator will reject strings different from format.
func IsDate(str string, opts IsDateOpts) bool {
	switch opts.Format {
	case ISO8601, USFormat, EUFormat, JapanFormat, LongForm, ShortForm, NoDelim, WeekDay, YearMonth, UnixTimestamp:
	default:
		opts.Format = ISO8601
	}

	if opts.StrictMode {
		return dateFormatRegex[opts.Format].MatchString(str)
	} else {
		return dateMatchesAnyFormat(str)
	}
}
