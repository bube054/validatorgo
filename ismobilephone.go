package validatorgo

// IsMobilePhoneOpts is used to configure IsMobilePhone
type IsMobilePhoneOpts struct {
	StrictMode *bool // matches format exactly
}

func (o *IsMobilePhoneOpts) mergeDefaults() {
	if o.StrictMode == nil {
		o.StrictMode = Bool(false)
	}
}

// A validator that checks if the string is a mobile phone number.
//
// locale is either a slice of locales e.g. []string{'sk-SK', 'sr-RS'} possible options are ("am-Am", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-EH", "ar-IQ", "ar-JO", "ar-KW", "ar-PS", "ar-SA", "ar-SD", "ar-SY", "ar-TN", "ar-YE", "az-AZ", "az-LB", "az-LY", "be-BY", "bg-BG", "bn-BD", "bs-BA", "ca-AD", "cs-CZ", "da-DK", "de-AT", "de-CH", "de-DE", "de-LU", "dv-MV", "dz-BT", "el-CY", "el-GR", "en-AG", "en-AI", "en-AU", "en-BM", "en-BS", "en-BW", "en-CA", "en-GB", "en-GG", "en-GH", "en-GY", "en-HK", "en-IE", "en-IN", "en-JM", "en-KE", "en-KI", "en-KN", "en-LS", "en-MO", "en-MT", "en-MU", "en-MW", "en-NG", "en-NZ", "en-PG", "en-PH", "en-PK", "en-RW", "en-SG", "en-SL", "en-SS", "en-TZ", "en-UG", "en-US", "en-ZA", "en-ZM", "en-ZW", "es-AR", "es-BO", "es-CL", "es-CO", "es-CR", "es-CU", "es-DO", "es-EC", "es-ES", "es-GT","es-HN", "es-MX", "es-NI", "es-PA", "es-PE", "es-PY", "es-SV", "es-UY", "es-VE", "et-EE", "fa-AF", "fa-IR", "fi-FI", "fj-FJ", "fo-FO", "fr-BE", "fr-BF", "fr-BJ", "fr-CD", "fr-CF", "fr-FR", "fr-GF", "fr-GP", "fr-MQ", "fr-PF", "fr-RE", "fr-WF", "ga-IE", "he-IL", "hu-HU", "id-ID", "ir-IR", "it-IT", "it-SM", "ja-JP", "ka-GE", "kk-KZ", "kl-GL", "ko-KR", "ky-KG", "lt-LT", "mg-MG", "mn-MN", "ms-MY", "my-MM", "mz-MZ", "nb-NO", "ne-NP", "nl-AW", "nl-BE", "nl-NL", "nn-NO", "pl-PL", "pt-AO", "pt-BR", "pt-PT", "ro-Md", "ro-RO", "ru-RU", "si-LK", "sk-SK", "sl-SI", "so-SO", "sq-AL", "sr-RS", "sv-SE", "tg-TJ", "th-TH", "tk-TM", "tr-TR", "uk-UA", "uz-UZ", "vi-VN", "zh-CN", "zh-HK", "zh-MO", "zh-TW"). If nil is provided any of the locales will be matched. If an unidentified value is used, function will return false.
//
// IsMobilePhoneOpts is a struct that can be supplied with the following keys:
//
// StrictMode, if this is set to true, the mobile phone number must be supplied with the country code and therefore must start with +.
//
//	ok, _ := validatorgo.IsMobilePhone("08070448986", []string{"en-US"} , nil)
//	fmt.Println(ok) // true
//	ok, _ = validatorgo.IsMobilePhone("090666666567", []string{"en-US"} , nil)
//	fmt.Println(ok) // false
func IsMobilePhone(str string, locales []string, opts *IsMobilePhoneOpts) (bool, error) {
	if opts == nil {
		opts = &IsMobilePhoneOpts{}
	}
	opts.mergeDefaults()

	strWithoutSpaces := stripDashesAndSpaces(str)

	if len(locales) == 0 {
		for _, valFunc := range mobilePhoneLocaleRegex {
			re := valFunc(*opts)
			match := re.MatchString(strWithoutSpaces)

			if match {
				return true, nil
			}
		}
	} else {
		for _, loc := range locales {
			valFunc, exists := mobilePhoneLocaleRegex[loc]

			if !exists {
				continue
			}

			re := valFunc(*opts)
			match := re.MatchString(strWithoutSpaces)

			if match {
				return true, nil
			}
		}
	}

	return false, newValidationError("IsMobilePhone", ErrInvalidFormat, "invalid mobilephone")
}

