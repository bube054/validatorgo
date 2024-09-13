// A go package for string validators and sanitizers.
package validatorgo

import "strings"

// ContainsOpt defaults to { ignoreCase: false, minOccurrences: 1 }.
type ContainsOpt struct {
	IgnoreCase     bool // Ignore case when doing comparison, default false.
	MinOccurrences int  // Minimum number of occurrences for the seed in the string. Defaults to 1.
}

// A validator that checks if the string contains the seed.
//
//	ok := govalidator.Contains("Hello, world! Hello, universe!", "hello", ContainsOpt{IgnoreCase: true})
//	fmt.Println(ok) // true
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
