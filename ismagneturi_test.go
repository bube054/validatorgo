package validatorgo

import "testing"

func TestIsMagnetURI(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid magnet links
		{name: "Simple Magnet Link with hash", param1: "magnet:?xt=urn:btih:123456789abcdef123456789abcdef123456789a", want: true},
		{name: "Magnet Link with multiple parameters", param1: "magnet:?xt=urn:btih:123456789abcdef123456789abcdef123456789a&dn=Example&tr=http://example.com/announce&xl=123456789", want: true},
		{name: "Magnet Link with different hash (ed2k)", param1: "magnet:?xt=urn:ed2k:123456789abcdef123456789abcdef12", want: true},
		{name: "Magnet Link with multiple trackers", param1: "magnet:?xt=urn:btih:123456789abcdef123456789abcdef123456789a&tr=http://example.com/announce&tr=http://example.org/announce", want: true},
		// Invalid magnet links
		{name: "Missing xt parameter", param1: "magnet:?dn=Example&tr=http://example.com/announce", want: false},
		{name: "Invalid hash length", param1: "magnet:?xt=urn:btih:12345", want: false},
		{name: "Missing the magnet scheme", param1: "?xt=urn:btih:123456789abcdef123456789abcdef123456789a", want: false},
		{name: "Invalid URN prefix", param1: "magnet:?xt=urn:invalid:123456789abcdef123456789abcdef123456789a", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMagnetURI(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
