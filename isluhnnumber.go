package validatorgo

import "strconv"

// A validator that checks if the string passes the [Luhn algorithm] check.
//
//	ok := validatorgo.IsLuhnNumber("4532015112830366")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsLuhnNumber("4532015112830367")
//	fmt.Println(ok) // false
//
// [Luhn algorithm]: https://en.wikipedia.org/wiki/Luhn_algorithm
func IsLuhnNumber(str string) bool {
	var (
		len      = len(str)
		sum      = 0
		isSecond = false
	)

	for i := len - 1; i >= 0; i-- {
		char := str[i]
		d, err := strconv.Atoi(string(char))

		if err != nil {
			return false
		}

		if isSecond {
			d = d * 2
			if d > 9 {
				d = digitSum(d)
			}
		}

		sum += d
		isSecond = !isSecond
	}

	return sum%10 == 0
}