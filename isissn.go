package validatorgo

import (
	"fmt"
	"regexp"
	"strconv"
)

type IsISSNOpts struct {
	RequireHyphen bool
	CaseSensitive bool
}

// A validator that checks if the string is an ISSN.
// IsISSNOpts is a struct which defaults to { CaseSensitive: false, RequireHyphen: false }.
// If CaseSensitive is true, ISSNs with a lowercase 'x' as the check digit are rejected.
func IsISSN(str string, opts IsISSNOpts) bool {
	var reqHypStr string
	if !opts.RequireHyphen {
		reqHypStr = "?"
	}

	var reqCasSenStr string
	if !opts.CaseSensitive {
		reqCasSenStr = "x"
	}

	re := regexp.MustCompile(fmt.Sprintf(`^[0-9]{4}-%s[0-9]{3}[0-9X%s]$`, reqHypStr, reqCasSenStr))
	capGrp := re.FindStringSubmatch(str)

	if capGrp == nil {
		return false
	}

	strWithoutDashes := stripDashesAndSpaces(str)
	len := len(strWithoutDashes)

	sum := 0
	checkVal := ""
	for ind, char := range strWithoutDashes {
		charVal := string(char)

		if len-1 == ind {
			checkVal = charVal
			break
		}

		pos := len - ind
		charInt, err := strconv.Atoi(charVal)

		if err != nil {
			return false
		}

		// fmt.Printf("%d x %d\n", charInt, pos)
		sum += pos * charInt
	}

	// fmt.Printf("Sum: %d\n", sum)

	checkValIsNotX := checkVal != "X" && checkVal != "x"
	checkDig, err := strconv.Atoi(checkVal)

	if checkValIsNotX && err != nil {
		return false
	}

	rem := sum % 11

	if rem == 0 {
		// fmt.Printf("There is no rem and check digit: %d", checkDig)
		return checkDig == 0
	}

	remSub11 := 11 - rem

	if remSub11 < 10 {
		// fmt.Printf("rem sub is less than 10 check digit: %d, remSub: %d\n", checkDig, remSub11)
		return remSub11 == checkDig
	}

	// fmt.Printf("rem sub is greater than 10 check val: %s", checkVal)
	return checkVal == "X" || checkVal == "x"
}
