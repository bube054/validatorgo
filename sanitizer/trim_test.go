package sanitizer

import (
	"testing"
)

func TestTrim(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   string
	}{
		{name: "Trim whitespace", param1: "  Hello World  ", param2: "", want: "Hello World"},
		{name: "Trim no whitespace", param1: "HelloWorld", param2: "", want: "HelloWorld"},
		{name: "Trim custom character '-' from both ends", param1: "--Hello World--", param2: "-", want: "Hello World"},
		{name: "Trim custom character '#' from both ends", param1: "###Hello World###", param2: "#", want: "Hello World"},
		{name: "Trim multiple custom characters '-_#'", param1: "-_-Hello World-_--", param2: "-_#", want: "Hello World"},
		{name: "Trim empty string", param1: "", param2: "", want: ""},
		{name: "Trim custom character '-' with no match", param1: "Hello World", param2: "-", want: "Hello World"},
		{name: "Trim only leading custom character '-'", param1: "--Hello World", param2: "-", want: "Hello World"},
		{name: "Trim only trailing custom character '-'", param1: "Hello World--", param2: "-", want: "Hello World"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Trim(test.param1, test.param2)

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
		param2 string
		want   string
	}{
		{name: "LTrim whitespace", param1: "  Hello World  ", param2: "", want: "Hello World  "},
		{name: "LTrim no whitespace", param1: "HelloWorld", param2: "", want: "HelloWorld"},
		{name: "LTrim custom character '-' from the left", param1: "--Hello World--", param2: "-", want: "Hello World--"},
		{name: "LTrim custom character '#' from the left", param1: "###Hello World###", param2: "#", want: "Hello World###"},
		{name: "LTrim multiple custom characters '-_#'", param1: "-_-Hello World-_--", param2: "-_#", want: "Hello World-_--"},
		{name: "LTrim empty string", param1: "", param2: "", want: ""},
		{name: "LTrim custom character '-' with no match", param1: "Hello World", param2: "-", want: "Hello World"},
		{name: "LTrim spaces with custom character '-' ignored", param1: "  --Hello World--", param2: "-", want: "  --Hello World--"},
		{name: "LTrim custom character '-' from left only", param1: "--Hello World", param2: "-", want: "Hello World"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := LTrim(test.param1, test.param2)

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
		param2 string
		want   string
	}{
		{name: "RTrim whitespace", param1: "  Hello World  ", param2: "", want: "  Hello World"},
		{name: "RTrim no whitespace", param1: "HelloWorld", param2: "", want: "HelloWorld"},
		{name: "RTrim custom character '-' from the right", param1: "--Hello World--", param2: "-", want: "--Hello World"},
		{name: "RTrim custom character '#' from the right", param1: "###Hello World###", param2: "#", want: "###Hello World"},
		{name: "RTrim multiple custom characters '-_#'", param1: "-_-Hello World-_--", param2: "-_#", want: "-_-Hello World"},
		{name: "RTrim empty string", param1: "", param2: "", want: ""},
		{name: "RTrim custom character '-' with no match", param1: "Hello World", param2: "-", want: "Hello World"},
		{name: "RTrim spaces with custom character '-' ignored", param1: "--Hello World--  ", param2: "-", want: "--Hello World--  "},
		{name: "RTrim custom character '-' from right only", param1: "Hello World--", param2: "-", want: "Hello World"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := RTrim(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%s`, wanted `%s`", result, test.want)
			}
		})
	}
}
