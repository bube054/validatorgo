package validatorgo

import (
	"strconv"
	"strings"
)

func ptr(i uint) *uint {
	return &i
}

func intPtr(i int) *int {
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

func validYearMonthDay(year, month, day string) bool {
	yr, err := strconv.Atoi(year)

	if err != nil {
		return false
	}

	mn, err := strconv.Atoi(month)

	if err != nil {
		return false
	}

	dy, err := strconv.Atoi(day)

	if err != nil {
		return false
	}

	monthLength := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Adjust for leap years
	if yr%400 == 0 || (yr%100 != 0 && yr%4 == 0) {
		monthLength[1] = 29
	}

	if !(mn > 0 && mn < 13) {
		return false
	}

	if dy < 0 || dy > monthLength[mn-1] {
		return false
	}

	return true
}