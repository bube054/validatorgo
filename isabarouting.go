package validatorgo

import (
	"strconv"
	"unicode/utf8"
)

// A validator that checks if the string is an ABA routing number for US bank account / cheque.
//
//	ok := govalidator.IsAbaRouting("123456789")
//	fmt.Println(ok) // false
//	ok = govalidator.IsAbaRouting("021000021")
//	fmt.Println(ok) // true
func IsAbaRouting(str string) bool {
	if utf8.RuneCountInString(str) != 9 || !IsNumeric(str, IsNumericOpts{NoSymbols: true}) {
		return false
	}

	digits := make([]int, 9)

	for i, char := range str {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digits[i] = digit
	}

	checksum := 3*(digits[0]+digits[3]+digits[6]) + 7*(digits[1]+digits[4]+digits[7]) + 1*(digits[2]+digits[5]+digits[8])

	return checksum%10 == 0
}
