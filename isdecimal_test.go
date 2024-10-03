package validatorgo

import "testing"

func TestIsDecimal(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsDecimalOpts
		want   bool
	}{
		// arabic_numerals_dot_decimal
		{name: "Valid_ar-IQ_ForceDecimal", param1: "١٢٣٤٫٥٦", param2: &IsDecimalOpts{Locale: "ar-IQ", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_ar-IQ_ForceDecimal", param1: "١٢٣٤", param2: &IsDecimalOpts{Locale: "ar-IQ", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_ar-IQ_Negative", param1: "-١٢٣٤٫٥٦", param2: &IsDecimalOpts{Locale: "ar-IQ"}, want: true},

		// comma_decimal_space_thousands
		{name: "Valid_bg-BG_ForceDecimal", param1: "1 234,56", param2: &IsDecimalOpts{Locale: "bg-BG", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_bg-BG_ForceDecimal", param1: "1 234", param2: &IsDecimalOpts{Locale: "bg-BG", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_bg-BG_Negative", param1: "-1 234,56", param2: &IsDecimalOpts{Locale: "bg-BG"}, want: true},

		// comma_decimal_dot_thousands
		{name: "Valid_de-DE_ForceDecimal", param1: "1.234,56", param2: &IsDecimalOpts{Locale: "de-DE", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_de-DE_ForceDecimal", param1: "1.234", param2: &IsDecimalOpts{Locale: "de-DE", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_de-DE_Negative", param1: "-1.234,56", param2: &IsDecimalOpts{Locale: "de-DE"}, want: true},

		// dot_decimal_comma_thousands (en-US default)
		{name: "Valid_en-US_ForceDecimal", param1: "1,234.56", param2: &IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_en-US_ForceDecimal", param1: "1,234", param2: &IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_en-US_Negative", param1: "-1,234.56", param2: nil, want: true}, // default locale is en-US

		// indian_numbering_system
		{name: "Valid_en-IN_ForceDecimal", param1: "1,23,456.78", param2: &IsDecimalOpts{Locale: "en-IN", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_en-IN_ForceDecimal", param1: "1,23,456", param2: &IsDecimalOpts{Locale: "en-IN", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_en-IN_Negative", param1: "-1,23,456.78", param2: &IsDecimalOpts{Locale: "en-IN"}, want: true},

		// dot_decimal_no_thousands
		{name: "Valid_id-ID_ForceDecimal", param1: "1234.56", param2: &IsDecimalOpts{Locale: "id-ID", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_id-ID_ForceDecimal", param1: "1234", param2: &IsDecimalOpts{Locale: "id-ID", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_id-ID_Negative", param1: "-1234.56", param2: &IsDecimalOpts{Locale: "id-ID"}, want: true},

		// arabic_numerals_comma_decimal
		{name: "Valid_fa-IR_ForceDecimal", param1: "۱۲۳۴٫۵۶", param2: &IsDecimalOpts{Locale: "fa-IR", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_fa-IR_ForceDecimal", param1: "۱۲۳۴,۵۶", param2: &IsDecimalOpts{Locale: "fa-IR", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},

		// comma_decimal_dot_thousands
		{name: "Valid_es-ES_ForceDecimal", param1: "1.234,56", param2: &IsDecimalOpts{Locale: "es-ES", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_es-ES_ForceDecimal", param1: "1.234", param2: &IsDecimalOpts{Locale: "es-ES", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_es-ES_Negative", param1: "-1.234,56", param2: &IsDecimalOpts{Locale: "es-ES"}, want: true},

		// dot_decimal_comma_thousands
		{name: "Valid_zh-CN_ForceDecimal", param1: "1,234.56", param2: &IsDecimalOpts{Locale: "zh-CN", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_zh-CN_ForceDecimal", param1: "1,234", param2: &IsDecimalOpts{Locale: "zh-CN", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_zh-CN_Negative", param1: "-1,234.56", param2: &IsDecimalOpts{Locale: "zh-CN"}, want: true},

		// comma_decimal_dot_thousands
		{name: "Valid_it-IT_ForceDecimal", param1: "1.234,56", param2: &IsDecimalOpts{Locale: "it-IT", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_it-IT_ForceDecimal", param1: "1.234", param2: &IsDecimalOpts{Locale: "it-IT", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_it-IT_Negative", param1: "-1.234,56", param2: &IsDecimalOpts{Locale: "it-IT"}, want: true},

		// dot_decimal_no_thousands
		{name: "Valid_th-TH_ForceDecimal", param1: "1234.56", param2: &IsDecimalOpts{Locale: "th-TH", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Invalid_th-TH_ForceDecimal", param1: "1234", param2: &IsDecimalOpts{Locale: "th-TH", ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Valid_th-TH_Negative", param1: "-1234.56", param2: &IsDecimalOpts{Locale: "th-TH"}, want: true},

		{name: "Invalid locale", param1: "-1234.56", param2: &IsDecimalOpts{Locale: "TX-nm"}, want: false},

		// Basic valid decimal numbers without forcing a decimal point
		{name: "Valid integer without forcing decimal", param1: "123", param2: &IsDecimalOpts{ForceDecimal: false}, want: true},
		{name: "Valid decimal number without forcing decimal", param1: "123.45", param2: &IsDecimalOpts{ForceDecimal: false}, want: true},

		// // Force decimal point presence
		{name: "Valid decimal number with forced decimal", param1: "123.45", param2: &IsDecimalOpts{ForceDecimal: true}, want: true},
		{name: "Integer but decimal point required", param1: "123", param2: &IsDecimalOpts{ForceDecimal: true}, want: false},

		// // Testing with different decimal digit ranges
		{name: "Decimal number with exact digits", param1: "123.45", param2: &IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: true},
		{name: "Decimal number with fewer digits", param1: "123.4", param2: &IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Decimal number with more digits", param1: "123.456", param2: &IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(2)}}, want: false},
		{name: "Decimal number with valid range of digits", param1: "123.456", param2: &IsDecimalOpts{ForceDecimal: true, DecimalDigits: DecimalDigits{Min: 2, Max: ptr(5)}}, want: true},

		// // Testing locale-specific decimal formats
		{name: "Valid decimal with French locale", param1: "123,45", param2: &IsDecimalOpts{ForceDecimal: true, Locale: "fr-FR"}, want: true},
		{name: "Invalid decimal for French locale", param1: "123.45", param2: &IsDecimalOpts{ForceDecimal: true, Locale: "fr-FR"}, want: false},
		{name: "Valid decimal with German locale", param1: "123,45", param2: &IsDecimalOpts{ForceDecimal: true, Locale: "de-DE"}, want: true},

		// // Invalid cases
		{name: "Non-numeric string", param1: "abc", param2: &IsDecimalOpts{ForceDecimal: false}, want: false},
		{name: "Invalid decimal number", param1: "123..45", param2: &IsDecimalOpts{ForceDecimal: true}, want: false},
		{name: "Negative decimal number", param1: "-123.45", param2: &IsDecimalOpts{ForceDecimal: true}, want: true},
		{name: "Positive number with locale", param1: "123.45", param2: &IsDecimalOpts{Locale: "en-US"}, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsDecimal(test.param1, test.param2)

			if result != test.want {
				t.Errorf("got `%t`, wanted `%t`", result, test.want)
			}
		})
	}
}
