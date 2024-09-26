package sanitizer

import (
	"regexp"
	"strings"
)

const (
	normalizeEmailOptsDefaultAllLowercase                  bool = true
	normalizeEmailOptsDefaultGmailLowercase                bool = true
	normalizeEmailOptsDefaultGmailRemoveDots               bool = true
	normalizeEmailOptsDefaultGmailRemoveSubaddress         bool = true
	normalizeEmailOptsDefaultGmailConvertGooglemaildotcom  bool = true
	normalizeEmailOptsDefaultOutlookdotcomLowercase        bool = true
	normalizeEmailOptsDefaultOutlookdotcomRemoveSubaddress bool = true
	normalizeEmailOptsDefaultYahooLowercase                bool = true
	normalizeEmailOptsDefaultYahooRemoveSubaddress         bool = true
	normalizeEmailOptsDefaultIcloudLowercase               bool = true
	normalizeEmailOptsDefaultIcloudRemoveSubaddress        bool = true
)

type NormalizeEmailOpts struct {
	AllLowercase                  bool
	GmailLowercase                bool
	GmailRemoveDots               bool
	GmailRemoveSubaddress         bool
	GmailConvertGooglemaildotcom  bool
	OutlookdotcomLowercase        bool
	OutlookdotcomRemoveSubaddress bool
	YahooLowercase                bool
	YahooRemoveSubaddress         bool
	IcloudLowercase               bool
	IcloudRemoveSubaddress        bool
}

