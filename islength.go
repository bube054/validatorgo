package validatorgo

import (
	"unicode/utf8"
)

// IsLengthOpts is used to configure IsLength
type IsLengthOpts struct {
	Min uint // Minimum character length
	Max *uint // Maximum character length
}

// A validator that checks if the string's length falls in a range.
//
// IsLengthOpts is a struct which defaults to { Min: 0, Max: nil }.
//
// Note: this function takes into account surrogate pairs.
//
//	ok := validatorgo.IsLength("hello", validatorgo.IsLengthOpts{Min: 3})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsLength("hi", validatorgo.IsLengthOpts{Min: 3})
//	fmt.Println(ok) // false
func IsLength(str string, opts IsLengthOpts) bool {
	len := uint(utf8.RuneCountInString(str))

	withinLimits := true

	if opts.Max != nil {
		isMax := *(opts.Max) >= len
		withinLimits = withinLimits && isMax
	}

	isMin := opts.Min <= len
	withinLimits = withinLimits && isMin

	return withinLimits
}
