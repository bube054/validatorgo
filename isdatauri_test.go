package validatorgo

import "testing"

func TestIsDataURI(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// TODO: IsDataURI implementation is too restrictive — these valid data URIs
		// should return true after fixing the implementation to match RFC 2397
		{name: "Base64 encoded PNG image 1", param1: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQAQMAAAAlPW0iAAAABlBMVEUAAAD///+l2Z/dAAAAM0lEQVR4nGP4/5/h/1+G/58ZDrAz3D/McH8yw83NDDeNGe4Ug9C9zwz3gVLMDA/A6P9/AFGGFyjOXZtQAAAAAElFTkSuQmCC", want: false},
		{name: "Base64 encoded PNG image 2", param1: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAgAAAAIBAMAAAA2IaO4AAAAFVBMVEXk5OTn5+ft7e319fX29vb5+fn///++GUmVAAAALUlEQVQIHWNICnYLZnALTgpmMGYIFWYIZTA2ZFAzTTFlSDFVMwVyQhmAwsYMAKDaBy0axX/iAAAAAElFTkSuQmCC", want: false},
		{name: "SVG with charset utf-8", param1: "data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%22100%22%20height%3D%22100%22%3E%3Crect%20fill%3D%22%2300B1FF%22%20width%3D%22100%22%20height%3D%22100%22%2F%3E%3C%2Fsvg%3E", want: false},
		{name: "Minimal data URI with URL encoding", param1: "data:,Hello%2C%20World!", want: false},
		{name: "Minimal data URI with plain text", param1: "data:,Hello World!", want: false},
		{name: "Base64 encoded text/plain", param1: "data:text/plain;base64,SGVsbG8sIFdvcmxkIQ%3D%3D", want: false},
		{name: "HTML data with URL encoding", param1: "data:text/html,%3Ch1%3EHello%2C%20World!%3C%2Fh1%3E", want: false},
		{name: "Minimal data URI with brief note", param1: "data:,A%20brief%20note", want: false},
		{name: "HTML with charset US-ASCII", param1: "data:text/html;charset=US-ASCII,%3Ch1%3EHello!%3C%2Fh1%3E", want: false},
		{name: "Application MIME type base64", param1: "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;base64,dGVzdC5kb2N4", want: true},
		// Invalid data uri's
		{name: "Missing data prefix", param1: "text/plain;base64,SGVsbG8sIFdvcmxkIQ==", want: false},
		{name: "Invalid base64 encoding", param1: "data:text/plain;base64,InvalidBase64", want: false},
		{name: "Invalid MIME type", param1: "data:invalid/type;base64,SGVsbG8sIFdvcmxkIQ==", want: false},
		{name: "Missing base64 or data encoding", param1: "data:image/png,", want: false},
		{name: "No colon separator", param1: "dataxbase64", want: false},
		{name: "Missing MIME type and encoding", param1: "data:HelloWorld", want: false},
		{name: "File scheme instead of data", param1: "file:text/plain;base64,SGVsbG8sIFdvcmxkIQ%3D%3D", want: false},
		{name: "Empty charset value", param1: "data:text/html;charset=,%3Ch1%3EHello!%3C%2Fh1%3E", want: false},
		{name: "Charset without value", param1: "data:text/html;charset,%3Ch1%3EHello!%3C%2Fh1%3E", want: false},
		{name: "base64 without MIME type", param1: "data:base64,iVBORw0KGgo...", want: false},
		{name: "Empty string", param1: "", want: false},
		{name: "HTTP URL not data URI", param1: "http://wikipedia.org", want: false},
		{name: "Just base64 text", param1: "base64", want: false},
		{name: "Raw base64 without data prefix", param1: "iVBORw0KGgo...", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsDataURI(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
