package validatorgo

import (
	"fmt"
	"regexp"
)

type IsMacAddressOpts struct {
	NoSeparators bool
	Type         *string
}

var (
	EUIType48 = "48"
	EUIType64 = "64"
)

// A validator that checks if the string is a MAC address.
// IsMacAddressOpts is a struct which defaults to { NoSeparators: false, Type: nil }.
// If NoSeparators is true, the validator will allow MAC addresses without separators.
// Also, it allows the use of hyphens, spaces or dots e.g. '01 02 03 04 05 ab', '01-02-03-04-05-ab' or '0102.0304.05ab'.
// The options also allow a eui property to specify if it needs to be validated against EUI-48 or EUI-64.
// The accepted values of eui are: "48", "64". I have exported EUIType48 and EUIType64. Defaults to "48" if invalid types are passed.
func IsMacAddress(str string, opts IsMacAddressOpts) bool {

	noSepReStr := `[\s:.-]?`
	if opts.NoSeparators {
		noSepReStr = ""
	}

	typeReStr := "5"
	if opts.Type == &EUIType64 {
		typeReStr = "7"
	}

	// fmt.Println(fmt.Sprintf(`^([[:xdigit:]]{2}%s){%s}[[:xdigit:]]{2}$`, noSepReStr, typeReStr))
	re := regexp.MustCompile(fmt.Sprintf(`^([[:xdigit:]]{2}%s){%s}[[:xdigit:]]{2}$`, noSepReStr, typeReStr))

	return re.MatchString(str)
}
