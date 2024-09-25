package sanitizer

import (
	"testing"
)

func TestEscape(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   string
	}{
		{name: "Basic case with script tags", param1: "<script>alert('XSS');</script>", want: "&lt;script&gt;alert(&#39;XSS&#39;);&lt;/script&gt;"},
		{name: "Basic case with HTML tags", param1: "<div>Hello World!</div>", want: "&lt;div&gt;Hello World!&lt;/div&gt;"},
		{name: "Case with special characters", param1: "Hello & welcome to <Validator>!", want: "Hello &amp; welcome to &lt;Validator&gt;!"},
		{name: "Case with quotes", param1: `He said, "Hello World!"`, want: "He said, &#34;Hello World!&#34;"},
		{name: "Case with apostrophes", param1: "It's a beautiful day!", want: "It&#39;s a beautiful day!"},
		{name: "Empty string case", param1: "", want: ""},
		{name: "Whitespace preservation", param1: "  <span>   ", want: "  &lt;span&gt;   "},
		{name: "Complex string with multiple types", param1: "Hello <b>World</b> & <i>Friends</i>", want: "Hello &lt;b&gt;World&lt;/b&gt; &amp; &lt;i&gt;Friends&lt;/i&gt;"},
		{name: "String with only blacklisted characters", param1: "<>", want: "&lt;&gt;"},
		{name: "String with no characters to escape", param1: "No special chars here.", want: "No special chars here."},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Escape(test.param1)

			if result != test.want {
				t.Errorf("got %s`, wanted %s`", result, test.want)
			}
		})
	}
}