// A sanitizer that canonicalizes an email address. (This doesn't validate that the input is an email, if you want to validate the email use IsEmail beforehand).
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
//
// YahooLowercase: true - Yahoo Mail addresses are known to be case-insensitive, so this switch allows lowercasing them even when AllLowercase is set to false. Please note that when AllLowercase is true, Yahoo Mail addresses are lowercased regardless of the value of this setting.
//
// YahooRemoveSubaddress: true: Normalizes addresses by removing "sub-addresses", which is the part following a "-" sign (e.g. "foo-bar@yahoo.com" becomes "foo@yahoo.com").
//
// IcloudLowercase: true - iCloud addresses (including MobileMe) are known to be case-insensitive, so this switch allows lowercasing them even when AllLowercase is set to false. Please note that when AllLowercase is true, iCloud addresses are lowercased regardless of the value of this setting.
//
// IcloudRemoveSubaddress: true: Normalizes addresses by removing "sub-addresses", which is the part following a "+" sign (e.g. "foo+bar@icloud.com" becomes "foo@icloud.com").
//
//	str := sanitizer.NormalizeEmail("Example@Example.com", &NormalizeEmailOpts{AllLowercase: true})
//	fmt.Println(str) // example@example.com
func NormalizeEmail(email string, opts *NormalizeEmailOpts) string {
	normEmail := email

	if opts == nil {
		opts = setNormalizedEmailOptsToDefault()
	}

	if opts.AllLowercase {
		normEmail = strings.ToLower(normEmail)
	}
	// dealing with googlemail starts here
	geMailRe := regexp.MustCompile(`^(.*)@(googlemail.com)$`)
	geCapGrp := geMailRe.FindStringSubmatch(normEmail)
	isGeMail := len(geCapGrp) != 0

	if opts.GmailConvertGooglemaildotcom && isGeMail {
		geLocal := geCapGrp[1]
		normEmail = geLocal + "@" + "gmail.com"
	}
	// dealing with googlemail ends here

	// dealing with gmail starts here
	gmailRe := regexp.MustCompile(`^(.*)@([gG][mM][aA][iI][lL]\.[cC][oO][mM])$`)
	gmailCapGr := gmailRe.FindStringSubmatch(normEmail)
	isGmail := len(gmailCapGr) != 0

	var (
		gLocal  string
		gDomain string
	)

	if isGmail {
		gLocal, gDomain = gmailCapGr[1], gmailCapGr[2]
	}

	if opts.GmailLowercase && isGmail {
		gDomain = "gmail.com"
		gLocal = strings.ToLower(gLocal)
		normEmail = gLocal + "@" + gDomain
	}

	if opts.GmailRemoveDots && isGmail {
		gLocal = Blacklist(gLocal, ".")
		normEmail = gLocal + "@" + gDomain
	}

	if opts.GmailRemoveSubaddress && isGmail {
		gLocal = strings.Split(gLocal, "+")[0]
		normEmail = gLocal + "@" + gDomain
	}

	// dealing with gmail ends here

	// dealing with outlook starts here
	outMailRe := regexp.MustCompile(`^(.*)@([oO][uU][tT][lL][oO][oO][kK]\.[cC][oO][mM])$`)
	outMailCapGr := outMailRe.FindStringSubmatch(normEmail)
	isOutMail := len(outMailCapGr) != 0

	var (
		outLocal  string
		outDomain string
	)

	if isOutMail {
		outLocal, outDomain = outMailCapGr[1], outMailCapGr[2]
	}

	if opts.OutlookdotcomLowercase && isOutMail {
		outDomain = "outlook.com"
		outLocal = strings.ToLower(outLocal)
		normEmail = outLocal + "@" + outDomain
	}

	if opts.OutlookdotcomRemoveSubaddress && isOutMail {
		outLocal = strings.Split(outLocal, "+")[0]
		normEmail = outLocal + "@" + outDomain
	}
	// dealing with outlook ends here

	// dealing with yahoo starts here
	yhMailRe := regexp.MustCompile(`^(.*)@([yY][aA][hH][oO][oO]\.[cC][oO][mM])$`)
	yhMailCapGr := yhMailRe.FindStringSubmatch(normEmail)
	isYhMail := len(yhMailCapGr) != 0

	var (
		yhLocal  string
		yhDomain string
	)

	if isYhMail {
		yhLocal, yhDomain = yhMailCapGr[1], yhMailCapGr[2]
	}

	if opts.YahooLowercase && isYhMail {
		yhDomain = "yahoo.com"
		yhLocal = strings.ToLower(yhLocal)
		normEmail = yhLocal + "@" + yhDomain
	}

	if opts.YahooRemoveSubaddress && isYhMail {
		yhLocal = strings.Split(yhLocal, "+")[0]
		normEmail = yhLocal + "@" + yhDomain
	}
	// dealing with yahoo ends here

	// dealing with icloud starts here
	icMailRe := regexp.MustCompile(`^(.*)@([iI][cC][lL][oO][uU][dD]\.[cC][oO][mM])$`)
	icMailCapGr := icMailRe.FindStringSubmatch(normEmail)
	isIcMail := len(icMailCapGr) != 0

	var (
		icLocal  string
		icDomain string
	)

	if isIcMail {
		icLocal, icDomain = icMailCapGr[1], icMailCapGr[2]
	}

	if opts.IcloudLowercase && isIcMail {
		icDomain = "icloud.com"
		icLocal = strings.ToLower(icLocal)
		normEmail = icLocal + "@" + icDomain
	}

	if opts.IcloudRemoveSubaddress && isIcMail {
		icLocal = strings.Split(icLocal, "+")[0]
		normEmail = icLocal + "@" + icDomain
	}

	// dealing with icloud starts here

	return normEmail
}

func setNormalizedEmailOptsToDefault() (opts *NormalizeEmailOpts) {
	opts = &NormalizeEmailOpts{}
	opts.AllLowercase = normalizeEmailOptsDefaultAllLowercase
	opts.GmailLowercase = normalizeEmailOptsDefaultGmailLowercase
	opts.GmailRemoveDots = normalizeEmailOptsDefaultGmailRemoveDots
	opts.GmailRemoveSubaddress = normalizeEmailOptsDefaultGmailRemoveSubaddress
	opts.GmailConvertGooglemaildotcom = normalizeEmailOptsDefaultGmailConvertGooglemaildotcom
	opts.OutlookdotcomLowercase = normalizeEmailOptsDefaultOutlookdotcomLowercase
	opts.OutlookdotcomRemoveSubaddress = normalizeEmailOptsDefaultOutlookdotcomRemoveSubaddress
	opts.YahooLowercase = normalizeEmailOptsDefaultYahooLowercase
	opts.YahooRemoveSubaddress = normalizeEmailOptsDefaultYahooRemoveSubaddress
	opts.IcloudLowercase = normalizeEmailOptsDefaultIcloudLowercase
	opts.IcloudRemoveSubaddress = normalizeEmailOptsDefaultIcloudRemoveSubaddress

	return
}
