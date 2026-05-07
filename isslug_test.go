package validatorgo

import "testing"

func TestIsSlug(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid slugs
		{name: "Valid Slug - Simple", param1: "my-first-blog-post", want: true},
		{name: "Valid Slug - Single Word", param1: "product", want: true},
		{name: "Valid Slug - Numeric", param1: "product-2024", want: true},
		{name: "Valid Slug - All Numeric", param1: "123456", want: true},
		{name: "Valid Slug - Lowercase with Hyphens", param1: "a-slug-with-lowercase", want: true},
		{name: "Valid Slug - Single Char", param1: "f", want: true},
		{name: "Valid Slug - Two Chars", param1: "fo", want: true},
		{name: "Valid Slug - Three Chars", param1: "foo", want: true},
		{name: "Valid Slug - Two Words", param1: "foo-bar", want: true},
		{name: "Valid Slug - Three Words", param1: "foo-bar-foo", want: true},
		{name: "Valid Slug - Alphanumeric", param1: "foo-75-b4r-foo", want: true},
		{name: "Valid Slug - Alphanumeric Mixed", param1: "a1-b2-c3", want: true},

		// Invalid slugs
		{name: "Invalid Slug - Contains Uppercase", param1: "My-First-Blog-Post", want: false},          // Uppercase letters
		{name: "Invalid Slug - Contains Special Characters", param1: "my@blog#post!", want: false},      // Special characters like @, #, !
		{name: "Invalid Slug - Contains Underscores", param1: "my_first_blog_post", want: false},        // Underscores instead of hyphens
		{name: "Invalid Slug - Leading Hyphen", param1: "-leading-hyphen", want: false},                 // Leading hyphen
		{name: "Invalid Slug - Trailing Hyphen", param1: "trailing-hyphen-", want: false},               // Trailing hyphen
		{name: "Invalid Slug - Multiple Consecutive Hyphens", param1: "multiple--hyphens", want: false}, // Multiple consecutive hyphens
		{name: "Invalid Slug - Contains Spaces", param1: "slug with spaces", want: false},               // Spaces in slug
		{name: "Invalid Slug - Empty String", param1: "", want: false},                                  // Empty string
		{name: "Invalid Slug - Only Hyphens", param1: "---", want: false},                                          // Only hyphens
		{name: "Invalid Slug - Many Consecutive Hyphens", param1: "not-----------slug", want: false},               // Many consecutive hyphens
		{name: "Invalid Slug - Special Characters Only", param1: "@#_$@", want: false},                             // Special characters
		{name: "Invalid Slug - Leading Hyphen Alt", param1: "-not-slug", want: false},                              // Leading hyphen
		{name: "Invalid Slug - Trailing Hyphen Alt", param1: "not-slug-", want: false},                             // Trailing hyphen
		{name: "Invalid Slug - Leading Underscore", param1: "_not-slug", want: false},                              // Leading underscore
		{name: "Invalid Slug - Trailing Underscore", param1: "not-slug_", want: false},                             // Trailing underscore
		{name: "Invalid Slug - Contains Spaces Alt", param1: "not slug", want: false},                              // Spaces
		{name: "Invalid Slug - Contains Dots", param1: "i.am.not.a.slug", want: false},                             // Dots
		{name: "Invalid Slug - Dots Between Words", param1: "slug.is.cool", want: false},                           // Dots
		{name: "Invalid Slug - Stars and Special Chars", param1: "foo-bar_foo*75-b4r-**_foo", want: false},          // Stars and underscores
		{name: "Invalid Slug - Stars Special Chars Ampersand", param1: "foo-bar_foo*75-b4r-**_foo-&&", want: false}, // Stars, underscores, ampersands
		{name: "Invalid Slug - Uppercase Letters", param1: "Foo-Bar", want: false},                                  // Uppercase
		{name: "Invalid Slug - Contains Colon", param1: "a:b", want: false},                                        // Colon
		{name: "Invalid Slug - Underscore Between Words", param1: "foo_bar", want: false},                           // Underscore (Go disallows)
		{name: "Invalid Slug - Mixed Hyphen Underscore", param1: "foo-bar_foo", want: false},                        // Mixed hyphen and underscore
		{name: "Invalid Slug - Alphanumeric With Underscore", param1: "a1-b2_c3", want: false},                     // Underscore in alphanumeric slug
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsSlug(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
