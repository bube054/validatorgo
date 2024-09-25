// A package for string sanitizers.
package sanitizer

import (
	"strings"
)

// A sanitizer that removes characters that appear in the blacklist.
//
// blacklistedChars are strings that will be removed from the input string.
//
//	str := sanitizer.Blacklist("Hello World!", "!@#$%^&*()")
//	fmt.Println(str) // Hello World
func Blacklist(str, blacklistedChars string) string {
	sanitizedStr := str

	for _, char := range blacklistedChars {
		sanitizedStr = strings.ReplaceAll(sanitizedStr, string(char), "")
	}

	return sanitizedStr
}
