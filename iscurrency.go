package validatorgo

import (
	"regexp"
	"strconv"
)

var (
	isCurrencyOptsDefaultSymbol                string = "$"
	isCurrencyOptsDefaultRequireSymbol         bool   = false
	isCurrencyOptsDefaultAllowSpaceAfterSymbol bool   = false
	isCurrencyOptsDefaultSymbolAfterDigits     bool   = false

	isCurrencyOptsDefaultAllowNegatives               bool = true
	isCurrencyOptsDefaultParensForNegatives           bool = false
	isCurrencyOptsDefaultNegativeSignBeforeDigits     bool = false
	isCurrencyOptsDefaultNegativeSignAfterDigits      bool = false
	isCurrencyOptsDefaultAllowNegativeSignPlaceholder bool = false

	isCurrencyOptsDefaultThousandSeparator     string = ","
	isCurrencyOptsDefaultDecimalSeparator      string = "."
	isCurrencyOptsDefaultAllowDecimal          bool   = true
	isCurrencyOptsDefaultRequireDecimal        bool   = false
	isCurrencyOptsDefaultAllowSpaceAfterDigits bool   = false
	isCurrencyOptsDefaultMaxDigitsAfterDecimal uint = 2
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
	MaxDigitsAfterDecimal uint   // $10.50
	AllowSpaceAfterDigits bool   // '$ 10.50
}

// A validator that checks if the string is a valid currency amount.
//
// IsCurrencyOpts is a struct which defaults to { Symbol: '$', RequireSymbol: false, AllowSpaceAfterSymbol: false, SymbolAfterDigits: false, AllowNegatives: true, ParensForNegatives: false, NegativeSignBeforeDigits: false, NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: false, ThousandsSeparator: ',', DecimalSeparator: '.', AllowDecimal: true, RequireDecimal: false, MaxDigitsAfterDecimal: 2, AllowSpaceAfterDigits: false }.
//
//	ok := validatorgo.IsCurrency("$100,000.00", nil)
//	fmt.Println(ok) // true
//	ok := validatorgo.IsCurrency("¥100", &validatorgo.IsCurrencyOpts{Symbol: "¥", RequireSymbol: true})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsCurrency("100.50", &validatorgo.IsCurrencyOpts{Symbol: "$", RequireSymbol: true})
//	fmt.Println(ok) // false
func IsCurrency(str string, opts *IsCurrencyOpts) bool {
	if opts == nil {
		opts = setIsCurrencyOptsToDefault()
	}

	escSymbol := regexp.QuoteMeta(opts.Symbol)

	if opts.ThousandSeparator == "" {
		opts.ThousandSeparator = isCurrencyOptsDefaultThousandSeparator
	}
	escThSep := regexp.QuoteMeta(opts.ThousandSeparator)

	if opts.DecimalSeparator == "" {
		opts.DecimalSeparator = isCurrencyOptsDefaultDecimalSeparator
	}
	escDecSep := regexp.QuoteMeta(opts.DecimalSeparator)

	decAftDigs := `(` + escDecSep + `\d{0,` + strconv.Itoa(int(opts.MaxDigitsAfterDecimal)) + `})?`

	reStr := `^(\(?)(-?)(` + escSymbol + `?)(-?)(\s*)(\d{1,3})(` + escThSep + `\d{3})*` + decAftDigs + `(` + escSymbol + `?)([-\s]?)(\)?)$`
	re, err := regexp.Compile(reStr)

	if err != nil {
		return false
	}

	capGrp := re.FindStringSubmatch(str)
	// fmt.Println(re.String())

	if len(capGrp) == 0 {
		return false
	}

	// fmt.Println(capGrp, len(capGrp))
	// fmt.Printf("%+v\n", opts)

	leftBrack := capGrp[1] // (
	beg1Op := capGrp[2]    // -/+
	begSym := capGrp[3]    // $/€...
	beg2Op := capGrp[4]    // -/+
	space := capGrp[5]     // " " or "	" or "		" ....
	// intPart1 := capGrp[6]    // 100
	// intPart2 := capGrp[7]    // ,000
	decPart := capGrp[8]     // .50
	endSym := capGrp[9]      // $/€...
	endOp := capGrp[10]      // -/+ or a space
	rightBrack := capGrp[11] // )

	if opts.RequireSymbol {
		if begSym != opts.Symbol && endSym != opts.Symbol {
			// fmt.Println("opts.RequireSymbol")
			return false
		}
	}

	if !opts.AllowSpaceAfterSymbol {
		if space != "" {
			// fmt.Println("opts.AllowSpaceAfterSymbol")
			return false
		}
	}

	if opts.SymbolAfterDigits {
		if endSym != opts.Symbol {
			// fmt.Println("opts.SymbolAfterDigits 1")
			return false
		}
	} else {
		// fmt.Println("uvyctuyi")
		if endSym == opts.Symbol {
			// fmt.Println("opts.SymbolAfterDigits 2")
			return false
		}
	}

	if !opts.AllowNegatives {
		if beg1Op == "-" || beg2Op == "-" || endOp == "-" {
			// fmt.Println("opts.AllowNegatives")
			return false
		}
	}

	if opts.ParensForNegatives {
		if leftBrack != "(" || rightBrack != ")" || beg1Op != "" || beg2Op != "" || endOp != "" {
			// fmt.Println("opts.ParensForNegatives")
			return false
		}
	}

	if opts.NegativeSignBeforeDigits {
		if beg1Op != "-" {
			// fmt.Println("opts.NegativeSignBeforeDigits")
			return false
		}
	}

	if opts.NegativeSignAfterDigits {
		if endOp != "-" {
			// fmt.Println("opts.NegativeSignAfterDigits", beg2Op)
			return false
		}
	}

	if !opts.AllowNegativeSignPlaceholder {
		if endOp != "" {
			// fmt.Println("opts.AllowNegativeSignPlaceholder", beg2Op)
			return false
		}
	}

	if !opts.AllowDecimal {
		if decPart != "" {
			// fmt.Println("opts.AllowDecimal")
			return false
		}
	}

	if opts.RequireDecimal {
		if decPart == "" {
			// fmt.Println("opts.RequireDecimal", decPart)
			return false
		}
	}

	if !opts.AllowSpaceAfterDigits && endOp != "-" {
		// fmt.Printf("`%s`\n", endOp)
		if !IsEmpty(endOp, &IsEmptyOpts{IgnoreWhitespace: false}) {
			// fmt.Println("opts.AllowSpaceAfterDigits")
			return false
		}
	}
	return true
}

