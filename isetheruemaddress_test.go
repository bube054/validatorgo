package validatorgo

import "testing"

func TestIsEthereumAddress(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{
			name:   "Valid address with mixed case letters and numbers",
			param1: "0xeA0B9657892321121287128712BC78A89F989AAA",
			want:   true,
		},
		{
			name:   "Valid address with all uppercase letters and numbers",
			param1: "0xBBAC6AABCFEBACEBCEABCEACB76767867676ACCC",
			want:   true,
		},
		{
			name:   "Invalid address with non-hex characters",
			param1: "0xiuahbsndakjsd",
			want:   false,
		},
		{
			name:   "Invalid address with incorrect length and special characters",
			param1: "0zzFGXD2E$",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsEthereumAddress(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
