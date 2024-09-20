package validatorgo

import "regexp"

// IsMobilePhoneOpts is used to configure IsMobilePhone
type IsMobilePhoneOpts struct {
	StrictMode bool // matches format exactly
}

// mobilePhoneLocaleRegex is the set of locales and their function validating mobile phone regex
var mobilePhoneLocaleRegex = map[string]func(opts IsMobilePhoneOpts) *regexp.Regexp{
	"am-AM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+374\d{8}$`) // Armenia
		}
		return regexp.MustCompile(`^0\d{8}$`)
	},
	"ar-AE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+971\d{8,9}$`) // UAE
		}
		return regexp.MustCompile(`^0\d{8,9}$`)
	},
	"ar-BH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+973\d{8}$`) // Bahrain
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"ar-DZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+213\d{9}$`) // Algeria
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"ar-EG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+20\d{10}$`) // Egypt
		}
		return regexp.MustCompile(`^01\d{9}$`)
	},
	"ar-IQ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+964\d{10}$`) // Iraq
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"ar-JO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+962\d{8,9}$`) // Jordan
		}
		return regexp.MustCompile(`^0\d{8,9}$`)
	},
	"ar-KW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+965\d{8}$`) // Kuwait
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"ar-PS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+970\d{8,9}$`) // Palestine
		}
		return regexp.MustCompile(`^0\d{8,9}$`)
	},
	"ar-SA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9665\d{8}$`) // Saudi Arabia
		}
		return regexp.MustCompile(`^05\d{8}$`)
	},
	"ar-SD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+249\d{9}$`) // Sudan
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"ar-SY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9639\d{8}$`) // Syria
		}
		return regexp.MustCompile(`^09\d{8}$`)
	},
	"ar-TN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+216\d{8}$`) // Tunisia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"ar-YE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+9677\d{7}$`) // Yemen
		}
		return regexp.MustCompile(`^07\d{7}$`)
	},
	"az-AZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+994\d{9,10}$`) // Azerbaijan
		}
		return regexp.MustCompile(`^0\d{9,10}$`)
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
			return regexp.MustCompile(`^\+375\d{9}$`) // Belarus
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"bg-BG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+359\d{9}$`) // Bulgaria
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"bn-BD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+8801\d{9}$`) // Bangladesh
		}
		return regexp.MustCompile(`^01\d{9}$`)
	},
	"bs-BA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+387\d{8}$`) // Bosnia and Herzegovina
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"ca-AD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+376\d{6}$`) // Andorra
		}
		return regexp.MustCompile(`^\d{6}$`)
	},
	"cs-CZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+420\d{9}$`) // Czech Republic
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"da-DK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+45\d{8}$`) // Denmark
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"de-AT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+43\d{10}$`) // Austria
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"de-CH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+41\d{9}$`) // Switzerland
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"de-DE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+49\d{10,11}$`) // Germany
		}
		return regexp.MustCompile(`^\d{10,11}$`)
	},
	"de-LU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+352\d{8}$`) // Luxembourg
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"dv-MV": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+960\d{7}$`) // Maldives
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"dz-BT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+975\d{8}$`) // Bhutan
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"el-CY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+357\d{8}$`) // Cyprus
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"el-GR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+30\d{10}$`) // Greece
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"en-AG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1268\d{7}$`) // Antigua and Barbuda
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-AI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1264\d{7}$`) // Anguilla
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-AU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+61\d{9}$`) // Australia
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-BM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1441\d{7}$`) // Bermuda
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-BS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1242\d{7}$`) // Bahamas
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-BW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+267\d{7,8}$`) // Botswana
		}
		return regexp.MustCompile(`^\d{7,8}$`)
	},
	"en-CA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1\d{10}$`) // Canada
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"en-GB": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+44\d{10}$`) // United Kingdom
		}
		return regexp.MustCompile(`^0\d{10}$`)
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
			return regexp.MustCompile(`^\+592\d{7}$`) // Guyana
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-HK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+852\d{8}$`) // Hong Kong
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"en-IE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+353\d{9}$`) // Ireland
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-IN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+91\d{10}$`) // India
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"en-JM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1876\d{7}$`) // Jamaica
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-KE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+254\d{9}$`) // Kenya
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-KI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+686\d{5}$`) // Kiribati
		}
		return regexp.MustCompile(`^\d{5}$`)
	},
	"en-KN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1869\d{7}$`) // Saint Kitts and Nevis
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-LS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+266\d{8}$`) // Lesotho
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"en-MO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+853\d{8}$`) // Macau
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"en-MT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+356\d{8}$`) // Malta
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"en-MU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+230\d{7}$`) // Mauritius
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-MW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+265\d{9}$`) // Malawi
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"en-NG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+234\d{10}$`) // Nigeria
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"en-NZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+64\d{9}$`) // New Zealand
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"en-PG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+675\d{7}$`) // Papua New Guinea
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"en-PH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+63\d{10}$`) // Philippines
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"en-PK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+92\d{10}$`) // Pakistan
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"en-RW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+250\d{9}$`) // Rwanda
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-SG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+65\d{8}$`) // Singapore
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"en-SL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+232\d{8}$`) // Sierra Leone
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"en-SS": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+211\d{9}$`) // South Sudan
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"en-TZ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+255\d{9}$`) // Tanzania
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-UG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+256\d{9}$`) // Uganda
		}
		return regexp.MustCompile(`^0\d{9}$`)
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
			return regexp.MustCompile(`^\+260\d{9}$`) // Zambia
		}
		return regexp.MustCompile(`^0\d{9}$`)
	},
	"en-ZW": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+263\d{9}$`) // Zimbabwe
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"es-AR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+54\d{10}$`) // Argentina
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"es-BO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+591\d{8}$`) // Bolivia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-CL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+56\d{9}$`) // Chile
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"es-CO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+57\d{10}$`) // Colombia
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"es-CR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+506\d{8}$`) // Costa Rica
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-CU": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+53\d{8}$`) // Cuba
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-DO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1809\d{7}$`) // Dominican Republic
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"es-EC": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+593\d{9}$`) // Ecuador
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"es-ES": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+34\d{9}$`) // Spain
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"es-GT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+502\d{8}$`) // Guatemala
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-HN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+504\d{8}$`) // Honduras
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-MX": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+52\d{10}$`) // Mexico
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"es-NI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+505\d{8}$`) // Nicaragua
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-PA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+507\d{8}$`) // Panama
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-PE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+51\d{9}$`) // Peru
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"es-PR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+1\d{10}$`) // Puerto Rico
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"es-PY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+595\d{9}$`) // Paraguay
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"es-SV": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+503\d{8}$`) // El Salvador
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"es-UY": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+598\d{9}$`) // Uruguay
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"es-VE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+58\d{10}$`) // Venezuela
		}
		return regexp.MustCompile(`^\d{10}$`)
	},
	"et-EE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+372\d{7,8}$`) // Estonia
		}
		return regexp.MustCompile(`^\d{7,8}$`)
	},
	"fa-IR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+98\d{10}$`) // Iran
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"fi-FI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+358\d{9}$`) // Finland
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fil-PH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+63\d{10}$`) // Philippines
		}
		return regexp.MustCompile(`^0\d{10}$`)
	},
	"fr-BE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+32\d{8}$`) // Belgium
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-BF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+226\d{8}$`) // Burkina Faso
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-BJ": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+229\d{8}$`) // Benin
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-CD": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+243\d{9}$`) // Democratic Republic of the Congo
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-CF": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+236\d{8}$`) // Central African Republic
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"fr-CG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+242\d{9}$`) // Congo
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-CH": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+41\d{9}$`) // Switzerland
		}
		return regexp.MustCompile(`^\d{9}$`)
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
			return regexp.MustCompile(`^\+33\d{9}$`) // France
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-GA": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+241\d{7}$`) // Gabon
		}
		return regexp.MustCompile(`^\d{7}$`)
	},
	"fr-GN": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+224\d{8}$`) // Guinea
		}
		return regexp.MustCompile(`^\d{8}$`)
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
		return regexp.MustCompile(`^\d{9}$`)
	},
	"fr-MG": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+261\d{9}$`) // Madagascar
		}
		return regexp.MustCompile(`^\d{9}$`)
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
			return regexp.MustCompile(`^\+228\d{8}$`) // Togo
		}
		return regexp.MustCompile(`^\d{8}$`)
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
		return regexp.MustCompile(`^\d{9}$`)
	},
	"he-IL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+972\d{8,9}$`) // Israel
		}
		return regexp.MustCompile(`^\d{8,9}$`)
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
		return regexp.MustCompile(`^\d{9}$`)
	},
	"hy-AM": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+374\d{8}$`) // Armenia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"id-ID": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+62\d{9,11}$`) // Indonesia
		}
		return regexp.MustCompile(`^\d{9,11}$`)
	},
	"it-IT": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+39\d{9,10}$`) // Italy
		}
		return regexp.MustCompile(`^\d{9,10}$`)
	},
	"ja-JP": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+81\d{9,10}$`) // Japan
		}
		return regexp.MustCompile(`^\d{9,10}$`)
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
			return regexp.MustCompile(`^\+32\d{8}$`) // Belgium (Dutch)
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"nl-NL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+31\d{9}$`) // Netherlands
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"no-NO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+47\d{8}$`) // Norway
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"pl-PL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+48\d{9}$`) // Poland
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"pt-AO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+244\d{9}$`) // Angola
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"pt-BR": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+55\d{10,11}$`) // Brazil
		}
		return regexp.MustCompile(`^\d{10,11}$`)
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
			return regexp.MustCompile(`^\+373\d{8}$`) // Moldova
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"ro-RO": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+40\d{9}$`) // Romania
		}
		return regexp.MustCompile(`^\d{9}$`)
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
		return regexp.MustCompile(`^\d{9}$`)
	},
	"sk-SK": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+421\d{9}$`) // Slovakia
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"sl-SI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+386\d{8}$`) // Slovenia
		}
		return regexp.MustCompile(`^\d{8}$`)
	},
	"sq-AL": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+355\d{9}$`) // Albania
		}
		return regexp.MustCompile(`^\d{9}$`)
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
			return regexp.MustCompile(`^\+381\d{8,9}$`) // Serbia
		}
		return regexp.MustCompile(`^\d{8,9}$`)
	},
	"sv-FI": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+358\d{9}$`) // Finland (Swedish-speaking)
		}
		return regexp.MustCompile(`^\d{9}$`)
	},
	"sv-SE": func(opts IsMobilePhoneOpts) *regexp.Regexp {
		if opts.StrictMode {
			return regexp.MustCompile(`^\+46\d{7,10}$`) // Sweden
		}
		return regexp.MustCompile(`^\d{7,10}$`)
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
		return regexp.MustCompile(`^\d{9}$`)
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
			return regexp.MustCompile(`^\+886\d{9,10}$`) // Taiwan
		}
		return regexp.MustCompile(`^\d{9,10}$`)
	},
}

