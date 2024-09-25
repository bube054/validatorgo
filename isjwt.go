package validatorgo

import (
	"encoding/base64"
	"regexp"
)

// A validator that checks if the string is valid JWT token.
//
//	ok := validatorgo.IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
//	fmt.Println(ok) // false
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
