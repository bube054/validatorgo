package validatorgo

import "fmt"

// Error code constants for validation failures.
const (
	ErrInvalidFormat      = "invalid_format"
	ErrTooShort           = "too_short"
	ErrTooLong            = "too_long"
	ErrMissingTLD         = "missing_tld"
	ErrBlacklistedHost    = "blacklisted_host"
	ErrNotWhitelistedHost = "not_whitelisted_host"
	ErrBlacklistedChar    = "blacklisted_char"
	ErrInvalidLocale      = "invalid_locale"
	ErrOutOfRange         = "out_of_range"
	ErrInvalidChecksum    = "invalid_checksum"
	ErrInvalidLength      = "invalid_length"
	ErrMissingRequired    = "missing_required"
	ErrInvalidDomain      = "invalid_domain"
	ErrUnsupportedVersion = "unsupported_version"
	ErrInvalidValue       = "invalid_value"
	ErrNotFound           = "not_found"
)

// ValidationError represents a validation failure with context about why it failed.
type ValidationError struct {
	// Validator is the name of the validator function that failed (e.g., "IsEmail").
	Validator string
	// Code is a machine-readable error code (e.g., ErrInvalidFormat).
	Code string
	// Message is a human-readable description of the failure.
	Message string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Validator, e.Message)
}

// newValidationError creates a new ValidationError.
func newValidationError(validator, code, message string) error {
	return &ValidationError{
		Validator: validator,
		Code:      code,
		Message:   message,
	}
}
