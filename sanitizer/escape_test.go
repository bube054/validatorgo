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
		{
			"test 1",
			`"Fran & Freddie's Diner" <tasty@example.com>`,
			"&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Escape(test.param1)

			if result != test.want {
				t.Errorf("got %s, wanted %s", result, test.want)
			}
		})
	}
}
