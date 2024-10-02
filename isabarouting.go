package validatorgo

import (
	"strconv"
	"unicode/utf8"
)

// A validator that checks if the string is an ABA routing number for US bank account / cheque.
//
//	ok := govalidator.IsAbaRouting("123456789")
//	fmt.Println(ok) // false
//	ok := govalidator.IsAbaRouting("021000021")
//	fmt.Println(ok) // true
func IsAbaRouting(str string) bool {
	strWithoutDashes := stripDashesAndSpaces(str)

	if utf8.RuneCountInString(strWithoutDashes) != 9 || !IsNumeric(strWithoutDashes, IsNumericOpts{NoSymbols: true}) {
		return false
	}

	digits := make([]int, 9)

	for i, char := range strWithoutDashes {
		digit, _ := strconv.Atoi(string(char))
		// err is ignore because we are assured IsNumericAbove to block all non numbers
		// if err != nil {
		// 	return false
		// }
		digits[i] = digit
	}

	checksum := 3*(digits[0]+digits[3]+digits[6]) + 7*(digits[1]+digits[4]+digits[7]) + 1*(digits[2]+digits[5]+digits[8])

	return checksum%10 == 0
}
