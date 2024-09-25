package validatorgo

import "testing"

func TestIsSemver(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "Valid semver",
			param1: "1.0.0",
			want:   true,
		},
		{
			name:   "Valid semver",
			param1: "2.1.3",
			want:   true,
		},
		{
			name:   "Valid semver",
			param1: "0.1.0",
			want:   true,
		},
		{
			name:   "Valid semver",
			param1: "1.0.0-alpha",
			want:   true,
		},
		{
			name:   "Valid semver",
			param1: "1.2.3-beta.1",
			want:   true,
		},
		{
			name:   "Valid semver",
			param1: "3.0.0-rc.1+build.001",
			want:   true,
		},
		{
			name:   "Valid semver",
			param1: "1.0.0+exp.sha.5114f85",
			want:   true,
		},
		{
			name:   "Invalid semver",
			param1: "1.0",
			want:   false,
		},
		// {
		// 	name:   "Invalid semver",
		// 	param1: "1.0.0-beta-",
		// 	want:   false,
		// },
		{
			name:   "Invalid semver",
			param1: "1.0.0.0",
			want:   false,
		},
		{
			name:   "Invalid semver",
			param1: "01.0.0",
			want:   false,
		},
		{
			name:   "Invalid semver",
			param1: "1.0.0+!build.123",
			want:   false,
		},
		{
			name:   "Invalid semver",
			param1: "v1.0.0",
			want:   false,
		},
		{
			name:   "Invalid semver",
			param1: "1.0.0-alpha@001",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsSemVer(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
