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
If you are unhappy using the long validatorgo, you can do something like this.
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
|  21  | **[IsDecimal(str string, opts IsDecimalOpts)](https://pkg.go.dev/github.com/bube054/validatorgo#IsDate)** | A Validator that checks if the string is a valid date. e.g. 2002-07-15. <br/> `IsDateOpts` is a struct which can contain the keys `Format`, `StrictMode`. <br/> Format: is a string and defaults to `validatorgo.ISO8601` if `"any"` or no value is provided. <br/> `StrictMode`: is a boolean and defaults to false. If `StrictMode` is set to true, the validator will reject strings different from Format. |

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