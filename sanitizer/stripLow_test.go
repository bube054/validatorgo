package sanitizer

import (
	"reflect"
	"testing"
)

func TestStripLow(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 bool
		want   string
	}{
		{name: "Remove low ASCII chars", param1: "Hello\x00World", param2: false, want: "HelloWorld"},
		{name: "Keep newlines, remove other low ASCII chars", param1: "Hello\x00\nWorld", param2: true, want: "Hello\nWorld"},
		{name: "No low ASCII chars to strip", param1: "HelloWorld", param2: false, want: "HelloWorld"},
		{name: "Remove low ASCII chars, keep newlines", param1: "Hi\x00\n\x01There", param2: true, want: "Hi\nThere"},
		{name: "Strip all low ASCII chars including newlines", param1: "Text\x00\nWith\x01\x02Low", param2: false, want: "TextWithLow"},
		{name: "Remove only control characters, keep newlines", param1: "\x01Hello\nWorld\x02", param2: true, want: "Hello\nWorld"},
		{name: "String with no low characters", param1: "NormalText123", param2: false, want: "NormalText123"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := StripLow(test.param1, test.param2)

			// if result != test.want {
			if !reflect.DeepEqual(result, test.want) {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}
