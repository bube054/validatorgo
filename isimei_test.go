package validatorgo

import "testing"

func TestIsIMEI(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsIMEIOpts
		want   bool
	}{
		{name: "Valid IMEI without hyphens", param1: "490154203237518", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens, hyphens allowed", param1: "49-015420-323751-8", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043376711187", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens, hyphens allowed", param1: "35-904337-671118-7", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043377500084", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-750008-4", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043370960806", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-096080-6", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043370473776", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-047377-6", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043375890982", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-589098-2", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043374488705", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-448870-5", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043374393277", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-439327-7", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043373011391", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-301139-1", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043371423317", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-142331-7", param2: IsIMEIOpts{AllowHyphens: true}, want: true},
		{name: "Valid IMEI without hyphens", param1: "359043370066711", param2: IsIMEIOpts{AllowHyphens: false}, want: true},
		{name: "Valid IMEI with hyphens", param1: "35-904337-006671-1", param2: IsIMEIOpts{AllowHyphens: true}, want: true},

		{"Invalid IMEI with hyphens, hyphens not allowed", "49-015420-323751-8", IsIMEIOpts{AllowHyphens: false}, false},
		{"Invalid IMEI too short", "490154203237", IsIMEIOpts{AllowHyphens: false}, false},
		{"Invalid IMEI too long", "4901542032375189", IsIMEIOpts{AllowHyphens: false}, false},
		{"Invalid IMEI with letters", "49015420323ABC", IsIMEIOpts{AllowHyphens: false}, false},
		{"Invalid IMEI with letters", "49015420323ABC", IsIMEIOpts{AllowHyphens: false}, false},
		{"Invalid IMEI with correct format but wrong checksum", "359043377500085", IsIMEIOpts{AllowHyphens: false}, false},
		{"Invalid IMEI with hyphens and wrong checksum", "35-904337-750008-5", IsIMEIOpts{AllowHyphens: true}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsIMEI(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}

// IMEI.info: 359043376711187
//  359043377500084
//  359043370960806
//  359043370473776
//  359043375890982
//  359043374488705
//  359043374393277
//  359043373011391
//  359043371423317
//  359043370066711
