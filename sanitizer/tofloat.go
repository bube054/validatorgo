package sanitizer

import "strconv"

// A sanitizer that converts the input string to a float64 and also returns an error if the input is not a float.
//
//	flt := sanitizer.ToFloat("123.45")
//	fmt.Println(flt) // 123.45
func ToFloat(str string) (float64, error) {
	float, err := strconv.ParseFloat(str, 64)

	return float, err
}