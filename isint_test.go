package validatorgo

import "testing"

func TestIsInt(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsIntOpts
		want   bool
	}{
		// Basic integer validation
		{name: "Valid int with nil config", param1: "123", param2: nil, want: true},
		{name: "Valid positive integer", param1: "123", param2: &IsIntOpts{}, want: true},
		{name: "Valid negative integer", param1: "-123", param2: &IsIntOpts{}, want: true},
		{name: "Invalid non-integer (float)", param1: "123.45", param2: &IsIntOpts{}, want: false},
		{name: "Invalid non-integer (alphabet)", param1: "abc", param2: &IsIntOpts{}, want: false},
		{name: "Invalid non-integer (alphanumeric)", param1: "12a3", param2: &IsIntOpts{}, want: false},

		// Leading zeroes validation
		{name: "Valid leading zeroes allowed", param1: "007", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "Invalid leading zeroes not allowed", param1: "007", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},

		// Min and Max validation
		{name: "Valid within Min and Max range", param1: "50", param2: &IsIntOpts{Min: intPtr(0), Max: intPtr(100)}, want: true},
		{name: "Invalid below Min range", param1: "-10", param2: &IsIntOpts{Min: intPtr(0)}, want: false},
		{name: "Invalid above Max range", param1: "150", param2: &IsIntOpts{Max: intPtr(100)}, want: false},

		// Gt and Lt validation
		{name: "Valid greater than Gt and less than Lt", param1: "3", param2: &IsIntOpts{Gt: intPtr(1), Lt: intPtr(5)}, want: true},
		{name: "Invalid not greater than Gt", param1: "1", param2: &IsIntOpts{Gt: intPtr(1)}, want: false},
		{name: "Invalid not less than Lt", param1: "5", param2: &IsIntOpts{Lt: intPtr(5)}, want: false},

		// // Combined Min, Max, Gt, Lt, and leading zeroes validation
		{name: "Valid leading zeroes, within Min, Max, Gt, and Lt", param1: "07", param2: &IsIntOpts{Min: intPtr(0), Max: intPtr(10), Gt: intPtr(0), Lt: intPtr(8), AllowLeadingZeroes: true}, want: true},
		{name: "Invalid below Gt", param1: "0", param2: &IsIntOpts{Min: intPtr(0), Max: intPtr(10), Gt: intPtr(0), Lt: intPtr(8), AllowLeadingZeroes: true}, want: false},
		{name: "Invalid above Lt", param1: "8", param2: &IsIntOpts{Min: intPtr(0), Max: intPtr(10), Gt: intPtr(0), Lt: intPtr(8), AllowLeadingZeroes: true}, want: false},
		{name: "Invalid leading zeroes, not allowed", param1: "07", param2: &IsIntOpts{Min: intPtr(0), Max: intPtr(10), Gt: intPtr(0), Lt: intPtr(8), AllowLeadingZeroes: false}, want: false},

		// // Edge cases
		{name: "Valid zero with Min and Max", param1: "0", param2: &IsIntOpts{Min: intPtr(0), Max: intPtr(0)}, want: true},
		{name: "Invalid zero with Min greater than zero", param1: "0", param2: &IsIntOpts{Min: intPtr(1)}, want: false},
		{name: "Valid single digit", param1: "5", param2: &IsIntOpts{}, want: true},
		{name: "Valid negative integer within Min and Max", param1: "-5", param2: &IsIntOpts{Min: intPtr(-10), Max: intPtr(0)}, want: true},
		{name: "Invalid negative integer below Min", param1: "-15", param2: &IsIntOpts{Min: intPtr(-10)}, want: false},

		// ---- Ported from validator.js: basic valid integers (no opts / default) ----
		{name: "JS: valid '13' no opts", param1: "13", param2: nil, want: true},
		{name: "JS: valid '0' no opts", param1: "0", param2: nil, want: true},
		{name: "JS: valid '-0' no opts", param1: "-0", param2: nil, want: true},
		{name: "JS: valid '+1' no opts", param1: "+1", param2: nil, want: true},
		// TODO: should be valid per validator.js — JS default allows leading zeroes, Go default does not
		{name: "JS: '01' no opts", param1: "01", param2: nil, want: false},
		// TODO: should be valid per validator.js — JS default allows leading zeroes, Go default does not
		{name: "JS: '-01' no opts", param1: "-01", param2: nil, want: false},
		// TODO: should be valid per validator.js — JS default allows leading zeroes, Go default does not
		{name: "JS: '000' no opts", param1: "000", param2: nil, want: false},

		// ---- Ported from validator.js: basic invalid integers (no opts / default) ----
		{name: "JS: invalid '100e10' no opts", param1: "100e10", param2: nil, want: false},
		{name: "JS: invalid '123.123' no opts", param1: "123.123", param2: nil, want: false},
		{name: "JS: invalid '   ' no opts", param1: "   ", param2: nil, want: false},
		{name: "JS: invalid '' no opts", param1: "", param2: nil, want: false},

		// ---- Ported from validator.js: allow_leading_zeroes: false ----
		{name: "JS: valid '13' leading zeroes false", param1: "13", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: true},
		{name: "JS: valid '123' leading zeroes false", param1: "123", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: true},
		{name: "JS: valid '0' leading zeroes false", param1: "0", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: true},
		{name: "JS: valid '-0' leading zeroes false", param1: "-0", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: true},
		{name: "JS: valid '+1' leading zeroes false", param1: "+1", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: true},
		{name: "JS: invalid '01' leading zeroes false", param1: "01", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},
		{name: "JS: invalid '-01' leading zeroes false", param1: "-01", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},
		{name: "JS: invalid '000' leading zeroes false", param1: "000", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},
		{name: "JS: invalid '100e10' leading zeroes false", param1: "100e10", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},
		{name: "JS: invalid '123.123' leading zeroes false", param1: "123.123", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},
		{name: "JS: invalid '   ' leading zeroes false", param1: "   ", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},
		{name: "JS: invalid '' leading zeroes false", param1: "", param2: &IsIntOpts{AllowLeadingZeroes: false}, want: false},

		// ---- Ported from validator.js: allow_leading_zeroes: true ----
		{name: "JS: valid '13' leading zeroes true", param1: "13", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '123' leading zeroes true", param1: "123", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '0' leading zeroes true", param1: "0", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '-0' leading zeroes true", param1: "-0", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '+1' leading zeroes true", param1: "+1", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '01' leading zeroes true", param1: "01", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '-01' leading zeroes true", param1: "-01", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '000' leading zeroes true", param1: "000", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '-000' leading zeroes true", param1: "-000", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: valid '+000' leading zeroes true", param1: "+000", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: true},
		{name: "JS: invalid '100e10' leading zeroes true", param1: "100e10", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: false},
		{name: "JS: invalid '123.123' leading zeroes true", param1: "123.123", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: false},
		{name: "JS: invalid '   ' leading zeroes true", param1: "   ", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: false},
		{name: "JS: invalid '' leading zeroes true", param1: "", param2: &IsIntOpts{AllowLeadingZeroes: true}, want: false},

		// ---- Ported from validator.js: min only ----
		{name: "JS: valid '15' min 10", param1: "15", param2: &IsIntOpts{Min: intPtr(10)}, want: true},
		{name: "JS: valid '80' min 10", param1: "80", param2: &IsIntOpts{Min: intPtr(10)}, want: true},
		{name: "JS: valid '99' min 10", param1: "99", param2: &IsIntOpts{Min: intPtr(10)}, want: true},
		{name: "JS: invalid '9' min 10", param1: "9", param2: &IsIntOpts{Min: intPtr(10)}, want: false},
		{name: "JS: invalid '6' min 10", param1: "6", param2: &IsIntOpts{Min: intPtr(10)}, want: false},
		{name: "JS: invalid '3.2' min 10", param1: "3.2", param2: &IsIntOpts{Min: intPtr(10)}, want: false},
		{name: "JS: invalid 'a' min 10", param1: "a", param2: &IsIntOpts{Min: intPtr(10)}, want: false},

		// ---- Ported from validator.js: min and max ----
		{name: "JS: valid '15' min 10 max 15", param1: "15", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: true},
		{name: "JS: valid '11' min 10 max 15", param1: "11", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: true},
		{name: "JS: valid '13' min 10 max 15", param1: "13", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: true},
		{name: "JS: invalid '9' min 10 max 15", param1: "9", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: false},
		{name: "JS: invalid '2' min 10 max 15", param1: "2", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: false},
		{name: "JS: invalid '17' min 10 max 15", param1: "17", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: false},
		{name: "JS: invalid '3.2' min 10 max 15", param1: "3.2", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: false},
		{name: "JS: invalid '33' min 10 max 15", param1: "33", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: false},
		{name: "JS: invalid 'a' min 10 max 15", param1: "a", param2: &IsIntOpts{Min: intPtr(10), Max: intPtr(15)}, want: false},

		// ---- Ported from validator.js: gt and lt ----
		{name: "JS: valid '14' gt 10 lt 15", param1: "14", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: true},
		{name: "JS: valid '11' gt 10 lt 15", param1: "11", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: true},
		{name: "JS: valid '13' gt 10 lt 15", param1: "13", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: true},
		{name: "JS: invalid '10' gt 10 lt 15", param1: "10", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: false},
		{name: "JS: invalid '15' gt 10 lt 15", param1: "15", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: false},
		{name: "JS: invalid '17' gt 10 lt 15", param1: "17", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: false},
		{name: "JS: invalid '3.2' gt 10 lt 15", param1: "3.2", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: false},
		{name: "JS: invalid '33' gt 10 lt 15", param1: "33", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: false},
		{name: "JS: invalid 'a' gt 10 lt 15", param1: "a", param2: &IsIntOpts{Gt: intPtr(10), Lt: intPtr(15)}, want: false},

		// ---- Ported from validator.js: nil min/max (undefined equivalent) ----
		{name: "JS: valid '143' nil min/max", param1: "143", param2: &IsIntOpts{Min: nil, Max: nil}, want: true},
		{name: "JS: valid '15' nil min/max", param1: "15", param2: &IsIntOpts{Min: nil, Max: nil}, want: true},
		{name: "JS: invalid '10.4' nil min/max", param1: "10.4", param2: &IsIntOpts{Min: nil, Max: nil}, want: false},
		{name: "JS: invalid 'bar' nil min/max", param1: "bar", param2: &IsIntOpts{Min: nil, Max: nil}, want: false},
		{name: "JS: invalid '10a' nil min/max", param1: "10a", param2: &IsIntOpts{Min: nil, Max: nil}, want: false},
		{name: "JS: invalid 'c44' nil min/max", param1: "c44", param2: &IsIntOpts{Min: nil, Max: nil}, want: false},

		// ---- Ported from validator.js: nil gt/lt (undefined equivalent) ----
		{name: "JS: valid '289373466' nil gt/lt", param1: "289373466", param2: &IsIntOpts{Gt: nil, Lt: nil}, want: true},
		{name: "JS: valid '55' nil gt/lt", param1: "55", param2: &IsIntOpts{Gt: nil, Lt: nil}, want: true},
		{name: "JS: valid '989' nil gt/lt", param1: "989", param2: &IsIntOpts{Gt: nil, Lt: nil}, want: true},
		{name: "JS: invalid '10.4' nil gt/lt", param1: "10.4", param2: &IsIntOpts{Gt: nil, Lt: nil}, want: false},
		{name: "JS: invalid 'baz' nil gt/lt", param1: "baz", param2: &IsIntOpts{Gt: nil, Lt: nil}, want: false},
		{name: "JS: invalid '66a' nil gt/lt", param1: "66a", param2: &IsIntOpts{Gt: nil, Lt: nil}, want: false},
		{name: "JS: invalid 'c21' nil gt/lt", param1: "c21", param2: &IsIntOpts{Gt: nil, Lt: nil}, want: false},

		// ---- Ported from validator.js: nil gt + nil max (null equivalent) ----
		{name: "JS: valid '1' nil gt/max", param1: "1", param2: &IsIntOpts{Gt: nil, Max: nil}, want: true},
		{name: "JS: valid '886' nil gt/max", param1: "886", param2: &IsIntOpts{Gt: nil, Max: nil}, want: true},
		{name: "JS: invalid '10.4' nil gt/max", param1: "10.4", param2: &IsIntOpts{Gt: nil, Max: nil}, want: false},
		{name: "JS: invalid 'h' nil gt/max", param1: "h", param2: &IsIntOpts{Gt: nil, Max: nil}, want: false},
		{name: "JS: invalid '1.2' nil gt/max", param1: "1.2", param2: &IsIntOpts{Gt: nil, Max: nil}, want: false},
		{name: "JS: invalid '+' nil gt/max", param1: "+", param2: &IsIntOpts{Gt: nil, Max: nil}, want: false},

		// ---- Ported from validator.js: nil lt + nil min (null equivalent) ----
		{name: "JS: valid '289373466' nil lt/min", param1: "289373466", param2: &IsIntOpts{Lt: nil, Min: nil}, want: true},
		{name: "JS: valid '55' nil lt/min", param1: "55", param2: &IsIntOpts{Lt: nil, Min: nil}, want: true},
		{name: "JS: valid '989' nil lt/min", param1: "989", param2: &IsIntOpts{Lt: nil, Min: nil}, want: true},
		{name: "JS: invalid ',' nil lt/min", param1: ",", param2: &IsIntOpts{Lt: nil, Min: nil}, want: false},
		{name: "JS: invalid '+11212+' nil lt/min", param1: "+11212+", param2: &IsIntOpts{Lt: nil, Min: nil}, want: false},
		{name: "JS: invalid 'fail' nil lt/min", param1: "fail", param2: &IsIntOpts{Lt: nil, Min: nil}, want: false},
		{name: "JS: invalid '111987234i' nil lt/min", param1: "111987234i", param2: &IsIntOpts{Lt: nil, Min: nil}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsInt(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
