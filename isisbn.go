package govalidator

import (
	// "fmt"
	"strconv"
	"strings"
)

// A validator that checks if the string is an ISBN.
// version: ISBN version to compare to. Accepted values are '10' and '13'. If none provided, both will be tested.
func IsISBN(str, version string) bool {
	strNum := stripDashesAndSpaces(str)

	if version == "10" {
		return valIsISBNv10(strNum)
	} else if version == "13" {
		return valIsISBNv13(strNum)
	} else {
		return valIsISBNv10(strNum) || valIsISBNv13(strNum)
	}
}

func valIsISBNv10(str string) bool {
	len := len(str)
	sum := 0

	if len != 10 {
		return false
	}

	for i, char := range str {
		pos := len - i
		num, err := strconv.Atoi(string(char))

		if err != nil {
			return false
		}

		sum += pos * num
	}

	rem := sum % 11

	return rem == 0
}

func valIsISBNv13(str string) bool {
	len := len(str)
	sum := 0

	if len != 13 {
		return false
	}

	for i, char := range str {
		pos := len - i
		num, err := strconv.Atoi(string(char))

		if err != nil {
			return false
		}

		if pos%2 == 0 {
			sum += 3 * num
			// fmt.Printf("3 * %d\n", num)
		} else {
			sum += 1 * num
			// fmt.Printf("1 * %d\n", num)
		}
	}

	// fmt.Println("Sum is", sum)

	rem := sum % 10

	return rem == 0
}

func stripDashesAndSpaces(str string) string {
	strWthOutDashes := strings.ReplaceAll(str, "-", "")
	strWthOutSpacesAndDashes := strings.ReplaceAll(strWthOutDashes, " ", "")

	return strWthOutSpacesAndDashes
}
