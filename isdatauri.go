package validatorgo

import (
	"regexp"
)

// A validator that checks if the string is a data uri format.
//
//	ok := validatorgo.IsDataURI("data:,Hello%2C%20World%21")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsDataURI("text/plain;base64,SGVsbG8sIFdvcmxkIQ==")
//	fmt.Println(ok) // false
func IsDataURI(str string) bool {
	re := regexp.MustCompile(`^data:([-\w]+\/[-+\w.]+)?((?:;?[\w]+=[-\w]+)*)(;base64)?,(.*)$`)

	capGrp := re.FindStringSubmatch(str)

	if len(capGrp) == 0 {
		return false
	}

	mimeTyp := capGrp[1]
	basePrt := capGrp[4]

	return IsBase64(basePrt, &IsBase64Opts{UrlSafe: true}) && IsMimeType(mimeTyp)
}
