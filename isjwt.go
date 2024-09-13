package validatorgo

import (
	"encoding/base64"
	"regexp"
)

// A validator that checks if the string is valid JWT token.
func IsJWT(str string) bool {
	re := regexp.MustCompile(`^([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`)
	capGrp := re.FindStringSubmatch(str)

	if capGrp == nil {
		return false
	}

	payload := capGrp[2]

	rawDecodedTxt, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return false
	}
	jsonDecodedTxt := string(rawDecodedTxt)

	return IsJSON(jsonDecodedTxt)
}
