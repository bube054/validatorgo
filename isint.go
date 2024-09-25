package validatorgo

import (
	"fmt"
	"regexp"
	"strconv"
)

type IsIntOpts struct {
	Min *int // minimum integer
	Max *int // maximum integer

	Gt *int // integer to exceeds
	Lt *int // integer to subceed

	AllowLeadingZeroes bool
}

// A validator that checks if the string is an integer.
//
// IsIntOpts is a struct which can contain the keys Min and/or Max to check the integer is within boundaries (e.g. { Min: nil, Max: nil }). 
//
// IsIntOpts can also contain the key AllowLeadingZeroes, which when set to false will disallow integer values with leading zeroes (e.g. { AllowLeadingZeroes: false }). 
//
// Finally, IsIntOpts can contain the keys Gt and/or Lt which will enforce integers being greater than or less than, respectively, the value provided (e.g. {Gt: ptr(1), Lt: ptr(4)} for a number between 1 and 4).
//
//	ok := validatorgo.IsInt("123", IsIntOpts{})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsInt("123.45", IsIntOpts{})
//	fmt.Println(ok) // false
func IsInt(str string, opts IsIntOpts) bool {
	var matches bool

	if opts.AllowLeadingZeroes {
		matches = regexp.MustCompile(`^([+-]?0*\d+)(\.0+)?$`).MatchString(str)
	} else {
		matches = regexp.MustCompile(`^([+-]?)((0|[1-9]\d*)(\.0*)?|0)$`).MatchString(str)
	}

	if !matches {
		return false
	}

	strInt, err := strconv.Atoi(str)

	if err != nil {
		// fmt.Println("failed parsing")
		return false
	}

	// fmt.Println("Passes regex and conversion")

	withinLimits := true

	if opts.Min != nil {
		// fmt.Println("within min")
		isMin := *(opts.Min) <= strInt
		withinLimits = withinLimits && isMin
	}

	if opts.Max != nil {
		// fmt.Println("within max")
		isMax := *(opts.Max) >= strInt
		withinLimits = withinLimits && isMax
	}

	fmt.Println(strInt, withinLimits)

	if opts.Lt != nil {
		isMin := strInt < *(opts.Lt)
		withinLimits = withinLimits && isMin
		// fmt.Println("within lt", withinLimits)
	}

	if opts.Gt != nil {
		isMax := strInt > *(opts.Gt)
		withinLimits = withinLimits && isMax
		// fmt.Println("within gt", withinLimits)
	}

	return withinLimits
}
