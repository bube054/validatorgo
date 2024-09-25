package validatorgo

type IsURLOpts struct {
	Protocols                 []string
	RequireTld                bool
	RequireProtocol           bool
	RequireHost               bool
	RequirePort               bool
	RequireValidProtocol      bool
	AllowUnderscores          bool
	HostWhitelist             bool
	HostBlacklist             bool
	AllowTrailingDot          bool
	AllowProtocolRelativeUrls bool
	AllowFragments            bool
	AllowQueryComponents      bool
	DisallowAuth              bool
	ValidateLength            bool
	MaxAllowedLength          bool
}

// A validator that checks if the string is a URL.
//
// IsURLOpts is a struct which defaults to { Protocols: ["http","https","ftp"], RequireTld: true, RequireProtocol: false, RequireHost: true, RequirePort: false, RequireValidProtocol: true, AllowUnderscores: false, HostWhitelist: false, HostBlacklist: false, AllowTrailingDot: false, AllowProtocolRelativeUrls: false, AllowFragments: true, AllowQueryComponents: true, DisallowAuth: false, ValidateLength: true }.
//
// RequireProtocol - if set to true isURL will return false if protocol is not present in the URL.
//
// RequireValidProtocol - isURL will check if the URL's protocol is present in the protocols option.
//
// protocols - valid protocols can be modified with this option.
//
// RequireHost - if set to false isURL will not check if host is present in the URL.
//
// RequirePort - if set to true isURL will check if port is present in the URL.
//
// AllowProtocolRelativeUrls - if set to true protocol relative URLs will be allowed.
//
// AllowFragments - if set to false isURL will return false if fragments are present.
//
// AllowQueryComponents - if set to false isURL will return false if query components are present.
//
// ValidateLength - if set to false isURL will skip string length validation. MaxAllowedLength will be ignored if this is set as false.
//
// MaxAllowedLength - if set isURL will not allow URLs longer than the specified value (default is 2084 that IE maximum URL length).
func IsURL(str string, opts *IsURLOpts) bool {
	return false
}
