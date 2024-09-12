package validatorgo

import (
	"testing"
)

func TestIsBase64(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsBase64Opts
		want   bool
	}{
		{
			name:   "Valid base 64 string",
			param1: "SGVsbG8gV29ybGQh",
			param2: IsBase64Opts{UrlSafe: false},
			want:   true,
		},
		{
			name:   "Invalid base 64 string",
			param1: "SGVsbG8gV29ybGQh!",
			param2: IsBase64Opts{UrlSafe: false},
			want:   false,
		},
		{
			name:   "Valid base 64 string thats url safe",
			param1: "SGVsbG8-V29ybGQ_",
			param2: IsBase64Opts{UrlSafe: true},
			want:   true,
		},
		{
			name:   "Invalid base 64 string thats url safe",
			param1: "SGVsbG8-V29ybGQ_!",
			param2: IsBase64Opts{UrlSafe: true},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBase64(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
