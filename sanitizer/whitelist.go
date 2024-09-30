package sanitizer

import (
	"regexp"
)

// A sanitizer that removes characters that do not appear in the whitelist.
//
// The characters are used in a RegExp and will escaped for you.
//
//	str := sanitizer.Whitelist("Hello123 World!", "a-zA-Z")
//	fmt.Println(str) // "HelloWorld"
func Whitelist(str, whitelistedChars string) string {
	if whitelistedChars == "." {
		return str
	}

	re := regexp.MustCompile("[^" + regexp.QuoteMeta(whitelistedChars) + "]+")
	return re.ReplaceAllString(str, "")
}
