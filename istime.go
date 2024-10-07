package validatorgo

import (
	"fmt"
	"regexp"
)

var (
	IsTimeOptsHourFormat24 string = "hour24"
	IsTimeOptsHourFormat12 string = "hour12"

	IsTimeOptsModeDefault     string = "default"
	IsTimeOptsModeWithSeconds string = "withSeconds"
)

type IsTimeOpts struct {
	HourFormat string // hour12 e.g 14:30 or hour24 e.g 02:30 PM
	Mode       string // default e.g 13:30 or withSeconds e.g 13:30:20
}

// A validator that checks if the string is a valid time e.g. 23:01:59
//
// IsTimeOpts is a struct which can contain the keys HourFormat and Mode.
//
// HourFormat is a key and defaults to "hour24".
//
// Mode is a key and defaults to "default".
//
// HourFormat can contain the values "hour12" or "hour24", "hour24" will validate hours in 24 format and "hour12" will validate hours in 12 format.
//
// Mode can contain the values "default" or "withSeconds", "default" will validate HH:MM or HH:MM:SS format, "withSeconds" will validate only HH:MM:SS format.
//
//	ok := validatorgo.IsTime("14:30", &validatorgo.IsTime{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsTime("09:5", &validatorgo.IsTime{})
//	fmt.Println(ok) // false
func IsTime(str string, opts *IsTimeOpts) bool {
	if opts == nil {
		opts = setIsTimeOptsToDefault()
	}

	if opts.HourFormat == "" {
		opts.HourFormat = IsTimeOptsHourFormat24
	}

	if opts.HourFormat != IsTimeOptsHourFormat12 && opts.HourFormat != IsTimeOptsHourFormat24 {
		return false
	}

	if opts.Mode == "" {
		opts.Mode = IsTimeOptsModeDefault
	}

	if opts.Mode != IsTimeOptsModeDefault && opts.Mode != IsTimeOptsModeWithSeconds {
		return false
	}

	hourFmtStr1, hourFmtStr2 := "", ""
	if opts.HourFormat == IsTimeOptsHourFormat24 {
		hourFmtStr1 = "(0[0-9]|1[0-9]|2[0-4])"
		hourFmtStr2 = ""
	} else {
		hourFmtStr1 = "(0[0-9]|1[0-2])"
		hourFmtStr2 = " (AM|PM|am|pm)"
	}

	minFmtStr := ""
	if opts.Mode == IsTimeOptsModeDefault {
		minFmtStr = "(:[0-5][0-9])?"
	} else {
		minFmtStr = ":[0-5][0-9]"
	}

	reStr := fmt.Sprintf(`^%s:[0-5][0-9]%s%s$`, hourFmtStr1, minFmtStr, hourFmtStr2)
	re := regexp.MustCompile(reStr)

	return re.MatchString(str)
}

func setIsTimeOptsToDefault() *IsTimeOpts {
	return &IsTimeOpts{
		HourFormat: IsTimeOptsHourFormat24,
		Mode:       IsTimeOptsModeDefault,
	}
}
