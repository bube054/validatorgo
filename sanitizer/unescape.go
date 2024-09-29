package sanitizer

import "html"

// A sanitizer that replaces HTML encoded entities with <, >, &, ', ", `, \ and /.
//
//	str := sanitizer.Unescape("&amp;")
//	fmt.Println(str) // "&"
func Unescape(str string) string {
	return html.UnescapeString(str)
}