package sanitizers

import (
	"regexp"
)

// A sanitizer that removes characters that appear in the blacklist.
// The characters are used in a RegExp and so you will need to escape some chars, e.g. blacklist(input, '\\[\\]').
func Blacklist(str, chars string) string {
	re := regexp.MustCompile(chars)
	newStr := re.ReplaceAllString(str, "")
	return newStr
}
