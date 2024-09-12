package validatorgo

import "testing"

func TestIsFloat(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsFloatOpts
		want   bool
	}{
		// Valid Inputs no options
		{name: "Typical float", param1: "123.456", param2: IsFloatOpts{}, want: true},
		{name: "A small float", param1: "0.1", param2: IsFloatOpts{}, want: true},
		{name: "A negative float", param1: "-42.42", param2: IsFloatOpts{}, want: true},
		{name: "A float with a zero decimal part", param1: "3.0", param2: IsFloatOpts{}, want: true},
		{name: "A positive float with an explicit plus sign", param1: "+0.99", param2: IsFloatOpts{}, want: true},
		{name: "A large float", param1: "1000000.000001", param2: IsFloatOpts{}, want: true},

		// Invalid Inputs n options
		{name: "A non numeric string", param1: "abc", param2: IsFloatOpts{}, want: false},
		{name: "A number with a decimal point but no digits after it", param1: "123.", param2: IsFloatOpts{}, want: false},
		{name: "A comma used instead of a dot for the decimal separator", param1: "123,456", param2: IsFloatOpts{}, want: false},
		{name: "An empty string", param1: "", param2: IsFloatOpts{}, want: false},
		{name: "Multiple decimal spaces", param1: "123..456", param2: IsFloatOpts{}, want: false},
		{name: "A comma used instead of a dot for the decimal separator", param1: "123,456", param2: IsFloatOpts{}, want: false},

		// Valid Inputs with Min/Max options
		{name: "Falls within range", param1: "2.3", param2: IsFloatOpts{Min: 1.5, Max: 5.5}, want: true},
		{name: "Equal to Min", param1: "1.5", param2: IsFloatOpts{Min: 1.5, Max: 5.5}, want: true},
		{name: "Equal to Max", param1: "5.5", param2: IsFloatOpts{Min: 1.5, Max: 5.5}, want: true},

		// Invalid Inputs with Min/Max options
		{name: "Less than Min", param1: "1.4", param2: IsFloatOpts{Min: 1.5, Max: 5.5}, want: false},
		{name: "Greater than Max", param1: "5.9", param2: IsFloatOpts{Min: 1.5, Max: 5.5}, want: false},
		{name: "Just below min", param1: "1.49", param2: IsFloatOpts{Min: 1.5, Max: 5.5}, want: false},
		{name: "Just above max", param1: "5.5001", param2: IsFloatOpts{Min: 1.5, Max: 5.5}, want: false},

		// Valid Inputs with Gt/Lt options
		{name: "Greater than 0", param1: "0.1", param2: IsFloatOpts{Gt: 0, Lt: 10}, want: true},
		{name: "Less than 10", param1: "9.999", param2: IsFloatOpts{Gt: 0, Lt: 10}, want: true},
		{name: "Not greater than 0", param1: "0", param2: IsFloatOpts{Gt: 0, Lt: 10}, want: false},
		{name: "Not less than 10", param1: "10", param2: IsFloatOpts{Gt: 0, Lt: 10}, want: false},

		// Invalid Inputs with Gt/Lt options
		{name: "A valid float with comma separator for german", param1: "1,23", param2: IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "A negative float with comma for german", param1: "-123,456", param2: IsFloatOpts{Locale: "de-DE"}, want: true},
		{name: "Uses . instead of , for german", param1: "1.23", param2: IsFloatOpts{Locale: "de-DE"}, want: false},

		// Valid/Invalid Inputs with Min/Max and locale options
		{name: "Meets meets min requirement", param1: "0.5", param2: IsFloatOpts{Locale: "en-US", Min: 0.5, Max: 100}, want: true},
		{name: "Within range and greater than 0", param1: "99.99", param2: IsFloatOpts{Locale: "en-US", Min: 0.5, Max: 100}, want: true},
		{name: "Below min requirement", param1: "0.45", param2: IsFloatOpts{Locale: "en-US", Min: 0.5, Max: 100}, want: false},
		{name: "Greater than max", param1: "120", param2: IsFloatOpts{Locale: "en-US", Min: 0.5, Max: 100}, want: false},
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
