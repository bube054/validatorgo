package validatorgo

import (
	"regexp"
	"strconv"
	"strings"
)

// A validator that checks if the string is an HSL (hue, saturation, lightness, optional alpha) color based on CSS Colors Level 4 specification.
//
//	ok := validatorgo.IsHSL("hsl(360, 100%, 50%)")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsHSL("hsl(360, 100%)")
//	fmt.Println(ok) // false
func IsHSL(str string) bool {
	re := regexp.MustCompile(`(?im)^hsla?\(\s*(?<h>[-+]?\d{1,3}(?:\.\d+)?)(deg|grad|rad|turn)?\s*,\s*(?<s>[-+]?\d{1,3}(?:\.\d+)?)%\s*,\s*(?<l>[-+]?\d{1,3}(?:\.\d+)?)%\s*(?:,\s*(?<alpha>[-+]?[\d.]+%?)\s*)?\)$
`)

	match := re.MatchString(str)

	if !match {
		return match
	}

	grps := re.FindStringSubmatch(str)
	hue, sat, light, alpha := grps[1], grps[3], grps[4], grps[5]

	const bitSize = 64
	hueNum, err := strconv.ParseFloat(hue, bitSize)

	if err != nil || hueNum < 0 || hueNum > 360 {
		return false
	}

	// satWthOutPer := strings.Replace(sat, "%", "", 1)
	satNum, err := strconv.ParseFloat(sat, bitSize)

	if err != nil || satNum < 0 || satNum > 100 {
		return false
	}

	// lightWthOutPer := strings.Replace(light, "%", "", 1)
	lightNum, err := strconv.ParseFloat(light, bitSize)

	if err != nil || lightNum < 0 || lightNum > 100 {
		return false
	}

	var (
		alphaNum float64
		alphaErr error
	)
	if strings.HasSuffix(alpha, "%") {
		alphaWthOutPer := strings.Replace(alpha, "%", "", 1)
		alphaNum, alphaErr = strconv.ParseFloat(alphaWthOutPer, bitSize)
		alphaNum = alphaNum / 100
	} else {
		alphaNum, alphaErr = strconv.ParseFloat(alpha, bitSize)
	}

	if alpha != "" && (alphaErr != nil || alphaNum < 0 || alphaNum > 1) {
		return false
	}

	return true
}
