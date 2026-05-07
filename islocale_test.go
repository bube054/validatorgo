package validatorgo

import "testing"

func TestIsLocale(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		want   bool
	}{
		// Valid Locales
		{name: "Is valid locale", param1: "ca_ES", want: true},
		{name: "Is valid locale", param1: "it_IT", want: true},
		{name: "Is valid locale", param1: "uk_UA", want: true},
		{name: "Is valid locale", param1: "zu_ZA", want: true},
		{name: "Valid locale uz_Latn_UZ", param1: "uz_Latn_UZ", want: true},
		{name: "Valid locale es_ES", param1: "es_ES", want: true},
		{name: "Valid locale sw_KE", param1: "sw_KE", want: true},
		// Invalid Locales
		{name: "Is invalid locale", param1: "en_XY", want: false},
		{name: "Is invalid locale", param1: "fr_QQ", want: false},
		{name: "Is invalid locale", param1: "fs_ZZ", want: false},
		{name: "Invalid locale empty string", param1: "", want: false},
		{name: "Invalid locale lo_POP", param1: "lo_POP", want: false},
		{name: "Invalid locale 12", param1: "12", want: false},
		{name: "Invalid locale 12_DD", param1: "12_DD", want: false},
		{name: "Invalid locale de-419-DE", param1: "de-419-DE", want: false},
		{name: "Invalid locale a-DE", param1: "a-DE", want: false},
		{name: "Invalid locale en only", param1: "en", want: false},
		{name: "Invalid locale gsw only", param1: "gsw", want: false},
		{name: "Invalid locale en-US", param1: "en-US", want: false},
		{name: "Invalid locale es-419", param1: "es-419", want: false},
		{name: "Invalid locale am_ET", param1: "am_ET", want: false},
		{name: "Invalid locale zh-CHS", param1: "zh-CHS", want: false},
		{name: "Invalid locale ca_ES_VALENCIA", param1: "ca_ES_VALENCIA", want: false},
		{name: "Invalid locale en_US_POSIX", param1: "en_US_POSIX", want: false},
		{name: "Invalid locale hak-CN", param1: "hak-CN", want: false},
		{name: "Invalid locale zh-Hant", param1: "zh-Hant", want: false},
		{name: "Invalid locale zh-Hans", param1: "zh-Hans", want: false},
		{name: "Invalid locale sr-Cyrl", param1: "sr-Cyrl", want: false},
		{name: "Invalid locale sr-Latn", param1: "sr-Latn", want: false},
		{name: "Invalid locale zh-cmn-Hans-CN", param1: "zh-cmn-Hans-CN", want: false},
		{name: "Invalid locale cmn-Hans-CN", param1: "cmn-Hans-CN", want: false},
		{name: "Invalid locale zh-yue-HK", param1: "zh-yue-HK", want: false},
		{name: "Invalid locale yue-HK", param1: "yue-HK", want: false},
		{name: "Invalid locale zh-Hans-CN", param1: "zh-Hans-CN", want: false},
		{name: "Invalid locale sr-Latn-RS", param1: "sr-Latn-RS", want: false},
		{name: "Invalid locale sl-rozaj", param1: "sl-rozaj", want: false},
		{name: "Invalid locale sl-rozaj-biske", param1: "sl-rozaj-biske", want: false},
		{name: "Invalid locale sl-nedis", param1: "sl-nedis", want: false},
		{name: "Invalid locale de-CH-1901", param1: "de-CH-1901", want: false},
		{name: "Invalid locale sl-IT-nedis", param1: "sl-IT-nedis", want: false},
		{name: "Invalid locale hy-Latn-IT-arevela", param1: "hy-Latn-IT-arevela", want: false},
		{name: "Invalid locale i-enochian", param1: "i-enochian", want: false},
		{name: "Invalid locale en-scotland-fonipa", param1: "en-scotland-fonipa", want: false},
		{name: "Invalid locale sl-IT-rozaj-biske-1994", param1: "sl-IT-rozaj-biske-1994", want: false},
		{name: "Invalid locale de-CH-x-phonebk", param1: "de-CH-x-phonebk", want: false},
		{name: "Invalid locale az-Arab-x-AZE-derbend", param1: "az-Arab-x-AZE-derbend", want: false},
		{name: "Invalid locale x-whatever", param1: "x-whatever", want: false},
		{name: "Invalid locale qaa-Qaaa-QM-x-southern", param1: "qaa-Qaaa-QM-x-southern", want: false},
		{name: "Invalid locale de-Qaaa", param1: "de-Qaaa", want: false},
		{name: "Invalid locale sr-Latn-QM", param1: "sr-Latn-QM", want: false},
		{name: "Invalid locale sr-Qaaa-RS", param1: "sr-Qaaa-RS", want: false},
		{name: "Invalid locale en-US-u-islamcal", param1: "en-US-u-islamcal", want: false},
		{name: "Invalid locale zh-CN-a-myext-x-private", param1: "zh-CN-a-myext-x-private", want: false},
		{name: "Invalid locale en-a-myext-b-another", param1: "en-a-myext-b-another", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsLocale(test.param1)

			assertValidation(t, result, test.want, err)
		})
	}
}
