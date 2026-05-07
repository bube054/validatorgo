package validatorgo

import (
	"regexp"
	"sync"
)

var regexCache sync.Map

// cachedCompile returns a pre-compiled *regexp.Regexp for the given key.
// If the key has not been seen before, it compiles the pattern and caches it.
// This is used for dynamic patterns where the regex depends on runtime values.
func cachedCompile(key, pattern string) *regexp.Regexp {
	if v, ok := regexCache.Load(key); ok {
		return v.(*regexp.Regexp)
	}
	re := regexp.MustCompile(pattern)
	regexCache.Store(key, re)
	return re
}
