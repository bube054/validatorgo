package govalidator

import (
	"regexp"
)

// ipVersionRegex is the set of versions and their validating regex
var ipVersionRegex = map[string]*regexp.Regexp{
	"4": regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`),
	"6": regexp.MustCompile(`^([[:xdigit:]]{1,4}(?::[[:xdigit:]]{1,4}){7}|::|:(?::[[:xdigit:]]{1,4}){1,6}|[[:xdigit:]]{1,4}:(?::[[:xdigit:]]{1,4}){1,5}|(?:[[:xdigit:]]{1,4}:){2}(?::[[:xdigit:]]{1,4}){1,4}|(?:[[:xdigit:]]{1,4}:){3}(?::[[:xdigit:]]{1,4}){1,3}|(?:[[:xdigit:]]{1,4}:){4}(?::[[:xdigit:]]{1,4}){1,2}|(?:[[:xdigit:]]{1,4}:){5}:[[:xdigit:]]{1,4}|(?:[[:xdigit:]]{1,4}:){1,6}:)$`),
}

func IsIP(str, version string) bool {
	if version != "4" && version != "6" {
		for _, re := range ipVersionRegex {
			matches := re.MatchString(str)
			if matches {
				return true
			}
		}
		return false
	} else {
		re := ipVersionRegex[version]
		return re.MatchString(str)
	}
}
