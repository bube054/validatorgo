package sanitizers

import (
	"strings"
)

// A sanitizer that trims characters from both the left-side and right-side of the input.
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// A sanitizer that trims characters from the left-side of the input.
func Ltrim(str string) string {
	return strings.TrimLeft(str, " ")
}

// A sanitizer that trims characters from the right-side of the input.
func Rtrim(str string) string {
	return strings.TrimRight(str, " ")
}
