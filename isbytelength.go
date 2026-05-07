package validatorgo

var (
	isIsByteLengthOptsDefaultMin uint  = 0
	isIsByteLengthOptsDefaultMax *uint = nil
)

// IsByteLengthOpts is used to configure IsByteLength
type IsByteLengthOpts struct {
	Min uint  // minimum byte length
	Max *uint // maximum byte length
}

// A validator that checks if the string's length (in UTF-8 bytes) falls in a range.
//
// IsByteLengthOpts is a struct which defaults to { Min: 0, Max: nil }.
//
//	ok, _ := validatorgo.IsByteLength("We♥Go", &validatorgo.IsByteLengthOpts{Min: 5})
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsByteLength("We♥Go", &validatorgo.IsByteLengthOpts{Min: 8})
//	fmt.Println(ok) // false
func IsByteLength(str string, opts *IsByteLengthOpts) (bool, error) {
	if opts == nil {
		opts = setIsByteLengthOptsToDefault()
	}

	lenInBytes := len(str)
	if opts.Max == nil {
		if lenInBytes >= int(opts.Min) {
			return true, nil
		}
		return false, newValidationError("IsByteLength", ErrInvalidFormat, "invalid bytelength")
	} else {
		if lenInBytes >= int(opts.Min) && lenInBytes <= int(*opts.Max) {
			return true, nil
		}
		return false, newValidationError("IsByteLength", ErrInvalidFormat, "invalid bytelength")
	}
}

func setIsByteLengthOptsToDefault() (opts *IsByteLengthOpts) {
	opts = &IsByteLengthOpts{}
	opts.Min = isIsByteLengthOptsDefaultMin
	opts.Max = isIsByteLengthOptsDefaultMax

	return
}
