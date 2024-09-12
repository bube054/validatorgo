package validatorgo

// IsByteLengthOpts is used to configure IsByteLength
type IsByteLengthOpts struct {
	Min uint
	Max *uint
}

// A validator that checks if the string's length (in UTF-8 bytes) falls in a range.
// options is a struct which defaults to { Min: 0, Max: nil }.
func IsByteLength(str string, opts IsByteLengthOpts) bool {
	lenInBytes := len(str)
	if opts.Max == nil {
		return lenInBytes >= int(opts.Min)
	} else {
		return lenInBytes >= int(opts.Min) && lenInBytes <= int(*opts.Max)
	}
}
