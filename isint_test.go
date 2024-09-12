package govalidator

import "testing"

func TestIsInt(t *testing.T) {
	zero := 0
	hundred := 100
	one := 1
	five := 5
	eight := 8
	ten := 10
	minusTen := -10

	tests := []struct {
		name   string
		param1 string
		param2 IsIntOpts
		want   bool
	}{
		// Basic integer validation
		{name: "Valid positive integer", param1: "123", param2: IsIntOpts{}, want: true},
		{name: "Valid negative integer", param1: "-123", param2: IsIntOpts{}, want: true},
		{name: "Invalid non-integer (float)", param1: "123.45", param2: IsIntOpts{}, want: false},
		{name: "Invalid non-integer (alphabet)", param1: "abc", param2: IsIntOpts{}, want: false},
		{name: "Invalid non-integer (alphanumeric)", param1: "12a3", param2: IsIntOpts{}, want: false},

		// Leading zeroes validation
		{name: "Valid leading zeroes allowed", param1: "007", param2: IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "Invalid leading zeroes not allowed", param1: "007", param2: IsIntOpts{AllowLeadingZeroes: false}, want: false},

		// Min and Max validation
		{name: "Valid within Min and Max range", param1: "50", param2: IsIntOpts{Min: &zero, Max: &hundred}, want: true},
		{name: "Invalid below Min range", param1: "-10", param2: IsIntOpts{Min: &zero}, want: false},
		{name: "Invalid above Max range", param1: "150", param2: IsIntOpts{Max: &hundred}, want: false},

		// Gt and Lt validation
		{name: "Valid greater than Gt and less than Lt", param1: "3", param2: IsIntOpts{Gt: &one, Lt: &five}, want: true},
		{name: "Invalid not greater than Gt", param1: "1", param2: IsIntOpts{Gt: &one}, want: false},
		{name: "Invalid not less than Lt", param1: "5", param2: IsIntOpts{Lt: &five}, want: false},

		// // Combined Min, Max, Gt, Lt, and leading zeroes validation
		{name: "Valid leading zeroes, within Min, Max, Gt, and Lt", param1: "07", param2: IsIntOpts{Min: &zero, Max: &ten, Gt: &zero, Lt: &eight, AllowLeadingZeroes: true}, want: true},
		{name: "Invalid below Gt", param1: "0", param2: IsIntOpts{Min: &zero, Max: &ten, Gt: &zero, Lt: &eight, AllowLeadingZeroes: true}, want: false},
		{name: "Invalid above Lt", param1: "8", param2: IsIntOpts{Min: &zero, Max: &ten, Gt: &zero, Lt: &eight, AllowLeadingZeroes: true}, want: false},
		{name: "Invalid leading zeroes, not allowed", param1: "07", param2: IsIntOpts{Min: &zero, Max: &ten, Gt: &zero, Lt: &eight, AllowLeadingZeroes: false}, want: false},

		// // Edge cases
		{name: "Valid zero with Min and Max", param1: "0", param2: IsIntOpts{Min: &zero, Max: &zero}, want: true},
		{name: "Invalid zero with Min greater than zero", param1: "0", param2: IsIntOpts{Min: &one}, want: false},
		{name: "Valid single digit", param1: "5", param2: IsIntOpts{}, want: true},
		{name: "Valid negative integer within Min and Max", param1: "-5", param2: IsIntOpts{Min: &minusTen, Max: &zero}, want: true},
		{name: "Invalid negative integer below Min", param1: "-15", param2: IsIntOpts{Min: &minusTen}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsInt(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
