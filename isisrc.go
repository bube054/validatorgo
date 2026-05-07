package validatorgo

import (
	"fmt"
	"regexp"
)

// A validator that checks if the string is an [ISRC].
//
// allowHyphens will allow codes with dashes present CC-XXX-YY-NNNNN
//
//	ok, _ := validatorgo.IsISRC("AASKG1912345", false)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsISRC("AA-SKG-19-12345", false)
//	fmt.Println(ok) // false
//
// [ISRC]: https://en.wikipedia.org/wiki/International_Standard_Recording_Code
func IsISRC(str string, allowHyphens bool) (bool, error) {
	var char string

	if allowHyphens {
		char = "-?"
	}

	re := regexp.MustCompile(fmt.Sprintf(`^([A-Z]{2})%s([A-Z0-9]{3})%s(\d{2})%s(\d{5})$`, char, char, char))
	capGrp := re.FindStringSubmatch(str)

	if capGrp == nil {
		return false, newValidationError("IsISRC", ErrInvalidFormat, "invalid isrc")
	}

	cntryCode := capGrp[1]

	okAlpha2, _ := IsISO31661Alpha2(cntryCode)
	if okAlpha2 {
		return true, nil
	}
	return false, newValidationError("IsISRC", ErrInvalidFormat, "invalid isrc")
}
