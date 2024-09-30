// A package for string sanitizers.
package sanitizer

import "regexp"

// A sanitizer that remove characters that appear in the blacklist.
//
// The characters are used in a RegExp and will escaped for you.
//
//	str := sanitizer.Blacklist("Hello123 World!", "0-9")
//	fmt.Println(str) // "Hello World!"
func Blacklist(str, blacklistedChars string) string {
	if blacklistedChars == "" {
		return str
	}

	// If blacklistedChars is ".", it means blacklist all characters
	if blacklistedChars == "." {
		return ""
	}

	re := regexp.MustCompile("[" + regexp.QuoteMeta(blacklistedChars) + "]+")
	return re.ReplaceAllString(str, "")
}
