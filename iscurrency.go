package validatorgo

import (
	"regexp"
	"strconv"
)

type IsCurrencyOpts struct {
	// deals with the options for currency
	Symbol                string // '$10.50'
	RequireSymbol         bool   // $10.50
	AllowSpaceAfterSymbol bool   // $ 10.50
	SymbolAfterDigits     bool   // 10.50$

	// deals with the numeric signs
	AllowNegatives               bool // -$10.50
	ParensForNegatives           bool // ($10.50)
	NegativeSignBeforeDigits     bool // $-10.50
	NegativeSignAfterDigits      bool // $10.50-
	AllowNegativeSignPlaceholder bool // $ 10.50

	// deals with the decimals/separator
	ThousandSeparator     string // 10,000.50
	DecimalSeparator      string // 10.50
	AllowDecimal          bool   // $10.00
	RequireDecimal        bool   // $10.00
	DigitsAfterDecimal    []int  // $10.50 []
	AllowSpaceAfterDigits bool   // '$ 10.50
}

func getMaxDigits(digits []int) string {
	if len(digits) == 0 {
		return "2"
	}
	return strconv.Itoa(digits[len(digits)-1])
}

func getMinDigits(digits []int) string {
	if len(digits) == 0 {
		return "0"
	}
	return strconv.Itoa(digits[0])
}

// A validator that checks if the string is a valid currency amount.
// options is a struct which defaults to { Symbol: '$', RequireSymbol: false, AllowSpaceAfterSymbol: false, SymbolAfterDigits: false, AllowNegatives: true, ParensForNegatives: false, NegativeSignBeforeDigits: false, NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: false, ThousandsSeparator: ',', DecimalSeparator: '.', AllowDecimal: true, RequireDecimal: false, DigitsAfterDecimal: [2], AllowSpaceAfterDigits: false }.
// Note: The array digits_after_decimal is filled with the exact number of digits allowed not a range, for example a range 1 to 3 will be given as [1, 2, 3].

func IsCurrency(str string, opts IsCurrencyOpts) bool {
	// Build the regex pattern based on options
	symbol := regexp.QuoteMeta(opts.Symbol)
	decimalSeparator := regexp.QuoteMeta(opts.DecimalSeparator)
	thousandSeparator := regexp.QuoteMeta(opts.ThousandSeparator)

	symbolPattern := ""
	if opts.RequireSymbol {
		symbolPattern = symbol
		if opts.AllowSpaceAfterSymbol {
			symbolPattern += "\\s?"
		}
	} else {
		symbolPattern = symbol + "?\\s*"
	}

	negativesPattern := ""
	if opts.AllowNegatives {
		if opts.ParensForNegatives {
			negativesPattern = "\\(?"
		} else if opts.NegativeSignBeforeDigits {
			negativesPattern = "[-"
		}
	}

	if opts.AllowNegativeSignPlaceholder {
		negativesPattern += " ?"
	} else if opts.NegativeSignAfterDigits {
		negativesPattern = ""
	}

	if opts.ParensForNegatives {
		negativesPattern += "\\)?"
	}

	decimalPattern := ""
	if opts.AllowDecimal {
		decimalPattern = decimalSeparator + "\\d{1," + getMaxDigits(opts.DigitsAfterDecimal) + "}"
		if opts.RequireDecimal {
			decimalPattern = decimalSeparator + "\\d{" + getMinDigits(opts.DigitsAfterDecimal) + "," + getMaxDigits(opts.DigitsAfterDecimal) + "}"
		}
	} else {
		decimalPattern = ""
	}

	thousandsPattern := ""
	if thousandSeparator != "" {
		thousandsPattern = "(?:(?:\\d{1,3}(?:\\" + thousandSeparator + "\\d{3})*)|\\d+)"
	} else {
		thousandsPattern = "\\d+"
	}

	suffix := ""
	if opts.SymbolAfterDigits {
		suffix = symbol
	}

	if opts.AllowSpaceAfterDigits {
		suffix += "\\s*"
	}

	pattern := "^" + symbolPattern + negativesPattern + thousandsPattern + decimalPattern + negativesPattern + suffix + "$"

	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}
