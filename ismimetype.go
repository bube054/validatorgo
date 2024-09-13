package validatorgo

import "regexp"

// A validator that checks if the string matches to a valid MIME type format.
func IsMimeType(str string) bool {
	return regexp.MustCompile(`^(application|audio|font|image|message|model|multipart|text|video|x-[\w.+-]+)\/[\w.+-]+(\s*;\s*[\w.+-]+=[\w.+-]+)*$`).MatchString(str)
}
