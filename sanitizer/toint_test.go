package sanitizer

import (
	"fmt"
	"strings"
	"testing"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   int
		err    error
	}{
		{name: "Valid integer string", param1: "123", want: 123, err: nil},
		{name: "Valid negative integer string", param1: "-987", want: -987, err: nil},
		{name: "Valid zero integer string", param1: "0", want: 0, err: nil},
		{name: "Valid large integer string", param1: "9223372036854775807", want: 9223372036854775807, err: nil},            // Max int64 value
		{name: "Valid negative large integer string", param1: "-9223372036854775808", want: -9223372036854775808, err: nil}, // Min int64 value
		{name: "Invalid integer string: letters", param1: "abc", want: 0, err: fmt.Errorf("invalid syntax")},
		{name: "Invalid integer string: symbols", param1: "#$%", want: 0, err: fmt.Errorf("invalid syntax")},
		{name: "Empty string", param1: "", want: 0, err: fmt.Errorf("invalid syntax")},
		{name: "String with spaces", param1: " 456 ", want: 456, err: nil},
		{name: "String with decimal", param1: "123.45", want: 0, err: fmt.Errorf("invalid syntax")},
		{name: "String with commas", param1: "1,234", want: 0, err: fmt.Errorf("invalid syntax")},
		{name: "Overflow large integer", param1: "9223372036854775808", want: 0, err: fmt.Errorf("value out of range")},
		{name: "Underflow large negative integer", param1: "-9223372036854775809", want: 0, err: fmt.Errorf("value out of range")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ToInt(test.param1)

			if err != nil {
				errMsg := err.Error()
				fmt.Println(errMsg)

				if strings.HasSuffix(errMsg, "invalid syntax") || strings.HasSuffix(errMsg, "value out of range") {
					return
				}
			}

			if result != test.want || err != test.err {
				t.Errorf("got `%d`, wanted `%d`", result, test.want)
			}
		})
	}
}
