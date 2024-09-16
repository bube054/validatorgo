// A go package for string validators and sanitizers.
package validatorgo

import "strings"

const defaultContainsMinOccurrence = 1

// ContainsOpt is used to configure Contains
type ContainsOpt struct {
	IgnoreCase     bool // ignore case when doing comparison, default false.
	MinOccurrences int  // minimum number of occurrences for the seed in the string. Defaults to 1.
}

// A validator that checks if the string contains the seed.
//
// ContainsOpt is a struct that defaults to { IgnoreCase: false, MinOccurrences: 1 }.
//
// ContainsOpt:
//
// IgnoreCase: Ignore case when doing comparison, default false.
//
// MinOccurrences: Minimum number of occurrences for the seed in the string. Defaults to 1.
//
//	ok := validatorgo.Contains("hello world", "world", ContainsOpt{})
//	fmt.Println(ok) // true
//
//	ok := validatorgo.Contains("hello world", "earth", ContainsOpt{})
//	fmt.Println(ok) // false
func Contains(str, seed string, opts ContainsOpt) bool {
	if opts.MinOccurrences <= 0 {
		opts.MinOccurrences = defaultContainsMinOccurrence
	}

	if opts.IgnoreCase {
		strLowerCase, seedLowerCase := strings.ToLower(str), strings.ToLower(seed)
		return strings.Contains(strLowerCase, seedLowerCase) && strings.Count(strLowerCase, seedLowerCase) >= opts.MinOccurrences
	} else {
		return strings.Contains(str, seed) && strings.Count(str, seed) >= opts.MinOccurrences
	}
}
