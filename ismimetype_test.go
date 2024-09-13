package validatorgo

import "testing"

func TestIsMimeType(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid mime types
		{name: "Valid MIME Type - JSON", param1: "application/json", want: true},
		{name: "Valid MIME Type - PNG Image", param1: "image/png", want: true},
		{name: "Valid MIME Type - HTML Text", param1: "text/html", want: true},
		{name: "Valid MIME Type - MPEG Audio", param1: "audio/mpeg", want: true},
		{name: "Valid MIME Type - MP4 Video", param1: "video/mp4", want: true},
		{name: "Valid MIME Type - Excel Spreadsheet", param1: "application/vnd.ms-excel", want: true},
		// Invalid mime types
		{name: "Invalid MIME Type - Trailing Slash", param1: "application/json/", want: false},
		{name: "Invalid MIME Type - Missing Subtype", param1: "text", want: false},
		{name: "Invalid MIME Type - Incomplete Image Type", param1: "image", want: false},
		{name: "Invalid MIME Type - Incorrect Format (Dash)", param1: "audio-mpeg", want: false},
		{name: "Invalid MIME Type - Invalid Character", param1: "application/vnd.ms@excel", want: false},
		{name: "Invalid MIME Type - Empty String", param1: "", want: false},
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
