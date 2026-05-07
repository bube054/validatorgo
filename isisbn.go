package validatorgo

import (
	// "fmt"
	"strconv"
)

// A validator that checks if the string is an [ISBN].
//
// version: ISBN version to compare to. Accepted values are "10" and "13". If none provided, both will be tested.
//
//	ok, _ := validatorgo.IsISBN("0-7167-0344-0", "10")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsISBN("0-7168-0344-0", "10")
//	fmt.Println(ok) // false
//
// [ISBN]: https://en.wikipedia.org/wiki/ISBN
func IsISBN(str, version string) (bool, error) {
	strNum := stripDashesAndSpaces(str)

	if version == "10" {
		ok, _ := valIsISBNv10(strNum)
		if ok {
			return true, nil
		}
		return false, newValidationError("IsISBN", ErrInvalidFormat, "invalid isbn")
	} else if version == "13" {
		ok, _ := valIsISBNv13(strNum)
		if ok {
			return true, nil
		}
		return false, newValidationError("IsISBN", ErrInvalidFormat, "invalid isbn")
	} else {
		ok10, _ := valIsISBNv10(strNum)
		ok13, _ := valIsISBNv13(strNum)
		if ok10 || ok13 {
			return true, nil
		}
		return false, newValidationError("IsISBN", ErrInvalidFormat, "invalid isbn")
	}
}

func valIsISBNv10(str string) (bool, error) {
	ln := len(str)
	sum := 0

	if ln != 10 {
		return false, newValidationError("valIsISBNv10", ErrInvalidFormat, "invalid valisisbnv10")
	}

	for i, char := range str {
		pos := ln - i
		num, err := strconv.Atoi(string(char))

		if err != nil {
			return false, newValidationError("valIsISBNv10", ErrInvalidFormat, "invalid valisisbnv10")
		}

		sum += pos * num
	}

	rem := sum % 11

	if rem == 0 {
		return true, nil
	}
	return false, newValidationError("valIsISBNv10", ErrInvalidFormat, "invalid valisisbnv10")
}

func valIsISBNv13(str string) (bool, error) {
	ln := len(str)
	sum := 0

	if ln != 13 {
		return false, newValidationError("valIsISBNv13", ErrInvalidFormat, "invalid valisisbnv13")
	}

	for i, char := range str {
		pos := ln - i
		num, err := strconv.Atoi(string(char))

		if err != nil {
			return false, newValidationError("valIsISBNv13", ErrInvalidFormat, "invalid valisisbnv13")
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

	if rem == 0 {
		return true, nil
	}
	return false, newValidationError("valIsISBNv13", ErrInvalidFormat, "invalid valisisbnv13")
}
