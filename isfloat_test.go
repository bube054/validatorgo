package validatorgo

import "testing"

func TestIsFloat(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsFloatOpts
		want   bool
	}{
		// Basic valid cases
		{name: "Valid float number", param1: "123.45", param2: &IsFloatOpts{}, want: true},
		{name: "Valid negative float", param1: "-123.45", param2: &IsFloatOpts{}, want: true},
		{name: "Valid float with positive sign", param1: "+123.45", param2: &IsFloatOpts{}, want: true},

		// Boundary conditions (Min and Max)
		{name: "Float equal to Min", param1: "100.0", param2: &IsFloatOpts{Min: floatPtr(100.0)}, want: true},
		{name: "Float greater than Min", param1: "150.0", param2: &IsFloatOpts{Min: floatPtr(100.0)}, want: true},
		{name: "Float less than Min", param1: "50.0", param2: &IsFloatOpts{Min: floatPtr(100.0)}, want: false},

		{name: "Float equal to Max", param1: "200.0", param2: &IsFloatOpts{Max: floatPtr(200.0)}, want: true},
		{name: "Float less than Max", param1: "150.0", param2: &IsFloatOpts{Max: floatPtr(200.0)}, want: true},
		{name: "Float greater than Max", param1: "250.0", param2: &IsFloatOpts{Max: floatPtr(200.0)}, want: false},

		// Greater than (Gt) and Less than (Lt)
		{name: "Float greater than Gt", param1: "150.0", param2: &IsFloatOpts{Gt: floatPtr(100.0)}, want: true},
		{name: "Float equal to Gt", param1: "100.0", param2: &IsFloatOpts{Gt: floatPtr(100.0)}, want: false},
		{name: "Float less than Gt", param1: "50.0", param2: &IsFloatOpts{Gt: floatPtr(100.0)}, want: false},

		{name: "Float less than Lt", param1: "50.0", param2: &IsFloatOpts{Lt: floatPtr(100.0)}, want: true},
		{name: "Float equal to Lt", param1: "100.0", param2: &IsFloatOpts{Lt: floatPtr(100.0)}, want: false},
		{name: "Float greater than Lt", param1: "150.0", param2: &IsFloatOpts{Lt: floatPtr(100.0)}, want: false},

		// Combining Min, Max, Gt, and Lt
		{name: "Float within Min and Max", param1: "150.0", param2: &IsFloatOpts{Min: floatPtr(100.0), Max: floatPtr(200.0)}, want: true},
		{name: "Float outside Min and Max", param1: "250.0", param2: &IsFloatOpts{Min: floatPtr(100.0), Max: floatPtr(200.0)}, want: false},

		{name: "Float greater than Gt and less than Lt", param1: "150.0", param2: &IsFloatOpts{Gt: floatPtr(100.0), Lt: floatPtr(200.0)}, want: true},
		{name: "Float outside Gt and Lt", param1: "250.0", param2: &IsFloatOpts{Gt: floatPtr(100.0), Lt: floatPtr(200.0)}, want: false},

		// Locale-specific cases
		{name: "Valid float with comma as decimal (German locale)", param1: "123,45", param2: &IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "Invalid float with comma as decimal (US locale)", param1: "123,45", param2: &IsFloatOpts{Locale: "en-US"}, want: false},
		{name: "Invalid Locale", param1: "123,45", param2: &IsFloatOpts{Locale: "nz-WE"}, want: false},

		// Invalid float inputs
		{name: "Invalid float with letters", param1: "123.45abc", param2: &IsFloatOpts{}, want: false},
		{name: "Invalid float without decimal", param1: "123", param2: &IsFloatOpts{}, want: false},
		{name: "Empty string", param1: "", param2: &IsFloatOpts{}, want: false},
		{name: "Only sign without number", param1: "+", param2: &IsFloatOpts{}, want: false},

		// Nil config cases
		{name: "Valid float number with nil config", param1: "123.45", param2: nil, want: true},
		{name: "Valid negative float with nil config", param1: "-123.45", param2: nil, want: true},
		{name: "Invalid float with letters and nil config", param1: "123.45abc", param2: nil, want: false},
		{name: "Empty string with nil config", param1: "", param2: nil, want: false},
		{name: "Only sign without number with nil config", param1: "+", param2: nil, want: false},

		// improving code cov
		{name: "Invalid float with comma as decimal (ar-IQ locale)", param1: "١٢٣,٤٥", param2: &IsFloatOpts{Locale: "ar-IQ"}, want: false},
		{name: "Invalid float with period as decimal (cs-CZ locale)", param1: "123.45", param2: &IsFloatOpts{Locale: "cs-CZ"}, want: false},
		{name: "Invalid float with comma as decimal (fa-AF locale)", param1: "۱۲۳,۴۵", param2: &IsFloatOpts{Locale: "fa-AF"}, want: false},
		{name: "Invalid float with comma as decimal (en-IN locale)", param1: "123,45", param2: &IsFloatOpts{Locale: "en-IN"}, want: false},

		// validator.js ported: default opts (no locale)
		// In validator.js, integers like "123" are valid floats; Go requires a decimal point.
		// TODO: should be valid per validator.js — Go regex requires decimal point
		{name: "JS default: integer 123", param1: "123", param2: &IsFloatOpts{}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digits after decimal
		{name: "JS default: trailing dot 123.", param1: "123.", param2: &IsFloatOpts{}, want: false},
		{name: "JS default: 123.123", param1: "123.123", param2: &IsFloatOpts{}, want: true},
		{name: "JS default: -123.123", param1: "-123.123", param2: &IsFloatOpts{}, want: true},
		{name: "JS default: -0.123", param1: "-0.123", param2: &IsFloatOpts{}, want: true},
		{name: "JS default: +0.123", param1: "+0.123", param2: &IsFloatOpts{}, want: true},
		{name: "JS default: 0.123", param1: "0.123", param2: &IsFloatOpts{}, want: true},
		// TODO: should be valid per validator.js — Go regex requires digit before dot for en-US
		{name: "JS default: .0", param1: ".0", param2: &IsFloatOpts{}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digit before dot for en-US
		{name: "JS default: -.123", param1: "-.123", param2: &IsFloatOpts{}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digit before dot for en-US
		{name: "JS default: +.123", param1: "+.123", param2: &IsFloatOpts{}, want: false},
		{name: "JS default: 01.123", param1: "01.123", param2: &IsFloatOpts{}, want: true},
		// TODO: should be valid per validator.js — Go regex does not support scientific notation
		{name: "JS default: scientific notation", param1: "-0.22250738585072011e-307", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: plus sign only", param1: "+", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: minus sign only", param1: "-", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: spaces", param1: "  ", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: empty", param1: "", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: dot only", param1: ".", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: comma only", param1: ",", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: foo", param1: "foo", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: 20.foo", param1: "20.foo", param2: &IsFloatOpts{}, want: false},
		{name: "JS default invalid: ISO date string", param1: "2020-01-06T14:31:00.135Z", param2: &IsFloatOpts{}, want: false},

		// validator.js ported: locale en-AU
		// TODO: should be valid per validator.js — Go regex requires decimal point
		{name: "JS en-AU: integer 123", param1: "123", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digits after decimal
		{name: "JS en-AU: trailing dot 123.", param1: "123.", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU: 123.123", param1: "123.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: true},
		{name: "JS en-AU: -123.123", param1: "-123.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: true},
		{name: "JS en-AU: -0.123", param1: "-0.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: true},
		{name: "JS en-AU: +0.123", param1: "+0.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: true},
		{name: "JS en-AU: 0.123", param1: "0.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: true},
		// TODO: should be valid per validator.js — Go regex requires digit before dot for en-AU
		{name: "JS en-AU: .0", param1: ".0", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digit before dot for en-AU
		{name: "JS en-AU: -.123", param1: "-.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digit before dot for en-AU
		{name: "JS en-AU: +.123", param1: "+.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU: 01.123", param1: "01.123", param2: &IsFloatOpts{Locale: "en-AU"}, want: true},
		// TODO: should be valid per validator.js — Go regex does not support scientific notation
		{name: "JS en-AU: scientific notation", param1: "-0.22250738585072011e-307", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU invalid: arabic decimal", param1: "123٫123", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU invalid: comma decimal", param1: "123,123", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU invalid: spaces", param1: "  ", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU invalid: empty", param1: "", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU invalid: dot only", param1: ".", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},
		{name: "JS en-AU invalid: foo", param1: "foo", param2: &IsFloatOpts{Locale: "en-AU"}, want: false},

		// validator.js ported: locale de-DE
		// TODO: should be valid per validator.js — Go regex requires comma-decimal for de-DE
		{name: "JS de-DE: integer 123", param1: "123", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digits after comma
		{name: "JS de-DE: trailing comma 123,", param1: "123,", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE: 123,123", param1: "123,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "JS de-DE: -123,123", param1: "-123,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "JS de-DE: -0,123", param1: "-0,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "JS de-DE: +0,123", param1: "+0,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "JS de-DE: 0,123", param1: "0,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: true},
		// TODO: should be valid per validator.js — Go regex requires digit before comma for de-DE
		{name: "JS de-DE: ,0", param1: ",0", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digit before comma for de-DE
		{name: "JS de-DE: -,123", param1: "-,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digit before comma for de-DE
		{name: "JS de-DE: +,123", param1: "+,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE: 01,123", param1: "01,123", param2: &IsFloatOpts{Locale: "de-DE"}, want: true},
		// TODO: should be valid per validator.js — Go regex does not support scientific notation
		{name: "JS de-DE: scientific notation", param1: "-0,22250738585072011e-307", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE invalid: dot decimal", param1: "123.123", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE invalid: arabic decimal", param1: "123٫123", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE invalid: spaces", param1: "  ", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE invalid: empty", param1: "", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE invalid: dot only", param1: ".", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},
		{name: "JS de-DE invalid: foo", param1: "foo", param2: &IsFloatOpts{Locale: "de-DE"}, want: false},

		// validator.js ported: min only
		{name: "JS min 3.7: 3.888", param1: "3.888", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: true},
		{name: "JS min 3.7: 3.92", param1: "3.92", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: true},
		{name: "JS min 3.7: 4.5", param1: "4.5", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: true},
		// TODO: should be valid per validator.js — Go regex requires decimal point, "50" is integer
		{name: "JS min 3.7: 50", param1: "50", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: false},
		{name: "JS min 3.7: 3.7", param1: "3.7", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: true},
		{name: "JS min 3.7: 3.71", param1: "3.71", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: true},
		{name: "JS min 3.7 invalid: 3.6", param1: "3.6", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: false},
		{name: "JS min 3.7 invalid: 3.69", param1: "3.69", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: false},
		// TODO: should be invalid per validator.js — Go regex requires decimal point, "3" is integer
		{name: "JS min 3.7 invalid: 3", param1: "3", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: false},
		{name: "JS min 3.7 invalid: 1.5", param1: "1.5", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: false},
		{name: "JS min 3.7 invalid: a", param1: "a", param2: &IsFloatOpts{Min: floatPtr(3.7)}, want: false},

		// validator.js ported: min and max
		{name: "JS min 0.1 max 1.0: 0.1", param1: "0.1", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: true},
		{name: "JS min 0.1 max 1.0: 1.0", param1: "1.0", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: true},
		{name: "JS min 0.1 max 1.0: 0.15", param1: "0.15", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: true},
		{name: "JS min 0.1 max 1.0: 0.33", param1: "0.33", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: true},
		{name: "JS min 0.1 max 1.0: 0.57", param1: "0.57", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: true},
		{name: "JS min 0.1 max 1.0: 0.7", param1: "0.7", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: true},
		// TODO: should be invalid per validator.js — Go regex requires decimal point, "0" is integer
		{name: "JS min 0.1 max 1.0 invalid: 0", param1: "0", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: false},
		{name: "JS min 0.1 max 1.0 invalid: 0.0", param1: "0.0", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: false},
		{name: "JS min 0.1 max 1.0 invalid: a", param1: "a", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: false},
		{name: "JS min 0.1 max 1.0 invalid: 1.3", param1: "1.3", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: false},
		{name: "JS min 0.1 max 1.0 invalid: 0.05", param1: "0.05", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: false},
		// TODO: should be invalid per validator.js — Go regex requires decimal point, "5" is integer
		{name: "JS min 0.1 max 1.0 invalid: 5", param1: "5", param2: &IsFloatOpts{Min: floatPtr(0.1), Max: floatPtr(1.0)}, want: false},

		// validator.js ported: gt and lt
		{name: "JS gt -5.5 lt 10: 9.9", param1: "9.9", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: true},
		{name: "JS gt -5.5 lt 10: 1.0", param1: "1.0", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: true},
		// TODO: should be valid per validator.js — Go regex requires decimal point, "0" is integer
		{name: "JS gt -5.5 lt 10: 0", param1: "0", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		// TODO: should be valid per validator.js — Go regex requires decimal point, "-1" is integer
		{name: "JS gt -5.5 lt 10: -1", param1: "-1", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		// TODO: should be valid per validator.js — Go regex requires decimal point, "7" is integer
		{name: "JS gt -5.5 lt 10: 7", param1: "7", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		{name: "JS gt -5.5 lt 10: -5.4", param1: "-5.4", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: true},
		// TODO: should be invalid per validator.js — Go regex requires decimal point, "10" is integer
		{name: "JS gt -5.5 lt 10 invalid: 10", param1: "10", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		{name: "JS gt -5.5 lt 10 invalid: -5.5", param1: "-5.5", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		{name: "JS gt -5.5 lt 10 invalid: a", param1: "a", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		{name: "JS gt -5.5 lt 10 invalid: -20.3", param1: "-20.3", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		// TODO: should be invalid per validator.js — Go regex does not support scientific notation
		{name: "JS gt -5.5 lt 10 invalid: 20e3", param1: "20e3", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		{name: "JS gt -5.5 lt 10 invalid: 10.00001", param1: "10.00001", param2: &IsFloatOpts{Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},

		// validator.js ported: combined min, max, gt, lt
		{name: "JS combined min/max/gt/lt: 9.99999", param1: "9.99999", param2: &IsFloatOpts{Min: floatPtr(-5.5), Max: floatPtr(10), Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: true},
		{name: "JS combined min/max/gt/lt: -5.499999", param1: "-5.499999", param2: &IsFloatOpts{Min: floatPtr(-5.5), Max: floatPtr(10), Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: true},
		// TODO: should be invalid per validator.js — Go regex requires decimal point, "10" is integer
		{name: "JS combined min/max/gt/lt invalid: 10", param1: "10", param2: &IsFloatOpts{Min: floatPtr(-5.5), Max: floatPtr(10), Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},
		{name: "JS combined min/max/gt/lt invalid: -5.5", param1: "-5.5", param2: &IsFloatOpts{Min: floatPtr(-5.5), Max: floatPtr(10), Gt: floatPtr(-5.5), Lt: floatPtr(10)}, want: false},

		// validator.js ported: de-DE locale with min
		{name: "JS de-DE min 3.1: 123,123", param1: "123,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: true},
		{name: "JS de-DE min 3.1: 3,1", param1: "3,1", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: true},
		{name: "JS de-DE min 3.1: 3,100001", param1: "3,100001", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: true},
		// TODO: should be valid per validator.js — Go regex requires comma-decimal, "123" is integer
		{name: "JS de-DE min 3.1: 123", param1: "123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digits after comma
		{name: "JS de-DE min 3.1: 123,", param1: "123,", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: 3,09", param1: "3,09", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: -,123", param1: "-,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: +,123", param1: "+,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: 01,123", param1: "01,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: -123,123", param1: "-123,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: -0,123", param1: "-0,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: +0,123", param1: "+0,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: 0,123", param1: "0,123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: ,0", param1: ",0", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: dot decimal", param1: "123.123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: arabic decimal", param1: "123٫123", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: spaces", param1: "  ", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: empty", param1: "", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: dot only", param1: ".", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},
		{name: "JS de-DE min 3.1 invalid: foo", param1: "foo", param2: &IsFloatOpts{Locale: "de-DE", Min: floatPtr(3.1)}, want: false},

		// validator.js ported: undefined min/max (nil in Go)
		{name: "JS nil min/max: 123.123", param1: "123.123", param2: &IsFloatOpts{Min: nil, Max: nil}, want: true},
		{name: "JS nil min/max: -767.767", param1: "-767.767", param2: &IsFloatOpts{Min: nil, Max: nil}, want: true},
		{name: "JS nil min/max: +111.111", param1: "+111.111", param2: &IsFloatOpts{Min: nil, Max: nil}, want: true},
		// TODO: should be valid per validator.js — Go regex requires decimal point
		{name: "JS nil min/max: 123", param1: "123", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digits after decimal
		{name: "JS nil min/max: 123.", param1: "123.", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},
		{name: "JS nil min/max invalid: ab565", param1: "ab565", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},
		{name: "JS nil min/max invalid: -,123", param1: "-,123", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},
		{name: "JS nil min/max invalid: +,123", param1: "+,123", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},
		{name: "JS nil min/max invalid: 7866.t", param1: "7866.t", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},
		{name: "JS nil min/max invalid: 123,123", param1: "123,123", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},
		// TODO: should be invalid per validator.js — Go regex requires digits after decimal
		{name: "JS nil min/max invalid: 123,", param1: "123,", param2: &IsFloatOpts{Min: nil, Max: nil}, want: false},

		// validator.js ported: undefined gt/lt (nil in Go)
		{name: "JS nil gt/lt: 14.34343", param1: "14.34343", param2: &IsFloatOpts{Gt: nil, Lt: nil}, want: true},
		{name: "JS nil gt/lt: 11.1", param1: "11.1", param2: &IsFloatOpts{Gt: nil, Lt: nil}, want: true},
		// TODO: should be valid per validator.js — Go regex requires decimal point
		{name: "JS nil gt/lt: 456", param1: "456", param2: &IsFloatOpts{Gt: nil, Lt: nil}, want: false},
		{name: "JS nil gt/lt invalid: ab565", param1: "ab565", param2: &IsFloatOpts{Gt: nil, Lt: nil}, want: false},
		{name: "JS nil gt/lt invalid: -,123", param1: "-,123", param2: &IsFloatOpts{Gt: nil, Lt: nil}, want: false},
		{name: "JS nil gt/lt invalid: +,123", param1: "+,123", param2: &IsFloatOpts{Gt: nil, Lt: nil}, want: false},
		{name: "JS nil gt/lt invalid: 7866.t", param1: "7866.t", param2: &IsFloatOpts{Gt: nil, Lt: nil}, want: false},

		// validator.js ported: locale ru-RU
		{name: "JS ru-RU: 11231554,34343", param1: "11231554,34343", param2: &IsFloatOpts{Locale: "ru-RU"}, want: true},
		{name: "JS ru-RU: 11,1", param1: "11,1", param2: &IsFloatOpts{Locale: "ru-RU"}, want: true},
		// TODO: should be valid per validator.js — Go regex requires comma-decimal, "456" is integer
		{name: "JS ru-RU: 456", param1: "456", param2: &IsFloatOpts{Locale: "ru-RU"}, want: false},
		// TODO: should be valid per validator.js — Go regex requires digit before comma for ru-RU
		{name: "JS ru-RU: ,311", param1: ",311", param2: &IsFloatOpts{Locale: "ru-RU"}, want: false},
		{name: "JS ru-RU invalid: ab565", param1: "ab565", param2: &IsFloatOpts{Locale: "ru-RU"}, want: false},
		{name: "JS ru-RU invalid: -.123", param1: "-.123", param2: &IsFloatOpts{Locale: "ru-RU"}, want: false},
		{name: "JS ru-RU invalid: +.123", param1: "+.123", param2: &IsFloatOpts{Locale: "ru-RU"}, want: false},
		{name: "JS ru-RU invalid: 7866.t", param1: "7866.t", param2: &IsFloatOpts{Locale: "ru-RU"}, want: false},
		{name: "JS ru-RU invalid: 22.3", param1: "22.3", param2: &IsFloatOpts{Locale: "ru-RU"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsFloat(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
