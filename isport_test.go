package validatorgo

import "testing"

func TestIsPort(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid port numbers
		{name: "Valid Port - Minimum", param1: "0", want: true},
		{name: "Valid Port - Common HTTP", param1: "80", want: true},
		{name: "Valid Port - Common HTTPS", param1: "443", want: true},
		{name: "Valid Port - FTP", param1: "21", want: true},
		{name: "Valid Port - SSH", param1: "22", want: true},
		{name: "Valid Port - Custom Port", param1: "8080", want: true},
		{name: "Valid Port - Upper Bound", param1: "65535", want: true},

		// Invalid port numbers
		{name: "Invalid Port - Below Minimum", param1: "-1", want: false},
		{name: "Invalid Port - Above Maximum", param1: "65536", want: false},
		{name: "Invalid Port - Non-numeric characters", param1: "abc123", want: false},
		{name: "Invalid Port - With special characters", param1: "123$", want: false},
		{name: "Invalid Port - Floating point", param1: "8080.1", want: false},
		{name: "Invalid Port - Empty string", param1: "", want: false},
		{name: "Invalid Port - Space character", param1: " 8080", want: false},
		{name: "Invalid Port - Leading zero", param1: "080", want: true},
		{name: "Invalid Port - Extra characters", param1: "443port", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPort(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
