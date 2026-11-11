package validatorgo

// IsBooleanOpts is used to configure IsBoolean
type IsBooleanOpts struct {
	Loose *bool // strictness of the equality
}

func (o *IsBooleanOpts) mergeDefaults() {
	if o.Loose == nil {
		o.Loose = Bool(false)
	}
}

// A validator that check if the string is a boolean.
//
// IsBooleanOpts is a struct which defaults to { Loose: false } and that can be supplied with the following key(s):
//
// Loose: If Loose is set to false, the validator will strictly match ['true', 'false', '0', '1'].
// If Loose is set to true, the validator will also match 'yes', 'no', and will match a valid boolean string of any case. (e.g.: ['true', 'True', 'TRUE', "false", "False", "FALSE"]).
//
//	ok, _ := validatorgo.IsBoolean("true", &validatorgo.IsBooleanOpts{Loose: false})
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsBoolean("bool", &validatorgo.IsBooleanOpts{Loose: false})
//	fmt.Println(ok) // false
func IsBoolean(str string, opts *IsBooleanOpts) (bool, error) {
	if opts == nil {
		opts = &IsBooleanOpts{}
	}
	optsCopy := *opts // fill in defaults on a copy, never on the caller's struct
	opts = &optsCopy
	opts.mergeDefaults()

	if *opts.Loose {
		if !compileOnce("^(true|True|TRUE|false|False|FALSE|yes|no|0|1)$").MatchString(str) {
			return false, newValidationError("IsBoolean", ErrInvalidValue, "string is not a valid boolean value")
		}
		return true, nil
	} else {
		if !compileOnce("^(true|false|0|1)$").MatchString(str) {
			return false, newValidationError("IsBoolean", ErrInvalidValue, "string is not a valid strict boolean value")
		}
		return true, nil
	}
}
