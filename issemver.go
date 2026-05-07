package validatorgo

import "regexp"

// A validator that checks if the string is a Semantic Versioning Specification (SemVer).
//
//	ok, _ := validatorgo.IsSemVer("1.0.0")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsSemVer("1.0.0.0")
//	fmt.Println(ok) // false
func IsSemVer(str string) (bool, error) {
	if regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`).MatchString(str) {
		return true, nil
	}
	return false, newValidationError("IsSemVer", ErrInvalidFormat, "invalid semver")
}
