package validatorgo

import (
	"unicode/utf8"
)

type IsLengthOpts struct {
	Min uint
	Max *uint
}

// A validator that checks if the string's length falls in a range.
// IsLengthOpts is a struct which defaults to { Min: 0, Max: nil }.
// Note: this function takes into account surrogate pairs.
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
