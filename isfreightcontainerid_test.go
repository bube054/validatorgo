package validatorgo

import "testing"

func TestIsFreightContainerID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
    {name: "Valid container ID", param1: "ABCU1234567", want: true},
    {name: "Valid container ID with Z category", param1: "XYZZ9876543", want: true},

    {name: "Invalid owner code (too short)", param1: "AB123456789", want: false},
    {name: "Invalid category identifier", param1: "ABCX1234567", want: false},
    {name: "Invalid serial number length", param1: "ABCJ12345", want: false},
    {name: "Invalid check digit", param1: "ABCU123456", want: false},
    {name: "Invalid characters in serial number", param1: "ABCU12A4567", want: false},
    {name: "Empty string", param1: "", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsFreightContainerID(test.param1)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
