package validatorgo

import "testing"

func TestIsHash(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   bool
	}{

		// Valid CRC32 hashes
		{name: "Valid CRC32 - all lowercase", param1: "d202ef8d", param2: "crc32", want: true},
		{name: "Valid CRC32 - mixed case", param1: "D202EF8D", param2: "crc32", want: true},
		{name: "Valid CRC32 - numeric", param1: "12345678", param2: "crc32", want: true},

		// Invalid CRC32 hashes
		{name: "Invalid CRC32 - too short", param1: "d202ef8", param2: "crc32", want: false},
		{name: "Invalid CRC32 - too long", param1: "d202ef8d9", param2: "crc32", want: false},
		{name: "Invalid CRC32 - contains non-hex character", param1: "d202ef8z", param2: "crc32", want: false},
		{name: "Invalid CRC32 - empty string", param1: "", param2: "crc32", want: false},
		{name: "Invalid CRC32 - non-hex characters", param1: "g202ef8d", param2: "crc32", want: false},

		// Valid CRC32b hashes
		{name: "Valid CRC32b - all lowercase", param1: "d202ef8d", param2: "crc32b", want: true},
		{name: "Valid CRC32b - mixed case", param1: "D202EF8D", param2: "crc32b", want: true},
		{name: "Valid CRC32b - numeric", param1: "12345678", param2: "crc32b", want: true},

		// Invalid CRC32b hashes
		{name: "Invalid CRC32b - too short", param1: "d202ef8", param2: "crc32b", want: false},
		{name: "Invalid CRC32b - too long", param1: "d202ef8d9", param2: "crc32b", want: false},
		{name: "Invalid CRC32b - contains non-hex character", param1: "d202ef8z", param2: "crc32b", want: false},
		{name: "Invalid CRC32b - empty string", param1: "", param2: "crc32b", want: false},
		{name: "Invalid CRC32b - non-hex characters", param1: "g202ef8d", param2: "crc32b", want: false},

		// Valid MD4 hashes
		{name: "Valid MD4 - all lowercase", param1: "c3fcd3d76192e4007dfb496cca67e13b", param2: "md4", want: true},
		{name: "Valid MD4 - mixed case", param1: "C3FCD3D76192E4007DFB496CCA67E13B", param2: "md4", want: true},
		{name: "Valid MD4 - numeric", param1: "12345678901234567890123456789012", param2: "md4", want: true},

		// Invalid MD4 hashes
		{name: "Invalid MD4 - too short", param1: "c3fcd3d76192e4007dfb496cca67e13", param2: "md4", want: false},
		{name: "Invalid MD4 - too long", param1: "c3fcd3d76192e4007dfb496cca67e13b1", param2: "md4", want: false},
		{name: "Invalid MD4 - contains non-hex character", param1: "c3fcd3d76192e4007dfb496cca67e13g", param2: "md4", want: false},
		{name: "Invalid MD4 - empty string", param1: "", param2: "md4", want: false},
		{name: "Invalid MD4 - non-hex characters", param1: "g3fcd3d76192e4007dfb496cca67e13b", param2: "md4", want: false},

		// Valid MD5 hashes
		{name: "Valid MD5 - all lowercase", param1: "9e107d9d372bb6826bd81d3542a419d6", param2: "md5", want: true},
		{name: "Valid MD5 - mixed case", param1: "9E107D9D372BB6826BD81D3542A419D6", param2: "md5", want: true},
		{name: "Valid MD5 - numeric", param1: "12345678901234567890123456789012", param2: "md5", want: true},

		// Invalid MD5 hashes
		{name: "Invalid MD5 - too short", param1: "9e107d9d372bb6826bd81d3542a419d", param2: "md5", want: false},
		{name: "Invalid MD5 - too long", param1: "9e107d9d372bb6826bd81d3542a419d61", param2: "md5", want: false},
		{name: "Invalid MD5 - contains non-hex character", param1: "9e107d9d372bb6826bd81d3542a419dz", param2: "md5", want: false},
		{name: "Invalid MD5 - empty string", param1: "", param2: "md5", want: false},
		{name: "Invalid MD5 - non-hex characters", param1: "g3fcd3d76192e4007dfb496cca67e13b", param2: "md5", want: false},

		// Valid ripemd128 hashes
		{name: "Valid ripemd128 - all lowercase", param1: "d202ef8d0f87f6a4b3eb58c49f1cb7d4", param2: "ripemd128", want: true},
		{name: "Valid ripemd128 - mixed case", param1: "D202EF8D0F87F6A4B3EB58C49F1CB7D4", param2: "ripemd128", want: true},
		{name: "Valid ripemd128 - numeric", param1: "1234567890abcdef1234567890abcdef", param2: "ripemd128", want: true},

		// Invalid ripemd128 hashes
		{name: "Invalid ripemd128 - too short", param1: "d202ef8d0f87f6a4b3eb58c49f1cb7d", param2: "ripemd128", want: false},
		{name: "Invalid ripemd128 - too long", param1: "d202ef8d0f87f6a4b3eb58c49f1cb7d41", param2: "ripemd128", want: false},
		{name: "Invalid ripemd128 - contains non-hex character", param1: "d202ef8d0f87f6a4b3eb58c49f1cb7dz", param2: "ripemd128", want: false},

		// Valid ripemd160 hashes
		{name: "Valid ripemd160 - all lowercase", param1: "f02e23a7e1e86f288f482144994429e7df36e14f", param2: "ripemd160", want: true},
		{name: "Valid ripemd160 - mixed case", param1: "F02E23A7E1E86F288F482144994429E7DF36E14F", param2: "ripemd160", want: true},
		{name: "Valid ripemd160 - numeric", param1: "1234567890123456789012345678901234567890", param2: "ripemd160", want: true},

		// Invalid ripemd160 hashes
		{name: "Invalid ripemd160 - too short", param1: "f02e23a7e1e86f288f482144994429e7df36e14", param2: "ripemd160", want: false},
		{name: "Invalid ripemd160 - too long", param1: "f02e23a7e1e86f288f482144994429e7df36e14f1", param2: "ripemd160", want: false},
		{name: "Invalid ripemd160 - contains non-hex character", param1: "f02e23a7e1e86f288f482144994429e7df36e14z", param2: "ripemd160", want: false},

		// Valid sha1 hashes
		{name: "Valid sha1 - all lowercase", param1: "a9993e364706816aba3e25717850c26c9cd0d89d", param2: "sha1", want: true},
		{name: "Valid sha1 - mixed case", param1: "A9993E364706816ABA3E25717850C26C9CD0D89D", param2: "sha1", want: true},
		{name: "Valid sha1 - numeric", param1: "1234567890123456789012345678901234567890", param2: "sha1", want: true},

		// Invalid sha1 hashes
		{name: "Invalid sha1 - too short", param1: "a9993e364706816aba3e25717850c26c9cd0d89", param2: "sha1", want: false},
		{name: "Invalid sha1 - too long", param1: "a9993e364706816aba3e25717850c26c9cd0d89d1", param2: "sha1", want: false},
		{name: "Invalid sha1 - contains non-hex character", param1: "a9993e364706816aba3e25717850c26c9cd0d89z", param2: "sha1", want: false},

		// Valid sha256 hashes
		{name: "Valid sha256 - all lowercase", param1: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", param2: "sha256", want: true},
		{name: "Valid sha256 - mixed case", param1: "BA7816BF8F01CFEA414140DE5DAE2223B00361A396177A9CB410FF61F20015AD", param2: "sha256", want: true},
		{name: "Valid sha256 - numeric", param1: "1234567890123456789012345678901234567890123456789012345678901234", param2: "sha256", want: true},

		// Invalid sha256 hashes
		{name: "Invalid sha256 - too short", param1: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015a", param2: "sha256", want: false},
		{name: "Invalid sha256 - too long", param1: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad1", param2: "sha256", want: false},
		{name: "Invalid sha256 - contains non-hex character", param1: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015az", param2: "sha256", want: false},

		// Valid sha384 hashes
		{name: "Valid sha384 - all lowercase", param1: "e8e4d9727695438c7f5c91347e50e3d68eaab5fe3f856685de5a80fbaafb3c1700776dea0eb7db09c940466ba270a4e4", param2: "sha384", want: true},
		{name: "Valid sha384 - mixed case", param1: "1f0ee3e77c92afdc0650e333b4b18816f7f54d7237ac766fefb5e0e9e39cd1f824a0f7de0481417cc106f2c7e73aa402", param2: "sha384", want: true},
		{name: "Valid sha384 - numeric", param1: "e4d966e14a785984f8c5d4789ebaf4f00ad3153e9e106ff3f896b356da0022fe88a6d60014fcf841516966fd6bcfbba4", param2: "sha384", want: true},

		// Invalid sha384 hashes
		{name: "Invalid sha384 - too short", param1: "cb00753f9e5e5dd58b2931b1e6c79b2bdc74b7f598380bda56dfadae4058b8c7e", param2: "sha384", want: false},
		{name: "Invalid sha384 - too long", param1: "cb00753f9e5e5dd58b2931b1e6c79b2bdc74b7f598380bda56dfadae4058b8c7ee1", param2: "sha384", want: false},
		{name: "Invalid sha384 - contains non-hex character", param1: "cb00753f9e5e5dd58b2931b1e6c79b2bdc74b7f598380bda56dfadae4058b8c7ez", param2: "sha384", want: false},

		// Invalid sha512 hashes
		{name: "Invalid sha512 - too short", param1: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3", param2: "sha512", want: false},
		{name: "Invalid sha512 - too long", param1: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e1", param2: "sha512", want: false},
		{name: "Invalid sha512 - contains non-hex character", param1: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3z", param2: "sha512", want: false},

		// Valid tiger128 hashes
		{name: "Valid tiger128 - all lowercase", param1: "329c83258a7469b6d21dd87dd0518e5d", param2: "tiger128", want: true},
		{name: "Valid tiger128 - mixed case", param1: "329C83258A7469B6D21DD87DD0518E5D", param2: "tiger128", want: true},
		{name: "Valid tiger128 - numeric", param1: "1234567890abcdef1234567890abcdef", param2: "tiger128", want: true},

		// Invalid tiger128 hashes
		{name: "Invalid tiger128 - too short", param1: "329c83258a7469b6d21dd87dd0518e5", param2: "tiger128", want: false},
		{name: "Invalid tiger128 - too long", param1: "329c83258a7469b6d21dd87dd0518e5d1", param2: "tiger128", want: false},
		{name: "Invalid tiger128 - contains non-hex character", param1: "329c83258a7469b6d21dd87dd0518e5z", param2: "tiger128", want: false},

		// Valid tiger160 hashes
		{name: "Valid tiger160 - all lowercase", param1: "329c83258a7469b6d21dd87dd0518e5d6b04c810", param2: "tiger160", want: true},
		{name: "Valid tiger160 - mixed case", param1: "329C83258A7469B6D21DD87DD0518E5D6B04C810", param2: "tiger160", want: true},

		// Invalid tiger160 hashes
		{name: "Invalid tiger160 - too short", param1: "329c83258a7469b6d21dd87dd0518e5d6b04c81", param2: "tiger160", want: false},
		{name: "Invalid tiger160 - too long", param1: "329c83258a7469b6d21dd87dd0518e5d6b04c8101", param2: "tiger160", want: false},
		{name: "Invalid tiger160 - contains non-hex character", param1: "329c83258a7469b6d21dd87dd0518e5d6b04c81z", param2: "tiger160", want: false},

		// Valid tiger192 hashes
		{name: "Valid tiger192 - all lowercase", param1: "abcdef0123456789abcdef0123456789abcdef0123456789", param2: "tiger192", want: true},
		{name: "Valid tiger192 - mixed case", param1: "ABCDEF0123456789abcdef0123456789ABCDEF0123456789", param2: "tiger192", want: true},
		{name: "Valid tiger192 - numeric", param1: "1234567890abcdef1234567890abcdef1234567890abcdef", param2: "tiger192", want: true},

		// Invalid tiger192 hashes
		{name: "Invalid tiger192 - too short", param1: "329c83258a7469b6d21dd87dd0518e5d6b04c810d8eabf64752abfabc5adf8ff2", param2: "tiger192", want: false},
		{name: "Invalid tiger192 - too long", param1: "329c83258a7469b6d21dd87dd0518e5d6b04c810d8eabf64752abfabc5adf8ff2e1", param2: "tiger192", want: false},
		{name: "Invalid tiger192 - contains non-hex character", param1: "329c83258a7469b6d21dd87dd0518e5d6b04c810d8eabf64752abfabc5adf8ff2e1g", param2: "tiger192", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsHash(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
