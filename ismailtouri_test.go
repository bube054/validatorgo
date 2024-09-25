package validatorgo

import "testing"

func TestIsMailtoURI(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsMailToURIOpts
		want   bool
	}{
		// Valid mailto URIs
		{name: "Basic mailto URI", param1: "mailto:someone@example.com", param2: IsMailToURIOpts{}, want: true},
		{name: "mailto URI with country code TLD", param1: "mailto:user@domain.co.uk", param2: IsMailToURIOpts{}, want: true},
		{name: "mailto URI with subject", param1: "mailto:admin@domain.com?subject=Hello", param2: IsMailToURIOpts{}, want: true},
		{name: "mailto URI with body", param1: "mailto:support@service.com?body=Help", param2: IsMailToURIOpts{}, want: true},
		{name: "mailto URI with cc and bcc", param1: "mailto:test@domain.com?cc=cc@domain.com&bcc=bcc@domain.com", param2: IsMailToURIOpts{}, want: true},
		// Invalid mailto URIs
		{name: "Missing mailto scheme", param1: "someone@example.com", param2: IsMailToURIOpts{}, want: false},
		{name: "Invalid email format", param1: "mailto:invalid@domain", param2: IsMailToURIOpts{}, want: false},
		{name: "Non-mailto URI", param1: "http://example.com", param2: IsMailToURIOpts{}, want: false},
		{name: "Invalid domain in email", param1: "mailto:user@.com", param2: IsMailToURIOpts{}, want: false},
		{name: "Missing domain extension", param1: "mailto:admin@domain?subject=Hello", param2: IsMailToURIOpts{}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsMailtoURI(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
