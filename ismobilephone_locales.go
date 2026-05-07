package validatorgo

import (
	"regexp"
)

// mobilePhoneLocaleRegex is the set of locales and their function validating mobile phone regex
var mobilePhoneLocaleRegex = map[string]func(opts IsMobilePhoneOpts) *regexp.Regexp{
	"am-AM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+374(33|4[134]|55|77|88|9[13-689])\d{6}$`) // Armenia
		}
		return regexp.MustCompile(`^(\+?374|0)(33|4[134]|55|77|88|9[13-689])\d{6}$`)
	},
	"ar-AE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9715[024568]\d{7}$`) // UAE
		}
		return regexp.MustCompile(`^((\+?971)|0)?5[024568]\d{7}$`)
	},
	"ar-BH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+973(3|6)\d{7}$`) // Bahrain
		}
		return regexp.MustCompile(`^(\+?973)?(3|6)\d{7}$`)
	},
	"ar-EH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+212[6-7]\d{8}$`) // Western Sahara
		}
		return regexp.MustCompile(`^(\+?212|0)?[6-7]\d{8}$`)
	},
	"ar-DZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+213(5|6|7)\d{8}$`) // Algeria
		}
		return regexp.MustCompile(`^(\+?213|0)(5|6|7)\d{8}$`)
	},
	"ar-EG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+201[0125]\d{8}$`) // Egypt
		}
		return regexp.MustCompile(`^((\+?20)|0)?1[0125]\d{8}$`)
	},
	"ar-IQ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9647[0-9]\d{8}$`) // Iraq
		}
		return regexp.MustCompile(`^(\+?964|0)?7[0-9]\d{8}$`)
	},
	"ar-JO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9627[789]\d{7}$`) // Jordan
		}
		return regexp.MustCompile(`^(\+?962|0)?7[789]\d{7}$`)
	},
	"ar-KW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+965([569]\d{7}|41\d{6})$`) // Kuwait
		}
		return regexp.MustCompile(`^(\+?965)?([569]\d{7}|41\d{6})$`)
	},
	"ar-PS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9705[69]\d{7}$`) // Palestine
		}
		return regexp.MustCompile(`^(\+?970|0)5[69]\d{7}$`)
	},
	"ar-SA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9665\d{8}$`) // Saudi Arabia
		}
		return regexp.MustCompile(`^(!?(\+?966)|0)?5\d{8}$`)
	},
	"ar-SD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+249(9[012369]|1[012])\d{7}$`) // Sudan
		}
		return regexp.MustCompile(`^((\+?249)|0)?(9[012369]|1[012])\d{7}$`)
	},
	"ar-SY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9639\d{8}$`) // Syria
		}
		return regexp.MustCompile(`^(!?(\+?963)|0)?9\d{8}$`)
	},
	"ar-TN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+216[2459]\d{7}$`) // Tunisia
		}
		return regexp.MustCompile(`^(\+?216)?[2459]\d{7}$`)
	},
	"ar-YE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+967(7[0137]\d{7}|[1-7]\d{6})$`) // Yemen
		}
		return regexp.MustCompile(`^(((\+|00)9677|0?7)[0137]\d{7}|((\+|00)967|0)[1-7]\d{6})$`)
	},
	"az-AZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+994(10|5[015]|7[07]|99)\d{7}$`) // Azerbaijan
		}
		return regexp.MustCompile(`^(\+994|0)(10|5[015]|7[07]|99)\d{7}$`)
	},
	"az-LB": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+961\d{7,8}$`) // Lebanon
		}
		return regexp.MustCompile(`^0\d{7,8}$`)
	},
	"az-LY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+2189\d{8}$`) // Libya
		}
		return regexp.MustCompile(`^09\d{8}$`)
	},
	// Example for 'be-BY'
	"be-BY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+375(24|25|29|33|44)\d{7}$`) // Belarus
		}
		return regexp.MustCompile(`^(\+?375|0)?(24|25|29|33|44)\d{7}$`)
	},
	"bg-BG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+3598[789]\d{7}$`) // Bulgaria
		}
		return regexp.MustCompile(`^(\+?359|0)?8[789]\d{7}$`)
	},
	"bn-BD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+8801[13456789][0-9]{8}$`) // Bangladesh
		}
		return regexp.MustCompile(`^(\+?880|0)1[13456789][0-9]{8}$`)
	},
	"bs-BA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+3876(([0-3]|[5-6])\d{6}|4\d{7})$`) // Bosnia and Herzegovina
		}
		return regexp.MustCompile(`^((((\+|00)3876)|06))(([0-3]|[5-6])\d{6}|4\d{7})$`)
	},
	"ca-AD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+376[346]\d{5}$`) // Andorra
		}
		return regexp.MustCompile(`^(\+376)?[346]\d{5}$`)
	},
	"cs-CZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+420[1-9][0-9]{8}$`) // Czech Republic
		}
		return regexp.MustCompile(`^(\+?420)?[1-9][0-9]{8}$`)
	},
	"da-DK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+45\d{8}$`) // Denmark
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"de-AT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+43\d{9}$`) // Austria
		}
		return regexp.MustCompile(`^(\+43|0)\d{9}$`)
	},
	"de-CH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+41[1-9]\d{8}$`) // Switzerland
		}
		return regexp.MustCompile(`^(\+41|0)[1-9]\d{8}$`)
	},
	"de-DE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+491(5[0-25-9]\d|6([23]|0\d?)|7([0-57-9]|6\d))\d{7}$`) // Germany
		}
		return regexp.MustCompile(`^((\+49|0)1)(5[0-25-9]\d|6([23]|0\d?)|7([0-57-9]|6\d))\d{7}$`)
	},
	"de-LU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+352(6\d1)\d{6}$`) // Luxembourg
		}
		return regexp.MustCompile(`^(\+352)?((6\d1)\d{6})$`)
	},
	"dv-MV": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+960(7[2-9]|9[1-9])\d{5}$`) // Maldives
		}
		return regexp.MustCompile(`^(\+?960)?(7[2-9]|9[1-9])\d{5}$`)
	},
	"dz-BT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+975(17|16|77|02)\d{6}$`) // Bhutan
		}
		return regexp.MustCompile(`^(\+?975|0)?(17|16|77|02)\d{6}$`)
	},
	"el-CY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+357?9[45679]\d{6}$`) // Cyprus
		}
		return regexp.MustCompile(`^(\+?357?)?(9[45679]\d{6})$`)
	},
	"el-GR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+306(8[5-9]|9[013-57-9])\d{7}$`) // Greece
		}
		return regexp.MustCompile(`^(\+?30|0)?6(8[5-9]|9[013-57-9])\d{7}$`)
	},
	"en-AG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1268\d{7}$`) // Antigua and Barbuda
		}
		return regexp.MustCompile(`^(\+?1)?268\d{7}$`)
	},
	"en-AI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1264\d{7}$`) // Anguilla
		}
		return regexp.MustCompile(`^(\+?1|0)?264\d{7}$`)
	},
	"en-AU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+614\d{8}$`) // Australia
		}
		return regexp.MustCompile(`^(\+?61|0)4\d{8}$`)
	},
	"en-BM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1441\d{7}$`) // Bermuda
		}
		return regexp.MustCompile(`^(\+?1)?441\d{7}$`)
	},
	"en-BS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1242\d{7}$`) // Bahamas
		}
		return regexp.MustCompile(`^(\+?1|0)?242\d{7}$`)
	},
	"en-BW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+2677[1-8]\d{6}$`) // Botswana
		}
		return regexp.MustCompile(`^(\+?267)?(7[1-8])\d{6}$`)
	},
	"en-CA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1\d{10}$`) // Canada
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"en-GB": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+447[1-9]\d{8}$`) // United Kingdom
		}
		return regexp.MustCompile(`^(\+?44|0)7[1-9]\d{8}$`)
	},
	"en-GG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+44\d{10}$`) // Guernsey (shares UK's country code)
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"en-GH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+233\d{9}$`) // Ghana
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-GY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+5926\d{6}$`) // Guyana
		}
		return regexp.MustCompile(`^(\+592|0)?6\d{6}$`)
	},
	"en-HK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+852[456789]\d{7}$`) // Hong Kong
		}
		return regexp.MustCompile(`^(\+?852)?[456789]\d{7}$`)
	},
	"en-IE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+3538[356789]\d{7}$`) // Ireland
		}
		return regexp.MustCompile(`^(\+?353|0)8[356789]\d{7}$`)
	},
	"en-IN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+91[6789]\d{9}$`) // India
		}
		return regexp.MustCompile(`^(\+?91|0)?[6789]\d{9}$`)
	},
	"en-JM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1876\d{7}$`) // Jamaica
		}
		return regexp.MustCompile(`^(\+?1)?876\d{7}$`)
	},
	"en-KE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+254(7|1)\d{8}$`) // Kenya
		}
		return regexp.MustCompile(`^(\+?254|0)(7|1)\d{8}$`)
	},
	"en-KI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+686[67][238]\d{6}$`) // Kiribati
		}
		return regexp.MustCompile(`^(\+?686)?[67][238]\d{6}$`)
	},
	"en-KN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1869\d{7}$`) // Saint Kitts and Nevis
		}
		return regexp.MustCompile(`^(\+?1)?869\d{7}$`)
	},
	"en-LS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+266\d{8}$`) // Lesotho
		}
		return regexp.MustCompile(`^(\+?266)?\d{8}$`)
	},
	"en-MO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+8536\d{7}$`) // Macau
		}
		return regexp.MustCompile(`^(\+?853)?6\d{7}$`)
	},
	"en-MT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+356(99|79|77|21|27|22|25)\d{6}$`) // Malta
		}
		return regexp.MustCompile(`^(\+?356|0)?(99|79|77|21|27|22|25)\d{6}$`)
	},
	"en-MU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+230\d{8}$`) // Mauritius
		}
		return regexp.MustCompile(`^(\+?230|0)?\d{8}$`)
	},
	"en-MW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+265(((77|88|31|99|98|21)\d{7})|(((111)|1)\d{6})|(32000\d{4}))$`) // Malawi
		}
		return regexp.MustCompile(`^(\+?265|0)(((77|88|31|99|98|21)\d{7})|(((111)|1)\d{6})|(32000\d{4}))$`)
	},
	"en-NG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+234[789]\d{9}$`) // Nigeria
		}
		return regexp.MustCompile(`^(\+?234|0)?[789]\d{9}$`)
	},
	"en-NZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+64[28]\d{8}$`) // New Zealand
		}
		return regexp.MustCompile(`^(\+?64|0)[28]\d{8}$`)
	},
	"en-PG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+675(7\d|8[18])\d{6}$`) // Papua New Guinea
		}
		return regexp.MustCompile(`^(\+?675|0)?(7\d|8[18])\d{6}$`)
	},
	"en-PH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+639\d{9}$`) // Philippines
		}
		return regexp.MustCompile(`^(09|\+639)\d{9}$`)
	},
	"en-PK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+923[0-6]\d{8}$`) // Pakistan
		}
		return regexp.MustCompile(`^((00|\+)?92|0)3[0-6]\d{8}$`)
	},
	"en-RW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+250[7]\d{8}$`) // Rwanda
		}
		return regexp.MustCompile(`^(\+?250|0)?[7]\d{8}$`)
	},
	"en-SG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+65[3689]\d{7}$`) // Singapore
		}
		return regexp.MustCompile(`^(\+65)?[3689]\d{7}$`)
	},
	"en-SL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+232\d{9}$`) // Sierra Leone
		}
		return regexp.MustCompile(`^(\+?232|0)\d{9}$`)
	},
	"en-SS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+211(9[1257])\d{7}$`) // South Sudan
		}
		return regexp.MustCompile(`^(\+?211|0)(9[1257])\d{7}$`)
	},
	"en-TZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+255[67]\d{8}$`) // Tanzania
		}
		return regexp.MustCompile(`^(\+?255|0)?[67]\d{8}$`)
	},
	"en-UG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+256[7]\d{8}$`) // Uganda
		}
		return regexp.MustCompile(`^(\+?256|0)?[7]\d{8}$`)
	},
	"en-US": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1\d{10}$`) // United States
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"en-ZA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+27\d{9}$`) // South Africa
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-ZM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+260[79][567]\d{7}$`) // Zambia
		}
		return regexp.MustCompile(`^(\+?26)?0[79][567]\d{7}$`)
	},
	"en-ZW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+263\d{9}$`) // Zimbabwe
		}
		return regexp.MustCompile(`^(\+263|0)\d{9}$`)
	},
	"es-AR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+549(11|[2368]\d)\d{8}$`) // Argentina
		}
		return regexp.MustCompile(`^(\+?549|0)?(11|[2368]\d)\d{8}$`)
	},
	"es-BO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+591(6|7)\d{7}$`) // Bolivia
		}
		return regexp.MustCompile(`^(\+?591)?(6|7)\d{7}$`)
	},
	"es-CL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+56\d{9}$`) // Chile
		}
		return regexp.MustCompile(`^(\+?56|0)\d{8}$`)
	},
	"es-CO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+573(0[01245]|1\d|2[0-4]|5[01])\d{7}$`) // Colombia
		}
		return regexp.MustCompile(`^(\+?57)?3(0[01245]|1\d|2[0-4]|5[01])\d{7}$`)
	},
	"es-CR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+506[2-8]\d{7}$`) // Costa Rica
		}
		return regexp.MustCompile(`^(\+506)?[2-8]\d{7}$`)
	},
	"es-CU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+535\d{7}$`) // Cuba
		}
		return regexp.MustCompile(`^(\+53|0053)?5\d{7}$`)
	},
	"es-DO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+18[024]9\d{7}$`) // Dominican Republic
		}
		return regexp.MustCompile(`^(\+?1)?8[024]9\d{7}$`)
	},
	"es-EC": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+593([2-7]|9[2-9])\d{7}$`) // Ecuador
		}
		return regexp.MustCompile(`^(\+?593|0)([2-7]|9[2-9])\d{7}$`)
	},
	"es-ES": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+34[67]\d{8}$`) // Spain
		}
		return regexp.MustCompile(`^(\+?34)?[67]\d{8}$`)
	},
	"es-GT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+502\d{8}$`) // Guatemala
		}
		return regexp.MustCompile(`^(\+?502)?\d{8}$`)
	},
	"es-HN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+504[9832]\d{7}$`) // Honduras
		}
		return regexp.MustCompile(`^(\+?504)?[9832]\d{7}$`)
	},
	"es-MX": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+52\d{10}$`) // Mexico
		}
		return regexp.MustCompile(`^(\+?52)?\d{10}$`)
	},
	"es-NI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+505\d{8}$`) // Nicaragua
		}
		return regexp.MustCompile(`^(\+?505)?\d{8}$`)
	},
	"es-PA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+507\d{8}$`) // Panama
		}
		return regexp.MustCompile(`^(\+?507)?\d{8}$`)
	},
	"es-PE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+519\d{8}$`) // Peru
		}
		return regexp.MustCompile(`^(\+?51)?9\d{8}$`)
	},
	"es-PR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1\d{10}$`) // Puerto Rico
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"es-PY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+5959[9876]\d{7}$`) // Paraguay
		}
		return regexp.MustCompile(`^(\+?595|0)9[9876]\d{7}$`)
	},
	"es-SV": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+503[67]\d{7}$`) // El Salvador
		}
		return regexp.MustCompile(`^(\+?503)?[67]\d{7}$`)
	},
	"es-UY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+5989[1-9]\d{6}$`) // Uruguay
		}
		return regexp.MustCompile(`^(\+598|0)?9[1-9]\d{6}$`)
	},
	"es-VE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+58(2|4)\d{9}$`) // Venezuela
		}
		return regexp.MustCompile(`^(\+?58|0)?(2|4)\d{9}$`)
	},
	"et-EE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+372(5\d{7}|8[1-4]\d{6,7})$`) // Estonia
		}
		return regexp.MustCompile(`^(\+?372)?(5\d{7}|8[1-4]\d{6,7})$`)
	},
	"fa-IR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+98\d{10}$`) // Iran
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"fi-FI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+358(4[0-6]|50)\d{7}$`) // Finland
		}
		return regexp.MustCompile(`^(\+?358|0)?(4[0-6]|50)\d{7}$`)
	},
	"fil-PH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+63\d{10}$`) // Philippines
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"fr-BE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+324\d{8}$`) // Belgium
		}
		return regexp.MustCompile(`^(\+?32|0)4\d{8}$`)
	},
	"fr-BF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+226[67]\d{8}$`) // Burkina Faso
		}
		return regexp.MustCompile(`^(\+226|0)[67]\d{8}$`)
	},
	"fr-BJ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+229\d{8}$`) // Benin
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-CD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+243(8|9)\d{8}$`) // Democratic Republic of the Congo
		}
		return regexp.MustCompile(`^(\+?243|0)?(8|9)\d{8}$`)
	},
	"fr-CF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+236(70|75|77|72|21|22)\d{7}$`) // Central African Republic
		}
		return regexp.MustCompile(`^(\+?236|0)?(70|75|77|72|21|22)\d{7}$`)
	},
	"fr-CG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+242\d{9}$`) // Congo
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-CH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+41[1-9]\d{8}$`) // Switzerland
		}
		return regexp.MustCompile(`^(\+41|0)[1-9]\d{8}$`)
	},
	"fr-CI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+225\d{8}$`) // Ivory Coast
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-CM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+237\d{8}$`) // Cameroon
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-DJ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+253\d{8}$`) // Djibouti
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-DZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+213\d{9}$`) // Algeria
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-FR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+33[67]\d{8}$`) // France
		}
		return regexp.MustCompile(`^(\+?33|0)[67]\d{8}$`)
	},
	"fr-GF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+594[67]\d{8}$`) // French Guiana
		}
		return regexp.MustCompile(`^(\+?594|0|00594)[67]\d{8}$`)
	},
	"fr-GP": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+590[67]\d{8}$`) // Guadeloupe
		}
		return regexp.MustCompile(`^(\+?590|0|00590)[67]\d{8}$`)
	},
	"fr-MQ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+596[67]\d{8}$`) // Martinique
		}
		return regexp.MustCompile(`^(\+?596|0|00596)[67]\d{8}$`)
	},
	"fr-RE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+262[67]\d{8}$`) // Réunion
		}
		return regexp.MustCompile(`^(\+?262|0|00262)[67]\d{8}$`)
	},
	"fr-GA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+241\d{7}$`) // Gabon
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"fr-GN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+224\d{9}$`) // Guinea
		}
		return regexp.MustCompile(`^(\+?224)?\d{9}$`)
	},
	"fr-GQ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+240\d{9}$`) // Equatorial Guinea
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-HT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+509\d{8}$`) // Haiti
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-LU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+352\d{9}$`) // Luxembourg
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-MA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+212\d{9}$`) // Morocco
		}
		return regexp.MustCompile(`^(\+?212|0)\d{9}$`)
	},
	"fr-MG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+261\d{9}$`) // Madagascar
		}
		return regexp.MustCompile(`^(\+?261|0)\d{9}$`)
	},
	"fr-ML": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+223\d{8}$`) // Mali
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-MR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+222\d{8}$`) // Mauritania
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-MU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+230\d{7}$`) // Mauritius
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"fr-NC": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+687\d{6}$`) // New Caledonia
		}
		return regexp.MustCompile(`^\d{6}$`)
	},
	"fr-NE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+227\d{8}$`) // Niger
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-PF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+689\d{6}$`) // French Polynesia
		}
		return regexp.MustCompile(`^\d{6}$`)
	},
	"fr-RW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+250\d{9}$`) // Rwanda
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-SN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+221\d{9}$`) // Senegal
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-SY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+963\d{9}$`) // Syria
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-TD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+235\d{8}$`) // Chad
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-TG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+228\d{7,8}$`) // Togo
		}
		return regexp.MustCompile(`^(\+?228)?\d{7,8}$`)
	},
	"fr-TN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+216\d{8}$`) // Tunisia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-VU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+678\d{7}$`) // Vanuatu
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"fr-WF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+681\d{6}$`) // Wallis and Futuna
		}
		return regexp.MustCompile(`^\d{6}$`)
	},
	"ga-IE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+353\d{9}$`) // Ireland
		}
		return regexp.MustCompile(`^(\+?353|0)\d{9}$`)
	},
	"he-IL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+972\d{9}$`) // Israel
		}
		return regexp.MustCompile(`^(\+?972|0)\d{9}$`)
	},
	"hi-IN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+91\d{10}$`) // India
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"hr-HR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+385\d{8,9}$`) // Croatia
		}
		return regexp.MustCompile(`^\d{8,9}$`)
	},
	"hu-HU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+36\d{9}$`) // Hungary
		}
		return regexp.MustCompile(`^(\+?36|06)\d{9}$`)
	},
	"hy-AM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+374\d{8}$`) // Armenia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"id-ID": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+62\d{10}$`) // Indonesia
		}
		return regexp.MustCompile(`^(\+?62|0)\d{10}$`)
	},
	"it-IT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+39\d{10}$`) // Italy
		}
		return regexp.MustCompile(`^(\+?39)?\d{10}$`)
	},
	"ja-JP": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+81\d{10}$`) // Japan
		}
		return regexp.MustCompile(`^(\+?81|0)\d{10}$`)
	},
	"ka-GE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+995\d{9}$`) // Georgia
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"kk-KZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+7\d{10}$`) // Kazakhstan
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"km-KH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+855\d{8,9}$`) // Cambodia
		}
		return regexp.MustCompile(`^\d{8,9}$`)
	},
	"ko-KR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+82\d{9,10}$`) // South Korea
		}
		return regexp.MustCompile(`^\d{9,10}$`)
	},
	"ky-KG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+996\d{9}$`) // Kyrgyzstan
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"lt-LT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+370\d{8}$`) // Lithuania
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"lv-LV": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+371\d{8}$`) // Latvia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"mn-MN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+976\d{8}$`) // Mongolia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"ms-MY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+60\d{9,10}$`) // Malaysia
		}
		return regexp.MustCompile(`^\d{9,10}$`)
	},
	"ne-NP": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+977\d{9,10}$`) // Nepal
		}
		return regexp.MustCompile(`^\d{9,10}$`)
	},
	"nl-BE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+324\d{8}$`) // Belgium (Dutch)
		}
		return regexp.MustCompile(`^(\+?32|0)4\d{8}$`)
	},
	"nl-NL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+31\d{9}$`) // Netherlands
		}
		return regexp.MustCompile(`^(\+?31|0)\d{9}$`)
	},
	"no-NO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+47\d{8}$`) // Norway
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"pl-PL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+48\d{8,9}$`) // Poland
		}
		return regexp.MustCompile(`^(\+?48)?\d{8,9}$`)
	},
	"pt-AO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+244\d{9}$`) // Angola
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"pt-BR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+55\d{11}$`) // Brazil
		}
		return regexp.MustCompile(`^(\+?55)?\d{11}$`)
	},
	"pt-CV": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+238\d{7}$`) // Cape Verde
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"pt-GW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+245\d{7}$`) // Guinea-Bissau
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"pt-MZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+258\d{9}$`) // Mozambique
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"pt-PT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+351\d{9}$`) // Portugal
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"ro-MD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+373((6[012689])|(7[6789]))\d{6}$`) // Moldova
		}
		return regexp.MustCompile(`^(\+?373|0)((6[012689])|(7[6789]))\d{6}$`)
	},
	"ro-Md": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+373((6[012689])|(7[6789]))\d{6}$`) // Moldova (alias)
		}
		return regexp.MustCompile(`^(\+?373|0)((6[012689])|(7[6789]))\d{6}$`)
	},
	"ro-RO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+40\d{9}$`) // Romania
		}
		return regexp.MustCompile(`^(\+?40|0)\d{9}$`)
	},
	"ru-BY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+375\d{9}$`) // Belarus
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"ru-KG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+996\d{9}$`) // Kyrgyzstan
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"ru-KZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+7\d{10}$`) // Kazakhstan
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"ru-RU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+7\d{10}$`) // Russia
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"ru-UA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+380\d{9}$`) // Ukraine
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"si-LK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+94\d{9}$`) // Sri Lanka
		}
		return regexp.MustCompile(`^(\+?94|0)\d{9}$`)
	},
	"sk-SK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+421\d{9}$`) // Slovakia
		}
		return regexp.MustCompile(`^(\+?421|0)\d{9}$`)
	},
	"sl-SI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+386\d{8}$`) // Slovenia
		}
		return regexp.MustCompile(`^(\+?386|0)\d{8}$`)
	},
	"sq-AL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+355\d{9}$`) // Albania
		}
		return regexp.MustCompile(`^(\+?355|0)\d{9}$`)
	},
	"sq-MK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+389\d{8}$`) // North Macedonia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"sr-BA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+387\d{8}$`) // Bosnia and Herzegovina
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"sr-ME": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+382\d{8}$`) // Montenegro
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"sr-RS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+381\d{9}$`) // Serbia
		}
		return regexp.MustCompile(`^(\+?381|0)\d{9}$`)
	},
	"sv-FI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+358\d{9}$`) // Finland (Swedish-speaking)
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"sv-SE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+46\d{9}$`) // Sweden
		}
		return regexp.MustCompile(`^(\+?46|0)\d{9}$`)
	},
	"ta-IN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+91\d{10}$`) // India (Tamil)
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"ta-LK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+94\d{9}$`) // Sri Lanka (Tamil)
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"th-TH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+66\d{8,9}$`) // Thailand
		}
		return regexp.MustCompile(`^\d{8,9}$`)
	},
	"tr-CY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+90\d{10}$`) // Cyprus (Turkish-speaking)
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"tr-TR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+90\d{10}$`) // Turkey
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"uk-UA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+380\d{9}$`) // Ukraine
		}
		return regexp.MustCompile(`^(\+?380|0)\d{9}$`)
	},
	"ur-PK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+92\d{10}$`) // Pakistan
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"uz-UZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+998\d{9}$`) // Uzbekistan
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"vi-VN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+84\d{9,10}$`) // Vietnam
		}
		return regexp.MustCompile(`^\d{9,10}$`)
	},
	"zh-CN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+86\d{11}$`) // China
		}
		return regexp.MustCompile(`^\d{11}$`)
	},
	"zh-HK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+852\d{8}$`) // Hong Kong
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"zh-MO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+853\d{8}$`) // Macau
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"zh-SG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+65\d{8}$`) // Singapore (Mandarin-speaking)
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"zh-TW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+8869\d{8}$`) // Taiwan
		}
		return regexp.MustCompile(`^(\+?886\-?|0)?9\d{8}$`)
	},
	"cy-GB": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+447[1-9]\d{8}$`) // Wales (same as en-GB)
		}
		return regexp.MustCompile(`^(\+?44|0)7[1-9]\d{8}$`)
	},
	"fa-AF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+93(2[0-8]|[3-5][0-4]|7[0-9])\d{7}$`) // Afghanistan
		}
		return regexp.MustCompile(`^(\+93|0)?(2[0-8]|[3-5][0-4]|7[0-9])\d{7}$`)
	},
	"fj-FJ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+679\d{7}$`) // Fiji
		}
		return regexp.MustCompile(`^(\+?679)?\d{7}$`)
	},
	"fo-FO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+298\d{6}$`) // Faroe Islands
		}
		return regexp.MustCompile(`^(\+?298)?\d{6}$`)
	},
	"fr-SC": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+248\d{7}$`) // Seychelles
		}
		return regexp.MustCompile(`^(\+?248)?\d{7}$`)
	},
	"ir-IR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+989\d{9}$`) // Iran
		}
		return regexp.MustCompile(`^(\+98|0)?9\d{9}$`)
	},
	"it-SM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+378\d{9}$`) // San Marino
		}
		return regexp.MustCompile(`^((\+378)|(0549)|(\+390549)|(\+3780549))?\d{9}$`)
	},
	"nl-AW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+297\d{9}$`) // Aruba
		}
		return regexp.MustCompile(`^(\+?297)?\d{9}$`)
	},
	"nn-NO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+47[49]\d{7}$`) // Norway (Nynorsk)
		}
		return regexp.MustCompile(`^(\+?47)?[49]\d{7}$`)
	},
	"so-SO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+252(6[0-9]\d{7}|7[1-9]\d{7})$`) // Somalia
		}
		return regexp.MustCompile(`^(\+?252|0)?((6[0-9])\d{7}|(7[1-9])\d{7})$`)
	},
}
