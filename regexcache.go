package validatorgo

import (
	"regexp"
	"sync"
)

var regexCache sync.Map // pattern -> *regexp.Regexp

// compileOnce returns the compiled regexp for pattern, compiling it only the first
// time a pattern is seen. Compiling is far slower than matching, so validators call
// this instead of regexp.MustCompile on every call. Patterns come from the library
// and from option values, never from the string being validated, so the cache stays
// small.
func compileOnce(pattern string) *regexp.Regexp {
	if v, ok := regexCache.Load(pattern); ok {
		return v.(*regexp.Regexp)
	}
	re := regexp.MustCompile(pattern)
	regexCache.Store(pattern, re)
	return re
}

// tryCompileOnce is compileOnce for patterns that may be invalid, such as ones built
// from option values. Only successful compiles are cached.
func tryCompileOnce(pattern string) (*regexp.Regexp, error) {
	if v, ok := regexCache.Load(pattern); ok {
		return v.(*regexp.Regexp), nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	regexCache.Store(pattern, re)
	return re, nil
}
