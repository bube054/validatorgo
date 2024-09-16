package validatorgo

import (
	"fmt"
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
//
// IsCurrencyOpts is a struct which defaults to { Symbol: '$', RequireSymbol: false, AllowSpaceAfterSymbol: false, SymbolAfterDigits: false, AllowNegatives: true, ParensForNegatives: false, NegativeSignBeforeDigits: false, NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: false, ThousandsSeparator: ',', DecimalSeparator: '.', AllowDecimal: true, RequireDecimal: false, DigitsAfterDecimal: [2], AllowSpaceAfterDigits: false }.
// Note: The slice DigitsAfterDecimal is filled with the exact number of digits allowed not a range, for example a range 1 to 3 will be given as [1, 2, 3].
//
//	ok := validatorgo.IsCurrency("£10.50", validatorgo.IsCurrencyOpts{Symbol: "£"})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsCurrency("£10.50", validatorgo.IsCurrencyOpts{Symbol: "$"})
//	fmt.Println(ok) // false
func IsCurrency(str string, opts IsCurrencyOpts) bool {
	escOpsSym := regexp.QuoteMeta(opts.Symbol)

	var reqSymReStr = "?"

	if opts.RequireSymbol {
		reqSymReStr = ""
	}

	alwSpcAftSymStr := ""

	if opts.AllowSpaceAfterSymbol {
		alwSpcAftSymStr = " ?"
	}

	re, err := regexp.Compile(fmt.Sprintf(`^(\(?)([+-]?(%s%s)%s[\d\,\.]*)(\)?)$`, escOpsSym, reqSymReStr, alwSpcAftSymStr))

	if err != nil {
		fmt.Println("re did not compile for reSym", err)
		return false
	}

	capGrp := re.FindStringSubmatch(str)

	if len(capGrp) == 0 {
		fmt.Println("re did not match for reSym", re.String())
		return false
	}

	// // check whether bracket match
	// lb, rb := capGrp[1], capGrp[4]
	// if lb != "" || rb != "" {
	// 	corBra := lb == "(" && rb == ")"
		
	// 	if !corBra {
	// 		fmt.Println("improper bracket format")
	// 		return false
	// 	}
	// }

	// // check whether symbols match
	// strSym := capGrp[3]
	// if strSym != opts.Symbol {
	// 	fmt.Println("symbols do not match")
	// 	return false
	// }

	// fmt.Println(1, capGrp[0])
	// fmt.Println(2, capGrp[1])
	// fmt.Println(3, capGrp[2])
	// fmt.Println(4, capGrp[3])
	// fmt.Println(5, capGrp[4])
	// fmt.Println(reSymCap, len(reSymCap))

	// if !reSym.MatchString(str) {
	// 	fmt.Println("re did not match for reSym", reSym.String())
	// 	return false
	// }

	fmt.Println(re.String())
	// fmt.Println(reSym.Find)
	return true
}
