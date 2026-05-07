package validatorgo

// A validator that checks if the string consists only of characters that appear in the whitelist chars.
//
//	ok, _ := validatorgo.IsWhitelisted("stop", "post")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsWhitelisted("bang", "take")
//	fmt.Println(ok) // false
func IsWhitelisted(str, chars string) (bool, error) {
	charsM := make(map[string]int)

	for _, char := range chars {
		val := string(char)
		_, exist := charsM[val]

		if exist {
			charsM[val] += 1
		} else {
			charsM[val] = 1
		}
	}

	for _, st := range str {
		val := string(st)

		_, exist := charsM[val]

		if !exist {
			return false, newValidationError("IsWhitelisted", ErrBlacklistedChar, "character not in whitelist")
		}

	}

	return true, nil
}
