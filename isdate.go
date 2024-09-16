package validatorgo

import (
	"regexp"
)

// IsDateOpts is used to configure IsDate
type IsDateOpts struct {
	Format     string
	StrictMode bool
}

// IsDate formats
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

// dateFormatRegex is the set of date formats and their validating regex
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
//
// IsDateOpts is a struct which can contain the keys Format, StrictMode.
//
// Format: is a string and defaults to validatorgo.ISO8601 if "any" or no value is provided.
//
// StrictMode: is a boolean and defaults to false. If StrictMode is set to true, the validator will reject strings different from Format.
//
//	ok := validatorgo.IsDate("2006-01-02", validatorgo.IsDateOpts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsDate("01/023/2006", validatorgo.IsDateOpts{})
//	fmt.Println(ok) // false
func IsDate(str string, opts IsDateOpts) bool {
	switch opts.Format {
	case ISO8601, USFormat, EUFormat, JapanFormat, LongForm, ShortForm, NoDelim, WeekDay, YearMonth, UnixTimestamp:
	case "", "any":
		opts.Format = ISO8601
	default:
		return false
	}

	if opts.StrictMode {
		return dateFormatRegex[opts.Format].MatchString(str)
	} else {
		return dateMatchesAnyFormat(str)
	}
}