func setIsCurrencyOptsToDefault() *IsCurrencyOpts {
	return &IsCurrencyOpts{
		Symbol:                isCurrencyOptsDefaultSymbol,
		RequireSymbol:         isCurrencyOptsDefaultRequireSymbol,
		AllowSpaceAfterSymbol: isCurrencyOptsDefaultAllowSpaceAfterSymbol,
		SymbolAfterDigits:     isCurrencyOptsDefaultSymbolAfterDigits,

		AllowNegatives:               isCurrencyOptsDefaultAllowNegatives,
		ParensForNegatives:           isCurrencyOptsDefaultParensForNegatives,
		NegativeSignBeforeDigits:     isCurrencyOptsDefaultNegativeSignBeforeDigits,
		NegativeSignAfterDigits:      isCurrencyOptsDefaultNegativeSignAfterDigits,
		AllowNegativeSignPlaceholder: isCurrencyOptsDefaultAllowNegativeSignPlaceholder,

		ThousandSeparator:     isCurrencyOptsDefaultThousandSeparator,
		DecimalSeparator:      isCurrencyOptsDefaultDecimalSeparator,
		AllowDecimal:          isCurrencyOptsDefaultAllowDecimal,
		RequireDecimal:        isCurrencyOptsDefaultRequireDecimal,
		MaxDigitsAfterDecimal: isCurrencyOptsDefaultMaxDigitsAfterDecimal,
		AllowSpaceAfterDigits: isCurrencyOptsDefaultAllowSpaceAfterDigits,
	}
}
