package validatorgo

import "testing"

func TestIsFreightContainerID(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// { name: "Valid FreightContainerID", param1: "MSCU1234567", want:   true},
		// {name: "Valid FreightContainerID", param1: "TGHU1234563", want: true},
		// {name: "Valid FreightContainerID", param1: "TLLU2841316", want: true},
		// {name: "Valid FreightContainerID", param1: "MAEU5612394", want: true},
		// {name: "Valid FreightContainerID", param1: "OOLU4621759", want: true},

		// {name: "Inalid FreightContainerID", param1: "ABCD1234561", want: false},
		// {name: "Inalid FreightContainerID", param1: "MSKU12345AB3", want: false},
		// {name: "Inalid FreightContainerID", param1: "CMAU8765439", want: false},
		// {name: "Inalid FreightContainerID", param1: "TGHU12345", want: false},
		// {name: "Inalid FreightContainerID", param1: "XYZU987654A", want: false},
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
