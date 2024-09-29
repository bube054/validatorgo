package sanitizer

import (
	"fmt"
	"strings"
	"testing"
)

func TestToFloat(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   float64
		err    error
	}{
		{name: "Valid float string", param1: "123.45", want: 123.45, err: nil},
		{name: "Valid negative float string", param1: "-987.65", want: -987.65, err: nil},
		{name: "Valid integer string", param1: "42", want: 42.0, err: nil},
		{name: "Valid float string with scientific notation", param1: "1.23e4", want: 12300.0, err: nil},
		{name: "Valid zero float string", param1: "0.0", want: 0.0, err: nil},
		{name: "Valid negative zero float string", param1: "-0.0", want: -0.0, err: nil},
		{name: "Invalid float string: letters", param1: "abc", want: 0.0, err: fmt.Errorf("invalid syntax")},
		{name: "Invalid float string: symbols", param1: "#$%", want: 0.0, err: fmt.Errorf("invalid syntax")},
		{name: "Empty string", param1: "", want: 0.0, err: fmt.Errorf("invalid syntax")},
		{name: "String with spaces", param1: " 123.45 ", want: 123.45, err: nil},
		{name: "String with commas", param1: "1,234.56", want: 0.0, err: fmt.Errorf("invalid syntax")},
		{name: "String with multiple dots", param1: "12.34.56", want: 0.0, err: fmt.Errorf("invalid syntax")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ToFloat(test.param1)

			if err != nil {
				errMsg := err.Error()

				if strings.HasSuffix(errMsg, "invalid syntax") {
					return
				}
			}

			if result != test.want || err != test.err {
				t.Errorf("got `%.2f`, wanted `%.2f`", result, test.want)
			}
		})
	}
}
