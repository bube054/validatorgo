package validatorgo

import (
	"regexp"
	"sort"
	"strings"
)

var (
	isURLOptsDefaultProtocols                 []string = []string{"http", "https", "ftp"}
	isURLOptsDefaultRequireTld                bool     = true
	isURLOptsDefaultRequireProtocol           bool     = false
	isURLOptsDefaultRequireHost               bool     = true
	isURLOptsDefaultRequirePort               bool     = false
	isURLOptsDefaultRequireValidProtocol      bool     = true
	isURLOptsDefaultAllowUnderscores          bool     = false
	isURLOptsDefaultHostWhitelist             []string = nil
	isURLOptsDefaultHostBlacklist             []string = nil
	isURLOptsDefaultAllowTrailingDot          bool     = false
	isURLOptsDefaultAllowProtocolRelativeUrls bool     = true
	isURLOptsDefaultAllowFragments            bool     = true
	isURLOptsDefaultAllowQueryComponents      bool     = true
	isURLOptsDefaultDisallowAuth              bool     = false
	isURLOptsDefaultValidateLength            bool     = true
	isURLOptsDefaultMaxLength                 int      = 2048
)

type IsURLOpts struct {
	Protocols                 []string
	RequireTld                bool
	RequireProtocol           bool
	RequireHost               bool
	RequirePort               bool
	RequireValidProtocol      bool
	AllowUnderscores          bool
	HostWhitelist             []string
	HostBlacklist             []string
	AllowTrailingDot          bool
	AllowProtocolRelativeUrls bool
	AllowFragments            bool
	AllowQueryComponents      bool
	DisallowAuth              bool
	ValidateLength            bool
	MaxAllowedLength          int
}

// A validator that checks if the string is a URL.
//
// IsURLOpts is a struct which defaults to { Protocols: []string{"https","http","ftp"}, RequireTld: true, RequireProtocol: false, RequireHost: true, RequirePort: false, RequireValidProtocol: true, AllowUnderscores: false, HostWhitelist: nil, HostBlacklist: nil, AllowTrailingDot: false, AllowProtocolRelativeUrls: true, AllowFragments: true, AllowQueryComponents: true, DisallowAuth: false, ValidateLength: true, MaxAllowedLength:: 2048 }.
//
// RequireProtocol - if set to true IsURL will return false if protocol is not present in the URL.
//
// RequireValidProtocol - IsURL will check if the URL's protocol is present in the protocols option.
//
// protocols - valid protocols can be modified with this option.
//
// RequireHost - if set to false IsURL will not check if host is present in the URL.
//
// RequirePort - if set to true IsURL will check if port is present in the URL.
//
// AllowProtocolRelativeUrls - if set to true protocol relative URLs will be allowed.
//
// AllowFragments - if set to false IsURL will return false if fragments are present.
//
// AllowQueryComponents - if set to false IsURL will return false if query components are present.
//
// ValidateLength - if set to false IsURL will skip string length validation. MaxAllowedLength will be ignored if this is set as false.
//
// MaxAllowedLength - if set IsURL will not allow URLs longer than the specified value (default is 2084 that IE maximum URL length).
//
//	ok := validatorgo.IsURL("http://localhost", &validatorgo.IsURLOpts{RequireTld: false, AllowProtocolRelativeUrls: true})
//	fmt.Println(ok) // true
//	ok := validatorgo.IsURL("example.com", &validatorgo.IsURLOpts{RequireProtocol: true, AllowProtocolRelativeUrls: true})
//	fmt.Println(ok) // false
func IsURL(str string, opts *IsURLOpts) bool {
	if opts == nil {
		opts = setIsURLOptsToDefault()
	}

	if len(opts.Protocols) == 0 {
		opts.Protocols = isURLOptsDefaultProtocols
	}

	if opts.MaxAllowedLength == 0 {
		opts.MaxAllowedLength = isURLOptsDefaultMaxLength
	}

	// make protocols with larger lengths come first
	sort.Slice(opts.Protocols, func(i, j int) bool {
		return len(opts.Protocols[i]) > len(opts.Protocols[j])
	})

	reStr := `^(\w+)?(:?\/\/)?(\w*:\w*@)?([\w\/]+)(\.[a-zA-Z]*(\.)?)?(:\d+)?(\/#[a-zA-Z0-9_]+)?(\/\?[a-zA-Z0-9_]+=[a-zA-Z0-9_]+(&[a-zA-Z0-9_]+=[a-zA-Z0-9_]+)*)?`
	re := regexp.MustCompile(reStr)

	capGrp := re.FindStringSubmatch(str)

	if len(capGrp) == 0 {
		return false
	}

	proto := capGrp[1]  // http
	slash := capGrp[2]  // ://
	auth := capGrp[3]   // user:pass@
	subdom := capGrp[4] // example
	dom := capGrp[5]    // .com
	trlDot := capGrp[6] // .
	port := capGrp[7]   // :8080
	hash := capGrp[8]   // /#section1
	param := capGrp[9]  // key1=value1&key2=value2&key3=value3

	if opts.RequireTld {
		if dom == "" {
			return false
		}
	}

	if opts.RequireProtocol {
		if proto == "" || slash == "" {
			return false
		}
	}

	if opts.RequireHost {
		if dom == "" {
			return false
		}
	}

	if opts.RequirePort {
		if port == "" {
			return false
		}
	}

	if opts.RequireValidProtocol {
		if !IsIn(proto, opts.Protocols) {
			return false
		}
	}

	if !opts.AllowUnderscores {
		if strings.Contains(subdom, "_") {
			return false
		}
	}

	if len(opts.HostWhitelist) > 0 {
		subdomPlusDom := subdom + dom
		isin := IsIn(subdomPlusDom, opts.HostWhitelist)

		if !isin {
			return false
		}
	}

	if len(opts.HostBlacklist) > 0 {
		subdomPlusDom := subdom + dom
		isin := IsIn(subdomPlusDom, opts.HostBlacklist)

		if isin {
			return false
		}
	}

	if !opts.AllowTrailingDot {
		if trlDot == "." {
			return false
		}
	}

	if !opts.AllowProtocolRelativeUrls {
		if proto != "" || slash != "//" {
			return false
		}
	}

	if !opts.AllowFragments {
		if hash != "" {
			return false
		}
	}

	if !opts.AllowQueryComponents {
		if param != "" {
			return false
		}
	}

	if opts.DisallowAuth {
		if auth != "" {
			return false
		}
	}

	if opts.ValidateLength {
		if len(str) > opts.MaxAllowedLength {
			return false
		}
	}

	return true
}

