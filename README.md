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
Simple example usage
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
| 1   | **[Contains(str string, seed string, options ContainsOpt)](https://github.com/bube054/validatorgo#Contains)** | A validator that checks if a string contains a seed. `ContainsOpt` defaults to `{ IgnoreCase: false, MinOccurrences: 1 }`. <br /> **[ContainsOpt](https://github.com/bube054/validatorgo#ContainsOpts):** <br /> - `IgnoreCase`: Ignore case during comparison (default: `false`). <br /> - `MinOccurrences`: Minimum number of occurrences for the seed in the string (default: `1`). |
| 2   | **[Equals(str string, comparison string) bool)](https://github.com/bube054/validatorgo#Equals)** | A validator that checks if the string matches the comparison. |
| 3   | **[IsAbaRouting(str string)](https://github.com/bube054/validatorgo#IsAbaRouting)** | A validator that checks if the string is an ABA routing number for US bank account / cheque. |
| 4   | **[IsAlpha(str string, opts IsAlphaOpts)](https://github.com/bube054/validatorgo#IsAlpha)** | A validator that checks if the string is an ABA routing number for US bank account / cheque. |

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
This project is licensed under the [MIT](https://github.com/bube054/validatorgo/blob/master/LICENSE). See the [LICENSE](https://github.com/bube054/validatorgo/blob/master/LICENSE) file for details.