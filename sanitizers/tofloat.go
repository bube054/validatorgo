package sanitizers

import "strconv"

func ToFloat(str string) float64 {
	feetFloat, _ := strconv.ParseFloat(str[:len(str)-1], 64)

	return feetFloat
}

// fmt.Sprintf("%.2f", fResult)
