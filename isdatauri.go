package validatorgo

import "regexp"

// A validator that checks if the string is a data uri format.
func IsDataURI(str string) bool {
	return regexp.MustCompile(`data:([-\w]+\/[-+\w.]+)?(;?\w+=[-\w]+)*(;base64)?,.*`).MatchString(str)
}
