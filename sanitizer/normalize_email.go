package sanitizer

const (
	normalizeEmailOptsDefaultAllLowercase bool = true
)

type NormalizeEmailOpts struct {
	AllLowercase                  bool
	GmailLowercase                bool
	GmailRemoveDots               bool
	GmailRemoveSubaddress         bool
	GmailConvertGooglemaildotcom  bool
	OutlookdotcomLowercase        bool
	OutlookdotcomRemoveSubaddress bool
	YahooRemoveSubaddress         bool
	IcloudLowercase               bool
	IcloudRemoveSubaddress        bool
}

// A sanitizer that canonicalizes an email address. (This doesn't validate that the input is an email, if you want to validate the email use isEmail beforehand).
//
// NormalizeEmailOpts is a struct with the following keys and default values:
//
// AllLowercase: true - Transforms the local part (before the @ symbol) of all email addresses to lowercase. Please note that this may violate RFC 5321, which gives providers the possibility to treat the local part of email addresses in a case sensitive way (although in practice most - yet not all - providers don't). The domain part of the email address is always lowercased, as it is case insensitive per RFC 1035.
//
// GmailLowercase: true - Gmail addresses are known to be case-insensitive, so this switch allows lowercasing them even when AllLowercase is set to false. Please note that when AllLowercase is true, Gmail addresses are lowercased regardless of the value of this setting.
//
// GmailRemoveDots: true: Removes dots from the local part of the email address, as Gmail ignores them (e.g. "john.doe" and "johndoe" are considered equal).
//
// GmailRemoveSubaddress: true: Normalizes addresses by removing "sub-addresses", which is the part following a "+" sign (e.g. "foo+bar@gmail.com" becomes "foo@gmail.com").
//
// GmailConvertGooglemaildotcom: true: Converts addresses with domain @googlemail.com to @gmail.com, as they're equivalent.
//
// OutlookdotcomLowercase: true - Outlook.com addresses (including Windows Live and Hotmail) are known to be case-insensitive, so this switch allows lowercasing them even when AllLowercase is set to false. Please note that when AllLowercase is true, Outlook.com addresses are lowercased regardless of the value of this setting.
//
// OutlookdotcomRemoveSubaddress: true: Normalizes addresses by removing "sub-addresses", which is the part following a "+" sign (e.g. "foo+bar@outlook.com" becomes "foo@outlook.com").
// yahoo_lowercase: true - Yahoo Mail addresses are known to be case-insensitive, so this switch allows lowercasing them even when AllLowercase is set to false. Please note that when AllLowercase is true, Yahoo Mail addresses are lowercased regardless of the value of this setting.
//
// YahooRemoveSubaddress: true: Normalizes addresses by removing "sub-addresses", which is the part following a "-" sign (e.g. "foo-bar@yahoo.com" becomes "foo@yahoo.com").
// IcloudLowercase: true - iCloud addresses (including MobileMe) are known to be case-insensitive, so this switch allows lowercasing them even when AllLowercase is set to false. Please note that when AllLowercase is true, iCloud addresses are lowercased regardless of the value of this setting.
//
// IcloudRemoveSubaddress: true: Normalizes addresses by removing "sub-addresses", which is the part following a "+" sign (e.g. "foo+bar@icloud.com" becomes "foo@icloud.com").
func NormalizeEmail(email string, opts *NormalizeEmailOpts) string {
	return ""
}
