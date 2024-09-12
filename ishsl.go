package validatorgo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// A validator that checks if the string is an HSL (hue, saturation, lightness, optional alpha) color based on CSS Colors Level 4 specification.
func IsHSL(str string) bool {
	re := regexp.MustCompile(`(?im)^hsla?\(\s*(?<h>[-+]?\d{1,3}(?:\.\d+)?)(deg|grad|rad|turn)?\s*,\s*(?<s>[-+]?\d{1,3}(?:\.\d+)?)%\s*,\s*(?<l>[-+]?\d{1,3}(?:\.\d+)?)%\s*(?:,\s*(?<alpha>[-+]?[\d.]+%?)\s*)?\)$
`)

	match := re.MatchString(str)

	if !match {
		return match
	}

	grps := re.FindStringSubmatch(str)
	fmt.Printf("GRPS %+v\n", grps)
	fmt.Println("0", grps[0])
	fmt.Println("1", grps[1])
	fmt.Println("2", grps[2])
	fmt.Println("3", grps[3])
	fmt.Println("4", grps[4])
	fmt.Println("5", grps[5])
	hue, sat, light, alpha := grps[1], grps[3], grps[4], grps[5]

	// fmt.Println("hue", hue)
	// fmt.Println("sat", sat)
	// fmt.Println("light", light)

	const bitSize = 64
	hueNum, err := strconv.ParseFloat(hue, bitSize)

	if err != nil || hueNum < 0 || hueNum > 360 {
		return false
	}

	// satWthOutPer := strings.Replace(sat, "%", "", 1)
	satNum, err := strconv.ParseFloat(sat, bitSize)
	fmt.Println(sat, satNum, err)

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
