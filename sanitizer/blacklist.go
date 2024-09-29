// A package for string sanitizers.
package sanitizer

import "regexp"

// A sanitizer that removes characters that appear in the blacklist.
//
// blacklistedChars are strings that will be removed from the input string.
//
//	str := sanitizer.Blacklist("Hello World!", "!@#$%^&*()")
//	fmt.Println(str) // "Hello World"
func Blacklist(str, blacklistedChars string) string {
	re := regexp.MustCompile("[" + regexp.QuoteMeta(blacklistedChars) + "]+")
	return re.ReplaceAllString(str, "")
}
