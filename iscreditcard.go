package validatorgo

import "regexp"

const (
	isCreditCardOptsDefaultProvider string = ""
)

// IsCreditCardOpts is used to configure IsCreditCard
type IsCreditCardOpts struct {
	Provider string // credit card issuer
}

var creditCardProviderRegex = map[string]string{
	"amex":           "^3[47][0-9]{13}$",
	"bcglobal":       "^(6541|6556)[0-9]{12}$",
	"carteblanche":   "^389[0-9]{11}$",
	"dinersclub":     "^3(?:0[0-5]|[68][0-9])[0-9]{11}$",
	"discover":       "^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$",
	"instapayment":   "^63[7-9][0-9]{13}$",
	"jcb":            "^(?:2131|1800|35\\d{3})\\d{11}$",
	"koreanlocal":    "^9[0-9]{15}$",
	"laser":          "^(6304|6706|6709|6771)[0-9]{12,15}$",
	"maestro":        "^(5018|5020|5038|6304|6759|6761|6763)[0-9]{8,15}$",
	"mastercard":     "^(5[1-5][0-9]{14}|2(22[1-9][0-9]{12}|2[3-9][0-9]{13}|[3-6][0-9]{14}|7[0-1][0-9]{13}|720[0-9]{12}))$",
	"solo":           "^(6334|6767)[0-9]{12}|(6334|6767)[0-9]{14}|(6334|6767)[0-9]{15}$",
	"switch":         "^(4903|4905|4911|4936|6333|6759)[0-9]{12}|(4903|4905|4911|4936|6333|6759)[0-9]{14}|(4903|4905|4911|4936|6333|6759)[0-9]{15}|564182[0-9]{10}|564182[0-9]{12}|564182[0-9]{13}|633110[0-9]{10}|633110[0-9]{12}|633110[0-9]{13}$",
	"unionpay":       "^(62[0-9]{14,17})$",
	"visa":           "^4[0-9]{12}(?:[0-9]{3})?$",
	"visamastercard": "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14})$",
}

// A validator that checks if the string is a credit card number.
// 
// IsCreditCardOpts is an struct that can be supplied with the following key(s): 
//
// Provider: is a key whose value should be a string, and defines the company issuing the credit card.
// Valid values include amex, bcglobal, carteblanche, dinersclub, discover, instapayment, jcb, koreanlocal, laser, maestro, mastercard, solo, switch, unionpay, visa, visamastercard or blank will check for any provider.
//
//	ok := validatorgo.IsCreditCard("378282246310005", &validatorgo.IsCreditCardOpts{Provider: "amex"})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsCreditCard("37828224631000", &validatorgo.IsCreditCardOpts{Provider: "amex"})
//	fmt.Println(ok) // false
func IsCreditCard(str string, opts *IsCreditCardOpts) bool {
	if opts == nil {
		opts =  setIsCreditCardOptsToDefault()
	}

	reStr, ok := creditCardProviderRegex[opts.Provider]

	if !ok {
		if opts.Provider != "" {
			return false
		}

		for _, val := range creditCardProviderRegex {
			if isVal := regexp.MustCompile(val).MatchString(str); isVal {
				return true
			}
		}
	} else {
		return regexp.MustCompile(reStr).MatchString(str)
	}

	return false
}

func setIsCreditCardOptsToDefault() (opts *IsCreditCardOpts) {
	opts = &IsCreditCardOpts{}
	opts.Provider = isCreditCardOptsDefaultProvider

	return
}