package sanitizer

import (
	"strings"
)

// StripLow removes characters with a numerical value < 32 and 127, mostly control characters.
// If keepNewLines is true, newline characters are preserved (\n and \r, hex 0xA and 0xD).
func StripLow(str string, keepNewLines bool) string {
	var newStr strings.Builder

	for _, r := range str {
		if r < 32 || r == 127 {
			if keepNewLines && (r == '\n' || r == '\r') {
				newStr.WriteRune(r)
			}
		} else {
			newStr.WriteRune(r)
		}
	}

	return newStr.String()
}
