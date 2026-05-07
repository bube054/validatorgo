package validatorgo

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	isISSNOptsDefaultRequireHyphen        bool = false
	isISSNOptsDefaultRequireCaseSensitive bool = false
)

func setIsISSNOptsToDefault() *IsISSNOpts {
	return &IsISSNOpts{
		RequireHyphen: isISSNOptsDefaultRequireHyphen,
		CaseSensitive: isISSNOptsDefaultRequireCaseSensitive,
	}
}

// IsISSNOpts is used to configure IsISSN
type IsISSNOpts struct {
	RequireHyphen bool // requires a hyphen
	CaseSensitive bool // must be exact case
}

// A validator that checks if the string is an [ISSN].
//
// IsISSNOpts is a struct which defaults to { CaseSensitive: false, RequireHyphen: false }.
//
// If CaseSensitive is true, ISSNs with a lowercase "x" as the check digit are rejected.
//
//	ok, _ := validatorgo.IsISSN("0378-5955", nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsISSN("1234567", nil)
//	fmt.Println(ok) // false
//
// [ISSN]: https://en.wikipedia.org/wiki/International_Standard_Serial_Number
func IsISSN(str string, opts *IsISSNOpts) (bool, error) {
	if opts == nil {
		opts = setIsISSNOptsToDefault()
	}

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

	if len(capGrp) == 0 {
		return false, newValidationError("IsISSN", ErrInvalidFormat, "invalid issn")
	}

	strWithoutDashes := stripDashesAndSpaces(str)
	length := len(strWithoutDashes)

	sum := 0
	checkVal := ""
	for ind, char := range strWithoutDashes {
		charVal := string(char)

		if length-1 == ind {
			checkVal = charVal
			break
		}

		pos := length - ind
		charInt, err := strconv.Atoi(charVal)

		if err != nil {
			return false, newValidationError("IsISSN", ErrInvalidFormat, "invalid issn")
		}

		// fmt.Printf("%d x %d\n", charInt, pos)
		sum += pos * charInt
	}

	// fmt.Printf("Sum: %d\n", sum)

	checkValIsNotX := checkVal != "X" && checkVal != "x"
	checkDig, err := strconv.Atoi(checkVal)

	if checkValIsNotX && err != nil {
		return false, newValidationError("IsISSN", ErrInvalidFormat, "invalid issn")
	}

	rem := sum % 11

	if rem == 0 {
		// fmt.Printf("There is no rem and check digit: %d", checkDig)
		if checkDig == 0 {
			return true, nil
		}
		return false, newValidationError("IsISSN", ErrInvalidFormat, "invalid issn")
	}

	remSub11 := 11 - rem

	if remSub11 < 10 {
		// fmt.Printf("rem sub is less than 10 check digit: %d, remSub: %d\n", checkDig, remSub11)
		if remSub11 == checkDig {
			return true, nil
		}
		return false, newValidationError("IsISSN", ErrInvalidFormat, "invalid issn")
	}

	// fmt.Printf("rem sub is greater than 10 check val: %s", checkVal)
	if checkVal == "X" || checkVal == "x" {
		return true, nil
	}
	return false, newValidationError("IsISSN", ErrInvalidFormat, "invalid issn")
}
