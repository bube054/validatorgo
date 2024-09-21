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

		// Invalid slugs
		{name: "Invalid Slug - Contains Uppercase", param1: "My-First-Blog-Post", want: false},          // Uppercase letters
		{name: "Invalid Slug - Contains Special Characters", param1: "my@blog#post!", want: false},      // Special characters like @, #, !
		{name: "Invalid Slug - Contains Underscores", param1: "my_first_blog_post", want: false},        // Underscores instead of hyphens
		{name: "Invalid Slug - Leading Hyphen", param1: "-leading-hyphen", want: false},                 // Leading hyphen
		{name: "Invalid Slug - Trailing Hyphen", param1: "trailing-hyphen-", want: false},               // Trailing hyphen
		{name: "Invalid Slug - Multiple Consecutive Hyphens", param1: "multiple--hyphens", want: false}, // Multiple consecutive hyphens
		{name: "Invalid Slug - Contains Spaces", param1: "slug with spaces", want: false},               // Spaces in slug
		{name: "Invalid Slug - Empty String", param1: "", want: false},                                  // Empty string
		{name: "Invalid Slug - Only Hyphens", param1: "---", want: false},                               // Only hyphens
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsSlug(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
