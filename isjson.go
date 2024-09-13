package validatorgo

import "encoding/json"

// type IsJSONOpts struct {
// AllowPrimitives bool
// }

// A validator that checks if the string is valid JSON (note: uses json.Valid).
func IsJSON(str string) bool {
	// var js json.RawMessage

	// return json.Unmarshal([]byte(str), &js) == nil
	return json.Valid([]byte(str))
}
