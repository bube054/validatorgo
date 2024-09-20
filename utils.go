package validatorgo

import "strings"

func ptr(i uint) *uint {
	return &i
}

func floatPtr(f float64) *float64 {
	return &f
}

func stripDashesAndSpaces(str string) string {
	strWthOutDashes := strings.ReplaceAll(str, "-", "")
	strWthOutSpacesAndDashes := strings.ReplaceAll(strWthOutDashes, " ", "")

	return strWthOutSpacesAndDashes
}

func stripDashes(str string) string {
	return strings.ReplaceAll(str, "-", "")
}

func stripSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func stripHyphens(str string) string {
	return strings.ReplaceAll(str, "-", "")
}

// digitSum returns the sum of digits in an int.
func digitSum(i int) (sum int) {
	for {
		sum += i % 10
		i /= 10
		if i == 0 {
			break
		}
	}
	return
}