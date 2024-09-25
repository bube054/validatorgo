package validatorgo

import "testing"

func TestIsDataURI(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid data uri's
		// {name: "The text/plain data Hello, World!", param1: "data:,Hello%2C%20World%21", want: true},
		// {name: "Base64 encoded Hello, World!", param1: "data:text/plain;base64,SGVsbG8sIFdvcmxkIQ==", want: true},
		// {name: "Base64 encoded PNG image", param1: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAUA", want: true},
		// {name: "HTML data with URL encoding", param1: "data:text/html,%3Ch1%3EHello%2C%20World!%3C%2Fh1%3E", want: true},
		// Invalid data uri's
		{name: "Missing data prefix", param1: "text/plain;base64,SGVsbG8sIFdvcmxkIQ==", want: false},
		{name: "Invalid base64 encoding", param1: "data:text/plain;base64,InvalidBase64", want: false},
		{name: "Invalid MIME type", param1: "data:invalid/type;base64,SGVsbG8sIFdvcmxkIQ==", want: false},
		{name: "Missing base64 or data encoding", param1: "data:image/png,", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsDataURI(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
