package validatorgo

import (
	"testing"
)

func TestIsCreditCard(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsCreditCardOpts
		want   bool
	}{
		// Valid test cases
		{name: "Valid Amex", param1: "378282246310005", param2: IsCreditCardOpts{Provider: "amex"}, want: true},
		{name: "Valid Visa", param1: "4111111111111111", param2: IsCreditCardOpts{Provider: "visa"}, want: true},
		{name: "Valid MasterCard", param1: "5555555555554444", param2: IsCreditCardOpts{Provider: "mastercard"}, want: true},
		{name: "Valid Discover", param1: "6011111111111117", param2: IsCreditCardOpts{Provider: "discover"}, want: true},
		{name: "Valid JCB", param1: "3530111333300000", param2: IsCreditCardOpts{Provider: "jcb"}, want: true},
		{name: "Valid UnionPay", param1: "6221260000000000", param2: IsCreditCardOpts{Provider: "unionpay"}, want: true},
		{name: "Valid Diners Club", param1: "30569309025904", param2: IsCreditCardOpts{Provider: "dinersclub"}, want: true},
		{name: "Valid any: Diners Club", param1: "30569309025904", param2: IsCreditCardOpts{Provider: ""}, want: true},

		// Invalid test cases
		{name: "Invalid Amex", param1: "37828224631000", param2: IsCreditCardOpts{Provider: "amex"}, want: false},
		{name: "Invalid Visa", param1: "41111111111111", param2: IsCreditCardOpts{Provider: "visa"}, want: false},
		{name: "Invalid MasterCard", param1: "55000055555555599", param2: IsCreditCardOpts{Provider: "mastercard"}, want: false},
		{name: "Invalid JCB", param1: "353011133330000", param2: IsCreditCardOpts{Provider: "jcb"}, want: false},
		{name: "Invalid UnionPay", param1: "62212600000000", param2: IsCreditCardOpts{Provider: "unionpay"}, want: false},
		{name: "Invalid Diners Club", param1: "3056930902590", param2: IsCreditCardOpts{Provider: "dinersclub"}, want: false},

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsCreditCard(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
