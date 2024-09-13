package validatorgo

import "strconv"

// A validator that checks if the string passes the Luhn algorithm check.
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
