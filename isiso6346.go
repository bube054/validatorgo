package validatorgo

import (
	"math"
	"regexp"
	"strconv"
)

// iso6346numValues is the set of alphabets and their iso6346 numerical values
var iso6346numValues = map[string]int{
	"A": 10, "B": 12, "C": 13, "D": 14, "E": 15, "F": 16, "G": 17, "H": 18, "I": 19, "J": 20, "K": 21, "L": 23, "M": 24,
	"N": 25, "O": 26, "P": 27, "Q": 28, "R": 29, "S": 30, "T": 31, "U": 32, "V": 34, "W": 35, "X": 36, "Y": 37, "Z": 38,
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

// A validator that checks if the string is a valid ISO 6346 shipping container identification.
func IsISO6346(str string) bool {
	re := regexp.MustCompile(`^([A-Z]{3})([UJZR])(\d{6})(\d)$`)
	capGrps := re.FindStringSubmatch(str)
	if capGrps == nil {
		return false
	}

	checkDig := capGrps[4]
	len := len(str)
	sum := 0
	for ind, char := range str {
		if len-1 == ind {
			break
		}

		numVal, ok := iso6346numValues[string(char)]

		if !ok {
			return false
		}

		sum += int(float64(numVal) * math.Pow(2.00, float64(ind)))
	}

	rem := sum % 11
	remStr := strconv.Itoa(rem)

	return remStr == checkDig
}
