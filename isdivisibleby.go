package validatorgo

import (
	"math"
	"strconv"
)

// A validator thats checks if the string is a number(integer not a floating point) that is divisible by another(integer not a floating point).
func IsDivisibleBy(str string, num int) bool {
	if num == 0 {
		return false
	}

	strInt, err := strconv.Atoi(str)

	if err != nil {
		return false
	}

	return math.Abs(float64(strInt%num)) == float64(0)
}
