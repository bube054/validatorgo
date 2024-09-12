package validatorgo

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"unicode/utf8"
)

var alphaFreightNumVal = map[string]int{
	"A": 10,
	"B": 12,
	"C": 13,
	"D": 14,
	"E": 15,
	"F": 16,
	"G": 17,
	"H": 18,
	"I": 19,
	"J": 20,
	"K": 21,
	"L": 23,
	"M": 24,
	"N": 25,
	"O": 26,
	"P": 27,
	"Q": 28,
	"R": 29,
	"S": 30,
	"T": 31,
	"U": 32,
	"V": 34,
	"W": 35,
	"X": 36,
	"Y": 37,
	"Z": 38,
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

// A validator that checks alias for IsISO6346, check if the string is a valid ISO 6346 shipping container identification.
func IsFreightContainerID(str string) bool {
	len := utf8.RuneCountInString(str)
	if len != 11 {
		return false
	}

	re := regexp.MustCompile(`^([A-Z]{3})([UJZ])([0-9]{6})([0-9])$`)
	match := re.MatchString(str)

	if !match {
		return false
	}

	sum := 0
	for ind, char := range str {
		if ind > 9 {
			break
		}

		wgh := int(math.Pow(2.0, float64(9-ind)))
		num, ok := alphaFreightNumVal[string(char)]

		if !ok {
			return false
		}
		fmt.Println(wgh * num)
		sum += wgh * num
	}

	grps := re.FindStringSubmatch(str)
	checkDig := grps[4]
	rem := sum % 11
	fmt.Println("GROUPS", grps)
	fmt.Println("CHECKDIG", checkDig)
	fmt.Println("SUM", sum)
	fmt.Println("REM", rem)

	if rem == 10 && checkDig == "0" {
		return true
	}

	checkDigInt, err := strconv.Atoi(checkDig)
	if err != nil {
		return false
	}

	return checkDigInt == rem
}
