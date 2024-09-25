package validatorgo

import (
	"testing"
)

func TestIsAbaRouting(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		{name: "Valid ABA routing", param1: "021000021", want: true},
		{name: "Valid ABA routing, with dashes", param1: "0260-0959-3", want: true},
    {name: "Invalid ABA routing (too short)", param1: "12345678", want: false},
    {name: "Invalid ABA routing (non-digit characters)", param1: "12A456789", want: false},
    {name: "Invalid ABA routing (checksum failure)", param1: "987654321", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsAbaRouting(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
