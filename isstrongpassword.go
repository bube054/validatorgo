package validatorgo

import (
	"fmt"
	"unicode"
)

// defaults values to configure IsStrongPassword, IsStrongPasswordOpts
const (
	defaultStrongPasswordMinimumLength    int = 8
	defaultStrongPasswordMinimumLowercase int = 1
	defaultStrongPasswordMinimumUppercase int = 1
	defaultStrongPasswordMinimumNumbers   int = 1
	defaultStrongPasswordMinimumSymbols   int = 1

	defaultStrongPasswordPointsPerUnique          int     = 1
	defaultStrongPasswordPointsPerRepeat          float64 = 0.5
	defaultStrongPasswordPointsForContainingLower int     = 10
	defaultStrongPasswordPointsForContainingUpper int     = 10
	defaultStrongPasswordPointsContainingNumber   int     = 10
	defaultStrongPasswordPointsContainingSymbol   int     = 10
)

// IsStrongPasswordOpts is used to configure IsStrongPassword
type IsStrongPasswordOpts struct {
	MinLength    *int // passwords minimum length
	MinLowercase *int // passwords minimum lowercase letters
	MinUppercase *int // passwords minimum uppercase letters
	MinNumbers   *int // passwords minimum numbers
	MinSymbols   *int // passwords minimum symbols

	PointsPerUnique           *int     // minimum points per unique character
	PointsPerRepeat           *float64 // minimum points per repeated character
	PointsForContainingLower  *int     // minimum points per lowercase character
	PointsForContainingUpper  *int     // minimum points per uppercase character
	PointsForContainingNumber *int     // minimum points per numeric character
	PointsForContainingSymbol *int     // minimum points per special character
}

// uppLwrSpecNumChars is used to count/differentiate different classes of characters
type uppLwrSpecNumChars struct {
	uppercase int
	lowercase int
	numbers   int
	symbol    int
	unique    int
	repeated  int
}

// A validator that checks if the string can be considered a strong password or not. Returns the validity and the score. Disregard score if password is not valid.
//
// Default IsStrongPasswordOpts: { MinLength: 8, MinLowercase: 1, MinUppercase: 1, MinNumbers: 1, MinSymbols: 1,  PointsPerUnique: 1, PointsPerRepeat: 0.5, PointsForContainingLower: 10, PointsForContainingUpper: 10, PointsForContainingNumber: 10, PointsForContainingSymbol: 10 }
//
//	ok, score := validatorgo.IsStrongPassword("Password123!", validatorgo.IsStrongPasswordOpts{})
//	fmt.Println(ok, score) // true, 1130
//	ok, score := validatorgo.IsStrongPassword("Password123@", validatorgo.IsStrongPasswordOpts{})
//	fmt.Println(ok, score) // false, 130
func IsStrongPassword(str string, opts IsStrongPasswordOpts) (bool, int) {
	optsWithDefaults := strongPasswordOptsToDefault(opts)

	score := 0
	valid := true

	if len(str) < *optsWithDefaults.MinLength {
		return false, score
	}

	ulsc := cntUppLwrSymNumChars(str)

	fmt.Printf("ulsc %+v\n", ulsc)

	if ulsc.lowercase < *optsWithDefaults.MinLowercase {
		return false, score
	}
	score += ulsc.lowercase * *optsWithDefaults.PointsForContainingLower // 7 * 10

	if ulsc.uppercase < *optsWithDefaults.MinUppercase {
		return false, score
	}
	score += ulsc.uppercase * *optsWithDefaults.PointsForContainingUpper // 1 * 10

	if ulsc.numbers < *optsWithDefaults.MinNumbers {
		return false, score
	}
	score += ulsc.numbers * *optsWithDefaults.PointsForContainingNumber // 3 * 10

	if ulsc.symbol < *optsWithDefaults.MinSymbols {
		return false, score
	}
	score += ulsc.symbol * *optsWithDefaults.PointsForContainingSymbol // 1 * 10

	score += ulsc.unique * *optsWithDefaults.PointsPerUnique // 10 * 1

	score += ulsc.repeated * int(*optsWithDefaults.PointsPerRepeat) // 1 * 0.5

	return valid, score
}

func strongPasswordOptsToDefault(opts IsStrongPasswordOpts) IsStrongPasswordOpts {
	if opts.MinLength == nil {
		opts.MinLength = intPtr(defaultStrongPasswordMinimumLength)
	}

	if opts.MinLowercase == nil {
		opts.MinLowercase = intPtr(defaultStrongPasswordMinimumLowercase)
	}

	if opts.MinUppercase == nil {
		opts.MinUppercase = intPtr(defaultStrongPasswordMinimumUppercase)
	}

	if opts.MinNumbers == nil {
		opts.MinNumbers = intPtr(defaultStrongPasswordMinimumNumbers)
	}

	if opts.MinSymbols == nil {
		opts.MinSymbols = intPtr(defaultStrongPasswordMinimumSymbols)
	}

	if opts.PointsPerUnique == nil {
		opts.PointsPerUnique = intPtr(defaultStrongPasswordPointsPerUnique)
	}

	if opts.PointsPerRepeat == nil {
		opts.PointsPerRepeat = floatPtr(defaultStrongPasswordPointsPerRepeat)
	}

	if opts.PointsForContainingLower == nil {
		opts.PointsForContainingLower = intPtr(defaultStrongPasswordPointsForContainingLower)
	}

	if opts.PointsForContainingUpper == nil {
		opts.PointsForContainingUpper = intPtr(defaultStrongPasswordPointsForContainingUpper)
	}

	if opts.PointsForContainingNumber == nil {
		opts.PointsForContainingNumber = intPtr(defaultStrongPasswordPointsContainingNumber)
	}

	if opts.PointsForContainingSymbol == nil {
		opts.PointsForContainingSymbol = intPtr(defaultStrongPasswordPointsContainingSymbol)
	}

	return opts
}

func cntUppLwrSymNumChars(str string) (ulsc uppLwrSpecNumChars) {
	runeCountMap := make(map[rune]int, len(str))

	for _, char := range str {
		_, ok := runeCountMap[char]

		if ok {
			runeCountMap[char] = runeCountMap[char] + 1
		} else {
			runeCountMap[char] = 1
		}

		if unicode.IsUpper(rune(char)) && unicode.IsLetter(rune(char)) {
			ulsc.uppercase++
		} else if unicode.IsLower(rune(char)) && unicode.IsLetter(rune(char)) {
			ulsc.lowercase++
		} else if unicode.IsNumber(rune(char)) {
			ulsc.numbers++
		} else {
			ulsc.symbol++
		}
	}

	for _, count := range runeCountMap {
		if count > 1 {
			ulsc.repeated++
			// fmt.Println("repeated", string(ru))
		} else {
			ulsc.unique++
			// fmt.Println("unique", string(ru))
		}
	}

	return
}
