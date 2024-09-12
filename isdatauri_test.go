package validatorgo

import "testing"

func TestIsDataURI(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "The text/plain data Hello, World!",
			param1: "data:,Hello%2C%20World%21",
			want:   true,
		},
		{
			name:   "base64-encoded version of the above",
			param1: "data:text/plain;base64,SGVsbG8sIFdvcmxkIQ==",
			want:   true,
		},
		{
			name:   "An HTML document with <h1>Hello, World!</h1>",
			param1: "data:text/html,%3Ch1%3EHello%2C%20World%21%3C%2Fh1%3E",
			want:   true,
		},
		{
			name:   "An HTML document with <script>alert('hi');</script>",
			param1: "data:text/html,%3Cscript%3Ealert%28%27hi%27%29%3B%3C%2Fscript%3E",
			want:   true,
		},
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
