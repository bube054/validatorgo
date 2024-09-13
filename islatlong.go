package validatorgo

import (
	"regexp"
)

type IsLatLongOpts struct {
	CheckDMS bool
}

// A validator that checks if the string is a valid latitude-longitude coordinate in the format lat,long or lat, long.
// options is an object that defaults to { CheckDMS: false }.
// Pass CheckDMS as true to validate DMS(degrees, minutes, and seconds) latitude-longitude format.
func IsLatLong(str string, opts IsLatLongOpts) bool {
	var re *regexp.Regexp
	if opts.CheckDMS {
		re = regexp.MustCompile(`^([+\-]?[0-8]?\d|90)[°˚º\s-]+([0-5]?\d)['′\s-]+([0-5]?\d(\.\d*)?)["¨˝\s-]*([NnSs])[\s,]+([+\-]?(0?\d?\d|1[0-7]\d|180))[°˚º\s-]+([0-5]?\d)['′\s-]+([0-5]?\d(\.\d*)?)["¨˝\s-]*([EeWw])$`)
		return re.MatchString(str)
	} else {
		re = regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)
		return re.MatchString(str)
	}
}
