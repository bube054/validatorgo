package validatorgo

import (
	"errors"
	"testing"
)

// assertValidation checks that a validator's bool result matches the expected value
// and that the error return is consistent:
//   - valid input (want=true): err must be nil
//   - invalid input (want=false): err must be a non-nil *ValidationError with populated fields
func assertValidation(t *testing.T, result, want bool, err error) {
	t.Helper()

	if result != want {
		t.Errorf("result: got %t, want %t", result, want)
	}

	if want {
		if err != nil {
			t.Errorf("expected nil error for valid input, got: %v", err)
		}
	} else {
		if err == nil {
			t.Error("expected non-nil error for invalid input, got nil")
			return
		}
		var ve *ValidationError
		if !errors.As(err, &ve) {
			t.Errorf("expected *ValidationError, got %T: %v", err, err)
			return
		}
		if ve.Validator == "" {
			t.Error("ValidationError.Validator is empty")
		}
		if ve.Code == "" {
			t.Error("ValidationError.Code is empty")
		}
		if ve.Message == "" {
			t.Error("ValidationError.Message is empty")
		}
	}
}
