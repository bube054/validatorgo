package validatorgo

import "regexp"

// A validator that checks if the string matches to a valid MIME type format.
//
//	ok := validatorgo.IsMimeType("text/plain")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsMimeType("textplain")
//	fmt.Println(ok) // false
func IsMimeType(str string) bool {
	return regexp.MustCompile(`^(application|audio|font|image|message|model|multipart|text|video|x-[\w.+-]+)\/[a-zA-Z][\w.+-]*(\s*;\s*[\w.+-]+=[\w.+-]+)*$`).MatchString(str)
}
