package sanitizer

import (
	"regexp"
	"sync"
)

var sanitizerRegexCache sync.Map

// cachedCompile returns a pre-compiled *regexp.Regexp for the given key.
// If the key has not been seen before, it compiles the pattern and caches it.
func cachedCompile(key, pattern string) *regexp.Regexp {
	if v, ok := sanitizerRegexCache.Load(key); ok {
		return v.(*regexp.Regexp)
	}
	re := regexp.MustCompile(pattern)
	sanitizerRegexCache.Store(key, re)
	return re
}
