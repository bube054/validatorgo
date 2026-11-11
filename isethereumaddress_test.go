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
			name:   "Valid address with zeros",
			param1: "0x0000000000000000000000000000000000000001",
			want:   true,
		},
		{
			name:   "Valid address 683E",
			param1: "0x683E07492fBDfDA84457C16546ac3f433BFaa128",
			want:   true,
		},
		{
			name:   "Valid address 88dA",
			param1: "0x88dA6B6a8D3590e88E0FcadD5CEC56A7C9478319",
			want:   true,
		},
		{
			name:   "Valid address 8a71",
			param1: "0x8a718a84ee7B1621E63E680371e0C03C417cCaF6",
			want:   true,
		},
		{
			name:   "Valid address FCb5",
			param1: "0xFCb5AFB808b5679b4911230Aa41FfCD0cd335b42",
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
		{
			name:   "Invalid address with non-hex GHIJK characters",
			param1: "0xGHIJK05pwm37asdf5555QWERZCXV2345AoEuIdHt",
			want:   false,
		},
		{
			name:   "Invalid address too long",
			param1: "0xFCb5AFB808b5679b4911230Aa41FfCD0cd335b422222",
			want:   false,
		},
		{
			name:   "Invalid address too short",
			param1: "0xFCb5AFB808b5679b4911230Aa41FfCD0cd33",
			want:   false,
		},
		{
			name:   "Invalid binary prefix instead of 0x",
			param1: "0b0110100001100101011011000110110001101111",
			want:   false,
		},
		{
			name:   "Invalid missing 0x prefix",
			param1: "683E07492fBDfDA84457C16546ac3f433BFaa128",
			want:   false,
		},
		{
			name:   "Invalid BTC address not ETH",
			param1: "1C6o5CDkLxjsVpnLSuqRs1UBFozXLEwYvU",
			want:   false,
		},
		{
			name:   "Empty string",
			param1: "",
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsEthereumAddress(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
