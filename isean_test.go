package validatorgo

import "testing"

func TestIsEAN(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "valid EAN-13",
			param1: "4006381333931",
			want:   true,
		},
		{
			name:   "valid EAN-8",
			param1: "73513537",
			want:   true,
		},
		{
			name:   "invalid EAN-13, checksum fails",
			param1: "4006381333932",
			want:   false,
		},
		{
			name:   "invalid EAN-8",
			param1: "12345678",
			want:   false,
		},
		{
			name:   "12 digits, not a valid EAN",
			param1: "123456789012",
			want:   false,
		},
		{
			name:   "contains a letter, invalid",
			param1: "40063A1333931",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsEAN(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
