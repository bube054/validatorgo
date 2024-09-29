package sanitizer

import (
	"html"
)

// A sanitizer that replaces <, >, &, ' and ". with HTML entities.
//
//	str := sanitizer.Escape("<script>alert('XSS');</script>")
//	fmt.Println(str) // "&lt;script&gt;alert(&#39;XSS&#39;);&lt;/script&gt;"
func Escape(str string) string {
	return html.EscapeString(str)
}
