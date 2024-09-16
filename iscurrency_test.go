package validatorgo

import (
	"testing"
)

func TestIsCurrency(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 IsCurrencyOpts
		want   bool
	}{
	// Valid test cases
	{name: "Test symbol", param1: "£10.50", param2: IsCurrencyOpts{Symbol: "£"}, want: true},
	{name: "Test require symbol", param1: "$10.50", param2: IsCurrencyOpts{Symbol: "$", RequireSymbol: true}, want: true},
	{name: "Test allow space after symbol", param1: "$ 10.50", param2: IsCurrencyOpts{Symbol: "$", AllowSpaceAfterSymbol: true}, want: true},
	{name: "Test symbol after digits", param1: "10.50$", param2: IsCurrencyOpts{Symbol: "$", SymbolAfterDigits: true}, want: true},
	{name: "Test allow negatives", param1: "-$10.50", param2: IsCurrencyOpts{Symbol: "$", AllowNegatives: true}, want: true},
	{name: "Test parens for negatives", param1: "($10.50)", param2: IsCurrencyOpts{Symbol: "$", AllowNegatives: true, ParensForNegatives: true}, want: true},
	{name: "Test negative sign before digits", param1: "$-10.50", param2: IsCurrencyOpts{Symbol: "$", AllowNegatives: true, NegativeSignBeforeDigits: true}, want: true},
	{name: "Test negative sign after digits", param1: "$10.50-", param2: IsCurrencyOpts{Symbol: "$", AllowNegatives: true, NegativeSignAfterDigits: true}, want: true},
	{name: "Test thousand separator", param1: "$10,000.50", param2: IsCurrencyOpts{Symbol: "$", ThousandSeparator: ","}, want: true},
	{name: "Test custom decimal separator", param1: "$10,000,50", param2: IsCurrencyOpts{Symbol: "$", ThousandSeparator: ",", DecimalSeparator: ","}, want: true},
	{name: "Test require decimal", param1: "$10.00", param2: IsCurrencyOpts{Symbol: "$", RequireDecimal: true}, want: true},
	{name: "Test digits after decimal", param1: "$10.50", param2: IsCurrencyOpts{Symbol: "$", AllowDecimal: true, DigitsAfterDecimal: []int{2}}, want: true},

	// Invalid test cases
	{name: "Test invalid symbol", param1: "€10.50", param2: IsCurrencyOpts{Symbol: "$"}, want: false},
	{name: "Test require symbol missing", param1: "10.50", param2: IsCurrencyOpts{Symbol: "$", RequireSymbol: true}, want: false},
	{name: "Test invalid space after symbol", param1: "$ 10.50", param2: IsCurrencyOpts{Symbol: "$", AllowSpaceAfterSymbol: false}, want: false},
	{name: "Test invalid symbol after digits", param1: "$10.50", param2: IsCurrencyOpts{Symbol: "$", SymbolAfterDigits: true}, want: false},
	{name: "Test negative sign disallowed", param1: "-$10.50", param2: IsCurrencyOpts{Symbol: "$", AllowNegatives: false}, want: false},
	{name: "Test parens for negatives disallowed", param1: "($10.50)", param2: IsCurrencyOpts{Symbol: "$", ParensForNegatives: false}, want: false},
	{name: "Test missing thousand separator", param1: "$10000.50", param2: IsCurrencyOpts{Symbol: "$", ThousandSeparator: ","}, want: false},
	{name: "Test incorrect decimal separator", param1: "$10.50", param2: IsCurrencyOpts{Symbol: "$", ThousandSeparator: ",", DecimalSeparator: ","}, want: false},
	{name: "Test missing required decimal", param1: "$10", param2: IsCurrencyOpts{Symbol: "$", RequireDecimal: true}, want: false},
	{name: "Test invalid digits after decimal", param1: "$10.5", param2: IsCurrencyOpts{Symbol: "$", AllowDecimal: true, DigitsAfterDecimal: []int{2}}, want: false},
}


	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsCurrency(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