// A validator that checks if the string is a mobile phone number.
//
// locale is either an array of locales e.g. []string{'sk-SK', 'sr-RS'} possible options are ('am-Am', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-EH', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-PS', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'az-AZ', 'az-LB', 'az-LY', 'be-BY', 'bg-BG', 'bn-BD', 'bs-BA', 'ca-AD', 'cs-CZ', 'da-DK', 'de-AT', 'de-CH', 'de-DE', 'de-LU', 'dv-MV', 'dz-BT', 'el-CY', 'el-GR', 'en-AG', 'en-AI', 'en-AU', 'en-BM', 'en-BS', 'en-BW', 'en-CA', 'en-GB', 'en-GG', 'en-GH', 'en-GY', 'en-HK', 'en-IE', 'en-IN', 'en-JM', 'en-KE', 'en-KI', 'en-KN', 'en-LS', 'en-MO', 'en-MT', 'en-MU', 'en-MW', 'en-NG', 'en-NZ', 'en-PG', 'en-PH', 'en-PK', 'en-RW', 'en-SG', 'en-SL', 'en-SS', 'en-TZ', 'en-UG', 'en-US', 'en-ZA', 'en-ZM', 'en-ZW', 'es-AR', 'es-BO', 'es-CL', 'es-CO', 'es-CR', 'es-CU', 'es-DO', 'es-EC', 'es-ES', 'es-GT','es-HN', 'es-MX', 'es-NI', 'es-PA', 'es-PE', 'es-PY', 'es-SV', 'es-UY', 'es-VE', 'et-EE', 'fa-AF', 'fa-IR', 'fi-FI', 'fj-FJ', 'fo-FO', 'fr-BE', 'fr-BF', 'fr-BJ', 'fr-CD', 'fr-CF', 'fr-FR', 'fr-GF', 'fr-GP', 'fr-MQ', 'fr-PF', 'fr-RE', 'fr-WF', 'ga-IE', 'he-IL', 'hu-HU', 'id-ID', 'ir-IR', 'it-IT', 'it-SM', 'ja-JP', 'ka-GE', 'kk-KZ', 'kl-GL', 'ko-KR', 'ky-KG', 'lt-LT', 'mg-MG', 'mn-MN', 'ms-MY', 'my-MM', 'mz-MZ', 'nb-NO', 'ne-NP', 'nl-AW', 'nl-BE', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-AO', 'pt-BR', 'pt-PT', 'ro-Md', 'ro-RO', 'ru-RU', 'si-LK', 'sk-SK', 'sl-SI', 'so-SO', 'sq-AL', 'sr-RS', 'sv-SE', 'tg-TJ', 'th-TH', 'tk-TM', 'tr-TR', 'uk-UA', 'uz-UZ', 'vi-VN', 'zh-CN', 'zh-HK', 'zh-MO', 'zh-TW'). If nil is provided any of the locales will be matched. If an unidentified value is used, function will return false.
//
// IsMobilePhoneOpts is a struct that can be supplied with the following keys:
//
// StrictMode, if this is set to true, the mobile phone number must be supplied with the country code and therefore must start with +.
//
//	ok := validatorg.IsMobilePhone("08070448986", []string{"en-US"} , validatorgo.IsMobilePhoneOpts{})
//	fmt.Println(ok) // true
//	ok := validatorg.IsMobilePhone("090666666567", []string{"en-US"} , validatorgo.IsMobilePhoneOpts{})
//	fmt.Println(ok) // false
func IsMobilePhone(str string, locales []string, opts IsMobilePhoneOpts) bool {
	strWithoutSpaces := stripDashesAndSpaces(str)

	if len(locales) == 0 {
		for _, valFunc := range mobilePhoneLocaleRegex {
			re := valFunc(opts)

			if !re.MatchString(strWithoutSpaces) {
				return false
			}
		}
	} else {

		for _, loc := range locales {
			valFunc, exists := mobilePhoneLocaleRegex[loc]

			if !exists {
				return false
			}

			re := valFunc(opts)

			if !re.MatchString(strWithoutSpaces) {
				return false
			}

		}
	}

	return true
}
