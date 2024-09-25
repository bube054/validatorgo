package validatorgo

import "testing"

func TestIsFloat(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsFloatOpts
		want   bool
	}{
		// Basic valid cases
		{name: "Valid float number", param1: "123.45", param2: IsFloatOpts{}, want: true},
		{name: "Valid negative float", param1: "-123.45", param2: IsFloatOpts{}, want: true},
		{name: "Valid float with positive sign", param1: "+123.45", param2: IsFloatOpts{}, want: true},

		// Boundary conditions (Min and Max)
		{name: "Float equal to Min", param1: "100.0", param2: IsFloatOpts{Min: floatPtr(100.0)}, want: true},
		{name: "Float greater than Min", param1: "150.0", param2: IsFloatOpts{Min: floatPtr(100.0)}, want: true},
		{name: "Float less than Min", param1: "50.0", param2: IsFloatOpts{Min: floatPtr(100.0)}, want: false},

		{name: "Float equal to Max", param1: "200.0", param2: IsFloatOpts{Max: floatPtr(200.0)}, want: true},
		{name: "Float less than Max", param1: "150.0", param2: IsFloatOpts{Max: floatPtr(200.0)}, want: true},
		{name: "Float greater than Max", param1: "250.0", param2: IsFloatOpts{Max: floatPtr(200.0)}, want: false},

		// Greater than (Gt) and Less than (Lt)
		{name: "Float greater than Gt", param1: "150.0", param2: IsFloatOpts{Gt: floatPtr(100.0)}, want: true},
		{name: "Float equal to Gt", param1: "100.0", param2: IsFloatOpts{Gt: floatPtr(100.0)}, want: false},
		{name: "Float less than Gt", param1: "50.0", param2: IsFloatOpts{Gt: floatPtr(100.0)}, want: false},

		{name: "Float less than Lt", param1: "50.0", param2: IsFloatOpts{Lt: floatPtr(100.0)}, want: true},
		{name: "Float equal to Lt", param1: "100.0", param2: IsFloatOpts{Lt: floatPtr(100.0)}, want: false},
		{name: "Float greater than Lt", param1: "150.0", param2: IsFloatOpts{Lt: floatPtr(100.0)}, want: false},

		// Combining Min, Max, Gt, and Lt
		{name: "Float within Min and Max", param1: "150.0", param2: IsFloatOpts{Min: floatPtr(100.0), Max: floatPtr(200.0)}, want: true},
		{name: "Float outside Min and Max", param1: "250.0", param2: IsFloatOpts{Min: floatPtr(100.0), Max: floatPtr(200.0)}, want: false},

		{name: "Float greater than Gt and less than Lt", param1: "150.0", param2: IsFloatOpts{Gt: floatPtr(100.0), Lt: floatPtr(200.0)}, want: true},
		{name: "Float outside Gt and Lt", param1: "250.0", param2: IsFloatOpts{Gt: floatPtr(100.0), Lt: floatPtr(200.0)}, want: false},

		// Locale-specific cases
		{name: "Valid float with comma as decimal (German locale)", param1: "123,45", param2: IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "Invalid float with comma as decimal (US locale)", param1: "123,45", param2: IsFloatOpts{Locale: "en-US"}, want: false},

		// Invalid float inputs
		{name: "Invalid float with letters", param1: "123.45abc", param2: IsFloatOpts{}, want: false},
		{name: "Invalid float without decimal", param1: "123", param2: IsFloatOpts{}, want: false},
		{name: "Empty string", param1: "", param2: IsFloatOpts{}, want: false},
		{name: "Only sign without number", param1: "+", param2: IsFloatOpts{}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsFloat(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
