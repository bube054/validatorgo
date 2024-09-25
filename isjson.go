package validatorgo

import "encoding/json"

// type IsJSONOpts struct {
// AllowPrimitives bool
// }

// A validator that checks if the string is valid JSON (note: uses json.Valid()).
//
//	ok := validatorgo.IsJSON(`{"name": "John", "age": 30, "city": "New York"}`)
//	fmt.Println(ok) // true
//	ok := validatorgo.IsJSON(`{'name': 'John', 'age': 30}`)
//	fmt.Println(ok) // false
func IsJSON(str string) bool {
	return json.Valid([]byte(str))
}
