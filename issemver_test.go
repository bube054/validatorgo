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
		// Ported from validator.js — valid semver cases
		{
			name:   "Valid semver - simple patch zero",
			param1: "0.0.4",
			want:   true,
		},
		{
			name:   "Valid semver - double digit components",
			param1: "10.20.30",
			want:   true,
		},
		{
			name:   "Valid semver - prerelease and meta",
			param1: "1.1.2-prerelease+meta",
			want:   true,
		},
		{
			name:   "Valid semver - meta only",
			param1: "1.1.2+meta",
			want:   true,
		},
		{
			name:   "Valid semver - meta with hyphen",
			param1: "1.1.2+meta-valid",
			want:   true,
		},
		{
			name:   "Valid semver - alpha.beta prerelease",
			param1: "1.0.0-alpha.beta",
			want:   true,
		},
		{
			name:   "Valid semver - alpha.beta.1 prerelease",
			param1: "1.0.0-alpha.beta.1",
			want:   true,
		},
		{
			name:   "Valid semver - alpha.1 prerelease",
			param1: "1.0.0-alpha.1",
			want:   true,
		},
		{
			name:   "Valid semver - alpha0.valid prerelease",
			param1: "1.0.0-alpha0.valid",
			want:   true,
		},
		{
			name:   "Valid semver - alpha.0valid prerelease",
			param1: "1.0.0-alpha.0valid",
			want:   true,
		},
		{
			name:   "Valid semver - complex prerelease and build",
			param1: "1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay",
			want:   true,
		},
		{
			name:   "Valid semver - rc.1 with build.1",
			param1: "1.0.0-rc.1+build.1",
			want:   true,
		},
		{
			name:   "Valid semver - rc.1 with build.123",
			param1: "2.0.0-rc.1+build.123",
			want:   true,
		},
		{
			name:   "Valid semver - DEV-SNAPSHOT prerelease",
			param1: "10.2.3-DEV-SNAPSHOT",
			want:   true,
		},
		{
			name:   "Valid semver - SNAPSHOT-123 prerelease",
			param1: "1.2.3-SNAPSHOT-123",
			want:   true,
		},
		{
			name:   "Valid semver - 2.0.0 plain",
			param1: "2.0.0",
			want:   true,
		},
		{
			name:   "Valid semver - 1.1.7 plain",
			param1: "1.1.7",
			want:   true,
		},
		{
			name:   "Valid semver - build.1848 metadata",
			param1: "2.0.0+build.1848",
			want:   true,
		},
		{
			name:   "Valid semver - alpha.1227 prerelease",
			param1: "2.0.1-alpha.1227",
			want:   true,
		},
		{
			name:   "Valid semver - alpha+beta",
			param1: "1.0.0-alpha+beta",
			want:   true,
		},
		{
			name:   "Valid semver - long dashes in prerelease and build",
			param1: "1.2.3----RC-SNAPSHOT.12.9.1--.12+788",
			want:   true,
		},
		{
			name:   "Valid semver - long dashes in prerelease with meta",
			param1: "1.2.3----R-S.12.9.1--.12+meta",
			want:   true,
		},
		{
			name:   "Valid semver - long dashes in prerelease no build",
			param1: "1.2.3----RC-SNAPSHOT.12.9.1--.12",
			want:   true,
		},
		{
			name:   "Valid semver - complex build metadata",
			param1: "1.0.0+0.build.1-rc.10000aaa-kk-0.1",
			want:   true,
		},
		{
			name:   "Valid semver - very large version numbers",
			param1: "99999999999999999999999.999999999999999999.99999999999999999",
			want:   true,
		},
		{
			name:   "Valid semver - 0A.is.legal prerelease",
			param1: "1.0.0-0A.is.legal",
			want:   true,
		},
		{
			name:   "Valid semver - beta prerelease",
			param1: "1.0.0-beta",
			want:   true,
		},
		// Ported from validator.js — invalid semver cases
		{
			name:   "Invalid semver - starts with dash",
			param1: "-invalid+invalid",
			want:   false,
		},
		{
			name:   "Invalid semver - starts with dash and leading zero",
			param1: "-invalid.01",
			want:   false,
		},
		{
			name:   "Invalid semver - alpha only",
			param1: "alpha",
			want:   false,
		},
		{
			name:   "Invalid semver - alpha.beta no version",
			param1: "alpha.beta",
			want:   false,
		},
		{
			name:   "Invalid semver - alpha.beta.1 no version",
			param1: "alpha.beta.1",
			want:   false,
		},
		{
			name:   "Invalid semver - alpha.1 no version",
			param1: "alpha.1",
			want:   false,
		},
		{
			name:   "Invalid semver - alpha+beta no version",
			param1: "alpha+beta",
			want:   false,
		},
		{
			name:   "Invalid semver - underscore in name",
			param1: "alpha_beta",
			want:   false,
		},
		{
			name:   "Invalid semver - alpha with trailing dot",
			param1: "alpha.",
			want:   false,
		},
		{
			name:   "Invalid semver - alpha with double trailing dots",
			param1: "alpha..",
			want:   false,
		},
		{
			name:   "Invalid semver - beta only",
			param1: "beta",
			want:   false,
		},
		{
			name:   "Invalid semver - underscore in prerelease",
			param1: "1.0.0-alpha_beta",
			want:   false,
		},
		{
			name:   "Invalid semver - starts with dash alpha",
			param1: "-alpha.",
			want:   false,
		},
		{
			name:   "Invalid semver - double dot in prerelease",
			param1: "1.0.0-alpha..",
			want:   false,
		},
		{
			name:   "Invalid semver - double dot in prerelease with id",
			param1: "1.0.0-alpha..1",
			want:   false,
		},
		{
			name:   "Invalid semver - triple dot in prerelease",
			param1: "1.0.0-alpha...1",
			want:   false,
		},
		{
			name:   "Invalid semver - four dots in prerelease",
			param1: "1.0.0-alpha....1",
			want:   false,
		},
		{
			name:   "Invalid semver - five dots in prerelease",
			param1: "1.0.0-alpha.....1",
			want:   false,
		},
		{
			name:   "Invalid semver - six dots in prerelease",
			param1: "1.0.0-alpha......1",
			want:   false,
		},
		{
			name:   "Invalid semver - seven dots in prerelease",
			param1: "1.0.0-alpha.......1",
			want:   false,
		},
		{
			name:   "Invalid semver - leading zero in minor",
			param1: "1.01.1",
			want:   false,
		},
		{
			name:   "Invalid semver - leading zero in patch",
			param1: "1.1.01",
			want:   false,
		},
		{
			name:   "Invalid semver - four part with DEV",
			param1: "1.2.3.DEV",
			want:   false,
		},
		{
			name:   "Invalid semver - two part with SNAPSHOT",
			param1: "1.2-SNAPSHOT",
			want:   false,
		},
		{
			name:   "Invalid semver - malformed concatenated version",
			param1: "1.2.31.2.3----RC-SNAPSHOT.12.09.1--..12+788",
			want:   false,
		},
		{
			name:   "Invalid semver - two part with RC-SNAPSHOT",
			param1: "1.2-RC-SNAPSHOT",
			want:   false,
		},
		{
			name:   "Invalid semver - negative major version",
			param1: "-1.0.3-gamma+b7718",
			want:   false,
		},
		{
			name:   "Invalid semver - just meta no version",
			param1: "+justmeta",
			want:   false,
		},
		{
			name:   "Invalid semver - double plus in build meta",
			param1: "9.8.7+meta+meta",
			want:   false,
		},
		{
			name:   "Invalid semver - double plus in prerelease build meta",
			param1: "9.8.7-whatever+meta+meta",
			want:   false,
		},
		{
			name:   "Invalid semver - trailing dash in prerelease",
			param1: "99999999999999999999999.999999999999999999.99999999999999999-",
			want:   false,
		},
		{
			name:   "Invalid semver - dashes and dots no version",
			param1: "---RC-SNAPSHOT.12.09.1--------------------------------..12",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsSemVer(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
