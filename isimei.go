package validatorgo

import (
	"fmt"
	"regexp"
	"strconv"
)

// IsIMEIOpts is used to configure IsIMEI
type IsIMEIOpts struct {
	AllowHyphens bool
}

// A validator that checks if the string is a valid [IMEI number]. IMEI should be of format ############### or ##-######-######-#.
//
// IsIMEIOpts is a struct which can contain the keys AllowHyphens. Defaults to first format.
//
// If AllowHyphens is set to true, the validator will validate the second format.
//
//	ok := validatorgo.IsIMEI("490154203237518", IsIMEIOpts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsIMEI("359043377500085", IsIMEIOpts{})
//	fmt.Println(ok) // false
//
// [IMEI number]: https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity
func IsIMEI(str string, opts IsIMEIOpts) bool {
	var re *regexp.Regexp

	if opts.AllowHyphens {
		re = regexp.MustCompile(`^\d{2}-?\d{6}-?\d{6}-?\d$`)
	} else {
		re = regexp.MustCompile(`^\d{15}$`)
	}

	if !re.MatchString(str) {
		return false
	}

	strWithoutHyphens := stripHyphens(str)

	fmt.Println(strWithoutHyphens)

	strLen := len(strWithoutHyphens)
	isSecond := false
	// checkDigit := 0
	sum := 0

	for i := strLen - 1; i >= 0; i-- {
		dig, _ := strconv.Atoi(string(strWithoutHyphens[i]))

		// if i == strLen-1 {
		// 	checkDigit = dig
		// }

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

	return sum%10 == 0
}
