package validatorgo

import (
	"regexp"
	"strconv"
)

// A validator that checks if the string is an HSL (hue, saturation, lightness, optional alpha) color based on CSS Colors Level 4 specification.
//
//	ok := validatorgo.IsHSL("hsl(360, 100%, 50%)")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsHSL("hsl(360, 100%)")
//	fmt.Println(ok) // false
func IsHSL(str string) bool {
	re := regexp.MustCompile(`^hsl(a)?\(\s?([^-+].*),\s?([^-+].*)%,\s?([^-+].*)%(,\s?([^-+].*))?\)$`)
	capGrp := re.FindStringSubmatch(str)

	if len(capGrp) == 0 {
		return false
	}

	a := capGrp[1]        // "a"
	hueVal := capGrp[2]   // "360"
	satVal := capGrp[3]   // "50"
	lightVal := capGrp[4] // "50"
	alphaVal := capGrp[6] // "0.5"

	if hueFlt, err := strconv.ParseFloat(hueVal, 64); err != nil {
		return false
	} else {
		if hueFlt < 0 || hueFlt > 360 {
			return false
		}
	}

	if satFlt, err := strconv.ParseFloat(satVal, 64); err != nil {
		return false
	} else {
		if satFlt < 0 || satFlt > 100 {
			return false
		}
	}

	if lightFlt, err := strconv.ParseFloat(lightVal, 64); err != nil {
		return false
	} else {
		if lightFlt < 0 || lightFlt > 100 {
			return false
		}
	}

	if alphaVal != "" {
		if a == "" {
			return false
		}

		if alphaVal, err := strconv.ParseFloat(alphaVal, 64); err != nil {
			return false
		} else {
			if alphaVal < 0 || alphaVal > 1 {
				return false
			}
		}
	}

	return true
}
