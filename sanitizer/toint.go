package sanitizer

import "strconv"

// A sanitizer that converts the input string to an int and also returns an error if the input is not a int. (Beware of octals)
//
//	num := sanitizer.ToInt("123")
//	fmt.Println(num) // 123
func ToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)

	return num, err
}