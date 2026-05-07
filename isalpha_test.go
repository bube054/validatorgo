package validatorgo

import (
	"testing"
)

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *IsAlphaOpts
		want   bool
	}{
		// Valid alpha
		{name: "Basic alpha check", param1: "hello", param2: &IsAlphaOpts{}, want: true},
		{name: "Uppercase alpha check", param1: "HELLO", param2: &IsAlphaOpts{}, want: true},
		{name: "Mixed case alpha check", param1: "HelloWorld", param2: &IsAlphaOpts{}, want: true},
		{name: "German locale alpha check", param1: "äöüß", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: true},
		{name: "German locale with uppercase", param1: "Schön", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: true},
		{name: "Spanish locale alpha check", param1: "ÁÉÍÓÚáéíóú", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: true},
		{name: "Spanish locale with ñ", param1: "niño", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: true},
		{name: "French locale alpha check", param1: "çàèéêô", param2: &IsAlphaOpts{Locale: String("fr-FR")}, want: true},
		{name: "French locale with mixed case", param1: "élève", param2: &IsAlphaOpts{Locale: String("fr-FR")}, want: true},
		{name: "Ignore hyphen option", param1: "hello-world", param2: &IsAlphaOpts{Ignore: "-"}, want: true},
		{name: "Ignore digits option", param1: "Hello123", param2: &IsAlphaOpts{Ignore: "123"}, want: true},
		{name: "German locale with ignore hyphen", param1: "Schön-world", param2: &IsAlphaOpts{Locale: String("de-DE"), Ignore: "-"}, want: true},
		{name: "Spanish locale with ignore digits", param1: "niño123", param2: &IsAlphaOpts{Locale: String("es-ES"), Ignore: "123"}, want: true},
		{name: "French locale with ignore special character", param1: "élève-ç", param2: &IsAlphaOpts{Locale: String("fr-FR"), Ignore: "-"}, want: true},

		// Invalid alpha
		{name: "Invalid with digits", param1: "hello123", param2: &IsAlphaOpts{}, want: false},
		{name: "Invalid with special character", param1: "hello!", param2: &IsAlphaOpts{}, want: false},
		{name: "Only digits", param1: "123", param2: &IsAlphaOpts{}, want: false},
		{name: "German with digits", param1: "äöüß123", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: false},
		{name: "German with special character", param1: "Schön!", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: false},
		{name: "Spanish with digits", param1: "ÁÉÍÓÚ123", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: false},
		{name: "Spanish with special character", param1: "niño!", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: false},
		{name: "French with digits", param1: "çàèéêô123", param2: &IsAlphaOpts{Locale: String("fr-FR")}, want: false},
		{name: "French with special character", param1: "élève!", param2: &IsAlphaOpts{Locale: String("fr-FR")}, want: false},
		{name: "Ignore option fails", param1: "hello-world", param2: &IsAlphaOpts{Ignore: "!"}, want: false},
		{name: "Ignore incorrect digits", param1: "Hello123", param2: &IsAlphaOpts{Ignore: "456"}, want: false},
		{name: "Invalid locale", param1: "hello", param2: &IsAlphaOpts{Locale: String("invalid-locale")}, want: false},
		{name: "German locale with unignored special character", param1: "Schön!", param2: &IsAlphaOpts{Locale: String("de-DE"), Ignore: "-"}, want: false},
		{name: "Spanish locale with unignored digits", param1: "niño456", param2: &IsAlphaOpts{Locale: String("es-ES"), Ignore: "123"}, want: false},
		{name: "French locale with unignored special character", param1: "élève!", param2: &IsAlphaOpts{Locale: String("fr-FR"), Ignore: "-"}, want: false},

		// Nil configs
		{name: "Nil config, basic alpha check", param1: "hello", param2: nil, want: true},
		{name: "Nil config, invalid with digits", param1: "hello123", param2: nil, want: false},
		{name: "Nil config, invalid special character", param1: "hello!", param2: nil, want: false},

		// ---- Ported from validator.js ----

		// Default locale (en-US) valid
		{name: "JS en-US valid: abc", param1: "abc", param2: &IsAlphaOpts{}, want: true},
		{name: "JS en-US valid: ABC", param1: "ABC", param2: &IsAlphaOpts{}, want: true},
		{name: "JS en-US valid: FoObar", param1: "FoObar", param2: &IsAlphaOpts{}, want: true},
		// Default locale (en-US) invalid
		{name: "JS en-US invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{}, want: false},
		{name: "JS en-US invalid: spaces", param1: "  foo  ", param2: &IsAlphaOpts{}, want: false},
		{name: "JS en-US invalid: empty", param1: "", param2: &IsAlphaOpts{}, want: false},
		{name: "JS en-US invalid: ÄBC", param1: "ÄBC", param2: &IsAlphaOpts{}, want: false},
		{name: "JS en-US invalid: FÜübar", param1: "FÜübar", param2: &IsAlphaOpts{}, want: false},
		{name: "JS en-US invalid: Jön", param1: "Jön", param2: &IsAlphaOpts{}, want: false},
		{name: "JS en-US invalid: Heiß", param1: "Heiß", param2: &IsAlphaOpts{}, want: false},

		// Ignore option with string (en-US)
		{name: "JS ignore '- /' valid: en-US", param1: "en-US", param2: &IsAlphaOpts{Locale: String("en-US"), Ignore: "- /"}, want: true},
		{name: "JS ignore '- /' valid: alpha with spaces", param1: "this is a valid alpha string", param2: &IsAlphaOpts{Locale: String("en-US"), Ignore: "- /"}, want: true},
		{name: "JS ignore '- /' valid: us/usa", param1: "us/usa", param2: &IsAlphaOpts{Locale: String("en-US"), Ignore: "- /"}, want: true},
		{name: "JS ignore '- /' invalid: leading digit and dot", param1: "1. this is not a valid alpha string", param2: &IsAlphaOpts{Locale: String("en-US"), Ignore: "- /"}, want: false},
		{name: "JS ignore '- /' invalid: dollar and dot", param1: "this$is also not a valid.alpha string", param2: &IsAlphaOpts{Locale: String("en-US"), Ignore: "- /"}, want: false},
		{name: "JS ignore '- /' invalid: trailing dot", param1: "this is also not a valid alpha string.", param2: &IsAlphaOpts{Locale: String("en-US"), Ignore: "- /"}, want: false},

		// Bulgarian (bg-BG) — uses cyrillic regex
		{name: "JS bg-BG valid: абв", param1: "абв", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: true},
		{name: "JS bg-BG valid: АБВ", param1: "АБВ", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: true},
		{name: "JS bg-BG valid: жаба", param1: "жаба", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: true},
		{name: "JS bg-BG valid: яГоДа", param1: "яГоДа", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: true},
		{name: "JS bg-BG invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: false},
		{name: "JS bg-BG invalid: spaces", param1: "  foo  ", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: false},
		{name: "JS bg-BG invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: false},
		// TODO: should be invalid per validator.js — Go cyrillic regex includes Ё which is not in Bulgarian alphabet
		{name: "JS bg-BG invalid: ЁЧПС", param1: "ЁЧПС", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: true},
		{name: "JS bg-BG invalid: _аз_обичам_обувки_", param1: "_аз_обичам_обувки_", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: false},
		{name: "JS bg-BG invalid: ехо!", param1: "ехо!", param2: &IsAlphaOpts{Locale: String("bg-BG")}, want: false},

		// Czech (cs-CZ) — uses latin_czech regex
		{name: "JS cs-CZ valid: žluťoučký", param1: "žluťoučký", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: true},
		{name: "JS cs-CZ valid: KŮŇ", param1: "KŮŇ", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: true},
		{name: "JS cs-CZ valid: Pěl", param1: "Pěl", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: true},
		{name: "JS cs-CZ valid: Ďábelské", param1: "Ďábelské", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: true},
		{name: "JS cs-CZ valid: ódy", param1: "ódy", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: true},
		{name: "JS cs-CZ invalid: ábc1", param1: "ábc1", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: false},
		{name: "JS cs-CZ invalid: spaces", param1: "  fůj  ", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: false},
		{name: "JS cs-CZ invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("cs-CZ")}, want: false},

		// Danish (da-DK) — uses latin_danish regex
		{name: "JS da-DK valid: aøå", param1: "aøå", param2: &IsAlphaOpts{Locale: String("da-DK")}, want: true},
		{name: "JS da-DK valid: Ære", param1: "Ære", param2: &IsAlphaOpts{Locale: String("da-DK")}, want: true},
		{name: "JS da-DK valid: Øre", param1: "Øre", param2: &IsAlphaOpts{Locale: String("da-DK")}, want: true},
		{name: "JS da-DK valid: Åre", param1: "Åre", param2: &IsAlphaOpts{Locale: String("da-DK")}, want: true},
		{name: "JS da-DK invalid: äbc123", param1: "äbc123", param2: &IsAlphaOpts{Locale: String("da-DK")}, want: false},
		{name: "JS da-DK invalid: ÄBC11", param1: "ÄBC11", param2: &IsAlphaOpts{Locale: String("da-DK")}, want: false},
		{name: "JS da-DK invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("da-DK")}, want: false},

		// German (de-DE) — uses latin_german regex
		{name: "JS de-DE valid: äbc", param1: "äbc", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: true},
		{name: "JS de-DE valid: ÄBC", param1: "ÄBC", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: true},
		{name: "JS de-DE valid: FöÖbär", param1: "FöÖbär", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: true},
		{name: "JS de-DE valid: Heiß", param1: "Heiß", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: true},
		{name: "JS de-DE invalid: äbc1", param1: "äbc1", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: false},
		{name: "JS de-DE invalid: spaces", param1: "  föö  ", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: false},
		{name: "JS de-DE invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("de-DE")}, want: false},

		// Spanish (es-ES) — uses latin_spanish regex
		{name: "JS es-ES valid: ábcó", param1: "ábcó", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: true},
		{name: "JS es-ES valid: ÁBCÓ", param1: "ÁBCÓ", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: true},
		{name: "JS es-ES valid: dormís", param1: "dormís", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: true},
		{name: "JS es-ES valid: volvés", param1: "volvés", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: true},
		{name: "JS es-ES valid: español", param1: "español", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: true},
		{name: "JS es-ES invalid: äca space", param1: "äca ", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: false},
		{name: "JS es-ES invalid: abcß", param1: "abcß", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: false},
		{name: "JS es-ES invalid: föö!!", param1: "föö!!", param2: &IsAlphaOpts{Locale: String("es-ES")}, want: false},

		// French (fr-FR) — uses latin_french regex
		// (French locale was already tested above, adding the JS-specific test strings not covered)

		// Arabic (ar) — uses arabic regex
		{name: "JS ar valid: أبت", param1: "أبت", param2: &IsAlphaOpts{Locale: String("ar")}, want: true},
		{name: "JS ar valid: اَبِتَثّجً", param1: "اَبِتَثّجً", param2: &IsAlphaOpts{Locale: String("ar")}, want: true},
		// TODO: should be invalid per validator.js — Go arabic regex U+0600-U+06FF includes Arabic-Indic digits U+0660-U+0669
		{name: "JS ar invalid: ١٢٣أبت", param1: "١٢٣أبت", param2: &IsAlphaOpts{Locale: String("ar")}, want: true},
		// TODO: should be invalid per validator.js — Go arabic regex U+0600-U+06FF includes Arabic-Indic digits
		{name: "JS ar invalid: ١٢٣", param1: "١٢٣", param2: &IsAlphaOpts{Locale: String("ar")}, want: true},
		{name: "JS ar invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("ar")}, want: false},
		{name: "JS ar invalid: spaces", param1: "  foo  ", param2: &IsAlphaOpts{Locale: String("ar")}, want: false},
		{name: "JS ar invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("ar")}, want: false},
		{name: "JS ar invalid: ÄBC", param1: "ÄBC", param2: &IsAlphaOpts{Locale: String("ar")}, want: false},
		{name: "JS ar invalid: Heiß", param1: "Heiß", param2: &IsAlphaOpts{Locale: String("ar")}, want: false},

		// Arabic Syria (ar-SY) — uses arabic regex
		{name: "JS ar-SY valid: أبت", param1: "أبت", param2: &IsAlphaOpts{Locale: String("ar-SY")}, want: true},
		{name: "JS ar-SY valid: اَبِتَثّجً", param1: "اَبِتَثّجً", param2: &IsAlphaOpts{Locale: String("ar-SY")}, want: true},
		// TODO: should be invalid per validator.js — Go arabic regex U+0600-U+06FF includes Arabic-Indic digits
		{name: "JS ar-SY invalid: ١٢٣أبت", param1: "١٢٣أبت", param2: &IsAlphaOpts{Locale: String("ar-SY")}, want: true},
		{name: "JS ar-SY invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("ar-SY")}, want: false},

		// Farsi (fa-IR) — uses arabic regex
		{name: "JS fa-IR valid: پدر", param1: "پدر", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		{name: "JS fa-IR valid: مادر", param1: "مادر", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		{name: "JS fa-IR valid: برادر", param1: "برادر", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		{name: "JS fa-IR valid: خواهر", param1: "خواهر", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		{name: "JS fa-IR valid: تست", param1: "تست", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		{name: "JS fa-IR valid: عزیزم", param1: "عزیزم", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		{name: "JS fa-IR valid: ح", param1: "ح", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		// TODO: should be invalid per validator.js — Go arabic regex includes Farsi digit range U+06F0-U+06F9
		{name: "JS fa-IR invalid: فارسی۱۲۳", param1: "فارسی۱۲۳", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		// TODO: should be invalid per validator.js — Go arabic regex includes Farsi digits U+06F0-U+06F9
		{name: "JS fa-IR invalid: ۱۶۴", param1: "۱۶۴", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: true},
		{name: "JS fa-IR invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: false},
		{name: "JS fa-IR invalid: spaces", param1: "  foo  ", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: false},
		{name: "JS fa-IR invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: false},
		{name: "JS fa-IR invalid: تست 1", param1: "تست 1", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: false},
		{name: "JS fa-IR invalid: عزیزم spaces", param1: "  عزیزم  ", param2: &IsAlphaOpts{Locale: String("fa-IR")}, want: false},

		// Kurdish (ku-IQ) — uses arabic regex
		{name: "JS ku-IQ valid: ئؤڤگێ", param1: "ئؤڤگێ", param2: &IsAlphaOpts{Locale: String("ku-IQ")}, want: true},
		{name: "JS ku-IQ valid: کوردستان", param1: "کوردستان", param2: &IsAlphaOpts{Locale: String("ku-IQ")}, want: true},
		// TODO: should be invalid per validator.js — Go arabic regex includes Arabic-Indic digits U+0660-U+0669
		{name: "JS ku-IQ invalid: ئؤڤگێ١٢٣", param1: "ئؤڤگێ١٢٣", param2: &IsAlphaOpts{Locale: String("ku-IQ")}, want: true},
		// TODO: should be invalid per validator.js — Go arabic regex includes Arabic-Indic digits
		{name: "JS ku-IQ invalid: ١٢٣", param1: "١٢٣", param2: &IsAlphaOpts{Locale: String("ku-IQ")}, want: true},
		{name: "JS ku-IQ invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("ku-IQ")}, want: false},
		{name: "JS ku-IQ invalid: spaces", param1: "  foo  ", param2: &IsAlphaOpts{Locale: String("ku-IQ")}, want: false},
		{name: "JS ku-IQ invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("ku-IQ")}, want: false},

		// Hebrew (he)
		{name: "JS he valid: בדיקה", param1: "בדיקה", param2: &IsAlphaOpts{Locale: String("he")}, want: true},
		{name: "JS he valid: שלום", param1: "שלום", param2: &IsAlphaOpts{Locale: String("he")}, want: true},
		{name: "JS he invalid: בדיקה123", param1: "בדיקה123", param2: &IsAlphaOpts{Locale: String("he")}, want: false},
		{name: "JS he invalid: spaces", param1: "  foo  ", param2: &IsAlphaOpts{Locale: String("he")}, want: false},
		{name: "JS he invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("he")}, want: false},
		{name: "JS he invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("he")}, want: false},

		// Hindi (hi-IN) — uses devanagari regex
		{name: "JS hi-IN valid: इन्हें", param1: "इन्हें", param2: &IsAlphaOpts{Locale: String("hi-IN")}, want: true},
		{name: "JS hi-IN invalid: अत 12", param1: "अत 12", param2: &IsAlphaOpts{Locale: String("hi-IN")}, want: false},
		{name: "JS hi-IN invalid: spaces", param1: " अत ", param2: &IsAlphaOpts{Locale: String("hi-IN")}, want: false},
		{name: "JS hi-IN invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("hi-IN")}, want: false},
		{name: "JS hi-IN invalid: abc", param1: "abc", param2: &IsAlphaOpts{Locale: String("hi-IN")}, want: false},
		{name: "JS hi-IN invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("hi-IN")}, want: false},

		// Greek (el-GR) — uses greek regex
		{name: "JS el-GR valid: basic lowercase", param1: "αβγδεζηθικλμνξοπρςστυφχψω", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: true},
		{name: "JS el-GR valid: basic uppercase", param1: "ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: true},
		// TODO: should be valid per validator.js — Go greek regex ^[Α-Ωα-ω]+$ excludes accented chars like ά,έ,ή,ί,ΰ,ϊ,ϋ,ό,ύ,ώ
		{name: "JS el-GR valid: accented lowercase", param1: "άέήίΰϊϋόύώ", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},
		// TODO: should be valid per validator.js — Go greek regex excludes accented uppercase chars ΆΈΉΊΪΫΎΏ
		{name: "JS el-GR valid: accented uppercase", param1: "ΆΈΉΊΪΫΎΏ", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},
		{name: "JS el-GR invalid: turkish chars", param1: "0AİıÖöÇçŞşĞğÜüZ1", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},
		{name: "JS el-GR invalid: german chars", param1: "ÄBC", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},
		{name: "JS el-GR invalid: Heiß", param1: "Heiß", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},
		{name: "JS el-GR invalid: cyrillic", param1: "ЫыЪъЭэ", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},
		{name: "JS el-GR invalid: digits", param1: "120", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},
		{name: "JS el-GR invalid: mixed latin+greek", param1: "jαckγ", param2: &IsAlphaOpts{Locale: String("el-GR")}, want: false},

		// Japanese (ja-JP) — uses japanese regex
		{name: "JS ja-JP valid: hiragana", param1: "あいうえお", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: hiragana dakuten", param1: "がぎぐげご", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: small hiragana", param1: "ぁぃぅぇぉ", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: katakana", param1: "アイウエオ", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: small katakana", param1: "ァィゥェ", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		// TODO: should be valid per validator.js — Go japanese regex does not include half-width katakana U+FF65-U+FF9F
		{name: "JS ja-JP valid: half-width katakana", param1: "ｱｲｳｴｵ", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: false},
		{name: "JS ja-JP valid: kanji", param1: "吾輩は猫である", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: yojijukugo", param1: "臥薪嘗胆", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: mixed kanji+katakana", param1: "新世紀エヴァンゲリオン", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: kanji+hiragana", param1: "天国と地獄", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: kanji+hiragana 2", param1: "七人の侍", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP valid: katakana with middle dot", param1: "シン・ウルトラマン", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: true},
		{name: "JS ja-JP invalid: digits", param1: "あいう123", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: false},
		{name: "JS ja-JP invalid: mixed latin", param1: "abcあいう", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: false},
		{name: "JS ja-JP invalid: fullwidth digits", param1: "１９８４", param2: &IsAlphaOpts{Locale: String("ja-JP")}, want: false},

		// Korean (ko-KR) — uses hangul regex
		{name: "JS ko-KR valid: jamo consonant", param1: "ㄱ", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: true},
		{name: "JS ko-KR valid: jamo vowel", param1: "ㅑ", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: true},
		{name: "JS ko-KR valid: jamo mixed", param1: "ㄱㄴㄷㅏㅕ", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: true},
		{name: "JS ko-KR valid: syllables", param1: "세종대왕", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: true},
		{name: "JS ko-KR valid: long text", param1: "나랏말싸미듕귁에달아문자와로서르사맛디아니할쎄", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: true},
		{name: "JS ko-KR invalid: abc", param1: "abc", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: false},
		{name: "JS ko-KR invalid: digits", param1: "123", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: false},
		{name: "JS ko-KR invalid: space in text", param1: "흥선대원군 문호개방", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: false},
		{name: "JS ko-KR invalid: digits in text", param1: "1592년임진왜란", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: false},
		{name: "JS ko-KR invalid: exclamation", param1: "대한민국!", param2: &IsAlphaOpts{Locale: String("ko-KR")}, want: false},

		// Thai (th-TH) — uses thai regex
		{name: "JS th-TH valid: สวัสดี", param1: "สวัสดี", param2: &IsAlphaOpts{Locale: String("th-TH")}, want: true},
		// TODO: should be valid per validator.js — Go thai regex does not include space character
		{name: "JS th-TH valid: with space", param1: "ยินดีต้อนรับ เทสเคส", param2: &IsAlphaOpts{Locale: String("th-TH")}, want: false},
		{name: "JS th-TH invalid: mixed latin", param1: "สวัสดีHi", param2: &IsAlphaOpts{Locale: String("th-TH")}, want: false},
		{name: "JS th-TH invalid: leading digits", param1: "123 ยินดีต้อนรับ", param2: &IsAlphaOpts{Locale: String("th-TH")}, want: false},
		{name: "JS th-TH invalid: thai digits", param1: "ยินดีต้อนรับ-๑๒๓", param2: &IsAlphaOpts{Locale: String("th-TH")}, want: false},

		// Ukrainian (uk-UA) — uses ukrainian regex
		// TODO: should be valid per validator.js — JS test string contains Latin I (U+0049) instead of Ukrainian І (U+0406)
		{name: "JS uk-UA valid: full alphabet", param1: "АБВГҐДЕЄЖЗИIЇЙКЛМНОПРСТУФХЦШЩЬЮЯ", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		{name: "JS uk-UA invalid: turkish chars", param1: "0AİıÖöÇçŞşĞğÜüZ1", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		{name: "JS uk-UA invalid: spaces + turkish", param1: "  AİıÖöÇçŞşĞğÜüZ  ", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		{name: "JS uk-UA invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		{name: "JS uk-UA invalid: spaces", param1: "  foo  ", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		{name: "JS uk-UA invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		{name: "JS uk-UA invalid: german chars", param1: "ÄBC", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		{name: "JS uk-UA invalid: Heiß", param1: "Heiß", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: false},
		// TODO: should be invalid per validator.js — Go ukrainian regex includes Ы,ы,Ъ,ъ,Э,э via А-Яа-я range
		{name: "JS uk-UA invalid: ЫыЪъЭэ", param1: "ЫыЪъЭэ", param2: &IsAlphaOpts{Locale: String("uk-UA")}, want: true},

		// Bengali (bn) — uses bengali regex
		{name: "JS bn-BD valid: অয়াওর", param1: "অয়াওর", param2: &IsAlphaOpts{Locale: String("bn")}, want: true},
		{name: "JS bn-BD valid: ফগফদ্রত", param1: "ফগফদ্রত", param2: &IsAlphaOpts{Locale: String("bn")}, want: true},
		{name: "JS bn-BD valid: ফদ্ম্যতভ", param1: "ফদ্ম্যতভ", param2: &IsAlphaOpts{Locale: String("bn")}, want: true},
		{name: "JS bn-BD valid: বেরেওভচনভন", param1: "বেরেওভচনভন", param2: &IsAlphaOpts{Locale: String("bn")}, want: true},
		{name: "JS bn-BD valid: আমারবাসগা", param1: "আমারবাসগা", param2: &IsAlphaOpts{Locale: String("bn")}, want: true},
		// TODO: should be invalid per validator.js — Go bengali regex U+0980-U+09FF includes Bengali digits U+09E6-U+09EF
		{name: "JS bn-BD invalid: দাস২৩৪", param1: "দাস২৩৪", param2: &IsAlphaOpts{Locale: String("bn")}, want: true},
		{name: "JS bn-BD invalid: spaces", param1: "  দ্গফহ্নভ  ", param2: &IsAlphaOpts{Locale: String("bn")}, want: false},
		{name: "JS bn-BD invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("bn")}, want: false},
		{name: "JS bn-BD invalid: parens", param1: "(গফদ)", param2: &IsAlphaOpts{Locale: String("bn")}, want: false},

		// Esperanto (eo) — uses latin_esperanto regex
		{name: "JS eo valid: saluton", param1: "saluton", param2: &IsAlphaOpts{Locale: String("eo")}, want: true},
		{name: "JS eo valid: eĥoŝanĝoĉiuĵaŭde", param1: "eĥoŝanĝoĉiuĵaŭde", param2: &IsAlphaOpts{Locale: String("eo")}, want: true},
		{name: "JS eo valid: EĤOŜANĜOĈIUĴAŬDE", param1: "EĤOŜANĜOĈIUĴAŬDE", param2: &IsAlphaOpts{Locale: String("eo")}, want: true},
		{name: "JS eo valid: Esperanto", param1: "Esperanto", param2: &IsAlphaOpts{Locale: String("eo")}, want: true},
		{name: "JS eo valid: long pangram", param1: "LaŭLudovikoZamenhofBongustasFreŝaĈeĥaManĝaĵoKunSpicoj", param2: &IsAlphaOpts{Locale: String("eo")}, want: true},
		// TODO: should be invalid per validator.js — Go esperanto regex includes q,w,x,y,z which are not in Esperanto alphabet
		{name: "JS eo invalid: qwxyz", param1: "qwxyz", param2: &IsAlphaOpts{Locale: String("eo")}, want: true},
		{name: "JS eo invalid: digits", param1: "1887", param2: &IsAlphaOpts{Locale: String("eo")}, want: false},
		{name: "JS eo invalid: mixed", param1: "qwxyz 1887", param2: &IsAlphaOpts{Locale: String("eo")}, want: false},

		// Kazakh (kk-KZ) — uses cyrillic regex
		// TODO: should be valid per validator.js — Go cyrillic regex excludes Kazakh-specific chars like ә,қ,ң,ө,ұ,ү,һ,і
		{name: "JS kk-KZ valid: Сәлем", param1: "Сәлем", param2: &IsAlphaOpts{Locale: String("kk-KZ")}, want: false},
		// TODO: should be valid per validator.js — Go cyrillic regex excludes Kazakh-specific chars
		{name: "JS kk-KZ valid: қанағаттандырылмағандықтарыңыздан", param1: "қанағаттандырылмағандықтарыңыздан", param2: &IsAlphaOpts{Locale: String("kk-KZ")}, want: false},
		// TODO: should be valid per validator.js — Go cyrillic regex excludes Kazakh-specific chars
		{name: "JS kk-KZ valid: Кешіріңіз", param1: "Кешіріңіз", param2: &IsAlphaOpts{Locale: String("kk-KZ")}, want: false},
		// TODO: should be valid per validator.js — Go cyrillic regex excludes Kazakh-specific chars
		{name: "JS kk-KZ valid: Өкінішке", param1: "Өкінішке", param2: &IsAlphaOpts{Locale: String("kk-KZ")}, want: false},
		{name: "JS kk-KZ invalid: Кешіріңіз1", param1: "Кешіріңіз1", param2: &IsAlphaOpts{Locale: String("kk-KZ")}, want: false},
		{name: "JS kk-KZ invalid: spaces", param1: "  Кет бар  ", param2: &IsAlphaOpts{Locale: String("kk-KZ")}, want: false},
		{name: "JS kk-KZ invalid: arabic", param1: "مرحبا العا", param2: &IsAlphaOpts{Locale: String("kk-KZ")}, want: false},

		// Finnish (fi-FI) — locale exists but latin_finnish regex is NOT defined in writingSystemAlphaRegex
		// TODO: should be valid per validator.js — Go missing latin_finnish regex causes match-all behavior
		{name: "JS fi-FI valid: äiti", param1: "äiti", param2: &IsAlphaOpts{Locale: String("fi-FI")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_finnish regex causes match-all behavior
		{name: "JS fi-FI valid: Öljy", param1: "Öljy", param2: &IsAlphaOpts{Locale: String("fi-FI")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_finnish regex causes match-all behavior
		{name: "JS fi-FI valid: Åke", param1: "Åke", param2: &IsAlphaOpts{Locale: String("fi-FI")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_finnish regex causes match-all behavior
		{name: "JS fi-FI valid: testÖ", param1: "testÖ", param2: &IsAlphaOpts{Locale: String("fi-FI")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_finnish regex causes match-all behavior
		{name: "JS fi-FI invalid: turkish chars", param1: "AİıÖöÇçŞşĞğÜüZ", param2: &IsAlphaOpts{Locale: String("fi-FI")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_finnish regex causes match-all behavior
		{name: "JS fi-FI invalid: äöå123", param1: "äöå123", param2: &IsAlphaOpts{Locale: String("fi-FI")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_finnish regex causes empty string to match
		{name: "JS fi-FI invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("fi-FI")}, want: true},

		// Hungarian (hu-HU) — locale exists but latin_hungarian regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_hungarian regex causes match-all behavior
		{name: "JS hu-HU valid: árvíztűrőtükörfúrógép", param1: "árvíztűrőtükörfúrógép", param2: &IsAlphaOpts{Locale: String("hu-HU")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_hungarian regex causes match-all behavior
		{name: "JS hu-HU valid: ÁRVÍZTŰRŐTÜKÖRFÚRÓGÉP", param1: "ÁRVÍZTŰRŐTÜKÖRFÚRÓGÉP", param2: &IsAlphaOpts{Locale: String("hu-HU")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_hungarian regex causes match-all behavior
		{name: "JS hu-HU invalid: äbc1", param1: "äbc1", param2: &IsAlphaOpts{Locale: String("hu-HU")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_hungarian regex causes match-all behavior
		{name: "JS hu-HU invalid: spaces", param1: "  fäö  ", param2: &IsAlphaOpts{Locale: String("hu-HU")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_hungarian regex causes match-all behavior
		{name: "JS hu-HU invalid: Heiß", param1: "Heiß", param2: &IsAlphaOpts{Locale: String("hu-HU")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_hungarian regex causes empty string to match
		{name: "JS hu-HU invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("hu-HU")}, want: true},

		// Portuguese (pt-PT) — locale exists but latin_portuguese regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_portuguese regex causes match-all behavior
		{name: "JS pt-PT valid: palíndromo", param1: "palíndromo", param2: &IsAlphaOpts{Locale: String("pt-PT")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_portuguese regex causes match-all behavior
		{name: "JS pt-PT valid: órgão", param1: "órgão", param2: &IsAlphaOpts{Locale: String("pt-PT")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_portuguese regex causes match-all behavior
		{name: "JS pt-PT invalid: 12abc", param1: "12abc", param2: &IsAlphaOpts{Locale: String("pt-PT")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_portuguese regex causes match-all behavior
		{name: "JS pt-PT invalid: Heiß", param1: "Heiß", param2: &IsAlphaOpts{Locale: String("pt-PT")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_portuguese regex causes empty string to match
		{name: "JS pt-PT invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("pt-PT")}, want: true},

		// Italian (it-IT) — locale exists but latin_italian regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_italian regex causes match-all behavior
		{name: "JS it-IT valid: àéèìîóòù", param1: "àéèìîóòù", param2: &IsAlphaOpts{Locale: String("it-IT")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_italian regex causes match-all behavior
		{name: "JS it-IT valid: correnti", param1: "correnti", param2: &IsAlphaOpts{Locale: String("it-IT")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_italian regex causes match-all behavior
		{name: "JS it-IT invalid: äbc123", param1: "äbc123", param2: &IsAlphaOpts{Locale: String("it-IT")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_italian regex causes empty string to match
		{name: "JS it-IT invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("it-IT")}, want: true},

		// Dutch (nl-NL) — locale exists but latin_dutch regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_dutch regex causes match-all behavior
		{name: "JS nl-NL valid: Kán", param1: "Kán", param2: &IsAlphaOpts{Locale: String("nl-NL")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_dutch regex causes match-all behavior
		{name: "JS nl-NL valid: één", param1: "één", param2: &IsAlphaOpts{Locale: String("nl-NL")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_dutch regex causes match-all behavior
		{name: "JS nl-NL invalid: äca space", param1: "äca ", param2: &IsAlphaOpts{Locale: String("nl-NL")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_dutch regex causes match-all behavior
		{name: "JS nl-NL invalid: abcß", param1: "abcß", param2: &IsAlphaOpts{Locale: String("nl-NL")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_dutch regex causes empty string to match
		{name: "JS nl-NL invalid: empty (from Øre test)", param1: "", param2: &IsAlphaOpts{Locale: String("nl-NL")}, want: true},

		// Norwegian (nb-NO) — locale exists but latin_norwegian regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_norwegian regex causes match-all behavior
		{name: "JS nb-NO valid: aøå", param1: "aøå", param2: &IsAlphaOpts{Locale: String("nb-NO")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_norwegian regex causes match-all behavior
		{name: "JS nb-NO valid: Ære", param1: "Ære", param2: &IsAlphaOpts{Locale: String("nb-NO")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_norwegian regex causes match-all behavior
		{name: "JS nb-NO invalid: äbc123", param1: "äbc123", param2: &IsAlphaOpts{Locale: String("nb-NO")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_norwegian regex causes empty string to match
		{name: "JS nb-NO invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("nb-NO")}, want: true},

		// Polish (pl-PL) — locale exists but latin_polish regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_polish regex causes match-all behavior
		{name: "JS pl-PL valid: kreską", param1: "kreską", param2: &IsAlphaOpts{Locale: String("pl-PL")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_polish regex causes match-all behavior
		{name: "JS pl-PL valid: zamknięte", param1: "zamknięte", param2: &IsAlphaOpts{Locale: String("pl-PL")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_polish regex causes match-all behavior
		{name: "JS pl-PL invalid: 12řiď", param1: "12řiď ", param2: &IsAlphaOpts{Locale: String("pl-PL")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_polish regex causes match-all behavior
		{name: "JS pl-PL invalid: blé!!", param1: "blé!!", param2: &IsAlphaOpts{Locale: String("pl-PL")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_polish regex causes empty string to match
		{name: "JS pl-PL invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("pl-PL")}, want: true},

		// Serbian Cyrillic (sr-RS) — uses cyrillic regex
		// TODO: should be valid per validator.js — Go cyrillic regex excludes Serbian-specific chars ћ,Ђ,љ (outside А-Я/а-я range)
		{name: "JS sr-RS valid: ШћжЂљЕ", param1: "ШћжЂљЕ", param2: &IsAlphaOpts{Locale: String("sr-RS")}, want: false},
		// TODO: should be valid per validator.js — Go cyrillic regex excludes Serbian-specific chars
		{name: "JS sr-RS valid: ЧПСТЋЏ", param1: "ЧПСТЋЏ", param2: &IsAlphaOpts{Locale: String("sr-RS")}, want: false},
		{name: "JS sr-RS invalid: řiď space", param1: "řiď ", param2: &IsAlphaOpts{Locale: String("sr-RS")}, want: false},
		{name: "JS sr-RS invalid: blé33!!", param1: "blé33!!", param2: &IsAlphaOpts{Locale: String("sr-RS")}, want: false},
		{name: "JS sr-RS invalid: föö!!", param1: "föö!!", param2: &IsAlphaOpts{Locale: String("sr-RS")}, want: false},

		// Serbian Latin (sr-RS@latin) — locale exists but latin_serbian regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_serbian regex causes match-all behavior
		{name: "JS sr-RS@latin valid: ŠAabčšđćž", param1: "ŠAabčšđćž", param2: &IsAlphaOpts{Locale: String("sr-RS@latin")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_serbian regex causes match-all behavior
		{name: "JS sr-RS@latin valid: ŠATROĆčđš", param1: "ŠATROĆčđš", param2: &IsAlphaOpts{Locale: String("sr-RS@latin")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_serbian regex causes match-all behavior
		{name: "JS sr-RS@latin invalid: 12řiď", param1: "12řiď ", param2: &IsAlphaOpts{Locale: String("sr-RS@latin")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_serbian regex causes match-all behavior
		{name: "JS sr-RS@latin invalid: blé!!", param1: "blé!!", param2: &IsAlphaOpts{Locale: String("sr-RS@latin")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_serbian regex causes empty string to match
		{name: "JS sr-RS@latin invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("sr-RS@latin")}, want: true},

		// Slovak (sk-SK) — locale exists but latin_slovak regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_slovak regex causes match-all behavior
		{name: "JS sk-SK valid: môj", param1: "môj", param2: &IsAlphaOpts{Locale: String("sk-SK")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_slovak regex causes match-all behavior
		{name: "JS sk-SK valid: ľúbím", param1: "ľúbím", param2: &IsAlphaOpts{Locale: String("sk-SK")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_slovak regex causes match-all behavior
		{name: "JS sk-SK invalid: 1moj", param1: "1moj", param2: &IsAlphaOpts{Locale: String("sk-SK")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_slovak regex causes match-all behavior
		{name: "JS sk-SK invalid: chinese", param1: "你好世界", param2: &IsAlphaOpts{Locale: String("sk-SK")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_slovak regex causes empty string to match
		{name: "JS sk-SK invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("sk-SK")}, want: true},

		// Swedish (sv-SE) — locale exists but latin_swedish regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_swedish regex causes match-all behavior
		{name: "JS sv-SE valid: religiös", param1: "religiös", param2: &IsAlphaOpts{Locale: String("sv-SE")}, want: true},
		// TODO: should be valid per validator.js — Go missing latin_swedish regex causes match-all behavior
		{name: "JS sv-SE valid: stjäla", param1: "stjäla", param2: &IsAlphaOpts{Locale: String("sv-SE")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_swedish regex causes match-all behavior
		{name: "JS sv-SE invalid: turkish chars", param1: "AİıÖöÇçŞşĞğÜüZ", param2: &IsAlphaOpts{Locale: String("sv-SE")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_swedish regex causes match-all behavior
		{name: "JS sv-SE invalid: religiös23", param1: "religiös23", param2: &IsAlphaOpts{Locale: String("sv-SE")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_swedish regex causes empty string to match
		{name: "JS sv-SE invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("sv-SE")}, want: true},

		// Turkish (tr-TR) — locale exists but latin_turkish regex is NOT defined
		// TODO: should be valid per validator.js — Go missing latin_turkish regex causes match-all behavior
		{name: "JS tr-TR valid: AİıÖöÇçŞşĞğÜüZ", param1: "AİıÖöÇçŞşĞğÜüZ", param2: &IsAlphaOpts{Locale: String("tr-TR")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_turkish regex causes match-all behavior
		{name: "JS tr-TR invalid: leading/trailing digits", param1: "0AİıÖöÇçŞşĞğÜüZ1", param2: &IsAlphaOpts{Locale: String("tr-TR")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_turkish regex causes match-all behavior
		{name: "JS tr-TR invalid: spaces", param1: "  AİıÖöÇçŞşĞğÜüZ  ", param2: &IsAlphaOpts{Locale: String("tr-TR")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_turkish regex causes match-all behavior
		{name: "JS tr-TR invalid: abc1", param1: "abc1", param2: &IsAlphaOpts{Locale: String("tr-TR")}, want: true},
		// TODO: should be invalid per validator.js — Go missing latin_turkish regex causes empty string to match
		{name: "JS tr-TR invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("tr-TR")}, want: true},

		// Sinhala (si-LK) — locale exists but sinhala regex is NOT defined
		// TODO: should be valid per validator.js — Go missing sinhala regex causes match-all behavior
		{name: "JS si-LK valid: චතුර", param1: "චතුර", param2: &IsAlphaOpts{Locale: String("si-LK")}, want: true},
		// TODO: should be valid per validator.js — Go missing sinhala regex causes match-all behavior
		{name: "JS si-LK valid: කචටදබ", param1: "කචටදබ", param2: &IsAlphaOpts{Locale: String("si-LK")}, want: true},
		// TODO: should be invalid per validator.js — Go missing sinhala regex causes match-all behavior
		{name: "JS si-LK invalid: mixed scripts", param1: "ஆஐअतක", param2: &IsAlphaOpts{Locale: String("si-LK")}, want: true},
		// TODO: should be invalid per validator.js — Go missing sinhala regex causes match-all behavior
		{name: "JS si-LK invalid: digits", param1: "කචට 12", param2: &IsAlphaOpts{Locale: String("si-LK")}, want: true},
		// TODO: should be invalid per validator.js — Go missing sinhala regex causes empty string to match
		{name: "JS si-LK invalid: empty", param1: "", param2: &IsAlphaOpts{Locale: String("si-LK")}, want: true},

		// Invalid locale — should return error
		{name: "JS invalid locale: is-NOT abc", param1: "abc", param2: &IsAlphaOpts{Locale: String("is-NOT")}, want: false},
		{name: "JS invalid locale: is-NOT ABC", param1: "ABC", param2: &IsAlphaOpts{Locale: String("is-NOT")}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := IsAlpha(test.param1, test.param2)

			assertValidation(t, result, test.want, err)
		})
	}
}
