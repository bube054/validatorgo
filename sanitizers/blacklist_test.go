package sanitizers

import (
	"testing"
)

func TestBlacklist(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 string
		want   string
	}{
		{
			"test 1",
			"The quick brown fox jumps over the lazy dog",
			"fox|dog",
			"The quick brown  jumps over the lazy ",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := Blacklist(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got %s, wanted %s", result, test.want)
			}
		})
	}
}
