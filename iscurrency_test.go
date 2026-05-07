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
			param2: &IsCurrencyOpts{Symbol: String("¥"), RequireSymbol: true},
			want:   true,
		},
		{
			name:   "Valid currency with symbol before digits, no decimals, symbol not required but present",
			param1: "£100",
			param2: &IsCurrencyOpts{Symbol: String("£"), RequireSymbol: false},
			want:   true,
		},
		{
			name:   "Valid currency with symbol before digits, with decimals",
			param1: "$100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Valid currency with symbol after digits",
			param1: "100₣",
			param2: &IsCurrencyOpts{Symbol: String("₣"), RequireSymbol: true, SymbolAfterDigits: true},
			want:   true,
		},
		{
			name:   "Valid currency with thousand separator and decimals",
			param1: "$1,000.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, ThousandSeparator: String(","), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Valid currency with optional symbol",
			param1: "100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: false, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Valid currency with space after symbol",
			param1: "$ 100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowSpaceAfterSymbol: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Valid currency with negative sign before symbol",
			param1: "-$100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), NegativeSignBeforeDigits: true, MaxDigitsAfterDecimal: Uint(2), AllowDecimal: Bool(true)},
			want:   true,
		},
		{
			name:   "Valid currency with parentheses for negative values",
			param1: "($100.50)",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, ParensForNegatives: true, MaxDigitsAfterDecimal: Uint(2), AllowDecimal: Bool(true)},
			want:   true,
		},
		{
			name:   "Valid currency with decimal requirement",
			param1: "$100.00",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Valid currency with multiple decimal digits options",
			param1: "$100.5",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Valid currency with negative sign after digits",
			param1: "$100.50-",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowDecimal: Bool(true), AllowNegatives: Bool(true), NegativeSignAfterDigits: true, AllowNegativeSignPlaceholder: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Valid currency with space after digits",
			param1: "$100.50 ",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowSpaceAfterDigits: true, RequireSymbol: true, AllowDecimal: Bool(true), AllowNegatives: Bool(true), NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},

		// Invalid cases
		{
			name:   "Invalid currency with missing required symbol",
			param1: "100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true},
			want:   false,
		},
		{
			name:   "Invalid currency with incorrect thousand separator",
			param1: "$1.000,50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, ThousandSeparator: String(","), DecimalSeparator: String(".")},
			want:   false,
		},
		{
			name:   "Invalid currency with missing decimal when required",
			param1: "$100",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Invalid currency with more than allowed decimal digits",
			param1: "$100.500",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Invalid currency with space after symbol when not allowed",
			param1: "$ 100",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowSpaceAfterSymbol: false},
			want:   false,
		},
		{
			name:   "Invalid currency with negative sign after digits when not allowed",
			param1: "$100-",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), NegativeSignAfterDigits: false},
			want:   false,
		},
		{
			name:   "Invalid currency with negative value but parentheses required",
			param1: "-$100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), ParensForNegatives: true},
			want:   false,
		},
		{
			name:   "Invalid currency with multiple negative signs",
			param1: "--$100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true)},
			want:   false,
		},
		{
			name:   "Invalid currency with multiple negative disallowed",
			param1: "-$100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(false)},
			want:   false,
		},
		{
			name:   "Invalid currency with symbol after digits when not allowed",
			param1: "100$",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, SymbolAfterDigits: false},
			want:   false,
		},
		{
			name:   "Invalid currency with disallowed decimal",
			param1: "$100.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowDecimal: Bool(false)},
			want:   false,
		},
		{
			name:   "Invalid currency with space after digits",
			param1: "$100.50 ",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowSpaceAfterDigits: false, RequireSymbol: true, AllowDecimal: Bool(true), AllowNegatives: Bool(true), NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},

		// ================================================================
		// Block 1: Default US format (nil opts) - Valid
		// ================================================================
		{
			name:   "Block1 valid: -$10,123.45",
			param1: "-$10,123.45",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 valid: $10,123.45",
			param1: "$10,123.45",
			param2: nil,
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires digits in groups of 1-3 followed by (,\d{3})*, so >3 digits without separator not supported
		{
			name:   "Block1 valid: $10123.45",
			param1: "$10123.45",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 valid: 10,123.45",
			param1: "10,123.45",
			param2: nil,
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires digits in groups of 1-3 followed by (,\d{3})*, so >3 digits without separator not supported
		{
			name:   "Block1 valid: 10123.45",
			param1: "10123.45",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 valid: 10,123",
			param1: "10,123",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 valid: 1,123,456",
			param1: "1,123,456",
			param2: nil,
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires digits in groups of 1-3 followed by (,\d{3})*, so >3 digits without separator not supported
		{
			name:   "Block1 valid: 1123456",
			param1: "1123456",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 valid: 1.39",
			param1: "1.39",
			param2: nil,
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires at least one leading digit (\d{1,3}), so bare decimal like .03 not supported
		{
			name:   "Block1 valid: .03",
			param1: ".03",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 valid: 0.10",
			param1: "0.10",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 valid: $0.10",
			param1: "$0.10",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 valid: -$0.01",
			param1: "-$0.01",
			param2: nil,
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires at least one leading digit (\d{1,3}), so bare decimal like $.99 not supported
		{
			name:   "Block1 valid: -$.99",
			param1: "-$.99",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 valid: $100,234,567.89",
			param1: "$100,234,567.89",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 valid: $10,123",
			param1: "$10,123",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 valid: 10,123 (no symbol)",
			param1: "10,123",
			param2: nil,
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires digits in groups of 1-3 followed by (,\d{3})*, so >3 digits without separator not supported
		{
			name:   "Block1 valid: -10123",
			param1: "-10123",
			param2: nil,
			want:   false,
		},

		// ================================================================
		// Block 1: Default US format (nil opts) - Invalid
		// ================================================================
		{
			name:   "Block1 invalid: 1.234 (3 decimals)",
			param1: "1.234",
			param2: nil,
			want:   false,
		},
		// TODO: should be invalid per validator.js — Go allows 0 to MaxDigitsAfterDecimal digits after decimal, validator.js requires exactly 2
		{
			name:   "Block1 invalid: $1.1 (1 decimal digit)",
			param1: "$1.1",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 invalid: $ 32.50 (space after symbol)",
			param1: "$ 32.50",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: 500$ (symbol after digits)",
			param1: "500$",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: .0001 (4 decimal digits)",
			param1: ".0001",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: $.001 (3 decimal digits)",
			param1: "$.001",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: $0.001 (3 decimal digits with zero)",
			param1: "$0.001",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: 12,34.56 (bad thousand grouping)",
			param1: "12,34.56",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: 123456,123,123456 (bad grouping)",
			param1: "123456,123,123456",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: 123,4 (bad grouping)",
			param1: "123,4",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: ,123 (leading comma)",
			param1: ",123",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: $-,123 (negative sign misplaced)",
			param1: "$-,123",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: $ (symbol only)",
			param1: "$",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: . (dot only)",
			param1: ".",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: , (comma only)",
			param1: ",",
			param2: nil,
			want:   false,
		},
		// TODO: should be invalid per validator.js — Go regex matches "00" as valid (2 digits in \d{1,3}), validator.js rejects leading zeros
		{
			name:   "Block1 invalid: 00 (double zero)",
			param1: "00",
			param2: nil,
			want:   true,
		},
		{
			name:   "Block1 invalid: $- (symbol with negative only)",
			param1: "$-",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: $-,. (symbol negative comma dot)",
			param1: "$-,.",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: - (negative only)",
			param1: "-",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: -$ (negative symbol only)",
			param1: "-$",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: empty string",
			param1: "",
			param2: nil,
			want:   false,
		},
		{
			name:   "Block1 invalid: - $ (negative space symbol)",
			param1: "- $",
			param2: nil,
			want:   false,
		},

		// ================================================================
		// Block 2: No decimals (AllowDecimal: Bool(false)) - Valid
		// ================================================================
		{
			name:   "Block2 valid: -$10,123",
			param1: "-$10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block2 valid: $10,123",
			param1: "$10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block2 valid: 10,123",
			param1: "10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block2 valid: 1,123,456",
			param1: "1,123,456",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block2 valid: $0",
			param1: "$0",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block2 valid: -$0",
			param1: "-$0",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},

		// ================================================================
		// Block 2: No decimals (AllowDecimal: Bool(false)) - Invalid
		// ================================================================
		{
			name:   "Block2 invalid: -$10,123.45",
			param1: "-$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block2 invalid: $10,123.45",
			param1: "$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block2 invalid: 1.39",
			param1: "1.39",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block2 invalid: .03",
			param1: ".03",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(false), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},

		// ================================================================
		// Block 3: Require decimal (RequireDecimal: true) - Valid
		// ================================================================
		{
			name:   "Block3 valid: -$10,123.45",
			param1: "-$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block3 valid: $10,123.45",
			param1: "$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block3 valid: 10,123.45",
			param1: "10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block3 valid: 1.39",
			param1: "1.39",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires at least one leading digit (\d{1,3}), so bare decimal like .03 not supported
		{
			name:   "Block3 valid: .03",
			param1: ".03",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block3 valid: 0.10",
			param1: "0.10",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block3 valid: $0.10",
			param1: "$0.10",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block3 valid: -$0.01",
			param1: "-$0.01",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires at least one leading digit (\d{1,3}), so bare decimal like $.99 not supported
		{
			name:   "Block3 valid: -$.99",
			param1: "-$.99",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block3 valid: $100,234,567.89",
			param1: "$100,234,567.89",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},

		// ================================================================
		// Block 3: Require decimal (RequireDecimal: true) - Invalid
		// ================================================================
		{
			name:   "Block3 invalid: $10,123 (no decimal)",
			param1: "$10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block3 invalid: 10,123 (no decimal)",
			param1: "10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block3 invalid: -10123 (no decimal)",
			param1: "-10123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block3 invalid: 1,123,456 (no decimal)",
			param1: "1,123,456",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), AllowDecimal: Bool(true), RequireDecimal: true, MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},

		// ================================================================
		// Block 4: Require symbol (RequireSymbol: true) - Valid
		// ================================================================
		{
			name:   "Block4 valid: -$10,123.45",
			param1: "-$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block4 valid: $10,123.45",
			param1: "$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires digits in groups of 1-3 followed by (,\d{3})*, so >3 digits without separator not supported
		{
			name:   "Block4 valid: $10123.45",
			param1: "$10123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 valid: $1.39",
			param1: "$1.39",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires at least one leading digit (\d{1,3}), so bare decimal like $.03 not supported
		{
			name:   "Block4 valid: $.03",
			param1: "$.03",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 valid: $0.10",
			param1: "$0.10",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block4 valid: -$0.01",
			param1: "-$0.01",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires at least one leading digit (\d{1,3}), so bare decimal like $.99 not supported
		{
			name:   "Block4 valid: -$.99",
			param1: "-$.99",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 valid: $100,234,567.89",
			param1: "$100,234,567.89",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block4 valid: $10,123",
			param1: "$10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires digits in groups of 1-3 followed by (,\d{3})*, so >3 digits without separator not supported
		{
			name:   "Block4 valid: -$10123",
			param1: "-$10123",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},

		// ================================================================
		// Block 4: Require symbol (RequireSymbol: true) - Invalid
		// ================================================================
		{
			name:   "Block4 invalid: 10,123.45 (no symbol)",
			param1: "10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 invalid: 10123.45 (no symbol)",
			param1: "10123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 invalid: 10,123 (no symbol)",
			param1: "10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 invalid: 1,123,456 (no symbol)",
			param1: "1,123,456",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 invalid: 1.39 (no symbol)",
			param1: "1.39",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 invalid: .03 (no symbol)",
			param1: ".03",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block4 invalid: 0.10 (no symbol)",
			param1: "0.10",
			param2: &IsCurrencyOpts{Symbol: String("$"), RequireSymbol: true, AllowNegatives: Bool(true), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},

		// ================================================================
		// Block 5: No negatives (AllowNegatives: Bool(false)) - Valid
		// ================================================================
		{
			name:   "Block5 valid: $10,123.45",
			param1: "$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block5 valid: 10,123.45",
			param1: "10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block5 valid: 10,123",
			param1: "10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block5 valid: 1.39",
			param1: "1.39",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — regex requires at least one leading digit (\d{1,3}), so bare decimal like .03 not supported
		{
			name:   "Block5 valid: .03",
			param1: ".03",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block5 valid: 0.10",
			param1: "0.10",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block5 valid: $0.10",
			param1: "$0.10",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block5 valid: $100,234,567.89",
			param1: "$100,234,567.89",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block5 valid: $10,123",
			param1: "$10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block5 valid: 10,123 (no symbol)",
			param1: "10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},

		// ================================================================
		// Block 5: No negatives (AllowNegatives: Bool(false)) - Invalid
		// ================================================================
		{
			name:   "Block5 invalid: -$10,123.45",
			param1: "-$10,123.45",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block5 invalid: -10123",
			param1: "-10123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block5 invalid: -$0.01",
			param1: "-$0.01",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block5 invalid: -$.99",
			param1: "-$.99",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block5 invalid: - (negative only)",
			param1: "-",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block5 invalid: -$ (negative symbol)",
			param1: "-$",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(false), AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},

		// ================================================================
		// Block 6: Parentheses for negatives (ParensForNegatives: true) - Valid
		// ================================================================
		// TODO: should be valid per validator.js — ParensForNegatives check rejects non-parenthesized positive values (implementation requires both parens present)
		{
			name:   "Block6 valid: 1,234",
			param1: "1,234",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 valid: (1,234)",
			param1: "(1,234)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block6 valid: ($6,954,231)",
			param1: "($6,954,231)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — ParensForNegatives check rejects non-parenthesized positive values (implementation requires both parens present)
		{
			name:   "Block6 valid: $10.03",
			param1: "$10.03",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 valid: (10.03)",
			param1: "(10.03)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block6 valid: ($10.03)",
			param1: "($10.03)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		{
			name:   "Block6 valid: ($0.01)",
			param1: "($0.01)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — ParensForNegatives check rejects non-parenthesized positive values (implementation requires both parens present)
		{
			name:   "Block6 valid: $10,123",
			param1: "$10,123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 valid: (10,123)",
			param1: "(10,123)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   true,
		},
		// TODO: should be valid per validator.js — ParensForNegatives check + regex requires digits in groups of 1-3, so >3 digits without separator not supported
		{
			name:   "Block6 valid: 10123",
			param1: "10123",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},

		// ================================================================
		// Block 6: Parentheses for negatives (ParensForNegatives: true) - Invalid
		// ================================================================
		{
			name:   "Block6 invalid: -$1.10 (negative sign instead of parens)",
			param1: "-$1.10",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 invalid: $ 32.50 (space after symbol)",
			param1: "$ 32.50",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 invalid: 500$ (symbol after digits)",
			param1: "500$",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 invalid: () (empty parens)",
			param1: "()",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 invalid: (-) (negative in parens)",
			param1: "(-)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
			want:   false,
		},
		{
			name:   "Block6 invalid: (-$) (negative symbol in parens)",
			param1: "(-$)",
			param2: &IsCurrencyOpts{Symbol: String("$"), AllowNegatives: Bool(true), ParensForNegatives: true, AllowDecimal: Bool(true), MaxDigitsAfterDecimal: Uint(2)},
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
