package validatorgo

import "regexp"

// hashAlgoRegex is the set of regex's for hashing algo
var hashAlgoRegex = map[string]*regexp.Regexp{
	"crc32":     regexp.MustCompile(`^[a-fA-F0-9]{8}$`),   // 8 hex characters
	"crc32b":    regexp.MustCompile(`^[a-fA-F0-9]{8}$`),   // 8 hex characters
	"md4":       regexp.MustCompile(`^[a-fA-F0-9]{32}$`),  // 32 hex characters
	"md5":       regexp.MustCompile(`^[a-fA-F0-9]{32}$`),  // 32 hex characters
	"ripemd128": regexp.MustCompile(`^[a-fA-F0-9]{32}$`),  // 32 hex characters
	"ripemd160": regexp.MustCompile(`^[a-fA-F0-9]{40}$`),  // 40 hex characters
	"sha1":      regexp.MustCompile(`^[a-fA-F0-9]{40}$`),  // 40 hex characters
	"sha256":    regexp.MustCompile(`^[a-fA-F0-9]{64}$`),  // 64 hex characters
	"sha384":    regexp.MustCompile(`^[a-fA-F0-9]{96}$`),  // 96 hex characters
	"sha512":    regexp.MustCompile(`^[a-fA-F0-9]{128}$`), // 128 hex characters
	"tiger128":  regexp.MustCompile(`^[a-fA-F0-9]{32}$`),  // 32 hex characters
	"tiger160":  regexp.MustCompile(`^[a-fA-F0-9]{40}$`),  // 40 hex characters
	"tiger192":  regexp.MustCompile(`^[a-fA-F0-9]{48}$`),  // 48 hex characters
}

// A validator that checks if the string is a hash of type algorithm.
//
// Algorithm is one of ('crc32', 'crc32b', 'md4', 'md5', 'ripemd128', 'ripemd160', 'sha1', 'sha256', 'sha384', 'sha512', 'tiger128', 'tiger160', 'tiger192'), No checksum are calculated.
//
//	ok := validatorgo.IsHash("d202ef8d", "crc32")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsHash("d202ef8", "crc32")
//	fmt.Println(ok) // false
func IsHash(str, algorithm string) bool {
	re, exist := hashAlgoRegex[algorithm]

	if !exist {
		return false
	}

	return re.MatchString(str)
}