func setIsURLOptsToDefault() *IsURLOpts {
	return &IsURLOpts{
		Protocols:                 isURLOptsDefaultProtocols,
		RequireTld:                isURLOptsDefaultRequireTld,
		RequireProtocol:           isURLOptsDefaultRequireProtocol,
		RequireHost:               isURLOptsDefaultRequireHost,
		RequirePort:               isURLOptsDefaultRequirePort,
		RequireValidProtocol:      isURLOptsDefaultRequireValidProtocol,
		AllowUnderscores:          isURLOptsDefaultAllowUnderscores,
		HostWhitelist:             isURLOptsDefaultHostWhitelist,
		HostBlacklist:             isURLOptsDefaultHostBlacklist,
		AllowTrailingDot:          isURLOptsDefaultAllowTrailingDot,
		AllowProtocolRelativeUrls: isURLOptsDefaultAllowProtocolRelativeUrls,
		AllowFragments:            isURLOptsDefaultAllowFragments,
		AllowQueryComponents:      isURLOptsDefaultAllowQueryComponents,
		DisallowAuth:              isURLOptsDefaultDisallowAuth,
		ValidateLength:            isURLOptsDefaultValidateLength, // isURLOptsDefaultMaxAllowedLength
		MaxAllowedLength:          isURLOptsDefaultMaxLength,
	}
}

// http://example.com
// https://example.com

// http://example.com
// http://localhost

// http://example.com
// example.com

// http://example.com
// /path/to/resource

// http://example.com:8080
// http://example.com

// http://example.com.

// //example.com

// http://example.com/#section1

// http://example.com/?key=value
// http://user:pass@example.com

// ^(https|http|ftp)?(:?\/\/)?(\w*:\w*@)?([\w\/]+)(\.[a-zA-Z]*(\.)?)?(:\d+)?(\/#[a-zA-Z0-9_]+)?(\/\?[a-zA-Z0-9_]+=[a-zA-Z0-9_]+)?((&[a-zA-Z0-9_]+=[a-zA-Z0-9_]+)*)?
