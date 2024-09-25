package validatorgo

import "regexp"

// A validator that checks if the string is a Semantic Versioning Specification (SemVer).
//
//	ok := validatorgo.IsSemVer("1.0.0")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsSemVer("1.0.0.0")
//	fmt.Println(ok) // false
func IsSemVer(str string) bool {
	return regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`).MatchString(str)
}
