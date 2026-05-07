package validatorgo

import (
	"time"
)

// IsDate formats
const (
	StandardDateLayout            = "2006-01-02"
	SlashDateLayout               = "2006/01/02"
	DateTimeLayout                = "2006-01-02 15:04:05"
	ISO8601Layout                 = "2006-01-02T15:04:05"
	ISO8601ZuluLayout             = "2006-01-02T15:04:05Z"
	ISO8601WithMillisecondsLayout = "2006-01-02T15:04:05.000Z"
)

var dateLayouts = [6]string{StandardDateLayout, SlashDateLayout, DateTimeLayout, ISO8601Layout, ISO8601ZuluLayout, ISO8601WithMillisecondsLayout}

var (
	isDateOptsDefaultFormat     string = StandardDateLayout
	isDateOptsDefaultStrictMode bool   = false
)

// IsDateOpts is used to configure IsDate
type IsDateOpts struct {
	Format     string
	StrictMode bool
}

func dateMatchesAnyFormat(str string) (bool, error) {
	for _, format := range dateLayouts {
		_, err := time.Parse(format, str)
		if err == nil {
			return true, nil
		}
	}
	return false, newValidationError("dateMatchesAnyFormat", ErrInvalidFormat, "invalid datematchesanyformat")
}

// A Validator that checks if the string is a valid date. e.g. 2002-07-15.
//
// IsDateOpts is a struct which can contain the keys Format, StrictMode.
//
// Format: is a string and defaults to validatorgo.StandardDateLayout if "any" or no value is provided.
//
// StrictMode: is a boolean and defaults to false. If StrictMode is set to true, the validator will reject strings different from Format.
//
//	ok, _ := validatorgo.IsDate("2006-01-02", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsDate("01/023/2006", nil)
//	fmt.Println(ok) // false
func IsDate(str string, opts *IsDateOpts) (bool, error) {
	if opts == nil {
		opts = setIsDateOptsToDefault()
	}

	switch opts.Format {
	case StandardDateLayout, SlashDateLayout, DateTimeLayout, ISO8601Layout, ISO8601ZuluLayout, ISO8601WithMillisecondsLayout:
	case "", "any":
		opts.Format = isDateOptsDefaultFormat
	default:
		return false, newValidationError("IsDate", ErrInvalidFormat, "invalid date")
	}

	if opts.StrictMode {
		_, err := time.Parse(opts.Format, str)
		if err == nil {
			return true, nil
		}
		return false, newValidationError("IsDate", ErrInvalidFormat, "invalid date")
	} else {
		ok, _ := dateMatchesAnyFormat(str)
		if ok {
			return true, nil
		}
		return false, newValidationError("IsDate", ErrInvalidFormat, "invalid date")
	}
}

func setIsDateOptsToDefault() *IsDateOpts {
	return &IsDateOpts{
		Format:     isDateOptsDefaultFormat,
		StrictMode: isDateOptsDefaultStrictMode,
	}
}
