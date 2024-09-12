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
		{
			name:   "Valid Amex",
			param1: "378282246310005",
			param2: IsCreditCardOpts{Provider: "amex"},
			want:   true,
		},
		{
			name:   "Invalid Amex",
			param1: "37828224631000", // one digit short
			param2: IsCreditCardOpts{Provider: "amex"},
			want:   false,
		},
		{
			name:   "Valid Visa",
			param1: "4111111111111111",
			param2: IsCreditCardOpts{Provider: "visa"},
			want:   true,
		},
		{
			name:   "Invalid Visa",
			param1: "41111111111111", // short number
			param2: IsCreditCardOpts{Provider: "visa"},
			want:   false,
		},
		{
			name:   "Valid MasterCard",
			param1: "5555555555554444",
			param2: IsCreditCardOpts{Provider: "mastercard"},
			want:   true,
		},
		{
			name:   "Invalid MasterCard",
			param1: "55000055555555599", // valid length but invalid number
			param2: IsCreditCardOpts{Provider: "mastercard"},
			want:   false,
		},
		{
			name:   "Valid Discover",
			param1: "6011111111111117",
			param2: IsCreditCardOpts{Provider: "discover"},
			want:   true,
		},
		// {
		// 	name:   "Invalid Discover",
		// 	param1: "60110009901394947",
		// 	param2: IsCreditCardOpts{Provider: "discover"},
		// 	want:   false,
		// },
		{
			name:   "Valid JCB",
			param1: "3530111333300000",
			param2: IsCreditCardOpts{Provider: "jcb"},
			want:   true,
		},
		{
			name:   "Invalid JCB",
			param1: "353011133330000", // one digit short
			param2: IsCreditCardOpts{Provider: "jcb"},
			want:   false,
		},
		{
			name:   "Valid UnionPay",
			param1: "6221260000000000",
			param2: IsCreditCardOpts{Provider: "unionpay"},
			want:   true,
		},
		{
			name:   "Invalid UnionPay",
			param1: "62212600000000", // short number
			param2: IsCreditCardOpts{Provider: "unionpay"},
			want:   false,
		},
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
