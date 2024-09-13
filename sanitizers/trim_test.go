package sanitizers

import (
	"testing"
)

func TestTrim(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{
			"test 1",
			"  Hello World!  ",
			"Hello World!",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Trim(test.param1)

			if result != test.want {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}

func TestLTrim(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{
			"test 1",
			"  Hello World!  ",
			"Hello World!  ",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Ltrim(test.param1)

			if result != test.want {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}

func TestRTrim(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{
			"test 1",
			"  Hello World!  ",
			"  Hello World!",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Rtrim(test.param1)

			if result != test.want {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}
