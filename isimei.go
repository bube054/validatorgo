package validatorgo

import (
	"regexp"
	"strconv"
)

// IsIMEIOpts is used to configure IsIMEI
type IsIMEIOpts struct {
	AllowHyphens *bool
}

func (o *IsIMEIOpts) mergeDefaults() {
	if o.AllowHyphens == nil {
		o.AllowHyphens = Bool(false)
	}
}

// A validator that checks if the string is a valid [IMEI number]. IMEI should be of format ############### or ##-######-######-#.
//
// IsIMEIOpts is a struct which can contain the keys AllowHyphens. Defaults to first format.
//
// If AllowHyphens is set to true, the validator will validate the second format.
//
//	ok, _ := validatorgo.IsIMEI("490154203237518", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsIMEI("359043377500085", nil)
//	fmt.Println(ok) // false
//
// [IMEI number]: https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity
func IsIMEI(str string, opts *IsIMEIOpts) (bool, error) {
	if opts == nil {
		opts = &IsIMEIOpts{}
	}
	opts.mergeDefaults()

	var re *regexp.Regexp

	if *opts.AllowHyphens {
		re = regexp.MustCompile(`^\d{2}-?\d{6}-?\d{6}-?\d$`)
	} else {
		re = regexp.MustCompile(`^\d{15}$`)
	}

	if !re.MatchString(str) {
		return false, newValidationError("IsIMEI", ErrInvalidFormat, "string does not match IMEI format")
	}

	strWithoutHyphens := stripHyphens(str)

	strLen := len(strWithoutHyphens)
	isSecond := false
	sum := 0

	for i := strLen - 1; i >= 0; i-- {
		dig, _ := strconv.Atoi(string(strWithoutHyphens[i]))

		if isSecond {
			dbDig := dig * 2

			if dbDig > 9 {
				dbDig = digitSum(dbDig)
			}

			sum += dbDig
		} else {
			sum += dig
		}

		isSecond = !isSecond
	}

	if sum%10 != 0 {
		return false, newValidationError("IsIMEI", ErrInvalidChecksum, "IMEI checksum is invalid")
	}
	return true, nil
}

