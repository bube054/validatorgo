package govalidator

import "testing"

func TestIsISSN(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsISSNOpts
		want   bool
	}{
		// Valid default options
		{name: "valid ISSN with hyphen", param1: "0378-5955", param2: IsISSNOpts{}, want: true},
		{name: "Valid ISSN with hyphen", param1: "0317-8471", param2: IsISSNOpts{}, want: true},
		{name: "Valid ISSN without hyphen", param1: "34244530", param2: IsISSNOpts{}, want: true},
		{name: "Valid ISSN without hyphen", param1: "92657087", param2: IsISSNOpts{}, want: true},
		// Valid CaseSensitive true
		{name: "Valid ISSN with uppercase 'X'", param1: "2434-561X", param2: IsISSNOpts{CaseSensitive: true}, want: true},
		{name: "Valid ISSN without hyphen", param1: "98765434", param2: IsISSNOpts{CaseSensitive: true}, want: true},
		// Valid RequireHyphen true
		{name: "Valid ISSN without hyphen", param1: "2049-3630", param2: IsISSNOpts{RequireHyphen: true}, want: true},
		// Valid CaseSensitive and RequireHyphen true
		{name: "Valid ISSN without hyphen", param1: "2434-561X", param2: IsISSNOpts{RequireHyphen: true, CaseSensitive: true}, want: true},
		// Invalid default options
		{name: "Too short", param1: "1234567", param2: IsISSNOpts{}, want: false},
		{name: "Too Long", param1: "1234567", param2: IsISSNOpts{}, want: false},
		{name: "Invalid character", param1: "0317-847A", param2: IsISSNOpts{}, want: false},
		// Invalid CaseSensitive true
		{name: "Invalid because 'X' is lowercase", param1: "2434-561x", param2: IsISSNOpts{CaseSensitive: true}, want: false},
		// Invalid RequireHyphen true
		{name: "Valid ISSN without hyphen", param1: "20493630", param2: IsISSNOpts{RequireHyphen: true}, want: false},
		// Valid CaseSensitive and RequireHyphen true
		{name: "Valid ISSN without hyphen", param1: "2434561x", param2: IsISSNOpts{RequireHyphen: true, CaseSensitive: true}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsISSN(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
