package sanitizer

import (
	"strings"
)

// A sanitizer that trim characters (whitespace by default) from both sides of the input.
//
//	str := sanitizer.Trim("  Hello World  ", "")
//	fmt.Println(str) // "Hello World"
func Trim(str, chars string) string {
	if chars == "" {
		return strings.Trim(str, " ")
	}
	return strings.Trim(str, chars)
}

// A sanitizer that trims characters (whitespace by default) from the left-side of the input.
//
//	str := sanitizer.LTrim("  Hello World  ", "")
//	fmt.Println(str) // "Hello World  "
func LTrim(str, chars string) string {
	if chars == "" {
		return strings.TrimLeft(str, " ")
	}
	return strings.TrimLeft(str, chars)
}

// A sanitizer that trims characters (whitespace by default) from the right-side of the input.
//
//	str := sanitizer.RTrim("  Hello World  ", "")
//	fmt.Println(str) // "  Hello World"
func RTrim(str, chars string) string {
	if chars == "" {
		return strings.TrimRight(str, " ")
	}
	return strings.TrimRight(str, chars)
}
