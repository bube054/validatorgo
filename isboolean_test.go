package validatorgo

import (
	"testing"
)

func TestIsBoolean(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsBooleanOpts
		want   bool
	}{
		{
			name:   "Test true with strict",
			param1: "true",
			param2: IsBooleanOpts{Loose: false},
			want:   true,
		},
		{
			name:   "Test True with strict",
			param1: "True",
			param2: IsBooleanOpts{Loose: false},
			want:   false,
		},
		{
			name:   "Test TRUE with strict",
			param1: "TRUE",
			param2: IsBooleanOpts{Loose: false},
			want:   false,
		},
		{
			name:   "Test false with strict",
			param1: "false",
			param2: IsBooleanOpts{Loose: false},
			want:   true,
		},
		{
			name:   "Test False with strict",
			param1: "False",
			param2: IsBooleanOpts{Loose: false},
			want:   false,
		},
		{
			name:   "Test FALSE with strict",
			param1: "FALSE",
			param2: IsBooleanOpts{Loose: false},
			want:   false,
		},
		{
			name:   "Test 1 with strict",
			param1: "1",
			param2: IsBooleanOpts{Loose: false},
			want:   true,
		},
		{
			name:   "Test 0 with strict",
			param1: "0",
			param2: IsBooleanOpts{Loose: false},
			want:   true,
		},
		{
			name:   "Test nothing with strict",
			param1: "",
			param2: IsBooleanOpts{Loose: false},
			want:   false,
		},
		{
			name:   "Test bool with strict",
			param1: "bool",
			param2: IsBooleanOpts{Loose: false},
			want:   false,
		},
		{
			name:   "Test true with loose",
			param1: "true",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test True with loose",
			param1: "True",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test TRUE with loose",
			param1: "TRUE",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test false with loose",
			param1: "false",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test False with loose",
			param1: "False",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test FALSE with loose",
			param1: "FALSE",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test 1 with loose",
			param1: "1",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test 0 with loose",
			param1: "0",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test nothing with loose",
			param1: "",
			param2: IsBooleanOpts{Loose: true},
			want:   false,
		},
		{
			name:   "Test bool with loose",
			param1: "bool",
			param2: IsBooleanOpts{Loose: true},
			want:   false,
		},
		{
			name:   "Test yes with loose",
			param1: "yes",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
		{
			name:   "Test no with loose",
			param1: "no",
			param2: IsBooleanOpts{Loose: true},
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBoolean(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
