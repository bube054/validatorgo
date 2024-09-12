package validatorgo

import "regexp"

// A validator checks if the string is an Ethereum address. Does not validate address checksums.
func IsEthereumAddress(str string) bool {
	return regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(str)
}
