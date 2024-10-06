package validatorgo

import "testing"

func TestIsLatLong(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsLatLongOpts
		want   bool
	}{
		// Valid latlong default options
		{name: "Standard latitude/longitude in decimal format", param1: "40.730610,-73.935242", want: true},
		{name: "Max valid latitude and longitude", param1: "90,180", want: true},
		{name: "Min valid latitude and longitude", param1: "-90,-180", want: true},
		{name: "Valid with spaces after the comma", param1: "40.730610, 73.935242", want: true},
		{name: "Equator and Prime Meridian", param1: "0,0", want: true},
		{name: "High precision, valid within range", param1: "47.1231231,179.99999999", want: true},
		// Invalid latlong default options
		{name: "Latitude and longitude out of range", param1: "91,181", want: false},
		{name: "Latitude and longitude out of range", param1: "-91,-181", want: false},
		{name: "Incorrect separator ; instead of a comma", param1: "40.730610;73.935242", want: false},
		{name: "Longitude exceeds valid range", param1: "40.730610, 200", want: false},
		{name: "Not a numeric pair", param1: "abc,xyz", want: false},
		{name: "Missing longitude value", param1: "40.730610, ", want: false},
		{name: "Missing latitude value", param1: " ,73.935242", want: false},
		{name: "Incomplete longitude", param1: "40.730610,-", want: false},
		{name: "Latitude exceeds the maximum allowed value", param1: "91,180", want: false},
		{name: "Longitude with too many decimal places", param1: "90.1,180", want: false},
		// Valid latlong withCheckDMS
		{name: "Simple valid DMS", param1: `40°44'55"N, 73°59'11"W`, param2: &IsLatLongOpts{CheckDMS: true}, want: true},
		{name: "With decimal seconds", param1: `40°44'55.5"N, 73°59'11.2"W`, param2: &IsLatLongOpts{CheckDMS: true}, want: true},
		{name: "South and East", param1: `40°44'55"S, 73°59'11"E`, param2: &IsLatLongOpts{CheckDMS: true}, want: true},
		// Invalid latlong withCheckDMS
		{name: "Minutes exceeding 59", param1: `40°60'00"N, 73°60'00"W`, param2: &IsLatLongOpts{CheckDMS: true}, want: false},
		{name: "Missing direction N/S or E/W", param1: `40°44'55, 73°59'11"`, param2: &IsLatLongOpts{CheckDMS: true}, want: false},
		{name: "Latitude exceeding valid range", param1: `100°00'00"N, 73°59'11"W`, param2: &IsLatLongOpts{CheckDMS: true}, want: false},
		{name: "Incorrect separator ;", param1: `40°44'55.5"N; 73°59'11.2"W`, param2: &IsLatLongOpts{CheckDMS: true}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsLatLong(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
