package validatorgo

import (
	"strconv"
)

// IsEAN checks if the string is a valid [EAN] (European Article Number).
//
//	ok := validatorgo.IsDecimal("4006381333931")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsDecimal("123456789012")
//	fmt.Println(ok) // false
//
// [EAN]: https://en.wikipedia.org/wiki/International_Article_Number
func IsEAN(str string) bool {
	length := len(str)
	if length != 8 && length != 13 {
		return false
	}

	// Check if all characters are digits
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
	}

	// Calculate checksum
	var sum int
	for i, r := range str[:length-1] {
		digit, _ := strconv.Atoi(string(r))
		if (length == 8 && i%2 == 0) || (length == 13 && i%2 != 0) {
			sum += digit * 3
		} else {
			sum += digit
		}
	}

	// The checksum digit is the last digit of the EAN
	checksum := (10 - (sum % 10)) % 10
	lastDigit, _ := strconv.Atoi(string(str[length-1]))

	return checksum == lastDigit
}
