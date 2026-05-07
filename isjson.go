package validatorgo

import "encoding/json"

// type IsJSONOpts struct {
// AllowPrimitives bool
// }

// A validator that checks if the string is valid JSON (note: uses json.Valid()).
//
//	ok, _ := validatorgo.IsJSON(`{"name": "John", "age": 30, "city": "New York"}`)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsJSON(`{'name': 'John', 'age': 30}`)
//	fmt.Println(ok) // false
func IsJSON(str string) (bool, error) {
	if json.Valid([]byte(str)) {
		return true, nil
	}
	return false, newValidationError("IsJSON", ErrInvalidFormat, "invalid json")
}
