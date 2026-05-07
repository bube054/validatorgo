package validatorgo

import (
	"math"
	"strconv"
)

// A validator thats checks if the string is a number(integer not a floating point) that is divisible by another(integer not a floating point).
//
//	ok, _ := validatorgo.IsDivisibleBy("10", 2)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsDivisibleBy("10", 3)
//	fmt.Println(ok) // false
func IsDivisibleBy(str string, num int) (bool, error) {
	if num == 0 {
		return false, newValidationError("IsDivisibleBy", ErrInvalidFormat, "invalid divisibleby")
	}

	strInt, err := strconv.Atoi(str)

	if err != nil {
		return false, newValidationError("IsDivisibleBy", ErrInvalidFormat, "invalid divisibleby")
	}

	if math.Abs(float64(strInt%num)) == float64(0) {
		return true, nil
	}
	return false, newValidationError("IsDivisibleBy", ErrInvalidFormat, "invalid divisibleby")
}
