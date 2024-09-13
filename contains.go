// A go package for string validators and sanitizers.
package validatorgo

import "strings"

// ContainsOpt is used to configure Contains
// defaults to { IgnoreCase: false, MinOccurrences: 1 }.
type ContainsOpt struct {
	IgnoreCase     bool // ignore case when doing comparison, default false.
	MinOccurrences int  // minimum number of occurrences for the seed in the string. Defaults to 1.
}

// A validator that checks if the string contains the seed.
//
// ContainsOpt is a struct that defaults to { IgnoreCase: false, MinOccurrences: 1 }.
//
// Options:
//
// IgnoreCase: Ignore case when doing comparison, default false.
//
// MinOccurrences: Minimum number of occurrences for the seed in the string. Defaults to 1.
//
//	v := validatorgo.Contains("hello world", "world", ContainsOpt{})
//	fmt.Println(v) // true
//
//	v := validatorgo.Contains("hello world", "earth", ContainsOpt{})
//	fmt.Println(v) // false
func Contains(str, seed string, options ContainsOpt) bool {
	if options.MinOccurrences <= 0 {
		options.MinOccurrences = 1
	}

	if options.IgnoreCase {
		strLowerCase, seedLowerCase := strings.ToLower(str), strings.ToLower(seed)
		return strings.Contains(strLowerCase, seedLowerCase) && strings.Count(strLowerCase, seedLowerCase) >= options.MinOccurrences
	} else {
		return strings.Contains(str, seed) && strings.Count(str, seed) >= options.MinOccurrences
	}
}
