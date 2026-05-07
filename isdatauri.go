package validatorgo

import (
	"regexp"
)

// A validator that checks if the string is a data uri format.
//
//	ok, _ := validatorgo.IsDataURI("data:,Hello%2C%20World%21")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsDataURI("text/plain;base64,SGVsbG8sIFdvcmxkIQ==")
//	fmt.Println(ok) // false
func IsDataURI(str string) (bool, error) {
	re := regexp.MustCompile(`^data:([-\w]+\/[-+\w.]+)?((?:;?[\w]+=[-\w]+)*)(;base64)?,(.*)$`)

	capGrp := re.FindStringSubmatch(str)

	if len(capGrp) == 0 {
		return false, newValidationError("IsDataURI", ErrInvalidFormat, "invalid datauri")
	}

	mimeTyp := capGrp[1]
	basePrt := capGrp[4]

	okBase64, _ := IsBase64(basePrt, &IsBase64Opts{UrlSafe: true})
	okMime, _ := IsMimeType(mimeTyp)

	if okBase64 && okMime {
		return true, nil
	}
	return false, newValidationError("IsDataURI", ErrInvalidFormat, "invalid datauri")
}
