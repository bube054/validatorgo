package validatorgo

import "regexp"

// A validator that checks if the string is a [Magnet URI] format.
//
//	ok := validatorgo.IsMagnetURI("magnet:?xt=urn:btih:123456789abcdef123456789abcdef123456789a")
//	fmt.Println(ok) // true
//	ok := validatorgo.IsMagnetURI("magnet:?dn=Example&tr=http://example.com/announce")
//	fmt.Println(ok) // false
//
// [Magnet URI]: https://en.wikipedia.org/wiki/Magnet_URI_scheme
func IsMagnetURI(str string) bool {
	return regexp.MustCompile(`(?:^magnet:\?|[^?&]&)xt(?:\.1)?=urn:(?:(?:aich|bitprint|btih|ed2k|ed2khash|kzhash|md5|sha1|tree:tiger):[a-z0-9]{32}(?:[a-z0-9]{8})?|btmh:1220[a-z0-9]{64})(?:$|&)`).MatchString(str)
}
