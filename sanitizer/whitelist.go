package sanitizer

import (
	"regexp"
)

// A sanitizer that removes characters that do not appear in the whitelist.
//
// The characters are used in a RegExp and so you will need to escape some chars, e.g. whitelist(input, '\\[\\]').
//
//	str := sanitizer.Whitelist("Hello World!", "!@#$%^&*()")
//	fmt.Println(str) // "Hello World"
func Whitelist(str, whitelistedChars string) string {
	re := regexp.MustCompile("[^" + regexp.QuoteMeta(whitelistedChars) + "]+")
	return re.ReplaceAllString(str, "")
}
