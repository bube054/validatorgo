package validatorgo

// IsByteLengthOpts is used to configure IsByteLength
type IsByteLengthOpts struct {
	Min uint  // minimum byte length
	Max *uint // maximum byte length
}

// A validator that checks if the string's length (in UTF-8 bytes) falls in a range.
//
// IsByteLengthOpts is a struct which defaults to { Min: 0, Max: nil }.
//
//	ok := validatorgo.IsByteLength("We♥Go", IsByteLengthOpts{Min: 5})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsByteLength("We♥Go", IsByteLengthOpts{Min: 8})
//	fmt.Println(ok) // false
func IsByteLength(str string, opts IsByteLengthOpts) bool {
	lenInBytes := len(str)
	if opts.Max == nil {
		return lenInBytes >= int(opts.Min)
	} else {
		return lenInBytes >= int(opts.Min) && lenInBytes <= int(*opts.Max)
	}
}
