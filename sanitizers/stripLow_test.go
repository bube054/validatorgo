package sanitizers

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
		{
			"test 1",
			"Hello World à Hello\x00World",
			false,
			"Hello World à HelloWorld",
		},
		{
			"test 2",
			"Hello\nWorld\rHello\x00World",
			true,
			"Hello\nWorld\rHelloWorld",
		},
		{
			"test 3",
			"Hello\nWorld\rHello\x00World",
			false,
			"HelloWorldHelloWorld",
		},
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
