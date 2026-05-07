package validatorgo

import (
	"net/url"
	"regexp"
	"strings"
)

var (
	validAttribute = regexp.MustCompile(`(?i)^[a-z\-]+=[\w\-]+$`)
	validData      = regexp.MustCompile(`(?i)^[a-z0-9!\$&'()*+,;=\-._~:@/?%\s]*$`)
)

// A validator that checks if the string is a data uri format.
//
//	ok, _ := validatorgo.IsDataURI("data:,Hello%2C%20World%21")
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsDataURI("text/plain;base64,SGVsbG8sIFdvcmxkIQ==")
//	fmt.Println(ok) // false
func IsDataURI(str string) (bool, error) {
	invalidErr := newValidationError("IsDataURI", ErrInvalidFormat, "invalid datauri")

	commaIdx := strings.Index(str, ",")
	if commaIdx < 0 {
		return false, invalidErr
	}

	head := str[:commaIdx]
	data := str[commaIdx+1:]

	attributes := strings.Split(head, ";")
	schemeAndMediaType := attributes[0]

	if !strings.HasPrefix(schemeAndMediaType, "data:") {
		return false, invalidErr
	}

	mediaType := schemeAndMediaType[5:]
	if mediaType != "" {
		okMime, _ := IsMimeType(mediaType)
		if !okMime {
			return false, invalidErr
		}
	}

	isBase64 := false
	for i := 1; i < len(attributes); i++ {
		if i == len(attributes)-1 && strings.EqualFold(attributes[i], "base64") {
			isBase64 = true
			continue
		}
		if !validAttribute.MatchString(attributes[i]) {
			return false, invalidErr
		}
	}

	if isBase64 {
		decodedData, err := url.PathUnescape(data)
		if err != nil {
			return false, invalidErr
		}
		okBase64, _ := IsBase64(decodedData, nil)
		if !okBase64 {
			return false, invalidErr
		}
	} else {
		if data == "" {
			return false, invalidErr
		}
		if !validData.MatchString(data) {
			return false, invalidErr
		}
	}

	return true, nil
}
