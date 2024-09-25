package validatorgo

import "regexp"

// A validator checks if the string is an [Ethereum address]. Does not validate address checksums.
//
//	ok := validatorgo.IsDivisibleBy("0xeA0B9657892321121287128712BC78A89F989AAA")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsDivisibleBy("0xiuahbsndakjsd")
//	fmt.Println(ok) // false
//
// [Ethereum address]: https://ethereum.org/
func IsEthereumAddress(str string) bool {
	return regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`).MatchString(str)
}
