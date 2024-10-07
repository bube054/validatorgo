package validatorgo

import (
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

// A validator that checks alias for IsISO6346, check if the string is a valid [ISO 6346] shipping container identification.
//
//	ok := validatorgo.IsFreightContainerID("ABCU1234567")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsFreightContainerID("AB123456789")
//	fmt.Println(ok) // false
//
// [ISO 6346]: https://en.wikipedia.org/wiki/ISO_6346
func IsFreightContainerID(str string) bool {
	// Check if length is 11
	if utf8.RuneCountInString(str) != 11 {
		return false
	}

	re := regexp.MustCompile(`^([A-Z]{3})([UJZ])([0-9]{6})([0-9])$`)
	match := re.MatchString(str)
	if !match {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		char := string(str[i])
		mag, ok := alphaFreightNumVal[char]

		if !ok {
			return false
		}

		mul := math.Pow(2.00, float64(i))

		sum += int(mul) * mag
	}

	actCheck := sum % 11
	givCheck, err := strconv.Atoi(string(str[10]))

	if err != nil {
		return false
	}

	return actCheck == givCheck
}
