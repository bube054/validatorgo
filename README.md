# validatorgo

<img alt="Tag" src="https://img.shields.io/badge/tag-v0.1.0-blue?labelColor=gray"> <img alt="Go Version" src="https://img.shields.io/badge/Go->=1.13-00ADD8?labelColor=gray"> <img alt="Reference" src="https://img.shields.io/badge/-reference-00ADD8?logo=go&labelColor=gray"> <img alt="Tests" src="https://img.shields.io/badge/tests-passing-brightgreen?logo=github&labelColor=gray"> <img alt="Go Report" src="https://img.shields.io/badge/go_report-A%2B-00ADD8"> <img alt="Coverage" src="https://img.shields.io/badge/coverage-100%25-brightgreen?logo=codecov"> <img alt="Contributors" src="https://img.shields.io/badge/contributors-125-blueviolet"> <img alt="License" src="https://img.shields.io/badge/license-MIT-yellow">

A library of string validators and sanitizers, based on the js library [validator.js](https://github.com/validatorjs/validator.js)

## Rationale
Why not use popular go libraries like [Package validator](https://github.com/go-playground/validator) or [govalidator](https://github.com/asaskevich/govalidator)? Because they focus on validating structs rather than standalone strings and do not have a robust library of validators as the original js [validator.js](https://github.com/validatorjs/validator.js) package it was based on.
I made this project to actually be a dependency of another go open source library called [ginvalidator](https://github.com/bube054/ginvalidator) used to validate http request, is also the go version of the popular node/express library [express-validator](https://express-validator.github.io/docs/). Seeing that the most popular libraries in the go ecosystem where either overkill, not as robust or not a good fit for my use case, I decided to make my own library.

## Installation
Using go get.
 ```go
  go get github.com/bube054/validatorgo
 ```
Then import the package into your own code.
 ```go
  import (
    "fmt"
    "github.com/bube054/validatorgo"
  )
 ```
If you are unhappy using the long validatorgo package name, you can do this.
 ```go
  import (
    "fmt"
    vgo "github.com/bube054/validatorgo"
  )
 ```
Simple example
 ```go
  func main(){
    id := "5f2a6c69e1d7a4e0077b4e6b"
    validId := vgo.IsMongoID(id)
    fmt.Println(validId) // true
  }
 ```

# Validators
Here is a list of the validators currently available.
| #  | Validator | Description |
| --- | -------- | ----------- |
| 1   | **[Contains(str, seed string, opts ContainsOpt)](https://pkg.go.dev/github.com/bube054/validatorgo#Contains)** | A validator that checks if a string contains a seed. `ContainsOpt` defaults to `{ IgnoreCase: false, MinOccurrences: 1 }`. <br /> **[ContainsOpt](https://pkg.go.dev/github.com/bube054/validatorgo#ContainsOpts):** <br /> - `IgnoreCase`: Ignore case during comparison (default: `false`). <br /> - `MinOccurrences`: Minimum number of occurrences for the seed in the string (default: `1`). |
| 2   | **[Equals(str string, comparison string))](https://pkg.go.dev/github.com/bube054/validatorgo#Equals)** | A validator that checks if the string matches the comparison. |
| 3   | **[IsAbaRouting(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsAbaRouting)** | A validator that checks if the string is an ABA routing number for US bank account / cheque. |
| 4   | **[IsAfter(str string, comparisonDate string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsAfter)** | A validator that checks if the string is a date that is after the specified date. <br/> `IsAfterOpts` is a struct that defaults to `{ ComparisonDate: "" }`. <br/> `IsAfterOpts`: <br/> `ComparisonDate`: defaults to the current time. <br/> string layouts for str and `ComparisonDate` can be different layout. <br/> these are the only valid layouts from the `time` package e.g `Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly`.  |
| 5   | **[IsAlpha(str string, opts IsAlphaOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsAlpha)** | A validator that checks if the string contains only letters (a-zA-Z). <br> `IsAlphaOpts` is an optional struct that can be supplied with the following key(s): <br> `Ignore`: is the string to be ignored e.g. " -" will ignore spaces and -'s. <br> `Locale`: one of `('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'bn', 'cs-CZ', 'da-DK', 'de-DE', 'el-GR', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fa-IR', 'fi-FI', 'fr-CA', 'fr-FR', 'he', 'hi-IN', 'hu-HU', 'it-IT', 'kk-KZ', 'ko-KR', 'ja-JP', 'ku-IQ', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-BR', 'pt-PT', 'ru-RU', 'si-LK', 'sl-SI', 'sk-SK', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'th-TH', 'tr-TR', 'uk-UA')` and defaults to `en-US` if none is provided. |
| 6   | **[IsAlphanumeric(str string, opts IsAlphanumericOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsAlphanumeric)** |A validator that checks if the string contains only letters and numbers (a-zA-Z0-9). <br> `IsAlphaOpts` is an optional struct that can be supplied with the following key(s): <br> `Ignore`: is the string to be ignored e.g. " -" will ignore spaces and -'s. <br> `Locale`: one of `('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'bn', 'cs-CZ', 'da-DK', 'de-DE', 'el-GR', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fa-IR', 'fi-FI', 'fr-CA', 'fr-FR', 'he', 'hi-IN', 'hu-HU', 'it-IT', 'kk-KZ', 'ko-KR', 'ja-JP', 'ku-IQ', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-BR', 'pt-PT', 'ru-RU', 'si-LK', 'sl-SI', 'sk-SK', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'th-TH', 'tr-TR', 'uk-UA')` and defaults to `en-US` if none is provided. |
| 7   | **[IsAscii(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsAscii)** |A validator that checks if the string contains ASCII chars only. |
| 8   | **[IsBase32(str string, opts IsBase32Opts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsBase32)** | A validator that checks if the string is base32 encoded. <br/> `IsBase32Opts` defaults to `{ Crockford: false }`. When Crockford is true it tests the given base32 encoded string using [crockford's](http://www.crockford.com/base32.html) base32 alternative. |
| 9   | **[IsBase58(str string) bool](https://pkg.go.dev/github.com/bube054/validatorgo#IsBase58)** |  A validator that checks if the string is base58 encoded. |
| 10   | **[IsBase64(str string, opts IsBase64Opts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsBase64)** |  A validator that checks if the string is base64 encoded. <br/> `IsBase64Opts` is an optional struct which defaults to `{ UrlSafe: false }`. <br> When `UrlSafe` is true it tests the given base64 encoded string is [url safe](https://base64.guru/standards/base64url). |
| 11   | **[IsBefore(str string, comparisonDate string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsBefore)** | A validator that checks if the string is a date that is before the specified date. <br/> `IsAfterOpts` is a struct that defaults to `{ ComparisonDate: "" }`. <br/> `IsAfterOpts`: <br/> `ComparisonDate`: defaults to the current time. <br/> string layouts for str and `ComparisonDate` can be different layout. <br/> these are the only valid layouts from the `time` package e.g `Layout, ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, Kitchen, Stamp, StampMilli, StampMicro, StampNano, DateTime, DateOnly, TimeOnly`.  |
| 12   | **[IsBic(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsBic)** | A validator that checks if the string is a BIC (Bank Identification Code) or SWIFT code.  |
| 13   | **[IsBoolean(str string, opts IsBooleanOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsBoolean)** | A validator that check if the string is a boolean.  |
| 14   | **[IsBTCAddress(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsBTCAddress)** | A validator that checks if the string is a valid BTC address.  |
| 15   | **[IsByteLength(str string, opts IsByteLengthOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsByteLength)** | A validator that checks if the string's length (in UTF-8 bytes) falls in a range. <br/> `IsByteLengthOpts` is a struct which defaults to `{ Min: 0, Max: nil }`.  |
|  16  | **[IsCreditCard(str string, opts IsCreditCardOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsCreditCard)** | A validator that checks if the string is a credit card number. <br/> `IsCreditCardOpts` is an struct that can be supplied with the following key(s): <br/> `Provider`: is a key whose value should be a string, and defines the company issuing the credit card. Valid values include `amex, bcglobal, carteblanche, dinersclub, discover, instapayment, jcb, koreanlocal, laser, maestro, mastercard, solo, switch, unionpay, visa, visamastercard` or blank will check for any provider.  |
|  17  | **[IsCurrency(str string, opts IsCurrencyOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsCurrency)** | A validator that checks if the string is a valid currency amount. <br/> IsCurrencyOpts is a struct which defaults to `{ Symbol: '$', RequireSymbol: false, AllowSpaceAfterSymbol: false, SymbolAfterDigits: false, AllowNegatives: true, ParensForNegatives: false, NegativeSignBeforeDigits: false, NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: false, ThousandsSeparator: ',', DecimalSeparator: '.', AllowDecimal: true, RequireDecimal: false, DigitsAfterDecimal: [2], AllowSpaceAfterDigits: false }`. <br/> Note: The slice `DigitsAfterDecimal` is filled with the exact number of digits allowed not a range, for example a range 1 to 3 will be given as [1, 2, 3]. |
|  18  | **[IsCurrency(str string, opts IsCurrencyOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsCurrency)** | A validator that checks if the string is a valid currency amount. <br/> IsCurrencyOpts is a struct which defaults to `{ Symbol: '$', RequireSymbol: false, AllowSpaceAfterSymbol: false, SymbolAfterDigits: false, AllowNegatives: true, ParensForNegatives: false, NegativeSignBeforeDigits: false, NegativeSignAfterDigits: false, AllowNegativeSignPlaceholder: false, ThousandsSeparator: ',', DecimalSeparator: '.', AllowDecimal: true, RequireDecimal: false, DigitsAfterDecimal: [2], AllowSpaceAfterDigits: false }`. <br/> Note: The slice `DigitsAfterDecimal` is filled with the exact number of digits allowed not a range, for example a range 1 to 3 will be given as [1, 2, 3]. |
|  19  | **[IsDataURI(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsDataURI)** | A validator that checks if the string is a data uri format. |
|  20  | **[IsDate(str string, opts IsDateOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsDate)** | A Validator that checks if the string is a valid date. e.g. 2002-07-15. <br/> `IsDateOpts` is a struct which can contain the keys `Format`, `StrictMode`. <br/> Format: is a string and defaults to `validatorgo.ISO8601` if `"any"` or no value is provided. <br/> `StrictMode`: is a boolean and defaults to false. If `StrictMode` is set to true, the validator will reject strings different from Format. |
|  21  | **[IsDecimal(str string, opts IsDecimalOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsDate)** | A validator that check if the string represents a decimal number, such as 0.1, .3, 1.1, 1.00003, 4.0, etc. <br/> `IsDecimalOpts` is a struct which defaults to `{ForceDecimal: false, DecimalDigits: {Min: 0, Max: nil}, locale: 'en-US'}`. <br/> `locale`: determines the decimal separator and is one of ('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'cs-CZ', 'da-DK', 'de-DE', 'el-GR', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fa', 'fa-AF', 'fa-IR', 'fr-FR', 'fr-CA', 'hu-HU', 'id-ID', 'it-IT', 'ku-IQ', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pl-Pl', 'pt-BR', 'pt-PT', 'ru-RU', 'sl-SI', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'tr-TR', 'uk-UA', 'vi-VN'). `Locale` defaults to `"en-US"` <br/> `ForceDecimal`: simply means a decimal must be present "123" will not pass but "123.45" will pass <br/> `DecimalDigits`: is the allowed decimal range e.g `{Min: 3, Max: nil}` 123.456 |
|  22  | **[IsDivisibleBy(str string, num int)](https://pkg.go.dev/github.com/bube054/validatorgo#IsDivisibleBy)** | A validator thats checks if the string is a number(integer not a floating point) that is divisible by another(integer not a floating point). |
|  23  | **[IsEAN(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsEAN)** | IsEAN checks if the string is a valid [EAN](https://en.wikipedia.org/wiki/International_Article_Number) (European Article Number). |
|  24  | **[IsEmail(str string, opts IsEmailOpts)](https://pkg.go.dev/github.com/bube054/validatorgoI#sEmail)** | A validator that checks if the string is an email. <br/>  `IsEmailOpts` is a struct which defaults to `{ AllowDisplayName: false, RequireDisplayName: false, AllowUTF8LocalPart: true, RequireTld: true, AllowIpDomain: false, allow_underscores: false, DomainSpecificValidation: false, BlacklistedChars: ”, HostBlacklist: [] }`. <br/> `AllowDisplayName`: if set to true, the validator will also match Display Name <email-address>. <br/> `RequireDisplayName`: if set to true, the validator will reject strings without the format Display Name <email-address>. <br/> `AllowUTF8LocalPart`: if set to false, the validator will not allow any non-English UTF8 character in email address' local part. <br/> `RequireTld`: if set to false, email addresses without a TLD in their domain will also be matched. <br/> `IgnoreMaxLength`: if set to true, the validator will not check for the standard max length of an email. <br/> `AllowIpDomain`: if set to true, the validator will allow IP addresses in the host part. <br/> `DomainSpecificValidation`: is true, some additional validation will be enabled, e.g. disallowing certain syntactically valid email addresses that are rejected by Gmail. <br/> `BlacklistedChars`: receives a string, then the validator will reject emails that include any of the characters in the string, in the name part. <br/> `HostBlacklist`: if set to an array of strings and the part of the email after the @ symbol matches one of the strings defined in it, the validation fails. <br/>`HostWhitelist`: if set to an array of strings and the part of the email after the @ symbol matches none of the strings defined in it, the validation fails. |
|  23  | **[IsEAN(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsEAN)** | IsEAN checks if the string is a valid [EAN](https://en.wikipedia.org/wiki/International_Article_Number) (European Article Number). |
|  24  | **[IsEmpty(str string, opts IsEmptyOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsEmpty)** | A validator check if the string has a length of zero. <br/> `IsEmptyOpts` is a struct which defaults to `{ IgnoreWhitespace: false }`. |
|  25  | **[IsEthereumAddress(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsEthereumAddress)** | A validator checks if the string is an [Ethereum address](https://ethereum.org/). Does not validate address checksums. |
|  26  | **[IsFloat(str string, opts IsFloatOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsFloat)** | A validator that checks if the string is a float. <br/> `IsFloatOpts` is a struct which can contain the fields `Min`, `Max`, `Gt`, and/or `Lt` to validate the float is within boundaries it also has `Locale` as an option. <br/> `Min` and `Max`: are equivalent to 'greater or equal' and 'less or equal' <br/> `Gt` and `Lt`: are their strict counterparts. <br/> `Locale` determines the decimal separator and is one of `('ar', 'ar-AE', 'ar-BH', 'ar-DZ', 'ar-EG', 'ar-IQ', 'ar-JO', 'ar-KW', 'ar-LB', 'ar-LY', 'ar-MA', 'ar-QA', 'ar-QM', 'ar-SA', 'ar-SD', 'ar-SY', 'ar-TN', 'ar-YE', 'bg-BG', 'cs-CZ', 'da-DK', 'de-DE', 'en-AU', 'en-GB', 'en-HK', 'en-IN', 'en-NZ', 'en-US', 'en-ZA', 'en-ZM', 'eo', 'es-ES', 'fr-CA', 'fr-FR', 'hu-HU', 'it-IT', 'nb-NO', 'nl-NL', 'nn-NO', 'pl-PL', 'pt-BR', 'pt-PT', 'ru-RU', 'sl-SI', 'sr-RS', 'sr-RS@latin', 'sv-SE', 'tr-TR', 'uk-UA')` |
|  27  | **[IsFQDN(str string, opts IsFQDNopts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsFQDN)** | A validator that checks if the string is a fully qualified domain name (e.g. domain.com). <br> `IsFQDNopts` is a struct which defaults to `{ RequireTld: false, AllowUnderscores: false, AllowTrailingDot: false, AllowNumericTld: false, allow_wildcard: false, IgnoreMaxLength: false }` |
|  28  | **[IsFreightContainerID(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsFreightContainerID)** | A validator that checks alias for IsISO6346, check if the string is a valid [ISO 6346](https://en.wikipedia.org/wiki/ISO_6346) shipping container identification. <br/> |
|  29  | **[IsFullWidth(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsFullWidth)** | A validator that checks if the string contains any full-width chars. |
|  30  | **[IsHalfWidth(str string)](https://pkg.go.dev/github.com/bube054/validatorgo#IsHalfWidth)** | A validator that checks if the string contains any half-width chars. |

# Sanitizers
Here is a list of the validators currently available.
| # | Sanitizer |  description  | 
|:-----|:--------:| :--------|
|  |  | |

# Maintainers
- [bube054](https://github.com/bube054) - Attah Gbubemi David (author)
- [Yusufsalami](https://github.com/Yusufsalami) - Yusuf Salami

# Other related projects
- [ginvalidator](https://github.com/bube054/ginvalidator)
- [echovalidator](https://github.com/bube054/echovalidator)
- [fibervalidator](https://github.com/bube054/fibervalidator)
- [chivalidator](https://github.com/bube054/chivalidator)

# License
This project is licensed under the [MIT](https://opensource.org/license/mit). See the [LICENSE](https://github.com/bube054/validatorgo/blob/master/LICENSE) file for details.