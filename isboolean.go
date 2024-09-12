package validatorgo

import "regexp"

// IsBooleanOpts is used to configure IsBoolean
type IsBooleanOpts struct {
	Loose bool
}

// A validator that check if the string is a boolean.
// IsBooleanOpts is a struct which defaults to { Loose: false }. If Loose is set to false, the validator will strictly match ['true', 'false', '0', '1'].
// If Loose is set to true, the validator will also match 'yes', 'no', and will match a valid boolean string of any case. (e.g.: ['true', 'True', 'TRUE', "false", "False", "FALSE"]).
func IsBoolean(str string, opts IsBooleanOpts) bool {
	if opts.Loose {
		return regexp.MustCompile("^(true|True|TRUE|false|False|FALSE|yes|no|0|1)$").MatchString(str)
	} else {
		return regexp.MustCompile("^(true|false|0|1)$").MatchString(str)
	}
}
