package validatorgo

// Bool returns a pointer to the given bool value.
// Use this when setting pointer fields on option structs.
//
//	IsEmail("x@y.com", &IsEmailOpts{AllowUTF8LocalPart: Bool(false)})
func Bool(v bool) *bool {
	return &v
}

// String returns a pointer to the given string value.
func String(v string) *string {
	return &v
}

// Int returns a pointer to the given int value.
func Int(v int) *int {
	return &v
}

// Uint returns a pointer to the given uint value.
func Uint(v uint) *uint {
	return &v
}

// Float64 returns a pointer to the given float64 value.
func Float64(v float64) *float64 {
	return &v
}
