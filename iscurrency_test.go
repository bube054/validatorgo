package validatorgo

import (
	"testing"
)

func TestIsCurrency(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsCurrencyOpts
		want   bool
	}{
		// Cases with nil config
		{
			name:   "Valid currency with nil/default IsCurrencyOpts",
			param1: "$100,000.00",
			param2: nil,
			want:   true,
		},
		{
			name:   "Invalid currency with nil/default IsCurrencyOpts",
			param1: "$10p",
			param2: nil,
			want:   false,
		},

		// Valid cases
		{
			name:   "Valid currency with symbol before digits, no decimals",
			param1: "¥100",
			param2: &IsCurrencyOpts{Symbol: "¥", RequireSymbol: true},
			want:   true,
		},
		{
			name:   "Valid currency with symbol before digits, no decimals, symbol not required but present",
			param1: "£100",
			param2: &IsCurrencyOpts{Symbol: "£", RequireSymbol: false},
			want:   true,
		},
		{
			name:   "Valid currency with symbol before digits, with decimals",
			param1: "$100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},
		{
			name:   "Valid currency with symbol after digits",
			param1: "100₣",
			param2: &IsCurrencyOpts{Symbol: "₣", RequireSymbol: true, SymbolAfterDigits: true},
			want:   true,
		},
		{
			name:   "Valid currency with thousand separator and decimals",
			param1: "$1,000.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, ThousandSeparator: ",", AllowDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},
		{
			name:   "Valid currency with optional symbol",
			param1: "100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: false, AllowDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},
		{
			name:   "Valid currency with space after symbol",
			param1: "$ 100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowSpaceAfterSymbol: true, AllowDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},
		{
			name:   "Valid currency with negative sign before symbol",
			param1: "-$100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowNegatives: true, NegativeSignBeforeDigits: true, MaxDigitsAfterDecimal: 2, AllowDecimal: true},
			want:   true,
		},
		{
			name:   "Valid currency with parentheses for negative values",
			param1: "($100.50)",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, ParensForNegatives: true, MaxDigitsAfterDecimal: 2, AllowDecimal: true},
			want:   true,
		},
		{
			name:   "Valid currency with decimal requirement",
			param1: "$100.00",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowDecimal: true, RequireDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},
		{
			name:   "Valid currency with multiple decimal digits options",
			param1: "$100.5",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},
		{
			name:   "Valid currency with negative sign after digits",
			param1: "$100.50-",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowDecimal: true, AllowNegatives: true, NegativeSignAfterDigits: true, AllowNegativeSignPlaceholder: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},
		{
			name:   "Valid currency with space after digits",
			param1: "$100.50 ",
			param2: &IsCurrencyOpts{Symbol: "$", AllowSpaceAfterDigits: true, RequireSymbol: true, AllowDecimal: true, AllowNegatives: true, NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: true, MaxDigitsAfterDecimal: 2},
			want:   true,
		},

		// Invalid cases
		{
			name:   "Invalid currency with missing required symbol",
			param1: "100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true},
			want:   false,
		},
		{
			name:   "Invalid currency with incorrect thousand separator",
			param1: "$1.000,50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, ThousandSeparator: ",", DecimalSeparator: "."},
			want:   false,
		},
		{
			name:   "Invalid currency with missing decimal when required",
			param1: "$100",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, RequireDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   false,
		},
		{
			name:   "Invalid currency with more than allowed decimal digits",
			param1: "$100.500",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowDecimal: true, MaxDigitsAfterDecimal: 2},
			want:   false,
		},
		{
			name:   "Invalid currency with space after symbol when not allowed",
			param1: "$ 100",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowSpaceAfterSymbol: false},
			want:   false,
		},
		{
			name:   "Invalid currency with negative sign after digits when not allowed",
			param1: "$100-",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowNegatives: true, NegativeSignAfterDigits: false},
			want:   false,
		},
		{
			name:   "Invalid currency with negative value but parentheses required",
			param1: "-$100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowNegatives: true, ParensForNegatives: true},
			want:   false,
		},
		{
			name:   "Invalid currency with multiple negative signs",
			param1: "--$100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowNegatives: true},
			want:   false,
		},
		{
			name:   "Invalid currency with multiple negative disallowed",
			param1: "-$100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowNegatives: false},
			want:   false,
		},
		{
			name:   "Invalid currency with symbol after digits when not allowed",
			param1: "100$",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, SymbolAfterDigits: false},
			want:   false,
		},
		{
			name:   "Invalid currency with disallowed decimal",
			param1: "$100.50",
			param2: &IsCurrencyOpts{Symbol: "$", RequireSymbol: true, AllowDecimal: false},
			want:   false,
		},
		{
			name:   "Invalid currency with space after digits",
			param1: "$100.50 ",
			param2: &IsCurrencyOpts{Symbol: "$", AllowSpaceAfterDigits: false, RequireSymbol: true, AllowDecimal: true, AllowNegatives: true, NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: true, MaxDigitsAfterDecimal: 2},
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsCurrency(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
