package sanitizer

// A sanitizer that converts the input string to a boolean.
//
// Everything except for "0", "false" and """ returns true.
//
// In strict mode only "1" and "true" return true.
//
//	bool := sanitizer.ToBoolean("", false)
//	fmt.Println(bool) // false
func ToBoolean(str string, strict bool) bool {
	if strict {
		if str == "1" || str == "true" {
			return true
		} else {
			return false
		}
	} else {
		if str == "0" || str == "false" || str == "" {
			return false
		} else {
			return true
		}
	}
}
