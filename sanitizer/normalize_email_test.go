package sanitizer

import (
	"testing"
)

func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		name   string
		param1 string
		param2 *NormalizeEmailOpts
		want   string
	}{
		{name: "Lowercase all email characters", param1: "Example@Example.com", param2: &NormalizeEmailOpts{AllLowercase: true}, want: "example@example.com"},
		{name: "Lowercase Gmail", param1: "JohnDoe@Gmail.COM", param2: &NormalizeEmailOpts{GmailLowercase: true}, want: "johndoe@gmail.com"},
		{name: "Remove dots in Gmail address", param1: "john.doe@gmail.com", param2: &NormalizeEmailOpts{GmailRemoveDots: true}, want: "johndoe@gmail.com"},
		{name: "Remove subaddress in Gmail", param1: "john.doe+newsletter@gmail.com", param2: &NormalizeEmailOpts{GmailRemoveSubaddress: true}, want: "john.doe@gmail.com"},
		{name: "Convert googlemail.com to gmail.com", param1: "johndoe@googlemail.com", param2: &NormalizeEmailOpts{GmailConvertGooglemaildotcom: true}, want: "johndoe@gmail.com"},
		{name: "Lowercase Outlook", param1: "JohnDoe@Outlook.COM", param2: &NormalizeEmailOpts{OutlookdotcomLowercase: true}, want: "johndoe@outlook.com"},
		{name: "Remove subaddress in Outlook", param1: "john.doe+notifications@outlook.com", param2: &NormalizeEmailOpts{OutlookdotcomRemoveSubaddress: true}, want: "john.doe@outlook.com"},
		{name: "Lowercase Yahoo", param1: "User+Label@yahoo.com", param2: &NormalizeEmailOpts{YahooLowercase: true}, want: "user+label@yahoo.com"},
		{name: "Remove subaddress in Yahoo", param1: "user+label@yahoo.com", param2: &NormalizeEmailOpts{YahooRemoveSubaddress: true}, want: "user@yahoo.com"},
		{name: "Lowercase iCloud domain only", param1: "user+label@Icloud.COM", param2: &NormalizeEmailOpts{IcloudLowercase: true}, want: "user+label@icloud.com"},
		{name: "Remove subaddress in iCloud", param1: "user+notification@icloud.com", param2: &NormalizeEmailOpts{IcloudRemoveSubaddress: true}, want: "user@icloud.com"},

		{name: "Lowercase and remove Gmail subaddress", param1: "JohnDoe+news@Gmail.COM", param2: &NormalizeEmailOpts{GmailLowercase: true, GmailRemoveSubaddress: true}, want: "johndoe@gmail.com"},
		{name: "Lowercase and remove dots in Gmail", param1: "john.doe+offers@gmail.com", param2: &NormalizeEmailOpts{GmailLowercase: true, GmailRemoveDots: true}, want: "johndoe+offers@gmail.com"},
		{name: "Remove dots and subaddress in Gmail", param1: "john.doe+test@gmail.com", param2: &NormalizeEmailOpts{GmailRemoveDots: true, GmailRemoveSubaddress: true}, want: "johndoe@gmail.com"},
		{name: "Convert googlemail.com and remove subaddress", param1: "johndoe+work@googlemail.com", param2: &NormalizeEmailOpts{GmailConvertGooglemaildotcom: true, GmailRemoveSubaddress: true}, want: "johndoe@gmail.com"},
		{name: "Lowercase and remove Outlook subaddress", param1: "User123+notes@Outlook.COM", param2: &NormalizeEmailOpts{OutlookdotcomLowercase: true, OutlookdotcomRemoveSubaddress: true}, want: "user123@outlook.com"},
		{name: "Lowercase and remove subaddress in Yahoo", param1: "MyEmail+promo@Yahoo.com", param2: &NormalizeEmailOpts{AllLowercase: true, YahooRemoveSubaddress: true}, want: "myemail@yahoo.com"},
		{name: "Lowercase, remove dots, and subaddress in Gmail", param1: "John.Doe+newsletter@Gmail.com", param2: &NormalizeEmailOpts{GmailLowercase: true, GmailRemoveDots: true, GmailRemoveSubaddress: true}, want: "johndoe@gmail.com"},
		{name: "Lowercase and remove iCloud subaddress", param1: "MyEmail+updates@iCloud.COM", param2: &NormalizeEmailOpts{IcloudLowercase: true, IcloudRemoveSubaddress: true}, want: "myemail@icloud.com"},
		{name: "Lowercase all and remove Gmail subaddress", param1: "User123+info@Gmail.com", param2: &NormalizeEmailOpts{AllLowercase: true, GmailRemoveSubaddress: true}, want: "user123@gmail.com"},
		{name: "Lowercase all and remove Outlook subaddress", param1: "John+Test@Outlook.com", param2: &NormalizeEmailOpts{AllLowercase: true, OutlookdotcomRemoveSubaddress: true}, want: "john@outlook.com"},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			result := NormalizeEmail(test.param1, test.param2)

			if result != test.want {
				t.Errorf("%s got `%s`, wanted `%s`", test.name, result, test.want)
			}
		})
	}
}
