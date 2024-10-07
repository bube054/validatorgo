package validatorgo

import "regexp"

var uuidRegex = map[string]*regexp.Regexp{
	"1": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
	"2": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[2][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
	"3": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[3][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
	"4": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[4][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
	"5": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
	// "6": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[6][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
	// "7": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[7][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
	// "8": regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[8][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`),
}

// A validator that checks if the string is an RFC9562 UUID.
//
// version is one of ("1"-"5"). if none is not provided, it will validate any of them.
//
//	ok := validatorgo.IsUUID("550e8400-e29b-11d4-a716-446655440000", "1")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsUUID("f47ac10b-58cc-4372-a567-0e02b2c3d479", "1")
//	fmt.Println(ok) // false
func IsUUID(str, version string) bool {
	re, exists := uuidRegex[version]

	if !exists {
		if version != "" {
			return false
		}

		for _, uiRe := range uuidRegex {
			match := uiRe.MatchString(str)

			if match {
				return true
			}
		}

		return false
	}

	return re.MatchString(str)
}
