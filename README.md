# validatorgo

[![Go Reference](https://pkg.go.dev/badge/github.com/bube054/validatorgo.svg)](https://pkg.go.dev/github.com/bube054/validatorgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/bube054/validatorgo)](https://goreportcard.com/report/github.com/bube054/validatorgo)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bube054/validatorgo)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> A comprehensive library of **string validators** and **sanitizers** for Go, inspired by the popular JavaScript library [validator.js](https://github.com/validatorjs/validator.js).

---

## Table of Contents

- [Features](#features)
- [Why validatorgo?](#why-validatorgo)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Validators](#validators)
  - [Strings & Text](#strings--text)
  - [Numbers & Booleans](#numbers--booleans)
  - [Dates & Times](#dates--times)
  - [Network, Web & URIs](#network-web--uris)
  - [Encoding, Hashes & JSON](#encoding-hashes--json)
  - [Identifiers, Codes & Versioning](#identifiers-codes--versioning)
  - [Geographic & Locale](#geographic--locale)
  - [Personal & Identity](#personal--identity)
  - [Finance & Crypto](#finance--crypto)
  - [Colors & Data Structures](#colors--data-structures)
- [Sanitizers](#sanitizers)
- [Error Handling](#error-handling)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **89+ validators** covering everything from emails and URLs to IBANs and ISO codes
- **13 sanitizers** for cleaning, normalizing, and converting input strings
- **Structured errors** — every validator returns `(bool, error)` with machine-readable error codes
- **Locale-aware** — many validators accept locale options for region-specific formats
- **Pure Go**, no heavy dependencies
- **Comprehensive tests** with ~87% coverage
- Drop-in mental model for anyone familiar with [validator.js](https://github.com/validatorjs/validator.js)

---

## Why validatorgo?

Other popular Go libraries like [go-playground/validator](https://github.com/go-playground/validator) and [asaskevich/govalidator](https://github.com/asaskevich/govalidator) primarily focus on validating **structs**, not standalone strings. They also lack the breadth of validators offered by [validator.js](https://github.com/validatorjs/validator.js).

validatorgo was originally built as a dependency for [ginvalidator](https://github.com/bube054/ginvalidator) — a Go equivalent of the popular Node.js library [express-validator](https://express-validator.github.io/docs/). The existing Go ecosystem options were either too complex, not robust enough, or unsuitable for that use case, so this library fills the gap.

---

## Installation

```sh
go get -u github.com/bube054/validatorgo
```

Then import it in your code:

```go
import (
    "fmt"
    "github.com/bube054/validatorgo"
)
```

Prefer a shorter alias?

```go
import (
    "fmt"
    vgo "github.com/bube054/validatorgo"
)
```

---

## Quick Start

**Validate a MongoDB ObjectID:**

```go
package main

import (
    "fmt"
    vgo "github.com/bube054/validatorgo"
)

func main() {
    id := "5f2a6c69e1d7a4e0077b4e6b"
    ok, err := vgo.IsMongoID(id)
    fmt.Println(ok)  // true
    fmt.Println(err) // <nil>
}
```

**Validate an email with options:**

```go
// Pass nil to use all defaults
ok, _ := vgo.IsEmail("user@example.com", nil)
fmt.Println(ok) // true

// Or customize only what you need
ok, err := vgo.IsEmail("user@example.com", &vgo.IsEmailOpts{
    RequireTld:               vgo.Bool(false), // override default (true)
    DomainSpecificValidation: true,
})
fmt.Println(ok, err) // true <nil>
```

**Sanitize a string:**

```go
import "github.com/bube054/validatorgo/sanitizer"

clean := sanitizer.Whitelist("Hello123 World!", "a-zA-Z")
fmt.Println(clean) // "HelloWorld"
```

For full API docs and parameter details, see the [Go Reference](https://pkg.go.dev/github.com/bube054/validatorgo).

---

## Validators

### Strings & Text

| Validator | Description |
| --- | --- |
| **[Contains](https://pkg.go.dev/github.com/bube054/validatorgo#Contains)**`(str, seed string, opts *ContainsOpt)` | Checks if a string contains a seed. `ContainsOpt` defaults to `{ IgnoreCase: false, MinOccurrences: 1 }`. |
| **[Equals](https://pkg.go.dev/github.com/bube054/validatorgo#Equals)**`(str, comparison string)` | Checks if the string matches the comparison. |
| **[Matches](https://pkg.go.dev/github.com/bube054/validatorgo#Matches)**`(str string, re *regexp.Regexp)` | Checks if the string matches the given regex. |
| **[IsEmpty](https://pkg.go.dev/github.com/bube054/validatorgo#IsEmpty)**`(str string, opts *IsEmptyOpts)` | Checks if the string is zero-length. `IsEmptyOpts` defaults to `{ IgnoreWhitespace: false }`. |
| **[IsLowerCase](https://pkg.go.dev/github.com/bube054/validatorgo#IsLowerCase)**`(str string)` | Checks if the string is lowercase. |
| **[IsUpperCase](https://pkg.go.dev/github.com/bube054/validatorgo#IsUpperCase)**`(str string)` | Checks if the string is uppercase. |
| **[IsAscii](https://pkg.go.dev/github.com/bube054/validatorgo#IsAscii)**`(str string)` | Checks if the string contains ASCII characters only. |
| **[IsAlpha](https://pkg.go.dev/github.com/bube054/validatorgo#IsAlpha)**`(str string, opts *IsAlphaOpts)` | Checks if the string contains only letters (a-zA-Z). Supports `Ignore` and `Locale` options. Defaults to `en-US`. |
| **[IsAlphanumeric](https://pkg.go.dev/github.com/bube054/validatorgo#IsAlphanumeric)**`(str string, opts *IsAlphanumericOpts)` | Checks if the string contains only letters and numbers (a-zA-Z0-9). Supports `Ignore` and `Locale`. |
| **[IsByteLength](https://pkg.go.dev/github.com/bube054/validatorgo#IsByteLength)**`(str string, opts *IsByteLengthOpts)` | Checks if the string's UTF-8 byte length falls within a range. Defaults to `{ Min: 0, Max: nil }`. |
| **[IsSlug](https://pkg.go.dev/github.com/bube054/validatorgo#IsSlug)**`(str string)` | Checks if the string is a valid slug. |
| **[IsFullWidth](https://pkg.go.dev/github.com/bube054/validatorgo#IsFullWidth)**`(str string)` | Checks if the string contains any full-width characters. |
| **[IsHalfWidth](https://pkg.go.dev/github.com/bube054/validatorgo#IsHalfWidth)**`(str string)` | Checks if the string contains any half-width characters. |
| **[IsMultibyte](https://pkg.go.dev/github.com/bube054/validatorgo#IsMultibyte)**`(str string)` | Checks if the string contains one or more multibyte characters. |
| **[IsSurrogatePair](https://pkg.go.dev/github.com/bube054/validatorgo#IsSurrogatePair)**`(str string)` | Checks if the string contains any surrogate pair characters. |
| **[IsVariableWidth](https://pkg.go.dev/github.com/bube054/validatorgo#IsVariableWidth)**`(str string)` | Checks if the string contains a mixture of full and half-width characters. |
| **[IsWhitelisted](https://pkg.go.dev/github.com/bube054/validatorgo#IsWhitelisted)**`(str, chars string)` | Checks if the string consists only of characters in the given whitelist. |

### Numbers & Booleans

| Validator | Description |
| --- | --- |
| **[IsBoolean](https://pkg.go.dev/github.com/bube054/validatorgo#IsBoolean)**`(str string, opts *IsBooleanOpts)` | Checks if the string is a boolean. |
| **[IsInt](https://pkg.go.dev/github.com/bube054/validatorgo#IsInt)**`(str string, opts *IsIntOpts)` | Checks if the string is an integer. Supports `Min`, `Max`, `Gt`, `Lt`, and `AllowLeadingZeroes`. |
| **[IsFloat](https://pkg.go.dev/github.com/bube054/validatorgo#IsFloat)**`(str string, opts *IsFloatOpts)` | Checks if the string is a float. Supports `Min`/`Max` (≤ ≥), `Gt`/`Lt` (strict), and `Locale`. |
| **[IsDecimal](https://pkg.go.dev/github.com/bube054/validatorgo#IsDecimal)**`(str string, opts *IsDecimalOpts)` | Checks if the string represents a decimal number (`0.1`, `.3`, `1.1`, etc.). Supports `ForceDecimal`, `DecimalDigits`, and `Locale`. |
| **[IsDivisibleBy](https://pkg.go.dev/github.com/bube054/validatorgo#IsDivisibleBy)**`(str string, num int)` | Checks if the string is an integer divisible by `num`. |
| **[IsNumeric](https://pkg.go.dev/github.com/bube054/validatorgo#IsNumeric)**`(str string, opts *IsNumericOpts)` | Checks if the string is numeric. `NoSymbols` rejects `+`, `-`, `.`. Supports `Locale`. |
| **[IsPort](https://pkg.go.dev/github.com/bube054/validatorgo#IsPort)**`(str string)` | Checks if the string is a valid port number (0–65535). |
| **[IsHexadecimal](https://pkg.go.dev/github.com/bube054/validatorgo#IsHexadecimal)**`(str string)` | Checks if the string is a hexadecimal number. |
| **[IsOctal](https://pkg.go.dev/github.com/bube054/validatorgo#IsOctal)**`(str string)` | Checks if the string is a valid octal number. |
| **[IsLuhnNumber](https://pkg.go.dev/github.com/bube054/validatorgo#IsLuhnNumber)**`(str string)` | Checks if the string passes the [Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm). |

### Dates & Times

| Validator | Description |
| --- | --- |
| **[IsDate](https://pkg.go.dev/github.com/bube054/validatorgo#IsDate)**`(str string, opts *IsDateOpts)` | Checks if the string is a valid date (e.g. `2002-07-15`). Supports `Format` and `StrictMode`. |
| **[IsAfter](https://pkg.go.dev/github.com/bube054/validatorgo#IsAfter)**`(str string, opts *IsAfterOpts)` | Checks if the string is a date after the given `ComparisonDate` (defaults to now). |
| **[IsBefore](https://pkg.go.dev/github.com/bube054/validatorgo#IsBefore)**`(str string, opts *IsBeforeOpts)` | Checks if the string is a date before the given `ComparisonDate` (defaults to now). |
| **[IsTime](https://pkg.go.dev/github.com/bube054/validatorgo#IsTime)**`(str string, opts *IsTimeOpts)` | Checks if the string is a valid time (e.g. `23:01:59`). Supports `HourFormat` (`hour12`/`hour24`) and `Mode` (`default`/`withSeconds`). |
| **[IsRFC3339](https://pkg.go.dev/github.com/bube054/validatorgo#IsRFC3339)**`(str string)` | Checks if the string is a valid [RFC 3339](https://tools.ietf.org/html/rfc3339) date. |
| **[IsISO8601](https://pkg.go.dev/github.com/bube054/validatorgo#IsISO8601)**`(str string, opts *IsISO8601Opts)` | Checks if the string is a valid [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date. Supports `Strict` and `StrictSeparator`. |

> **Supported time layouts:** `Layout`, `ANSIC`, `UnixDate`, `RubyDate`, `RFC822(Z)`, `RFC850`, `RFC1123(Z)`, `Kitchen`, `Stamp`/`StampMilli`/`StampMicro`/`StampNano`, `DateTime`, `DateOnly`, `TimeOnly`, plus custom: `StandardDateLayout`, `SlashDateLayout`, `DateTimeLayout`, `ISO8601Layout`, `ISO8601ZuluLayout`, `ISO8601WithMillisecondsLayout`.

### Network, Web & URIs

| Validator | Description |
| --- | --- |
| **[IsEmail](https://pkg.go.dev/github.com/bube054/validatorgo#IsEmail)**`(str string, opts *IsEmailOpts)` | Checks if the string is an email. Supports `AllowDisplayName`, `RequireDisplayName`, `AllowUTF8LocalPart`, `RequireTld`, `AllowIpDomain`, `DomainSpecificValidation`, `BlacklistedChars`, `HostBlacklist`, `HostWhitelist`, `IgnoreMaxLength`. |
| **[IsURL](https://pkg.go.dev/github.com/bube054/validatorgo#IsURL)**`(str string, opts *IsURLOpts)` | Checks if the string is a URL. Highly configurable — see godoc for full options including `Protocols`, `RequireTld`, `RequireProtocol`, `RequireHost`, `RequirePort`, `HostWhitelist`/`HostBlacklist`, `AllowFragments`, `MaxAllowedLength`, etc. |
| **[IsIP](https://pkg.go.dev/github.com/bube054/validatorgo#IsIP)**`(str, version string)` | Checks if the string is an IP address. `version` is `"4"`, `"6"`, or empty for both. |
| **[IsIPRange](https://pkg.go.dev/github.com/bube054/validatorgo#IsIPRange)**`(str, version string)` | Checks if the string is an IP range. `version` is `"4"`, `"6"`, or empty for both. |
| **[IsFQDN](https://pkg.go.dev/github.com/bube054/validatorgo#IsFQDN)**`(str string, opts *IsFQDNOpts)` | Checks if the string is a fully qualified domain name. Supports `RequireTld`, `AllowUnderscores`, `AllowTrailingDot`, `AllowNumericTld`, `AllowWildcard`, `IgnoreMaxLength`. |
| **[IsMagnetURI](https://pkg.go.dev/github.com/bube054/validatorgo#IsMagnetURI)**`(str string)` | Checks if the string is a [Magnet URI](https://en.wikipedia.org/wiki/Magnet_URI_scheme). |
| **[IsMailtoURI](https://pkg.go.dev/github.com/bube054/validatorgo#IsMailtoURI)**`(str string, opts *IsMailToURIOpts)` | Checks if the string is a [Mailto URI](https://en.wikipedia.org/wiki/Mailto). Embeds `IsEmailOpts` for inner email validation. |
| **[IsDataURI](https://pkg.go.dev/github.com/bube054/validatorgo#IsDataURI)**`(str string)` | Checks if the string is in `data:` URI format. |
| **[IsMimeType](https://pkg.go.dev/github.com/bube054/validatorgo#IsMimeType)**`(str string)` | Checks if the string matches a valid MIME type format. |
| **[IsMacAddress](https://pkg.go.dev/github.com/bube054/validatorgo#IsMacAddress)**`(str string, opts *IsMacAddressOpts)` | Checks if the string is a MAC address. Supports `NoSeparators` and `Type` (`"48"` or `"64"`). |

### Encoding, Hashes & JSON

| Validator | Description |
| --- | --- |
| **[IsBase32](https://pkg.go.dev/github.com/bube054/validatorgo#IsBase32)**`(str string, opts *IsBase32Opts)` | Checks if the string is Base32-encoded. `Crockford` enables [Crockford's variant](http://www.crockford.com/base32.html). |
| **[IsBase58](https://pkg.go.dev/github.com/bube054/validatorgo#IsBase58)**`(str string)` | Checks if the string is Base58-encoded. |
| **[IsBase64](https://pkg.go.dev/github.com/bube054/validatorgo#IsBase64)**`(str string, opts *IsBase64Opts)` | Checks if the string is Base64-encoded. `UrlSafe` enables [URL-safe Base64](https://base64.guru/standards/base64url). |
| **[IsHash](https://pkg.go.dev/github.com/bube054/validatorgo#IsHash)**`(str, algorithm string)` | Checks if the string is a hash of the given algorithm: `crc32`, `crc32b`, `md4`, `md5`, `ripemd128`, `ripemd160`, `sha1`, `sha256`, `sha384`, `sha512`, `tiger128`, `tiger160`, `tiger192`. |
| **[IsMD5](https://pkg.go.dev/github.com/bube054/validatorgo#IsMD5)**`(str string)` | Convenience for `IsHash(str, "md5")`. |
| **[IsJSON](https://pkg.go.dev/github.com/bube054/validatorgo#IsJSON)**`(str string)` | Checks if the string is valid JSON (uses `json.Valid()`). |

### Identifiers, Codes & Versioning

| Validator | Description |
| --- | --- |
| **[IsUUID](https://pkg.go.dev/github.com/bube054/validatorgo#IsUUID)**`(str, version string)` | Checks if the string is an RFC 9562 UUID. `version` is `"1"`–`"5"`, or empty for any. |
| **[IsMongoID](https://pkg.go.dev/github.com/bube054/validatorgo#IsMongoID)**`(str string)` | Checks if the string is a valid MongoDB ObjectId (hex-encoded). |
| **[IsULID](https://pkg.go.dev/github.com/bube054/validatorgo#IsULID)**`(str string)` | Checks if the string is a [ULID](https://github.com/ulid/spec). |
| **[IsSemVer](https://pkg.go.dev/github.com/bube054/validatorgo#IsSemVer)**`(str string)` | Checks if the string follows [Semantic Versioning](https://semver.org/). |
| **[IsISIN](https://pkg.go.dev/github.com/bube054/validatorgo#IsISIN)**`(str string)` | Checks if the string is an [ISIN](https://en.wikipedia.org/wiki/International_Securities_Identification_Number) (stock/security identifier). |
| **[IsISBN](https://pkg.go.dev/github.com/bube054/validatorgo#IsISBN)**`(str, version string)` | Checks if the string is an [ISBN](https://en.wikipedia.org/wiki/ISBN). `version` is `"10"`, `"13"`, or empty for both. |
| **[IsISSN](https://pkg.go.dev/github.com/bube054/validatorgo#IsISSN)**`(str string, opts *IsISSNOpts)` | Checks if the string is an [ISSN](https://en.wikipedia.org/wiki/International_Standard_Serial_Number). Supports `CaseSensitive` and `RequireHyphen`. |
| **[IsISRC](https://pkg.go.dev/github.com/bube054/validatorgo#IsISRC)**`(str string, allowHyphens bool)` | Checks if the string is an [ISRC](https://en.wikipedia.org/wiki/International_Standard_Recording_Code). |
| **[IsEAN](https://pkg.go.dev/github.com/bube054/validatorgo#IsEAN)**`(str string)` | Checks if the string is a valid [EAN](https://en.wikipedia.org/wiki/International_Article_Number) (European Article Number). |
| **[IsIMEI](https://pkg.go.dev/github.com/bube054/validatorgo#IsIMEI)**`(str string, opts *IsIMEIOpts)` | Checks if the string is a valid [IMEI](https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity). `AllowHyphens` switches between formats. |
| **[IsFreightContainerID](https://pkg.go.dev/github.com/bube054/validatorgo#IsFreightContainerID)**`(str string)` | Alias for `IsISO6346`. |
| **[IsISO6346](https://pkg.go.dev/github.com/bube054/validatorgo#IsISO6346)**`(str string)` | Checks if the string is a valid [ISO 6346](https://en.wikipedia.org/wiki/ISO_6346) shipping container ID. |

### Geographic & Locale

| Validator | Description |
| --- | --- |
| **[IsLocale](https://pkg.go.dev/github.com/bube054/validatorgo#IsLocale)**`(str string)` | Checks if the string is a valid locale identifier. |
| **[IsISO6391](https://pkg.go.dev/github.com/bube054/validatorgo#IsISO6391)**`(str string)` | Checks if the string is a valid [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) language code. |
| **[IsISO31661Alpha2](https://pkg.go.dev/github.com/bube054/validatorgo#IsISO31661Alpha2)**`(str string)` | Checks if the string is a valid [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code. |
| **[IsISO31661Alpha3](https://pkg.go.dev/github.com/bube054/validatorgo#IsISO31661Alpha3)**`(str string)` | Checks if the string is a valid [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) country code. |
| **[IsISO31661Numeric](https://pkg.go.dev/github.com/bube054/validatorgo#IsISO31661Numeric)**`(str string)` | Checks if the string is a valid [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_numeric) numeric country code. |
| **[IsPostalCode](https://pkg.go.dev/github.com/bube054/validatorgo#IsPostalCode)**`(str, locale string)` | Checks if the string is a valid postal code. Pass `"any"` or empty to match any supported locale. |
| **[IsLatLong](https://pkg.go.dev/github.com/bube054/validatorgo#IsLatLong)**`(str string, opts *IsLatLongOpts)` | Checks if the string is a valid `lat,long` pair. `CheckDMS` validates degrees-minutes-seconds format. |
| **[IsLicensePlate](https://pkg.go.dev/github.com/bube054/validatorgo#IsLicensePlate)**`(str, locale string)` | Checks if the string matches a country's license plate format. Locales: `cs-CZ`, `de-DE`, `de-LI`, `en-IN`, `en-SG`, `en-PK`, `es-AR`, `hu-HU`, `pt-BR`, `pt-PT`, `sq-AL`, `sv-SE`, or `any`. |

### Personal & Identity

| Validator | Description |
| --- | --- |
| **[IsMobilePhone](https://pkg.go.dev/github.com/bube054/validatorgo#IsMobilePhone)**`(str string, locales []string, opts *IsMobilePhoneOpts)` | Checks if the string is a mobile phone number. Pass a slice of locales (e.g. `[]string{"sk-SK", "sr-RS"}`) or `nil` for any. `StrictMode` requires a leading `+`. Supports 150+ locales. |
| **[IsIdentityCard](https://pkg.go.dev/github.com/bube054/validatorgo#IsIdentityCard)**`(str, locale string)` | Checks if the string is a valid identity card code. Locales include `LK`, `PL`, `ES`, `FI`, `IN`, `IT`, `IR`, `MZ`, `NO`, `TH`, `zh-TW`, `he-IL`, `ar-LY`, `ar-TN`, `zh-CN`, `zh-HK`, `PK`, or `any`. |
| **[IsPassportNumber](https://pkg.go.dev/github.com/bube054/validatorgo#IsPassportNumber)**`(str, countryCode string)` | Checks if the string is a valid passport number for the given country. |
| **[IsTaxID](https://pkg.go.dev/github.com/bube054/validatorgo#IsTaxID)**`(str, locale string)` | Checks if the string is a valid Tax Identification Number. Defaults to `en-US`; `any` matches any supported locale. |
| **[IsVAT](https://pkg.go.dev/github.com/bube054/validatorgo#IsVAT)**`(str, countryCode string)` | Checks if the string is a valid VAT number for the given ISO 3166-1 alpha-2 country code (or `any`). |
| **[IsStrongPassword](https://pkg.go.dev/github.com/bube054/validatorgo#IsStrongPassword)**`(str string, opts *IsStrongPasswordOpts)` | Checks password strength against configurable rules (min length, casing, numbers, symbols, scoring). |

### Finance & Crypto

| Validator | Description |
| --- | --- |
| **[IsCreditCard](https://pkg.go.dev/github.com/bube054/validatorgo#IsCreditCard)**`(str string, opts *IsCreditCardOpts)` | Checks if the string is a credit card number. `Provider` can restrict to: `amex`, `bcglobal`, `carteblanche`, `dinersclub`, `discover`, `instapayment`, `jcb`, `koreanlocal`, `laser`, `maestro`, `mastercard`, `solo`, `switch`, `unionpay`, `visa`, `visamastercard`. |
| **[IsCurrency](https://pkg.go.dev/github.com/bube054/validatorgo#IsCurrency)**`(str string, opts *IsCurrencyOpts)` | Checks if the string is a valid currency amount. Highly configurable (symbol, separators, negatives, decimals — see godoc). |
| **[IsBic](https://pkg.go.dev/github.com/bube054/validatorgo#IsBic)**`(str string)` | Checks if the string is a BIC (Bank Identifier Code) / SWIFT code. |
| **[IsIBAN](https://pkg.go.dev/github.com/bube054/validatorgo#IsIBAN)**`(str, countryCode string)` | Checks if the string is a valid IBAN. Supports 70+ country codes (no checksum). |
| **[IsAbaRouting](https://pkg.go.dev/github.com/bube054/validatorgo#IsAbaRouting)**`(str string)` | Checks if the string is a valid ABA routing number (US bank). |
| **[IsIso4217](https://pkg.go.dev/github.com/bube054/validatorgo#IsIso4217)**`(str string)` | Checks if the string is a valid [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code. |
| **[IsBTCAddress](https://pkg.go.dev/github.com/bube054/validatorgo#IsBTCAddress)**`(str string)` | Checks if the string is a valid Bitcoin address. |
| **[IsEthereumAddress](https://pkg.go.dev/github.com/bube054/validatorgo#IsEthereumAddress)**`(str string)` | Checks if the string is an [Ethereum address](https://ethereum.org/) (no checksum validation). |

### Colors & Data Structures

| Validator | Description |
| --- | --- |
| **[IsHexColor](https://pkg.go.dev/github.com/bube054/validatorgo#IsHexColor)**`(str string)` | Checks if the string is a hex color. |
| **[IsHSL](https://pkg.go.dev/github.com/bube054/validatorgo#IsHSL)**`(str string)` | Checks if the string is an HSL color (CSS Colors Level 4). |
| **[IsRgbColor](https://pkg.go.dev/github.com/bube054/validatorgo#IsRgbColor)**`(str string, opts *IsRgbOpts)` | Checks if the string is an RGB or RGBA color. Supports `IncludePercentValues` and `AllowSpaces`. |
| **[IsArray](https://pkg.go.dev/github.com/bube054/validatorgo#IsArray)**`(str string, opts *IsArrayOpts)` | Checks if the value is an array. `Min`/`Max` constrain the length. |
| **[IsObject](https://pkg.go.dev/github.com/bube054/validatorgo#IsObject)**`(str string, opts *IsObjectOpts)` | Checks if the value is a JSON object (e.g. `{}`, `{"foo":"bar"}`). With `Strict: false`, arrays and `null` also pass. |
| **[IsIn](https://pkg.go.dev/github.com/bube054/validatorgo#IsIn)**`(str string, values []string)` | Checks if the string appears in a list of allowed values. |

---

## Sanitizers

```go
import "github.com/bube054/validatorgo/sanitizer"

clean := sanitizer.Whitelist("Hello123 World!", "a-zA-Z")
fmt.Println(clean) // "HelloWorld"
```

| Sanitizer | Description |
| --- | --- |
| **[Blacklist](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#Blacklist)**`(str, blacklistedChars string)` | Removes characters that match the blacklist (auto-escaped for regex). |
| **[Whitelist](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#Whitelist)**`(str, whitelistedChars string)` | Removes characters not in the whitelist (auto-escaped for regex). |
| **[Escape](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#Escape)**`(str string)` | Replaces `<`, `>`, `&`, `'`, and `"` with HTML entities. |
| **[Unescape](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#Unescape)**`(str string)` | Reverses `Escape` — replaces HTML entities with their characters. |
| **[LTrim](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#LTrim)**`(str, chars string)` | Trims characters (whitespace by default) from the left side. |
| **[RTrim](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#RTrim)**`(str, chars string)` | Trims characters (whitespace by default) from the right side. |
| **[Trim](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#Trim)**`(str, chars string)` | Trims characters (whitespace by default) from both sides. |
| **[StripLow](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#StripLow)**`(str string, keepNewLines bool)` | Removes characters with code < 32 and code 127 (mostly control chars). Preserves newlines if `keepNewLines` is true. |
| **[NormalizeEmail](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#NormalizeEmail)**`(email string, opts *NormalizeEmailOpts)` | Canonicalizes an email address. Provider-aware (Gmail, Outlook, Yahoo, iCloud) — handles dots, sub-addresses, and casing per provider rules. |
| **[ToBoolean](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#ToBoolean)**`(str string, strict bool)` | Converts the string to a bool. Everything except `"0"`, `"false"`, and `""` returns `true`. In strict mode, only `"1"` and `"true"` return `true`. |
| **[ToDate](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#ToDate)**`(str string)` | Converts the string to `*time.Time`, or `nil` if unparseable. |
| **[ToFloat](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#ToFloat)**`(str string)` | Converts the string to `float64`, returning an error on failure. |
| **[ToInt](https://pkg.go.dev/github.com/bube054/validatorgo/sanitizer#ToInt)**`(str string)` | Converts the string to `int`, returning an error on failure. (Beware of octals.) |

---

## Error Handling

All validators return `(bool, error)`. On success, `err` is `nil`. On failure, `err` is a `*ValidationError` with machine-readable context:

```go
ok, err := validatorgo.IsEmail("bad-email", nil)
if err != nil {
    var ve *validatorgo.ValidationError
    if errors.As(err, &ve) {
        fmt.Println(ve.Validator) // "IsEmail"
        fmt.Println(ve.Code)     // "INVALID_FORMAT"
        fmt.Println(ve.Message)  // "invalid email format"
    }
}
```

**Error codes:** `INVALID_FORMAT`, `TOO_SHORT`, `TOO_LONG`, `MISSING_TLD`, `BLACKLISTED_HOST`, `NOT_WHITELISTED_HOST`, `BLACKLISTED_CHAR`, `INVALID_LOCALE`, `OUT_OF_RANGE`, `INVALID_CHECKSUM`, `INVALID_LENGTH`, `MISSING_REQUIRED`, `INVALID_DOMAIN`, `UNSUPPORTED_VERSION`, `INVALID_VALUE`, `NOT_FOUND`.

### Pointer-typed option fields

Option fields whose defaults differ from Go's zero value use pointer types (e.g., `RequireTld *bool` defaults to `true`). Pass `nil` for the struct to get all defaults, or use the exported helpers to set specific fields:

```go
validatorgo.Bool(true)     // *bool
validatorgo.String("foo")  // *string
validatorgo.Int(42)        // *int
validatorgo.Uint(10)       // *uint
validatorgo.Float64(3.14)  // *float64
```

Example:

```go
ok, _ := validatorgo.IsFQDN("example.com", &validatorgo.IsFQDNOpts{
    RequireTld:      validatorgo.Bool(false), // override default (true)
    AllowTrailingDot: true,
})
```

---

## Maintainers

- [**Attah Gbubemi David** (@bube054)](https://github.com/bube054) — author

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) and the [Code of Conduct](CODE_OF_CONDUCT.md) before opening a PR.

A few quick guidelines:
- Match the style and structure of existing validators.
- Add tests for any new functionality (we aim for high coverage).
- Update this README's table when adding a new validator or sanitizer.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit). See [LICENSE](LICENSE) for details.