package validatorgo

import (
	"unicode/utf8"
)

// IsLengthOpts is used to configure IsLength
type IsLengthOpts struct {
	Min *uint // Minimum character length
	Max *uint // Maximum character length
}

func (o *IsLengthOpts) mergeDefaults() {
	if o.Min == nil {
		o.Min = Uint(0)
	}
}

// A validator that checks if the string's length falls in a range.
//
// IsLengthOpts is a struct which defaults to { Min: 0, Max: nil }.
//
// Note: this function takes into account surrogate pairs.
//
//	ok, _ := validatorgo.IsLength("hello", &validatorgo.IsLengthOpts{Min: 3})
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsLength("hi", &validatorgo.IsLengthOpts{Min: 3})
//	fmt.Println(ok) // false
func IsLength(str string, opts *IsLengthOpts) (bool, error) {
	if opts == nil {
		opts = &IsLengthOpts{}
	}
	opts.mergeDefaults()

	length := uint(utf8.RuneCountInString(str))

	withinLimits := true

	if opts.Max != nil {
		isMax := *(opts.Max) >= length
		withinLimits = withinLimits && isMax
	}

	isMin := *opts.Min <= length
	withinLimits = withinLimits && isMin

	if withinLimits {
		return true, nil
	}
	return false, newValidationError("IsLength", ErrInvalidFormat, "invalid length")
}

