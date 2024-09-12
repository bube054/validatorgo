package validatorgo

import (
	"testing"
)

func TestIsBase32(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsBase32Opts
		want   bool
	}{
		{
			name:   "Valid base 32 string",
			param1: "MZXW6YTBOI======",
			param2: IsBase32Opts{Crockford: false},
			want:   true,
		},
		{
			name:   "Valid base 32 string without padding",
			param1: "MZXW6YTBOI",
			param2: IsBase32Opts{Crockford: false},
			want:   true,
		},
		{
			name:   "Invalid base 32 string with incorrect padding",
			param1: "MZXW6YTBOI===",
			param2: IsBase32Opts{Crockford: false},
			want:   false,
		},
		{
			name:   "Invalid base 32 string",
			param1: "MZXW6YTBOI#=====",
			param2: IsBase32Opts{Crockford: false},
			want:   false,
		},
		{
			name:   "Valid base 32 string with Crockford variant",
			param1: "C4RG69S8NP1X",
			param2: IsBase32Opts{Crockford: true},
			want:   true,
		},
		{
			name:   "Invalid base 32 string with Crockford variant",
			param1: "C4RG69S8NP1O",
			param2: IsBase32Opts{Crockford: true},
			want:   false,
		},
		{
			name:   "Valid Crockford string with lowercase letters",
			param1: "c4rg69s8np1x",
			param2: IsBase32Opts{Crockford: true},
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBase32(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
