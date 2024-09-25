package sanitizer

import (
	"html"
)

// A sanitizer that replaces <, >, &, ' and ". with HTML entities.
func Escape(str string) string {
	return html.EscapeString(str)
}
