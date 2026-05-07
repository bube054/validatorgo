package validatorgo

import (
	"fmt"
	"regexp"
)

// IsMacAddressOpts is used to configure IsMacAddress
type IsMacAddressOpts struct {
	NoSeparators *bool   // will not allow separators
	Type         *string // mac address type
}

func (o *IsMacAddressOpts) mergeDefaults() {
	if o.NoSeparators == nil {
		o.NoSeparators = Bool(false)
	}
}

// A validator that checks if the string is a MAC address.
//
// IsMacAddressOpts is a struct which defaults to { NoSeparators: false, Type: nil }.
//
// If NoSeparators is true, the validator will allow MAC addresses without separators.
//
// Also, it allows the use of hyphens, spaces or dots e.g. "01 02 03 04 05 ab", "01-02-03-04-05-ab" or "0102.0304.05ab".
//
// The Type is a pointer to "48" or "64", defaults to "48"
//
//	ok, _ := validatorgo.IsMacAddress("00:1A:2B:3C:4D:5E", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsMacAddress("00:1A:2B:3C:4D:ZZ", nil)
//	fmt.Println(ok) // false
func IsMacAddress(str string, opts *IsMacAddressOpts) (bool, error) {
	if opts == nil {
		opts = &IsMacAddressOpts{}
	}
	opts.mergeDefaults()

	eu48, eu64 := "48", "64"

	if opts.Type == nil {
		opts.Type = strPtr(eu48)
	}

	if *opts.Type != eu48 && *opts.Type != eu64 {
		return false, newValidationError("IsMacAddress", ErrInvalidFormat, "invalid macaddress")
	}

	noSepReStr := `[\s:.-]?`
	if *opts.NoSeparators {
		noSepReStr = ""
	}

	typeReStr := "5"
	if *opts.Type == *strPtr(eu64) {
		typeReStr = "7"
	}

	re := regexp.MustCompile(fmt.Sprintf(`^([[:xdigit:]]{2}%s){%s}[[:xdigit:]]{2}$`, noSepReStr, typeReStr))

	if re.MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsMacAddress", ErrInvalidFormat, "invalid macaddress")
}

