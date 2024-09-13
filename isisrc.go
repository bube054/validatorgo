package validatorgo

import (
	"fmt"
	"regexp"
)

// A validator that checks if the string is an ISRC.
// allowHyphens will allow codes with dashes present CC-XXX-YY-NNNNN
func IsISRC(str string, allowHyphens bool) bool {
	var char string

	if allowHyphens {
		char = "-?"
	}

	re := regexp.MustCompile(fmt.Sprintf(`^([A-Z]{2})%s([A-Z]{3})%s(\d{2})%s(\d{5})$`, char, char, char))
	capGrp := re.FindStringSubmatch(str)

	if capGrp == nil {
		return false
	}

	cntryCode := capGrp[1]

	return IsISO31661Alpha2(cntryCode)
}
