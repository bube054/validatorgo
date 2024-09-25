package validatorgo

import (
	"regexp"
)

// A validator that checks if the string is a valid [RFC 3339] date.
//
//	ok := validatorgo.IsRFC3339("2024-09-21T14:30:00Z")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsRFC3339("2024-09-21 14:30:00Z")
//	fmt.Println(ok) // false
//
// [RFC 3339]: https://tools.ietf.org/html/rfc3339
func IsRFC3339(str string) bool {
	re := regexp.MustCompile(`^(\d{4})-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])T([01]\d|2[0-3]):[0-5]\d:[0-5]\d(Z|([+-](0[0-9]|1[0-3]):[0-5]\d))$`)

	capGrp := re.FindStringSubmatch(str)

	if len(capGrp) == 0 {
		return false
	}

	year, month, day := capGrp[1], capGrp[2], capGrp[3]

	return validYearMonthDay(year, month, day)
}
