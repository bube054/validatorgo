package validatorgo

import "testing"

func TestIsMimeType(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid mime types
		{name: "Valid text MIME type", param1: "text/plain", want: true},
		{name: "Valid image MIME type", param1: "image/jpeg", want: true},
		{name: "Valid application MIME type", param1: "application/json", want: true},
		{name: "Valid audio MIME type", param1: "audio/mpeg", want: true},
		{name: "Valid video MIME type", param1: "video/mp4", want: true},
		{name: "Vendor-specific MIME type", param1: "application/vnd.ms-excel", want: true},
		{name: "Vendor-specific with personal prefix", param1: "application/x-my-custom", want: true},
		{name: "MIME type with charset", param1: "text/html; charset=UTF-8", want: true},
		{name: "MIME type with boundary", param1: "multipart/form-data; boundary=something", want: true},
		{name: "MIME type with plus symbol", param1: "application/ld+json", want: true},
		{name: "MIME type with dot", param1: "image/vnd.adobe.photoshop", want: true},

		// Invalid mime types
		{name: "Missing slash", param1: "textplain", want: false},
		{name: "Multiple slashes", param1: "text//plain", want: false},
		{name: "Invalid characters in type", param1: "text/pla!n", want: false},
		{name: "Invalid characters in subtype", param1: "image/jp@g", want: false},
		{name: "Missing type", param1: "/plain", want: false},
		{name: "Missing subtype", param1: "image/", want: false},
		{name: "Invalid parameter format", param1: "text/html;charset", want: false},
		{name: "Malformed parameter key-value", param1: "application/json; charset", want: false},
		{name: "Extra semicolon", param1: "application/json;", want: false},
		{name: "Non-alphanumeric subtype", param1: "application/123", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMimeType(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
