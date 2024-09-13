package validatorgo

// A validator that checks if the string consists only of characters that appear in the whitelist chars.
func IsWhitelisted(str, chars string) bool {
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
			return false
		}

	}

	return true
}
