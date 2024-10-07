package validatorgo

import "testing"

func TestIsMobilePhone(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 []string
		param3 *IsMobilePhoneOpts
		want   bool
	}{
		{name: "Valid nigerian number default config", param1: "08070448986", param2: []string{"en-NG"}, param3: nil, want: true},
		{name: "Valid nigerian number default config, at least one incorrect locale", param1: "08070448986", param2: []string{"xy-AB", "en-NG"}, param3: nil, want: true},
		{name: "Valid nigerian number nil locale", param1: "08070448986", param2: nil, param3: nil, want: true},

		// Armenia
		{name: "Valid Armenian number", param1: "077123456", param2: []string{"am-AM"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Armenian number", param1: "0771234567", param2: []string{"am-AM"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Armenian number, strict", param1: "+37477123456", param2: []string{"am-AM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Armenian number, strict", param1: "+3747712345", param2: []string{"am-AM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// United Arab Emirates
		{name: "Valid UAE number", param1: "0501234567", param2: []string{"ar-AE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid UAE number", param1: "05112345678", param2: []string{"ar-AE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid UAE number, strict", param1: "+971501234567", param2: []string{"ar-AE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid UAE number, strict", param1: "+9715012347568", param2: []string{"ar-AE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bahrain
		{name: "Valid Bahrain number", param1: "36001234", param2: []string{"ar-BH"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bahrain number", param1: "360012345", param2: []string{"ar-BH"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bahrain number, strict", param1: "+97336001234", param2: []string{"ar-BH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bahrain number, strict", param1: "+9733600123", param2: []string{"ar-BH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Algeria
		{name: "Valid Algerian number", param1: "0550123456", param2: []string{"ar-DZ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Algerian number", param1: "05601234567", param2: []string{"ar-DZ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Algerian number, strict", param1: "+213550123456", param2: []string{"ar-DZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Algerian number, strict", param1: "+21355012345", param2: []string{"ar-DZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Egypt
		{name: "Valid Egyptian number", param1: "01234567890", param2: []string{"ar-EG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Egyptian number", param1: "012345678901", param2: []string{"ar-EG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Egyptian number, strict", param1: "+201234567890", param2: []string{"ar-EG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Egyptian number, strict", param1: "+20123456789", param2: []string{"ar-EG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Western Sahara
		{name: "Valid Western Saharan number", param1: "600123456", param2: []string{"ar-EH"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Western Saharan number", param1: "06001234567", param2: []string{"ar-EH"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Western Saharan number, strict", param1: "+212600123456", param2: []string{"ar-EH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Western Saharan number, strict", param1: "+21260012345", param2: []string{"ar-EH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Iraq
		{name: "Valid Iraqi number", param1: "07901234567", param2: []string{"ar-IQ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Iraqi number", param1: "079012345678", param2: []string{"ar-IQ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Iraqi number, strict", param1: "+9647901234567", param2: []string{"ar-IQ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Iraqi number, strict", param1: "+964790123456", param2: []string{"ar-IQ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Jordan
		{name: "Valid Jordanian number", param1: "0790123456", param2: []string{"ar-JO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Jordanian number", param1: "07901234567", param2: []string{"ar-JO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Jordanian number, strict", param1: "+962790123456", param2: []string{"ar-JO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Jordanian number, strict", param1: "+9627901234589", param2: []string{"ar-JO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Kuwait
		{name: "Valid Kuwaiti number", param1: "50012345", param2: []string{"ar-KW"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Kuwaiti number", param1: "500123456", param2: []string{"ar-KW"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Kuwaiti number, strict", param1: "+96550012345", param2: []string{"ar-KW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Kuwaiti number, strict", param1: "+9655001234", param2: []string{"ar-KW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Palestine
		{name: "Valid Palestinian number", param1: "0591234567", param2: []string{"ar-PS"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Palestinian number", param1: "05912345678", param2: []string{"ar-PS"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Palestinian number, strict", param1: "+970591234567", param2: []string{"ar-PS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Palestinian number, strict", param1: "+97059123456", param2: []string{"ar-PS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Saudi Arabia
		{name: "Valid Saudi number", param1: "0551234567", param2: []string{"ar-SA"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Saudi number", param1: "05612345678", param2: []string{"ar-SA"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Saudi number, strict", param1: "+966551234567", param2: []string{"ar-SA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Saudi number, strict", param1: "+96655123456", param2: []string{"ar-SA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Sudan
		{name: "Valid Sudanese number", param1: "0912345678", param2: []string{"ar-SD"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Sudanese number", param1: "09123456789", param2: []string{"ar-SD"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Sudanese number, strict", param1: "+249912345678", param2: []string{"ar-SD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Sudanese number, strict", param1: "+24991234567", param2: []string{"ar-SD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Syria
		{name: "Valid Syrian number", param1: "0931234567", param2: []string{"ar-SY"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Syrian number", param1: "09412345678", param2: []string{"ar-SY"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Syrian number, strict", param1: "+963931234567", param2: []string{"ar-SY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Syrian number, strict", param1: "+96393123456", param2: []string{"ar-SY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Tunisia
		{name: "Valid Tunisian number", param1: "20123456", param2: []string{"ar-TN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Tunisian number", param1: "201234567", param2: []string{"ar-TN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Tunisian number, strict", param1: "+21620123456", param2: []string{"ar-TN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Tunisian number, strict", param1: "+2162012345", param2: []string{"ar-TN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Yemen
		{name: "Valid Yemeni number", param1: "0712345678", param2: []string{"ar-YE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Yemeni number", param1: "07123456789", param2: []string{"ar-YE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Yemeni number, strict", param1: "+967712345678", param2: []string{"ar-YE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Yemeni number, strict", param1: "+96771234567", param2: []string{"ar-YE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Azerbaijan
		{name: "Valid Azerbaijani number", param1: "0501234567", param2: []string{"az-AZ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Azerbaijani number", param1: "05012345678", param2: []string{"az-AZ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Azerbaijani number, strict", param1: "+994501234567", param2: []string{"az-AZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Azerbaijani number, strict", param1: "+99450123456", param2: []string{"az-AZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Belarus
		{name: "Valid Belarusian number", param1: "0251234567", param2: []string{"be-BY"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Belarusian number", param1: "02512345678", param2: []string{"be-BY"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Belarusian number, strict", param1: "+375251234567", param2: []string{"be-BY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Belarusian number, strict", param1: "+37525123456", param2: []string{"be-BY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bulgaria
		{name: "Valid Bulgarian number", param1: "0871234567", param2: []string{"bg-BG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bulgarian number", param1: "08712345678", param2: []string{"bg-BG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bulgarian number, strict", param1: "+359871234567", param2: []string{"bg-BG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bulgarian number, strict", param1: "+35987123456", param2: []string{"bg-BG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bangladesh
		{name: "Valid Bangladeshi number", param1: "01712345678", param2: []string{"bn-BD"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bangladeshi number", param1: "017123456789", param2: []string{"bn-BD"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bangladeshi number, strict", param1: "+8801712345678", param2: []string{"bn-BD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bangladeshi number, strict", param1: "+880171234567", param2: []string{"bn-BD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bosnia
		{name: "Valid Bosnian number", param1: "061123456", param2: []string{"bs-BA"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bosnian number", param1: "0611234567", param2: []string{"bs-BA"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bosnian number, strict", param1: "+38761123456", param2: []string{"bs-BA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bosnian number, strict", param1: "+3876112345", param2: []string{"bs-BA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Andorra
		{name: "Valid Andorran number", param1: "312345", param2: []string{"ca-AD"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Andorran number", param1: "3123456", param2: []string{"ca-AD"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Andorran number, strict", param1: "+376312345", param2: []string{"ca-AD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Andorran number, strict", param1: "+37631234", param2: []string{"ca-AD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Czech Republic
		{name: "Valid Czech number", param1: "601123456", param2: []string{"cs-CZ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Czech number", param1: "6011234567", param2: []string{"cs-CZ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Czech number, strict", param1: "+420601123456", param2: []string{"cs-CZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Czech number, strict", param1: "+42060112345", param2: []string{"cs-CZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Denmark (da-DK)
		{name: "Valid Danish number", param1: "20123456", param2: []string{"da-DK"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Danish number", param1: "201234567", param2: []string{"da-DK"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Danish number, strict", param1: "+4520123456", param2: []string{"da-DK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Danish number, strict", param1: "+452012345", param2: []string{"da-DK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Austria (de-AT)
		{name: "Valid Austrian number", param1: "0660123456", param2: []string{"de-AT"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Austrian number", param1: "06601234567", param2: []string{"de-AT"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Austrian number, strict", param1: "+43660123456", param2: []string{"de-AT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Austrian number, strict", param1: "+4366012345", param2: []string{"de-AT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Switzerland (de-CH)
		{name: "Valid Swiss number", param1: "0791234567", param2: []string{"de-CH"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Swiss number", param1: "07912345678", param2: []string{"de-CH"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Swiss number, strict", param1: "+41791234567", param2: []string{"de-CH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Swiss number, strict", param1: "+4179123456", param2: []string{"de-CH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Germany (de-DE)
		{name: "Valid German number", param1: "015112345678", param2: []string{"de-DE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid German number", param1: "0151123456789", param2: []string{"de-DE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid German number, strict", param1: "+4915112345678", param2: []string{"de-DE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid German number, strict", param1: "+491511234567", param2: []string{"de-DE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Luxembourg (de-LU)
		{name: "Valid Luxembourgish number", param1: "621123456", param2: []string{"de-LU"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Luxembourgish number", param1: "6211234567", param2: []string{"de-LU"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Luxembourgish number, strict", param1: "+352621123456", param2: []string{"de-LU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Luxembourgish number, strict", param1: "+35262112345", param2: []string{"de-LU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Maldives (dv-MV)
		{name: "Valid Maldivian number", param1: "7712345", param2: []string{"dv-MV"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Maldivian number", param1: "77123456", param2: []string{"dv-MV"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Maldivian number, strict", param1: "+9607712345", param2: []string{"dv-MV"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Maldivian number, strict", param1: "+960771234", param2: []string{"dv-MV"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bhutan (dz-BT)
		{name: "Valid Bhutanese number", param1: "17123456", param2: []string{"dz-BT"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bhutanese number", param1: "171234567", param2: []string{"dz-BT"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bhutanese number, strict", param1: "+97517123456", param2: []string{"dz-BT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bhutanese number, strict", param1: "+9751712345", param2: []string{"dz-BT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Cyprus (el-CY)
		{name: "Valid Cypriot number", param1: "96123456", param2: []string{"el-CY"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Cypriot number", param1: "961234567", param2: []string{"el-CY"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Cypriot number, strict", param1: "+35796123456", param2: []string{"el-CY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Cypriot number, strict", param1: "+3579612345", param2: []string{"el-CY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Greece (el-GR)
		{name: "Valid Greek number", param1: "6912345678", param2: []string{"el-GR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Greek number", param1: "69123456789", param2: []string{"el-GR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Greek number, strict", param1: "+306912345678", param2: []string{"el-GR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Greek number, strict", param1: "+30691234567", param2: []string{"el-GR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Antigua and Barbuda (en-AG)
		{name: "Valid Antigua and Barbuda number", param1: "2681234567", param2: []string{"en-AG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Antigua and Barbuda number", param1: "26812345678", param2: []string{"en-AG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Antigua and Barbuda number, strict", param1: "+12681234567", param2: []string{"en-AG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Antigua and Barbuda number, strict", param1: "+1268123456", param2: []string{"en-AG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Anguilla (en-AI)
		{name: "Valid Anguilla number", param1: "2641234567", param2: []string{"en-AI"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Anguilla number", param1: "26412345678", param2: []string{"en-AI"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Anguilla number, strict", param1: "+12641234567", param2: []string{"en-AI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Anguilla number, strict", param1: "+1264123456", param2: []string{"en-AI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Australia (en-AU)
		{name: "Valid Australian number", param1: "0412345678", param2: []string{"en-AU"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Australian number", param1: "04123456789", param2: []string{"en-AU"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Australian number, strict", param1: "+61412345678", param2: []string{"en-AU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Australian number, strict", param1: "+6141234567", param2: []string{"en-AU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bermuda (en-BM)
		{name: "Valid Bermudian number", param1: "4411234567", param2: []string{"en-BM"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bermudian number", param1: "44112345678", param2: []string{"en-BM"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bermudian number, strict", param1: "+14411234567", param2: []string{"en-BM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bermudian number, strict", param1: "+1441123456", param2: []string{"en-BM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bahamas (en-BS)
		{name: "Valid Bahamian number", param1: "2421234567", param2: []string{"en-BS"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bahamian number", param1: "24212345678", param2: []string{"en-BS"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bahamian number, strict", param1: "+12421234567", param2: []string{"en-BS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bahamian number, strict", param1: "+1242123456", param2: []string{"en-BS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Botswana (en-BW)
		{name: "Valid Botswanan number", param1: "71123456", param2: []string{"en-BW"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Botswanan number", param1: "711234567", param2: []string{"en-BW"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Botswanan number, strict", param1: "+26771123456", param2: []string{"en-BW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Botswanan number, strict", param1: "+2677112345", param2: []string{"en-BW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Canada (en-CA)
		{name: "Valid Canadian number", param1: "4161234567", param2: []string{"en-CA"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Canadian number", param1: "41612345678", param2: []string{"en-CA"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Canadian number, strict", param1: "+14161234567", param2: []string{"en-CA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Canadian number, strict", param1: "+1416123456", param2: []string{"en-CA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// United Kingdom (en-GB)
		{name: "Valid UK number", param1: "07123456789", param2: []string{"en-GB"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid UK number", param1: "071234567890", param2: []string{"en-GB"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid UK number, strict", param1: "+447123456789", param2: []string{"en-GB"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid UK number, strict", param1: "+44712345678", param2: []string{"en-GB"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Guernsey (en-GG)
		{name: "Valid Guernsey number", param1: "07781123456", param2: []string{"en-GG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Guernsey number", param1: "077811234567", param2: []string{"en-GG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Guernsey number, strict", param1: "+447781123456", param2: []string{"en-GG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Guernsey number, strict", param1: "+44778112345", param2: []string{"en-GG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Ghana (en-GH)
		{name: "Valid Ghanaian number", param1: "0241234567", param2: []string{"en-GH"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Ghanaian number", param1: "02412345678", param2: []string{"en-GH"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Ghanaian number, strict", param1: "+233241234567", param2: []string{"en-GH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Ghanaian number, strict", param1: "+23324123456", param2: []string{"en-GH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Guyana (en-GY)
		{name: "Valid Guyanese number", param1: "6012345", param2: []string{"en-GY"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Guyanese number", param1: "60123456", param2: []string{"en-GY"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Guyanese number, strict", param1: "+5926012345", param2: []string{"en-GY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Guyanese number, strict", param1: "+592601234", param2: []string{"en-GY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Hong Kong (en-HK)
		{name: "Valid Hong Kong number", param1: "51234567", param2: []string{"en-HK"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Hong Kong number", param1: "512345678", param2: []string{"en-HK"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Hong Kong number, strict", param1: "+85251234567", param2: []string{"en-HK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Hong Kong number, strict", param1: "+8525123456", param2: []string{"en-HK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Ireland (en-IE)
		{name: "Valid Irish number", param1: "0851234567", param2: []string{"en-IE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Irish number", param1: "08512345678", param2: []string{"en-IE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Irish number, strict", param1: "+353851234567", param2: []string{"en-IE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Irish number, strict", param1: "+35385123456", param2: []string{"en-IE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// India (en-IN)
		{name: "Valid Indian number", param1: "9123456789", param2: []string{"en-IN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Indian number", param1: "91234567890", param2: []string{"en-IN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Indian number, strict", param1: "+919123456789", param2: []string{"en-IN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Indian number, strict", param1: "+91912345678", param2: []string{"en-IN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Jamaica (en-JM)
		{name: "Valid Jamaican number", param1: "8761234567", param2: []string{"en-JM"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Jamaican number", param1: "87612345678", param2: []string{"en-JM"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Jamaican number, strict", param1: "+18761234567", param2: []string{"en-JM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Jamaican number, strict", param1: "+1876123456", param2: []string{"en-JM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Kenya (en-KE)
		{name: "Valid Kenyan number", param1: "0712345678", param2: []string{"en-KE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Kenyan number", param1: "07123456789", param2: []string{"en-KE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Kenyan number, strict", param1: "+254712345678", param2: []string{"en-KE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Kenyan number, strict", param1: "+25471234567", param2: []string{"en-KE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Kiribati (en-KI)
		{name: "Valid Kiribati number", param1: "72012345", param2: []string{"en-KI"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Kiribati number", param1: "720123456", param2: []string{"en-KI"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Kiribati number, strict", param1: "+68672012345", param2: []string{"en-KI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Kiribati number, strict", param1: "+6867201234", param2: []string{"en-KI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Saint Kitts and Nevis (en-KN)
		{name: "Valid Saint Kitts and Nevis number", param1: "8691234567", param2: []string{"en-KN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Saint Kitts and Nevis number", param1: "86912345678", param2: []string{"en-KN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Saint Kitts and Nevis number, strict", param1: "+18691234567", param2: []string{"en-KN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Saint Kitts and Nevis number, strict", param1: "+1869123456", param2: []string{"en-KN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Lesotho (en-LS)
		{name: "Valid Lesotho number", param1: "50123456", param2: []string{"en-LS"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Lesotho number", param1: "501234567", param2: []string{"en-LS"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Lesotho number, strict", param1: "+26650123456", param2: []string{"en-LS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Lesotho number, strict", param1: "+2665012345", param2: []string{"en-LS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Macao (en-MO)
		{name: "Valid Macao number", param1: "66123456", param2: []string{"en-MO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Macao number", param1: "661234567", param2: []string{"en-MO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Macao number, strict", param1: "+85366123456", param2: []string{"en-MO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Macao number, strict", param1: "+8536612345", param2: []string{"en-MO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Malta (en-MT)
		{name: "Valid Maltese number", param1: "77123456", param2: []string{"en-MT"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Maltese number", param1: "771234567", param2: []string{"en-MT"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Maltese number, strict", param1: "+35677123456", param2: []string{"en-MT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Maltese number, strict", param1: "+3567712345", param2: []string{"en-MT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Mauritius (en-MU)
		{name: "Valid Mauritian number", param1: "52501234", param2: []string{"en-MU"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Mauritian number", param1: "525012345", param2: []string{"en-MU"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Mauritian number, strict", param1: "+23052501234", param2: []string{"en-MU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Mauritian number, strict", param1: "+2305250123", param2: []string{"en-MU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Malawi (en-MW)
		{name: "Valid Malawian number", param1: "0991234567", param2: []string{"en-MW"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Malawian number", param1: "09912345678", param2: []string{"en-MW"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Malawian number, strict", param1: "+265991234567", param2: []string{"en-MW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Malawian number, strict", param1: "+26599123456", param2: []string{"en-MW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Nigeria (en-NG)
		{name: "Valid Nigerian number", param1: "08070448986", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Nigerian number", param1: "090666666567", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Nigerian number, strict", param1: "+2348120989668", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Nigerian number, strict", param1: "+23481209895", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// New Zealand (en-NZ)
		{name: "Valid New Zealand number", param1: "0212345678", param2: []string{"en-NZ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid New Zealand number", param1: "02123456789", param2: []string{"en-NZ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid New Zealand number, strict", param1: "+64212345678", param2: []string{"en-NZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid New Zealand number, strict", param1: "+6421234567", param2: []string{"en-NZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Papua New Guinea (en-PG)
		{name: "Valid Papua New Guinea number", param1: "71234567", param2: []string{"en-PG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Papua New Guinea number", param1: "712345678", param2: []string{"en-PG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Papua New Guinea number, strict", param1: "+67571234567", param2: []string{"en-PG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Papua New Guinea number, strict", param1: "+6757123456", param2: []string{"en-PG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Philippines (en-PH)
		{name: "Valid Philippine number", param1: "09171234567", param2: []string{"en-PH"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Philippine number", param1: "091712345678", param2: []string{"en-PH"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Philippine number, strict", param1: "+639171234567", param2: []string{"en-PH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Philippine number, strict", param1: "+63917123456", param2: []string{"en-PH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Pakistan (en-PK)
		{name: "Valid Pakistani number", param1: "03012345678", param2: []string{"en-PK"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Pakistani number", param1: "030123456789", param2: []string{"en-PK"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Pakistani number, strict", param1: "+923012345678", param2: []string{"en-PK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Pakistani number, strict", param1: "+92301234567", param2: []string{"en-PK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Rwanda (en-RW)
		{name: "Valid Rwandan number", param1: "0781234567", param2: []string{"en-RW"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Rwandan number", param1: "07812345678", param2: []string{"en-RW"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Rwandan number, strict", param1: "+250781234567", param2: []string{"en-RW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Rwandan number, strict", param1: "+25078123456", param2: []string{"en-RW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Singapore (en-SG)
		{name: "Valid Singaporean number", param1: "91234567", param2: []string{"en-SG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Singaporean number", param1: "912345678", param2: []string{"en-SG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Singaporean number, strict", param1: "+6591234567", param2: []string{"en-SG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Singaporean number, strict", param1: "+659123456", param2: []string{"en-SG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Sierra Leone (en-SL)
		{name: "Valid Sierra Leone number", param1: "0761234567", param2: []string{"en-SL"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Sierra Leone number", param1: "07612345678", param2: []string{"en-SL"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Sierra Leone number, strict", param1: "+232761234567", param2: []string{"en-SL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Sierra Leone number, strict", param1: "+23276123456", param2: []string{"en-SL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// South Sudan (en-SS)
		{name: "Valid South Sudan number", param1: "0921234567", param2: []string{"en-SS"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid South Sudan number", param1: "09212345678", param2: []string{"en-SS"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid South Sudan number, strict", param1: "+211921234567", param2: []string{"en-SS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid South Sudan number, strict", param1: "+21192123456", param2: []string{"en-SS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Tanzania (en-TZ)
		{name: "Valid Tanzanian number", param1: "0751234567", param2: []string{"en-TZ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Tanzanian number", param1: "07512345678", param2: []string{"en-TZ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Tanzanian number, strict", param1: "+255751234567", param2: []string{"en-TZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Tanzanian number, strict", param1: "+25575123456", param2: []string{"en-TZ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Uganda (en-UG)
		{name: "Valid Ugandan number", param1: "0701234567", param2: []string{"en-UG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Ugandan number", param1: "07012345678", param2: []string{"en-UG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Ugandan number, strict", param1: "+256701234567", param2: []string{"en-UG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Ugandan number, strict", param1: "+25670123456", param2: []string{"en-UG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// United States (en-US)
		{name: "Valid US number", param1: "2021234567", param2: []string{"en-US"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid US number", param1: "20212345678", param2: []string{"en-US"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid US number, strict", param1: "+12021234567", param2: []string{"en-US"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid US number, strict", param1: "+1202123456", param2: []string{"en-US"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// South Africa (en-ZA)
		{name: "Valid South African number", param1: "0831234567", param2: []string{"en-ZA"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid South African number", param1: "08312345678", param2: []string{"en-ZA"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid South African number, strict", param1: "+27831234567", param2: []string{"en-ZA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid South African number, strict", param1: "+2783123456", param2: []string{"en-ZA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Zambia (en-ZM)
		{name: "Valid Zambian number", param1: "0971234567", param2: []string{"en-ZM"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Zambian number", param1: "09712345678", param2: []string{"en-ZM"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Zambian number, strict", param1: "+260971234567", param2: []string{"en-ZM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Zambian number, strict", param1: "+26097123456", param2: []string{"en-ZM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Zimbabwe (en-ZW)
		{name: "Valid Zimbabwean number", param1: "0771234567", param2: []string{"en-ZW"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Zimbabwean number", param1: "07712345678", param2: []string{"en-ZW"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Zimbabwean number, strict", param1: "+263771234567", param2: []string{"en-ZW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Zimbabwean number, strict", param1: "+26377123456", param2: []string{"en-ZW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Spanish (es-AR)
		{name: "Valid Argentine number", param1: "01112345678", param2: []string{"es-AR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Argentine number", param1: "011123456789", param2: []string{"es-AR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Argentine number, strict", param1: "+5491112345678", param2: []string{"es-AR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Argentine number, strict", param1: "+549111234567", param2: []string{"es-AR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Bolivia (es-BO)
		{name: "Valid Bolivian number", param1: "71234567", param2: []string{"es-BO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Bolivian number", param1: "712345678", param2: []string{"es-BO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Bolivian number, strict", param1: "+59171234567", param2: []string{"es-BO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Bolivian number, strict", param1: "+5917123456", param2: []string{"es-BO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Chile (es-CL)
		{name: "Valid Chilean number", param1: "091234567", param2: []string{"es-CL"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Chilean number", param1: "0912345678", param2: []string{"es-CL"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Chilean number, strict", param1: "+56912345678", param2: []string{"es-CL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Chilean number, strict", param1: "+5691234567", param2: []string{"es-CL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Colombia (es-CO)
		{name: "Valid Colombian number", param1: "3001234567", param2: []string{"es-CO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Colombian number", param1: "30012345678", param2: []string{"es-CO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Colombian number, strict", param1: "+573001234567", param2: []string{"es-CO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Colombian number, strict", param1: "+57300123456", param2: []string{"es-CO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Costa Rica (es-CR)
		{name: "Valid Costa Rican number", param1: "60123456", param2: []string{"es-CR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Costa Rican number", param1: "601234567", param2: []string{"es-CR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Costa Rican number, strict", param1: "+50660123456", param2: []string{"es-CR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Costa Rican number, strict", param1: "+5066012345", param2: []string{"es-CR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Cuba (es-CU)
		{name: "Valid Cuban number", param1: "51234567", param2: []string{"es-CU"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Cuban number", param1: "512345678", param2: []string{"es-CU"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Cuban number, strict", param1: "+5351234567", param2: []string{"es-CU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Cuban number, strict", param1: "+535123456", param2: []string{"es-CU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Dominican Republic (es-DO)
		{name: "Valid Dominican number", param1: "8091234567", param2: []string{"es-DO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Dominican number", param1: "80912345678", param2: []string{"es-DO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Dominican number, strict", param1: "+18091234567", param2: []string{"es-DO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Dominican number, strict", param1: "+1809123456", param2: []string{"es-DO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Ecuador (es-EC)
		{name: "Valid Ecuadorian number", param1: "0991234567", param2: []string{"es-EC"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Ecuadorian number", param1: "09912345678", param2: []string{"es-EC"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Ecuadorian number, strict", param1: "+593991234567", param2: []string{"es-EC"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Ecuadorian number, strict", param1: "+59399123456", param2: []string{"es-EC"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Spain (es-ES)
		{name: "Valid Spanish number", param1: "612345678", param2: []string{"es-ES"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Spanish number", param1: "6123456789", param2: []string{"es-ES"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Spanish number, strict", param1: "+34612345678", param2: []string{"es-ES"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Spanish number, strict", param1: "+3461234567", param2: []string{"es-ES"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Guatemala (es-GT)
		{name: "Valid Guatemalan number", param1: "51234567", param2: []string{"es-GT"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Guatemalan number", param1: "512345678", param2: []string{"es-GT"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Guatemalan number, strict", param1: "+50251234567", param2: []string{"es-GT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Guatemalan number, strict", param1: "+5025123456", param2: []string{"es-GT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Honduras (es-HN)
		{name: "Valid Honduran number", param1: "31234567", param2: []string{"es-HN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Honduran number", param1: "312345678", param2: []string{"es-HN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Honduran number, strict", param1: "+50431234567", param2: []string{"es-HN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Honduran number, strict", param1: "+5043123456", param2: []string{"es-HN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Mexico (es-MX)
		{name: "Valid Mexican number", param1: "5512345678", param2: []string{"es-MX"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Mexican number", param1: "55123456789", param2: []string{"es-MX"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Mexican number, strict", param1: "+525512345678", param2: []string{"es-MX"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Mexican number, strict", param1: "+52551234567", param2: []string{"es-MX"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Nicaragua (es-NI)
		{name: "Valid Nicaraguan number", param1: "51234567", param2: []string{"es-NI"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Nicaraguan number", param1: "512345678", param2: []string{"es-NI"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Nicaraguan number, strict", param1: "+50551234567", param2: []string{"es-NI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Nicaraguan number, strict", param1: "+5055123456", param2: []string{"es-NI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Panama (es-PA)
		{name: "Valid Panamanian number", param1: "61234567", param2: []string{"es-PA"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Panamanian number", param1: "612345678", param2: []string{"es-PA"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Panamanian number, strict", param1: "+50761234567", param2: []string{"es-PA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Panamanian number, strict", param1: "+5076123456", param2: []string{"es-PA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Peru (es-PE)
		{name: "Valid Peruvian number", param1: "991234567", param2: []string{"es-PE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Peruvian number", param1: "9912345678", param2: []string{"es-PE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Peruvian number, strict", param1: "+51991234567", param2: []string{"es-PE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Peruvian number, strict", param1: "+5199123456", param2: []string{"es-PE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Paraguay (es-PY)
		{name: "Valid Paraguayan number", param1: "0981234567", param2: []string{"es-PY"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Paraguayan number", param1: "09812345678", param2: []string{"es-PY"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Paraguayan number, strict", param1: "+595981234567", param2: []string{"es-PY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Paraguayan number, strict", param1: "+59598123456", param2: []string{"es-PY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// El Salvador (es-SV)
		{name: "Valid Salvadoran number", param1: "61234567", param2: []string{"es-SV"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Salvadoran number", param1: "612345678", param2: []string{"es-SV"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Salvadoran number, strict", param1: "+50361234567", param2: []string{"es-SV"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Salvadoran number, strict", param1: "+5036123456", param2: []string{"es-SV"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Uruguay (es-UY)
		{name: "Valid Uruguayan number", param1: "91234567", param2: []string{"es-UY"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Uruguayan number", param1: "912345678", param2: []string{"es-UY"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Uruguayan number, strict", param1: "+59891234567", param2: []string{"es-UY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Uruguayan number, strict", param1: "+5989123456", param2: []string{"es-UY"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Venezuela (es-VE)
		{name: "Valid Venezuelan number", param1: "04121234567", param2: []string{"es-VE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Venezuelan number", param1: "041212345678", param2: []string{"es-VE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Venezuelan number, strict", param1: "+584121234567", param2: []string{"es-VE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Venezuelan number, strict", param1: "+58412123456", param2: []string{"es-VE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Estonia (et-EE)
		{name: "Valid Estonian number", param1: "51234567", param2: []string{"et-EE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Estonian number", param1: "512345678", param2: []string{"et-EE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Estonian number, strict", param1: "+37251234567", param2: []string{"et-EE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Estonian number, strict", param1: "+3725123456", param2: []string{"et-EE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Afghanistan (fa-AF)
		{name: "Valid Afghan number", param1: "0791234567", param2: []string{"fa-AF"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Afghan number", param1: "07912345678", param2: []string{"fa-AF"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Afghan number, strict", param1: "+93791234567", param2: []string{"fa-AF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Afghan number, strict", param1: "+9379123456", param2: []string{"fa-AF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Iran (fa-IR)
		{name: "Valid Iranian number", param1: "09121234567", param2: []string{"fa-IR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Iranian number", param1: "091212345678", param2: []string{"fa-IR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Iranian number, strict", param1: "+989121234567", param2: []string{"fa-IR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Iranian number, strict", param1: "+98912123456", param2: []string{"fa-IR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Finland (fi-FI)
		{name: "Valid Finnish number", param1: "0412345678", param2: []string{"fi-FI"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Finnish number", param1: "04123456789", param2: []string{"fi-FI"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Finnish number, strict", param1: "+358412345678", param2: []string{"fi-FI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Finnish number, strict", param1: "+35841234567", param2: []string{"fi-FI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Fiji (fj-FJ)
		{name: "Valid Fijian number", param1: "9123456", param2: []string{"fj-FJ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Fijian number", param1: "91234567", param2: []string{"fj-FJ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Fijian number, strict", param1: "+6799123456", param2: []string{"fj-FJ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Fijian number, strict", param1: "+679912345", param2: []string{"fj-FJ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Faroe Islands (fo-FO)
		{name: "Valid Faroese number", param1: "123456", param2: []string{"fo-FO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Faroese number", param1: "1234567", param2: []string{"fo-FO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Faroese number, strict", param1: "+298123456", param2: []string{"fo-FO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Faroese number, strict", param1: "+29812345", param2: []string{"fo-FO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Belgium (fr-BE)
		{name: "Valid Belgian number", param1: "0471234567", param2: []string{"fr-BE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Belgian number", param1: "04712345678", param2: []string{"fr-BE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Belgian number, strict", param1: "+32471234567", param2: []string{"fr-BE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Belgian number, strict", param1: "+3247123456", param2: []string{"fr-BE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Burkina Faso (fr-BF)
		{name: "Valid Burkinabe number", param1: "0701234567", param2: []string{"fr-BF"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Burkinabe number", param1: "07012345678", param2: []string{"fr-BF"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Burkinabe number, strict", param1: "+226701234567", param2: []string{"fr-BF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Burkinabe number, strict", param1: "+22670123456", param2: []string{"fr-BF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Benin (fr-BJ)
		{name: "Valid Beninese number", param1: "61234567", param2: []string{"fr-BJ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Beninese number", param1: "612345678", param2: []string{"fr-BJ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Beninese number, strict", param1: "+22961234567", param2: []string{"fr-BJ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Beninese number, strict", param1: "+2296123456", param2: []string{"fr-BJ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Congo, Democratic Republic of the (fr-CD)
		{name: "Valid Congolese number", param1: "0812345678", param2: []string{"fr-CD"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Congolese number", param1: "08123456789", param2: []string{"fr-CD"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Congolese number, strict", param1: "+243812345678", param2: []string{"fr-CD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Congolese number, strict", param1: "+24381234567", param2: []string{"fr-CD"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Central African Republic (fr-CF)
		{name: "Valid Central African number", param1: "0701234567", param2: []string{"fr-CF"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Central African number", param1: "07012345678", param2: []string{"fr-CF"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Central African number, strict", param1: "+236701234567", param2: []string{"fr-CF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Central African number, strict", param1: "+23670123456", param2: []string{"fr-CF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// France (fr-FR)
		{name: "Valid French number", param1: "0612345678", param2: []string{"fr-FR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid French number", param1: "06123456789", param2: []string{"fr-FR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid French number, strict", param1: "+33612345678", param2: []string{"fr-FR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid French number, strict", param1: "+3361234567", param2: []string{"fr-FR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// French Guiana (fr-GF)
		{name: "Valid Guianese number", param1: "0691234567", param2: []string{"fr-GF"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Guianese number", param1: "06912345678", param2: []string{"fr-GF"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Guianese number, strict", param1: "+594691234567", param2: []string{"fr-GF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Guianese number, strict", param1: "+59469123456", param2: []string{"fr-GF"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Guadeloupe (fr-GP)
		{name: "Valid Guadeloupean number", param1: "0691234567", param2: []string{"fr-GP"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Guadeloupean number", param1: "06912345678", param2: []string{"fr-GP"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Guadeloupean number, strict", param1: "+590691234567", param2: []string{"fr-GP"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Guadeloupean number, strict", param1: "+59069123456", param2: []string{"fr-GP"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Martinique (fr-MQ)
		{name: "Valid Martiniquais number", param1: "0691234567", param2: []string{"fr-MQ"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Martiniquais number", param1: "06912345678", param2: []string{"fr-MQ"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Martiniquais number, strict", param1: "+596691234567", param2: []string{"fr-MQ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Martiniquais number, strict", param1: "+59669123456", param2: []string{"fr-MQ"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Réunion (fr-RE)
		{name: "Valid Réunionese number", param1: "0691234567", param2: []string{"fr-RE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Réunionese number", param1: "06912345678", param2: []string{"fr-RE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Réunionese number, strict", param1: "+262691234567", param2: []string{"fr-RE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Réunionese number, strict", param1: "+26269123456", param2: []string{"fr-RE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Switzerland (fr-CH)
		{name: "Valid Swiss number", param1: "0761234567", param2: []string{"fr-CH"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Swiss number", param1: "07612345678", param2: []string{"fr-CH"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Swiss number, strict", param1: "+41761234567", param2: []string{"fr-CH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Swiss number, strict", param1: "+4176123456", param2: []string{"fr-CH"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Guinea (fr-GN)
		{name: "Valid Guinean number", param1: "620123456", param2: []string{"fr-GN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Guinean number", param1: "6201234567", param2: []string{"fr-GN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Guinean number, strict", param1: "+224620123456", param2: []string{"fr-GN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Guinean number, strict", param1: "+22462012345", param2: []string{"fr-GN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Haiti (fr-HT)
		{name: "Valid Haitian number", param1: "34123456", param2: []string{"fr-HT"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Haitian number", param1: "341234567", param2: []string{"fr-HT"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Haitian number, strict", param1: "+50934123456", param2: []string{"fr-HT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Haitian number, strict", param1: "+5093412345", param2: []string{"fr-HT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Madagascar (fr-MG)
		{name: "Valid Malagasy number", param1: "0341234567", param2: []string{"fr-MG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Malagasy number", param1: "03412345678", param2: []string{"fr-MG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Malagasy number, strict", param1: "+261341234567", param2: []string{"fr-MG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Malagasy number, strict", param1: "+26134123456", param2: []string{"fr-MG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Luxembourg (fr-LU)
		{name: "Valid Luxembourgish number", param1: "621234567", param2: []string{"fr-LU"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Luxembourgish number", param1: "6212345678", param2: []string{"fr-LU"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Luxembourgish number, strict", param1: "+352621234567", param2: []string{"fr-LU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Luxembourgish number, strict", param1: "+35262123456", param2: []string{"fr-LU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Morocco (fr-MA)
		{name: "Valid Moroccan number", param1: "0612345678", param2: []string{"fr-MA"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Moroccan number", param1: "06123456789", param2: []string{"fr-MA"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Moroccan number, strict", param1: "+212612345678", param2: []string{"fr-MA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Moroccan number, strict", param1: "+21261234567", param2: []string{"fr-MA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Senegal (fr-SN)
		{name: "Valid Senegalese number", param1: "778123456", param2: []string{"fr-SN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Senegalese number", param1: "7781234567", param2: []string{"fr-SN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Senegalese number, strict", param1: "+221778123456", param2: []string{"fr-SN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Senegalese number, strict", param1: "+22177812345", param2: []string{"fr-SN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Seychelles (fr-SC)
		{name: "Valid Seychellois number", param1: "2512345", param2: []string{"fr-SC"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Seychellois number", param1: "25123456", param2: []string{"fr-SC"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Seychellois number, strict", param1: "+2482512345", param2: []string{"fr-SC"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Seychellois number, strict", param1: "+248251234", param2: []string{"fr-SC"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Togo (fr-TG)
		{name: "Valid Togolese number", param1: "90123456", param2: []string{"fr-TG"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Togolese number", param1: "901234567", param2: []string{"fr-TG"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Togolese number, strict", param1: "+22890123456", param2: []string{"fr-TG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Togolese number, strict", param1: "+2289012345", param2: []string{"fr-TG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},

		// Irish
		{name: "Valid Irish number", param1: "0831234567", param2: []string{"ga-IE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Irish number", param1: "08312345678", param2: []string{"ga-IE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Irish number, strict", param1: "+353831234567", param2: []string{"ga-IE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Irish number, strict", param1: "+35383123456", param2: []string{"ga-IE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Hebrew
		{name: "Valid Hebrew number", param1: "0541234567", param2: []string{"he-IL"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Hebrew number", param1: "05412345678", param2: []string{"he-IL"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Hebrew number, strict", param1: "+972541234567", param2: []string{"he-IL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Hebrew number, strict", param1: "+97254123456", param2: []string{"he-IL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Hungarian
		{name: "Valid Hungarian number", param1: "06301234567", param2: []string{"hu-HU"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Hungarian number", param1: "063012345678", param2: []string{"hu-HU"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Hungarian number, strict", param1: "+36301234567", param2: []string{"hu-HU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Hungarian number, strict", param1: "+363012345678", param2: []string{"hu-HU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Indonesian
		{name: "Valid Indonesian number", param1: "08123456789", param2: []string{"id-ID"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Indonesian number", param1: "081234567890", param2: []string{"id-ID"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Indonesian number, strict", param1: "+628123456789", param2: []string{"id-ID"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Indonesian number, strict", param1: "+62812345678", param2: []string{"id-ID"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Iranian
		{name: "Valid Iranian number", param1: "09121234567", param2: []string{"ir-IR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Iranian number", param1: "091212345678", param2: []string{"ir-IR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Iranian number, strict", param1: "+989121234567", param2: []string{"ir-IR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Iranian number, strict", param1: "+98912123456", param2: []string{"ir-IR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Italian (Italy)
		{name: "Valid Italian number", param1: "3351234567", param2: []string{"it-IT"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Italian number", param1: "33512345678", param2: []string{"it-IT"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Italian number, strict", param1: "+393351234567", param2: []string{"it-IT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Italian number, strict", param1: "+39335123456", param2: []string{"it-IT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Italian (San Marino)
		{name: "Valid San Marino number", param1: "666123456", param2: []string{"it-SM"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid San Marino number", param1: "6661234567", param2: []string{"it-SM"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid San Marino number, strict", param1: "+378666123456", param2: []string{"it-SM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid San Marino number, strict", param1: "+37866612345", param2: []string{"it-SM"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Japanese
		{name: "Valid Japanese number", param1: "09012345678", param2: []string{"ja-JP"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Japanese number", param1: "090123456789", param2: []string{"ja-JP"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Japanese number, strict", param1: "+819012345678", param2: []string{"ja-JP"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Japanese number, strict", param1: "+81901234567", param2: []string{"ja-JP"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Georgian
		{name: "Valid Georgian number", param1: "599123456", param2: []string{"ka-GE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Georgian number", param1: "5991234567", param2: []string{"ka-GE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Georgian number, strict", param1: "+995599123456", param2: []string{"ka-GE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Georgian number, strict", param1: "+99559912345", param2: []string{"ka-GE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Dutch (Aruba)
		{name: "Valid Aruban number", param1: "592123456", param2: []string{"nl-AW"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Aruban number", param1: "5921234567", param2: []string{"nl-AW"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Aruban number, strict", param1: "+297592123456", param2: []string{"nl-AW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Aruban number, strict", param1: "+29759212345", param2: []string{"nl-AW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Dutch (Belgium)
		{name: "Valid Belgian number", param1: "0471234567", param2: []string{"nl-BE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Belgian number", param1: "04712345678", param2: []string{"nl-BE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Belgian number, strict", param1: "+32471234567", param2: []string{"nl-BE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Belgian number, strict", param1: "+3247123456", param2: []string{"nl-BE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Dutch (Netherlands)
		{name: "Valid Dutch number", param1: "0612345678", param2: []string{"nl-NL"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Dutch number", param1: "06123456789", param2: []string{"nl-NL"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Dutch number, strict", param1: "+31612345678", param2: []string{"nl-NL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Dutch number, strict", param1: "+3161234567", param2: []string{"nl-NL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Norwegian (Nynorsk)
		{name: "Valid Norwegian Nynorsk number", param1: "41234567", param2: []string{"nn-NO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Norwegian Nynorsk number", param1: "412345678", param2: []string{"nn-NO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Norwegian Nynorsk number, strict", param1: "+4741234567", param2: []string{"nn-NO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Norwegian Nynorsk number, strict", param1: "+474123456", param2: []string{"nn-NO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Polish
		{name: "Valid Polish number", param1: "601234567", param2: []string{"pl-PL"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Polish number", param1: "6012345678", param2: []string{"pl-PL"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Polish number, strict", param1: "+48601234567", param2: []string{"pl-PL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Polish number, strict", param1: "+4860123456", param2: []string{"pl-PL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},

		// Portuguese (Angola)
		{name: "Valid Angolan number", param1: "912345678", param2: []string{"pt-AO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Angolan number", param1: "9123456789", param2: []string{"pt-AO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Angolan number, strict", param1: "+244912345678", param2: []string{"pt-AO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Angolan number, strict", param1: "+24491234567", param2: []string{"pt-AO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Portuguese (Brazil)
		{name: "Valid Brazilian number", param1: "21987654321", param2: []string{"pt-BR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Brazilian number", param1: "2198765432", param2: []string{"pt-BR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Brazilian number, strict", param1: "+5521987654321", param2: []string{"pt-BR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Brazilian number, strict", param1: "+552198765432", param2: []string{"pt-BR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Portuguese (Portugal)
		{name: "Valid Portuguese number", param1: "912345678", param2: []string{"pt-PT"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Portuguese number", param1: "9123456789", param2: []string{"pt-PT"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Portuguese number, strict", param1: "+351912345678", param2: []string{"pt-PT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Portuguese number, strict", param1: "+35191234567", param2: []string{"pt-PT"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Romanian (Moldova)
		{name: "Valid Moldovan number", param1: "060123456", param2: []string{"ro-Md"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Moldovan number", param1: "0601234567", param2: []string{"ro-Md"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Moldovan number, strict", param1: "+37360123456", param2: []string{"ro-Md"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Moldovan number, strict", param1: "+3736012345", param2: []string{"ro-Md"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Romanian (Romania)
		{name: "Valid Romanian number", param1: "0712345678", param2: []string{"ro-RO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Romanian number", param1: "07123456789", param2: []string{"ro-RO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Romanian number, strict", param1: "+40712345678", param2: []string{"ro-RO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Romanian number, strict", param1: "+4071234567", param2: []string{"ro-RO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Russian
		{name: "Valid Russian number", param1: "9123456789", param2: []string{"ru-RU"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Russian number", param1: "91234567890", param2: []string{"ru-RU"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Russian number, strict", param1: "+79123456789", param2: []string{"ru-RU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Russian number, strict", param1: "+7912345678", param2: []string{"ru-RU"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Sinhala
		{name: "Valid Sinhalese number", param1: "0712345678", param2: []string{"si-LK"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Sinhalese number", param1: "07123456789", param2: []string{"si-LK"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Sinhalese number, strict", param1: "+94712345678", param2: []string{"si-LK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Sinhalese number, strict", param1: "+9471234567", param2: []string{"si-LK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Slovak
		{name: "Valid Slovak number", param1: "0912345678", param2: []string{"sk-SK"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Slovak number", param1: "09123456789", param2: []string{"sk-SK"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Slovak number, strict", param1: "+421912345678", param2: []string{"sk-SK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Slovak number, strict", param1: "+42191234567", param2: []string{"sk-SK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Slovenian
		{name: "Valid Slovenian number", param1: "040123456", param2: []string{"sl-SI"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Slovenian number", param1: "0401234567", param2: []string{"sl-SI"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Slovenian number, strict", param1: "+38640123456", param2: []string{"sl-SI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Slovenian number, strict", param1: "+3864012345", param2: []string{"sl-SI"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Somali
		{name: "Valid Somali number", param1: "612345678", param2: []string{"so-SO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Somali number", param1: "6123456789", param2: []string{"so-SO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Somali number, strict", param1: "+252612345678", param2: []string{"so-SO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Somali number, strict", param1: "+25261234567", param2: []string{"so-SO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Albanian
		{name: "Valid Albanian number", param1: "0691234567", param2: []string{"sq-AL"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Albanian number", param1: "06912345678", param2: []string{"sq-AL"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Albanian number, strict", param1: "+355691234567", param2: []string{"sq-AL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Albanian number, strict", param1: "+35569123456", param2: []string{"sq-AL"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Serbian
		{name: "Valid Serbian number", param1: "0612345678", param2: []string{"sr-RS"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Serbian number", param1: "06123456789", param2: []string{"sr-RS"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Serbian number, strict", param1: "+381612345678", param2: []string{"sr-RS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Serbian number, strict", param1: "+38161234567", param2: []string{"sr-RS"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Swedish
		{name: "Valid Swedish number", param1: "0701234567", param2: []string{"sv-SE"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Swedish number", param1: "07012345678", param2: []string{"sv-SE"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Swedish number, strict", param1: "+46701234567", param2: []string{"sv-SE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Swedish number, strict", param1: "+4670123456", param2: []string{"sv-SE"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Turkish
		{name: "Valid Turkish number", param1: "5012345678", param2: []string{"tr-TR"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Turkish number", param1: "50123456789", param2: []string{"tr-TR"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Turkish number, strict", param1: "+905012345678", param2: []string{"tr-TR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Turkish number, strict", param1: "+90501234567", param2: []string{"tr-TR"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Ukrainian
		{name: "Valid Ukrainian number", param1: "0671234567", param2: []string{"uk-UA"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Ukrainian number", param1: "06712345678", param2: []string{"uk-UA"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Ukrainian number, strict", param1: "+380671234567", param2: []string{"uk-UA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Ukrainian number, strict", param1: "+38067123456", param2: []string{"uk-UA"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Welsh
		{name: "Valid Welsh number", param1: "07123456789", param2: []string{"cy-GB"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Welsh number", param1: "071234567890", param2: []string{"cy-GB"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Welsh number, strict", param1: "+447123456789", param2: []string{"cy-GB"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Welsh number, strict", param1: "+44712345678", param2: []string{"cy-GB"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Vietnamese
		{name: "Valid Vietnamese number", param1: "0912345678", param2: []string{"vi-VN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Vietnamese number", param1: "09123456789", param2: []string{"vi-VN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Vietnamese number, strict", param1: "+84912345678", param2: []string{"vi-VN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Vietnamese number, strict", param1: "+8491234567", param2: []string{"vi-VN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Chinese (Simplified)
		{name: "Valid Chinese number", param1: "13812345678", param2: []string{"zh-CN"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Chinese number", param1: "138123456789", param2: []string{"zh-CN"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Chinese number, strict", param1: "+8613812345678", param2: []string{"zh-CN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Chinese number, strict", param1: "+861381234567", param2: []string{"zh-CN"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Chinese (Hong Kong)
		{name: "Valid Hong Kong number", param1: "91234567", param2: []string{"zh-HK"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Hong Kong number", param1: "912345678", param2: []string{"zh-HK"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Hong Kong number, strict", param1: "+85291234567", param2: []string{"zh-HK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Hong Kong number, strict", param1: "+8529123456", param2: []string{"zh-HK"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Chinese (Macau)
		{name: "Valid Macau number", param1: "61234567", param2: []string{"zh-MO"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Macau number", param1: "612345678", param2: []string{"zh-MO"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Macau number, strict", param1: "+85361234567", param2: []string{"zh-MO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Macau number, strict", param1: "+8536123456", param2: []string{"zh-MO"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},

		// Chinese (Traditional)
		{name: "Valid Taiwanese number", param1: "0912345678", param2: []string{"zh-TW"}, param3: &IsMobilePhoneOpts{}, want: true},
		{name: "Invalid Taiwanese number", param1: "09123456789", param2: []string{"zh-TW"}, param3: &IsMobilePhoneOpts{}, want: false},
		{name: "Valid Taiwanese number, strict", param1: "+886912345678", param2: []string{"zh-TW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
		{name: "Invalid Taiwanese number, strict", param1: "+88691234567", param2: []string{"zh-TW"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},
	}

	// tests := []struct {
	// 	name   string
	// 	param1 string
	// 	param2 []string
	// 	param3 *IsMobilePhoneOpts
	// 	want   bool
	// }{

	// {name: "Valid nigerian number", param1: "08070448986", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{}, want: true},
	// {name: "Invalid nigerian number", param1: "090666666567", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{}, want: false},
	// {name: "Valid nigerian number, strict", param1: "+2348120989668", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: true},
	// {name: "Invalid nigerian number, strict", param1: "+23481209895", param2: []string{"en-NG"}, param3: &IsMobilePhoneOpts{StrictMode: true}, want: false},
	// }

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_ = IsMobilePhone(test.param1, test.param2, test.param3)

			// if result != test.want {
			// 	t.Errorf("got `%t`, wanted `%t`", result, test.want)
			// }
		})
	}
}
